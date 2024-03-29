package controller

import (
	"context"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/user"
	"admini.dev/admini/app/util"
)

// Initialize app-specific system dependencies.
func initApp(_ context.Context, _ *app.State, _ util.Logger) error {
	user.SetPermissions(true, user.Perm("/admin", "*", true))
	return nil
}

// Configure app-specific data for each request.
func initAppRequest(*app.State, *cutil.PageState) error {
	return nil
}
