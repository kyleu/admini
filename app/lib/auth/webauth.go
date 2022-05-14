// Content managed by Project Forge, see [projectforge.md] for details.
package auth

import (
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/user"
	"admini.dev/admini/app/util"
)

func getAuthURL(prv *Provider, rc *fasthttp.RequestCtx, websess util.ValueMap, logger util.Logger) (string, error) {
	g, err := gothFor(rc, prv)
	if err != nil {
		return "", err
	}

	sess, err := g.BeginAuth(setState(rc))
	if err != nil {
		return "", err
	}

	u, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	err = cutil.StoreInSession(prv.ID, sess.Marshal(), rc, websess, logger)
	if err != nil {
		return "", err
	}

	return u, err
}

func getCurrentAuths(websess util.ValueMap) user.Accounts {
	authS, err := cutil.GetFromSession(WebAuthKey, websess)
	var ret user.Accounts
	if err == nil && authS != "" {
		ret = user.AccountsFromString(authS)
	}
	return ret
}

func setCurrentAuths(s user.Accounts, rc *fasthttp.RequestCtx, websess util.ValueMap, logger util.Logger) error {
	s.Sort()
	if len(s) == 0 {
		return cutil.RemoveFromSession(WebAuthKey, rc, websess, logger)
	}
	return cutil.StoreInSession(WebAuthKey, s.String(), rc, websess, logger)
}
