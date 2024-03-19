// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/http"

	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	x := util.ValueMap{"status": "OK"}
	_, _ = cutil.RespondJSON(w, "", x)
}
