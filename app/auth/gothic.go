package auth

import (
	"encoding/base64"
	"net/url"

	"github.com/go-gem/sessions"
	"github.com/kyleu/admini/app/util"
	"github.com/markbates/goth"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

const SessKey = "auth"

func BeginAuthHandler(prv *Provider, ctx *fasthttp.RequestCtx, websess *sessions.Session) (string, error) {
	u, err := GetAuthURL(prv, ctx, websess)
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

func GetAuthURL(prv *Provider, ctx *fasthttp.RequestCtx, websess *sessions.Session) (string, error) {
	sess, err := prv.goth.BeginAuth(SetState(ctx))
	if err != nil {
		return "", err
	}

	u, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	err = StoreInSession(prv.ID, sess.Marshal(), ctx, websess)
	if err != nil {
		return "", err
	}

	return u, err
}

type Params struct {
	q *fasthttp.Args
}

func (p *Params) Get(key string) string {
	b := p.q.Peek(key)
	if len(b) > 0 {
		return string(b)
	}
	return ""
}

func CompleteUserAuth(prv *Provider, ctx *fasthttp.RequestCtx, websess *sessions.Session) (Sessions, error) {
	value, err := GetFromSession(prv.ID, websess)
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

	_, err = sess.Authorize(prv.goth, &Params{q: ctx.Request.URI().QueryArgs()})
	if err != nil {
		return nil, err
	}

	err = StoreInSession(prv.ID, sess.Marshal(), ctx, websess)
	if err != nil {
		return nil, err
	}

	gu, err := prv.goth.FetchUser(sess)
	if err != nil {
		return nil, err
	}

	return addToSession(gu.Provider, gu.Email, ctx, websess)
}

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

func getCurrentAuths(websess *sessions.Session) Sessions {
	authS, err := GetFromSession(SessKey, websess)
	var ret Sessions
	if err == nil && authS != "" {
		ret = SessionsFromString(authS)
	}
	return ret
}

func setCurrentAuths(websess *sessions.Session, s Sessions, ctx *fasthttp.RequestCtx) error {
	s.Sort()
	return StoreInSession(SessKey, s.String(), ctx, websess)
}

func GetFromSession(key string, websess *sessions.Session) (string, error) {
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

func StoreInSession(k string, v string, ctx *fasthttp.RequestCtx, websess *sessions.Session) error {
	websess.Values[k] = v
	return websess.Save(ctx)
}

func SetState(ctx *fasthttp.RequestCtx) string {
	state := ctx.Request.URI().QueryArgs().Peek("state")
	if len(state) > 0 {
		return string(state)
	}

	nonceBytes := util.RandomBytes(64)

	return base64.URLEncoding.EncodeToString(nonceBytes)
}

func validateState(ctx *fasthttp.RequestCtx, sess goth.Session) error {
	rawAuthURL, err := sess.GetAuthURL()
	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuthURL)
	if err != nil {
		return err
	}

	originalState := authURL.Query().Get("state")
	qs := string(ctx.Request.URI().QueryArgs().Peek("state"))
	if originalState != "" && (originalState != qs) {
		return errors.New("state token mismatch")
	}
	return nil
}

func removeProviderData(ctx *fasthttp.RequestCtx, websess *sessions.Session, prvKeys ...string) error {
	for _, k := range prvKeys {
		delete(websess.Values, k)
	}
	return websess.Save(ctx)
}

func Logout(ctx *fasthttp.RequestCtx, websess *sessions.Session, prvKeys ...string) error {
	a := getCurrentAuths(websess)
	a = a.Purge(prvKeys...)
	for _, k := range prvKeys {
		delete(websess.Values, k)
	}
	return websess.Save(ctx)
}
