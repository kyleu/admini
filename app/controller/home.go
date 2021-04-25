package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return render(r, w, as, &views.Home{}, ps)
	})
}
