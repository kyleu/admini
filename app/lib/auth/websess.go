// Content managed by Project Forge, see [projectforge.md] for details.
package auth

import (
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app/controller/csession"
	"admini.dev/admini/app/lib/user"
	"admini.dev/admini/app/util"
)

const WebAuthKey = "auth"

func addToSession(
	provider string, email string, token string, rc *fasthttp.RequestCtx, websess util.ValueMap, logger util.Logger,
) (*user.Account, user.Accounts, error) {
	ret := getCurrentAuths(websess)
	s := &user.Account{Provider: provider, Email: email, Token: token}
	for _, x := range ret {
		if x.Provider == s.Provider && x.Email == s.Email {
			return s, ret, nil
		}
	}
	ret = append(ret, s)
	err := setCurrentAuths(ret, rc, websess, logger)
	if err != nil {
		return nil, nil, err
	}
	return s, ret, nil
}

func removeProviderData(rc *fasthttp.RequestCtx, websess util.ValueMap, logger util.Logger) error {
	dirty := false
	for s := range websess {
		if isProvider(s) {
			logger.Debug("removing auth info for provider [" + s + "]")
			dirty = true
			delete(websess, s)
		}
	}
	if dirty {
		return csession.SaveSession(rc, websess, logger)
	}
	return nil
}

func isProvider(k string) bool {
	for _, x := range AvailableProviderKeys {
		if x == k {
			return true
		}
	}
	return false
}
