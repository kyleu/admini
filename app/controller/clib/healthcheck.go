// Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app/controller/cutil"
)

func Healthcheck(rc *fasthttp.RequestCtx) {
	x := map[string]string{"status": "OK"}
	_, _ = cutil.RespondJSON(rc, "", x)
}
