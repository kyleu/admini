package auth

import (
	"github.com/go-gem/sessions"
	"github.com/valyala/fasthttp"
)

func getAuthURL(prv *Provider, ctx *fasthttp.RequestCtx, websess *sessions.Session) (string, error) {
	sess, err := prv.goth.BeginAuth(setState(ctx))
	if err != nil {
		return "", err
	}

	u, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	err = storeInSession(prv.ID, sess.Marshal(), ctx, websess)
	if err != nil {
		return "", err
	}

	return u, err
}

func getCurrentAuths(websess *sessions.Session) Sessions {
	authS, err := getFromSession(SessKey, websess)
	var ret Sessions
	if err == nil && authS != "" {
		ret = SessionsFromString(authS)
	}
	return ret
}

func setCurrentAuths(websess *sessions.Session, s Sessions, ctx *fasthttp.RequestCtx) error {
	s.Sort()
	return storeInSession(SessKey, s.String(), ctx, websess)
}
