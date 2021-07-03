package controller

import (
	"fmt"
	"time"

	"github.com/go-gem/sessions"
	"github.com/gorilla/securecookie"
	"github.com/kyleu/admini/app/auth"
	"github.com/kyleu/admini/app/user"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/menu"

	"go.uber.org/zap"

	"github.com/kyleu/admini/views/verror"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
)

const (
	defaultSearchPath  = "/search"
	defaultProfilePath = "/profile"
)

func act(key string, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	ps := actPrepare(ctx)
	clean(ps)
	actComplete(key, ps, ctx, f)
}

func actWorkspace(key string, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	ps := actPrepare(ctx)
	actComplete(key, ps, ctx, f)
}

func actPrepare(ctx *fasthttp.RequestCtx) *cutil.PageState {
	logger := currentApp.RootLogger.With(zap.String("path", string(ctx.Request.URI().Path())))

	if store == nil {
		store = initStore()
	}
	session, err := store.Get(ctx, util.AppKey)
	if err != nil {
		logger.Warnf("error retrieving session: %+v", err)
	}
	flashes := util.StringArrayFromInterfaces(session.Flashes())
	if session.IsNew || len(flashes) > 0 {
		session.Options = &sessions.Options{Path: "/", HttpOnly: true /* , SameSite: http.SameSiteStrictMode */}
		err = session.Save(ctx)
		if err != nil {
			logger.Warnf("can't save session: %+v", err)
		}
	}

	prof, err := loadProfile(session)
	if err != nil {
		logger.Warnf("can't load profile: %+v", err)
	}

	var a auth.Sessions
	authX, ok := session.Values["auth"]
	if ok {
		authS, ok := authX.(string)
		if ok {
			a = auth.SessionsFromString(authS)
		}
	}

	return &cutil.PageState{
		Method:  string(ctx.Method()),
		URI:     ctx.Request.URI(),
		Flashes: flashes,
		Session: session,
		Profile: prof,
		Auth:    a,
		Icons:   initialIcons,
		Logger:  logger,
	}
}

func initStore(keyPairs ...[]byte) *sessions.CookieStore {
	ret := sessions.NewCookieStore([]byte(sessionKey))
	for _, x := range ret.Codecs {
		c, ok := x.(*securecookie.SecureCookie)
		if ok {
			c.MaxLength(65536)
		}
	}
	return ret
}

func loadProfile(session *sessions.Session) (*user.Profile, error) {
	return &user.Profile{Name: "Test"}, nil
}

func actComplete(key string, ps *cutil.PageState, ctx *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	status := fasthttp.StatusOK
	cutil.WriteCORS(ctx)
	startNanos := time.Now().UnixNano()
	redir, err := f(currentApp, ps)
	if err != nil {
		status = fasthttp.StatusInternalServerError
		ctx.SetStatusCode(status)

		ps.Logger.Errorf("error running action [%s]: %+v", key, err)

		if len(ps.Breadcrumbs) == 0 {
			ps.Breadcrumbs = []string{"Error"}
		}
		errDetail := util.GetErrorDetail(err)
		page := &verror.Error{Err: errDetail}

		clean(ps)
		redir, err = render(ctx, currentApp, page, ps)
		if err != nil {
			_, _ = ctx.Write([]byte(fmt.Sprintf("error while running error handler: %+v", err)))
		}
	}
	if redir != "" {
		ctx.Response.Header.Set("Location", redir)
		status = fasthttp.StatusFound
		ctx.SetStatusCode(status)
	}
	elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	l := ps.Logger.With(zap.String("method", ps.Method), zap.Int("status", status), zap.Float64("elapsed", elapsedMillis))
	l.Debugf("processed request in [%.3fms] (render: %.3fms)", elapsedMillis, ps.RenderElapsed)
}

func clean(ps *cutil.PageState) {
	if ps.RootIcon == "" {
		ps.RootIcon = "app"
	}
	if ps.RootPath == "" {
		ps.RootPath = "/"
	}
	if ps.RootTitle == "" {
		ps.RootTitle = util.AppName
	}
	if ps.SearchPath == "" {
		ps.SearchPath = defaultSearchPath
	}
	if ps.ProfilePath == "" {
		ps.ProfilePath = defaultProfilePath
	}
	if len(ps.Menu) == 0 {
		ps.Menu = menu.For(currentApp)
	}
}
