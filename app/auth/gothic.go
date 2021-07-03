package auth

import (
	"github.com/go-gem/sessions"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func BeginAuthHandler(prv *Provider, ctx *fasthttp.RequestCtx, websess *sessions.Session, logger *zap.SugaredLogger) (string, error) {
	u, err := getAuthURL(prv, ctx, websess, logger)
	if err != nil {
		return "", err
	}
	refer := string(ctx.Request.URI().QueryArgs().Peek("refer"))
	if refer != "" {
		_ = StoreInSession("auth-refer", refer, ctx, websess, logger)
	}
	return u, nil
}

func CompleteUserAuth(prv *Provider, ctx *fasthttp.RequestCtx, websess *sessions.Session, logger *zap.SugaredLogger) (*Session, Sessions, error) {
	value, err := getFromSession(prv.ID, websess)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = removeProviderData(ctx, websess, logger)
	}()

	sess, err := prv.goth.UnmarshalSession(value)
	if err != nil {
		return nil, nil, err
	}

	err = validateState(ctx, sess)
	if err != nil {
		return nil, nil, err
	}

	user, err := prv.goth.FetchUser(sess)
	if err == nil {
		return addToSession(user.Provider, user.Email, ctx, websess, logger)
	}

	_, err = sess.Authorize(prv.goth, &params{q: ctx.Request.URI().QueryArgs()})
	if err != nil {
		return nil, nil, err
	}

	err = StoreInSession(prv.ID, sess.Marshal(), ctx, websess, logger)
	if err != nil {
		return nil, nil, err
	}

	gu, err := prv.goth.FetchUser(sess)
	if err != nil {
		return nil, nil, err
	}

	return addToSession(gu.Provider, gu.Email, ctx, websess, logger)
}

func Logout(ctx *fasthttp.RequestCtx, websess *sessions.Session, logger *zap.SugaredLogger, prvKeys ...string) error {
	a := getCurrentAuths(websess)
	n := a.Purge(prvKeys...)
	dirty := false
	if len(a) != len(n) {
		dirty = true
		err := setCurrentAuths(n, ctx, websess, logger)
		if err != nil {
			return err
		}
	}
	for _, k := range prvKeys {
		dirty = true
		delete(websess.Values, k)
	}
	if dirty {
		return SaveSession(ctx, websess, logger)
	}
	return nil
}

