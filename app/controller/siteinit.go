// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"context"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
)

// Initialize system dependencies for the marketing site.
func initSite(_ context.Context, _ *app.State, _ util.Logger) error {
	return nil
}

// Configure marketing site data for each request.
func initSiteRequest(*app.State, *cutil.PageState) error {
	return nil
}
