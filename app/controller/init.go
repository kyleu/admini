// Package controller $PF_IGNORE$
package controller

import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/user"
)

// Initialize app-specific system dependencies.
func initApp(*app.State) {
	user.SetPermissions(true, user.Perm("/admin", "*", true))
}

// Configure app-specific data for each request.
func initAppRequest(*app.State, *cutil.PageState) error {
	return nil
}

// Initialize system dependencies for the marketing site.
func initSite(*app.State) {
}

// Configure marketing site data for each request.
func initSiteRequest(*app.State, *cutil.PageState) error {
	return nil
}
