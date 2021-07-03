package auth

import (
	"github.com/go-gem/sessions"
	"github.com/valyala/fasthttp"
)

func BeginAuthHandler(prv *Provider, ctx *fasthttp.RequestCtx, websess *sessions.Session) (string, error) {
	u, err := getAuthURL(prv, ctx, websess)
	if err != nil {
		return "", err
	}
	refer := string(ctx.Request.URI().QueryArgs().Peek("refer"))
	if refer != "" {
		websess.Values["auth-refer"] = refer
		_ = websess.Save(ctx)
	}
	return u, nil
}

func CompleteUserAuth(prv *Provider, ctx *fasthttp.RequestCtx, websess *sessions.Session) (Sessions, error) {
	value, err := getFromSession(prv.ID, websess)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = removeProviderData(ctx, websess, prv.ID)
	}()

	sess, err := prv.goth.UnmarshalSession(value)
	if err != nil {
		return nil, err
	}

	err = validateState(ctx, sess)
	if err != nil {
		return nil, err
	}

	user, err := prv.goth.FetchUser(sess)
	if err == nil {
		return addToSession(user.Provider, user.Email, ctx, websess)
	}

	_, err = sess.Authorize(prv.goth, &params{q: ctx.Request.URI().QueryArgs()})
	if err != nil {
		return nil, err
	}

	err = storeInSession(prv.ID, sess.Marshal(), ctx, websess)
	if err != nil {
		return nil, err
	}

	gu, err := prv.goth.FetchUser(sess)
	if err != nil {
		return nil, err
	}

	return addToSession(gu.Provider, gu.Email, ctx, websess)
}

func Logout(ctx *fasthttp.RequestCtx, websess *sessions.Session, prvKeys ...string) error {
	a := getCurrentAuths(websess)
	a = a.Purge(prvKeys...)
	for _, k := range prvKeys {
		delete(websess.Values, k)
	}
	return websess.Save(ctx)
}

