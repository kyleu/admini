package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/user"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vprofile"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	act("profile", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		profile := &user.Profile{Name: "Test"}
		ps.Title = "Profile"
		ps.Data = profile
		return render(r, w, as, &vprofile.Profile{Profile: profile}, ps, "Profile")
	})
}
