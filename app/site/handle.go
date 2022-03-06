// Content managed by Project Forge, see [projectforge.md] for details.
package site

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"admini.dev/app"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/site/download"
	"admini.dev/app/util"
	"admini.dev/doc"
	"admini.dev/views/layout"
	"admini.dev/views/vsite"
)

func Handle(path []string, rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (string, layout.Page, []string, error) {
	if len(path) == 0 {
		msg := "\n  " +
			"<meta name=\"go-import\" content=\"admini.dev git %s\">\n  " +
			"<meta name=\"go-source\" content=\"admini.dev %s %s/tree/master{/dir} %s/blob/master{/dir}/{file}#L{line}\">"
		ps.HeaderContent = fmt.Sprintf(msg, util.AppSource, util.AppSource, util.AppSource, util.AppSource)
		ps.Data = siteData("Welcome to the marketing site!")
		return "", &vsite.Index{}, path, nil
	}

	var page layout.Page
	var err error
	switch path[0] {
	case keyDownload:
		dls := download.GetLinks(as.BuildInfo.Version)
		ps.Data = map[string]interface{}{"base": "https://admini.dev/releases/download/v" + as.BuildInfo.Version, "links": dls}
		page = &vsite.Download{Links: dls}
	case keyInstall:
		page, err = mdTemplate("Installation", "This static page contains installation instructions", "installation.md", ps)
	case keyContrib:
		page, err = mdTemplate("Contributing", "This static page describes how to build "+util.AppName, "contributing.md", ps)
	case keyTech:
		page, err = mdTemplate("Technology", "This static page describes the technology used in "+util.AppName, "technology.md", ps)
	default:
		page, err = mdTemplate("Documentation", "Documentation for "+util.AppName, path[0]+".md", ps)
	}
	return "", page, path, err
}

func siteData(result string, kvs ...string) map[string]interface{} {
	ret := map[string]interface{}{"app": util.AppName, "url": util.AppURL, "result": result}
	for i := 0; i < len(kvs); i += 2 {
		ret[kvs[i]] = kvs[i+1]
	}
	return ret
}

func mdTemplate(title string, description string, path string, ps *cutil.PageState) (layout.Page, error) {
	ps.Data = siteData(title, "description", description)
	ps.Title = title
	html, err := doc.HTML(path)
	if err != nil {
		return nil, err
	}
	page := &vsite.MarkdownPage{Title: title, HTML: html}
	return page, nil
}
