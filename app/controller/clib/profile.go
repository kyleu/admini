package clib

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/csession"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/theme"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/vprofile"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	controller.Act("profile", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(w, r, as, ps)
	})
}

func ProfileSite(w http.ResponseWriter, r *http.Request) {
	controller.ActSite("profile", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(w, r, as, ps)
	})
}

func profileAction(w http.ResponseWriter, r *http.Request, as *app.State, ps *cutil.PageState) (string, error) {
	ps.SetTitleAndData("Profile", ps.Profile)
	thm := as.Themes.Get(ps.Profile.Theme, ps.Logger)

	prvs, err := as.Auth.Providers(ps.Logger)
	if err != nil {
		return "", errors.Wrap(err, "can't load providers")
	}

	redir := "/"
	ref := r.Header.Get("Referer")
	if ref != "" {
		u, err := url.Parse(ref)
		if err == nil && u != nil && u.Path != cutil.DefaultProfilePath {
			redir = u.Path
		}
	}
	ps.DefaultNavIcon = "profile"
	page := &vprofile.Profile{Profile: ps.Profile, Theme: thm, Providers: prvs, Referrer: redir}
	return controller.Render(r, as, page, ps, "Profile")
}

func ProfileSave(w http.ResponseWriter, r *http.Request) {
	controller.Act("profile.save", w, r, func(_ *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(r, ps.RequestBody)
		if err != nil {
			return "", err
		}

		n := ps.Profile.Clone()

		referrerDefault := util.OrDefault(frm.GetStringOpt("referrer"), cutil.DefaultProfilePath)

		n.Name = frm.GetStringOpt("name")
		n.Mode = frm.GetStringOpt("mode")
		n.Theme = frm.GetStringOpt("theme")
		if n.Theme == theme.Default.Key {
			n.Theme = ""
		}

		err = csession.SaveProfile(n, w, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}

		return controller.ReturnToReferrer("Saved profile", referrerDefault, ps)
	})
}
