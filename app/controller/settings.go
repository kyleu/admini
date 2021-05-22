package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsettings"
)

func Settings(w http.ResponseWriter, r *http.Request) {
	act("settings", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		settings := map[string]string{"settings": "TODO"}
		ps.Title = "Settings"
		ps.Data = settings
		return render(r, w, as, &vsettings.Settings{Settings: settings}, ps, "settings")
	})
}
