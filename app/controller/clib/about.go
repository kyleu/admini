package clib

import (
	"net/http"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views"
)

func About(w http.ResponseWriter, r *http.Request) {
	controller.Act("about", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.SetTitleAndData("About "+util.AppName, util.AppName+" v"+as.BuildInfo.Version)
		return controller.Render(r, as, &views.About{Version: as.BuildInfo.Version, Started: as.Started}, ps, "about")
	})
}
