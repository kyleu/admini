package controller

import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/user"
	"admini.dev/admini/app/util"
)

// Initialize app-specific system dependencies.
func initApp(*app.State, util.Logger) {
	user.SetPermissions(true, user.Perm("/admin", "*", true))
}

// Configure app-specific data for each request.
func initAppRequest(*app.State, *cutil.PageState) error {
	return nil
}
