package controller

import (
	"fmt"

	"github.com/kyleu/admini/app/auth"
	"github.com/kyleu/admini/views/vauth"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
)

func AuthDetail(ctx *fasthttp.RequestCtx) {
	act("auth.detail", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		prv, err := getProvider(ctx)
		if err != nil {
			return "", err
		}
		u, _, err := auth.CompleteUserAuth(prv, ctx, ps.Session, ps.Logger)
		if err == nil {
			return render(ctx, as, &vauth.Detail{Provider: prv, Session: u}, ps)
		}
		return auth.BeginAuthHandler(prv, ctx, ps.Session)
	})
}

func AuthCallback(ctx *fasthttp.RequestCtx) {
	act("auth.callback", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		prv, err := getProvider(ctx)
		if err != nil {
			return "", err
		}
		u, _, err := auth.CompleteUserAuth(prv, ctx, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}
		referX, ok := ps.Session.Values["auth-refer"]
		if ok {
			refer, ok := referX.(string)
			if ok && refer != "" {
				msg := fmt.Sprintf("signed in to %s as [%s]", auth.AvailableProviderNames[prv.ID], u.Email)
				delete(ps.Session.Values, "auth-refer")
				_ = ps.Session.Save(ctx)
				return flashAndRedir(true, msg, refer, ctx, ps)
			}
		}
		return render(ctx, as, &vauth.Detail{Provider: prv, Session: u}, ps)
	})
}

func getProvider(ctx *fasthttp.RequestCtx) (*auth.Provider, error) {
	key, err := ctxRequiredString(ctx, "key", false)
	if err != nil {
		return nil, err
	}
	prvs, err := currentApp.Auth.Providers()
	if err != nil {
		return nil, errors.Wrap(err, "can't load providers")
	}
	prv := prvs.Get(key)
	if prv == nil {
		return nil, errors.Errorf("no provider available with id [%s]", key)
	}
	return prv, nil
}

func AuthLogout(ctx *fasthttp.RequestCtx) {
	act("auth.logout", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		err = auth.Logout(ctx, ps.Session, key)
		if err != nil {
			return "", err
		}

		return "/profile", nil
	})
}
