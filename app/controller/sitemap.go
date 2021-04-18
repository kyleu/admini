package controller

import (
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/ctx"
	"net/http"
	"strings"
)

func SitemapXML(w http.ResponseWriter, r *http.Request) {
	act("sitemap", w, r, func(st *ctx.PageState) (string, error) {
		ret := make([]string, 0)
		ret = append(ret, `<?xml version="1.0" encoding="UTF-8"?>`)
		ret = append(ret, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
		for _, rt := range cutil.ExtractRoutes(ctx.App.Router) {
			if routeMatches(rt) {
				url := rt.Path
				ret = append(ret, `  <url>`)
				ret = append(ret, `     <loc>`+url+`</loc>`)
				ret = append(ret, `     <changefreq>always</changefreq>`)
				ret = append(ret, `  </url>`)
			}
		}
		ret = append(ret, `</urlset>`)
		_, _ = w.Write([]byte(strings.Join(ret, "\n")))
		return "", nil
	})
}

func routeMatches(rt *cutil.RouteDescription) bool {
	pathCheck := func(s ...string) bool {
		for _, x := range s {
			if strings.Contains(rt.Path, x) {
				return false
			}
		}
		return true
	}
	if !pathCheck("admin", "assets", "sitemap", "robots", "{") {
		return false
	}
	if !strings.Contains(rt.Methods, http.MethodGet) {
		return false
	}
	return true
}
