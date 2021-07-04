package controller

import (
	"net/url"

	"github.com/go-gem/sessions"
	"github.com/kyleu/admini/app/auth"
	"github.com/kyleu/admini/app/user"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vuser"
)

func Profile(ctx *fasthttp.RequestCtx) {
	act("profile", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Profile"
		ps.Data = ps.Profile
		thm := as.Themes.Get(ps.Profile.Theme)

		prvs, err := as.Auth.Providers()
		if err != nil {
			return "", errors.Wrap(err, "can't load providers")
		}

		redir := "/"
		ref := string(ctx.Request.Header.Peek("Referer"))
		if ref != "" {
			u, err := url.Parse(ref)
			if err == nil && u != nil {
				redir = u.Path
			}
		}

		page := &vuser.Profile{Profile: ps.Profile, Theme: thm, Providers: prvs, Referrer: redir}
		return render(ctx, as, page, ps, "Profile")
	})
}

func ProfileSave(ctx *fasthttp.RequestCtx) {
	act("profile.save", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		dirty := false

		def := user.DefaultProfile.Clone()
		n := ps.Profile.Clone()

		name, _ := frm.GetString("name", true)
		n.Name = name
		if name != def.Name {
			dirty = true
		}
		mode, _ := frm.GetString("mode", true)
		n.Mode = mode
		if mode != def.Mode {
			dirty = true
		}
		theme, _ := frm.GetString("theme", true)
		if theme == "default" {
			theme = ""
		}
		n.Theme = theme
		if theme != def.Theme {
			dirty = true
		}

		if dirty {
			err := auth.StoreInSession("profile", util.ToJSON(n), ctx, ps.Session, ps.Logger)
			if err != nil {
				return "", errors.Wrap(err, "unable to save profile in session")
			}
		}

		return flashAndRedir(true, "profile saved", "/profile", ctx, ps)
	})
}

func loadProfile(session *sessions.Session) (*user.Profile, error) {
	x, ok := session.Values["profile"]
	if !ok {
		return user.DefaultProfile.Clone(), nil
	}
	s, ok := x.(string)
	if !ok {
		return user.DefaultProfile.Clone(), nil
	}
	p := &user.Profile{}
	err := util.FromJSON([]byte(s), p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
