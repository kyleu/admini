package controller

import (
	"fmt"

	"github.com/go-gem/sessions"
	"github.com/gorilla/securecookie"
	"github.com/kyleu/admini/app/user"
)

func initStore(keyPairs ...[]byte) *sessions.CookieStore {
	ret := sessions.NewCookieStore([]byte(sessionKey))
	for _, x := range ret.Codecs {
		c, ok := x.(*securecookie.SecureCookie)
		if ok {
			c.MaxLength(65536)
		}
	}
	return ret
}

func loadProfile(session *sessions.Session) (*user.Profile, error) {
	x, ok := session.Values["profile"]
	if !ok {
		return user.DefaultProfile, nil
	}
	println(fmt.Sprintf("#################: %v (%T)", x, x))
	return &user.Profile{Name: "Test", Mode: "dark"}, nil
}
