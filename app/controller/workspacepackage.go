package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kyleu/admini/app/source"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vworkspace"

	"github.com/kyleu/admini/app"
)

func handlePackage(w http.ResponseWriter, r *http.Request, as *app.State, ps *cutil.PageState, src *source.Source, mp *schema.ModelPackage, remaining []string) (string, error) {
	ps.Data = mp
	if len(remaining) == 0 {
		page := &vworkspace.PackageDetail{Pkg: mp, Source: src.Key}
		return render(r, w, as, page, ps, mp.Path()...)
	}
	page := &views.TODO{Message: fmt.Sprintf("unhandled package action [%v]", strings.Join(remaining, "/"))}
	return render(r, w, as, page, ps, mp.Path()...)
}
