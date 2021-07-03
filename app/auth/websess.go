package auth

import (
	"github.com/go-gem/sessions"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

const SessKey = "auth"

func addToSession(provider string, email string, ctx *fasthttp.RequestCtx, websess *sessions.Session) (Sessions, error) {
	ret := getCurrentAuths(websess)
	s := &Session{Provider: provider, Email: email}
	for _, x := range ret {
		if x.Provider == s.Provider && x.Email == s.Email {
			return ret, nil
		}
	}
	ret = append(ret, s)
	err := setCurrentAuths(websess, ret, ctx)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func getFromSession(key string, websess *sessions.Session) (string, error) {
	value, ok := websess.Values[key]
	if !ok {
		return "", errors.Errorf("could not find a matching session value with key [%s] for this request", key)
	}
	s, ok := value.(string)
	if !ok {
		return "", errors.Errorf("session value with key [%s] is of type [%T], not [string]", key, value)
	}
	return s, nil
}

func storeInSession(k string, v string, ctx *fasthttp.RequestCtx, websess *sessions.Session) error {
	websess.Values[k] = v
	return websess.Save(ctx)
}

func removeProviderData(ctx *fasthttp.RequestCtx, websess *sessions.Session, prvKeys ...string) error {
	for _, k := range prvKeys {
		delete(websess.Values, k)
	}
	return websess.Save(ctx)
}
