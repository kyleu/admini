package controller

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views/verror"
	"github.com/valyala/fasthttp"

	"github.com/pkg/errors"

	"github.com/go-gem/sessions"
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/layout"
)

var (
	currentApp   *app.State
	initialIcons = []string{"search"}
)

var sessionKey = func() string {
	x := os.Getenv("SESSION_KEY")
	if x == "" {
		x = "random_secret_key"
	}
	return x
}()

var store = sessions.NewCookieStore([]byte(sessionKey))

func SetState(a *app.State) {
	currentApp = a
}

func ctxRequiredString(ctx *fasthttp.RequestCtx, key string, allowEmpty bool) (string, error) {
	v, ok := ctx.UserValue(key).(string)
	if !ok || ((!allowEmpty) && v == "") {
		return v, errors.Errorf("must provide [%s] in path", key)
	}
	return v, nil
}

func render(ctx *fasthttp.RequestCtx, appState *app.State, page layout.Page, ps *cutil.PageState, breadcrumbs ...string) (string, error) {
	defer func() {
		x := recover()
		if x != nil {
			ps.Logger.Error(fmt.Sprintf("Error processing template: %+v", x))
			switch t := x.(type) {
			case error:
				ed := util.GetErrorDetail(t)
				verror.WriteDetail(ctx, ed, currentApp, ps)
			default:
				ed := &util.ErrorDetail{Message: fmt.Sprintf("%v", t)}
				verror.WriteDetail(ctx, ed, currentApp, ps)
			}
		}
	}()
	ps.Breadcrumbs = append(ps.Breadcrumbs, breadcrumbs...)
	ct := cutil.GetContentType(ctx)
	if ps.Data != nil {
		if cutil.IsContentTypeJSON(ct) {
			return cutil.RespondJSON(ctx, "", ps.Data)
		} else if cutil.IsContentTypeXML(ct) {
			return cutil.RespondXML(ctx, "", ps.Data)
		}
	}
	startNanos := time.Now().UnixNano()
	ctx.Response.Header.SetContentType("text/html; charset=UTF-8")
	views.WriteRender(ctx, page, appState, ps)
	ps.RenderElapsed = float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	return "", nil
}

func renderWS(req *cutil.WorkspaceRequest, page layout.Page, bc ...string) (string, error) {
	return render(req.Ctx, req.AS, page, req.PS, bc...)
}

func ersp(msg string, args ...interface{}) (string, error) {
	return "", errors.Errorf(msg, args...)
}

func flashAndRedir(success bool, msg string, redir string, ctx *fasthttp.RequestCtx, ps *cutil.PageState) (string, error) {
	status := "error"
	if success {
		status = "success"
	}
	ps.Session.AddFlash(fmt.Sprintf("%s:%s", status, msg))
	err := ps.Session.Save(ctx)
	if err != nil {
		return "", errors.Wrap(err, "unable to save flash session")
	}

	if strings.HasPrefix(redir, "/") {
		return redir, nil
	}
	if strings.HasPrefix(redir, "http") {
		ps.Logger.Warn("flash redirect attempted for non-local request")
		return "/", nil
	}
	return redir, nil
}

func flashError(err error, redir string, ctx *fasthttp.RequestCtx, ps *cutil.PageState) (string, error) {
	return flashAndRedir(false, err.Error(), redir, ctx, ps)
}
