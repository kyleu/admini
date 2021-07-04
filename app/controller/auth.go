package controller

import (
	"fmt"

	"github.com/kyleu/admini/app/auth"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
)

func AuthDetail(ctx *fasthttp.RequestCtx) {
	act("auth.detail", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		prv, err := getProvider(as, ctx)
		if err != nil {
			return "", err
		}
		u, _, err := auth.CompleteUserAuth(prv, ctx, ps.Session, ps.Logger)
		if err == nil {
			return signInReturn(prv, u, ctx, as, ps)
		}
		return auth.BeginAuthHandler(prv, ctx, ps.Session, ps.Logger)
	})
}

func AuthCallback(ctx *fasthttp.RequestCtx) {
	act("auth.callback", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		prv, err := getProvider(as, ctx)
		if err != nil {
			return "", err
		}
		u, _, err := auth.CompleteUserAuth(prv, ctx, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}
		return signInReturn(prv, u, ctx, as, ps)
	})
}

func AuthLogout(ctx *fasthttp.RequestCtx) {
	act("auth.logout", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		err = auth.Logout(ctx, ps.Session, ps.Logger, key)
		if err != nil {
			return "", err
		}

		return "/profile", nil
	})
}

func signInReturn(prv *auth.Provider, u *auth.Session, ctx *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (string, error) {
	refer := ""
	referX, ok := ps.Session.Values[auth.ReferKey]
	if ok {
		refer, ok = referX.(string)
		if ok {
			_ = auth.RemoveFromSession(auth.ReferKey, ctx, ps.Session, ps.Logger)
		}
	}
	if refer == "" {
		refer = "/profile"
	}
	msg := fmt.Sprintf("signed in to %s as [%s]", auth.AvailableProviderNames[prv.ID], u.Email)
	return flashAndRedir(true, msg, refer, ctx, ps)
}

func getProvider(as *app.State, ctx *fasthttp.RequestCtx) (*auth.Provider, error) {
	key, err := ctxRequiredString(ctx, "key", false)
	if err != nil {
		return nil, err
	}
	prvs, err := as.Auth.Providers()
	if err != nil {
		return nil, errors.Wrap(err, "can't load providers")
	}
	prv := prvs.Get(key)
	if prv == nil {
		return nil, errors.Errorf("no provider available with id [%s]", key)
	}
	return prv, nil
}
