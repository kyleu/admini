// Content managed by Project Forge, see [projectforge.md] for details.
package site

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/site/download"
	"admini.dev/admini/app/util"
	"admini.dev/admini/doc"
	"admini.dev/admini/views/layout"
	"admini.dev/admini/views/vsite"
)

func Handle(path []string, rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (string, layout.Page, []string, error) {
	if len(path) == 0 {
		ps.Data = siteData("Welcome to the marketing site!")
		return "", &vsite.Index{}, path, nil
	}

	var page layout.Page
	var err error
	switch path[0] {
	case util.AppKey:
		msg := "\n  " +
			"<meta name=\"go-import\" content=\"admini.dev/admini git %s\">\n  " +
			"<meta name=\"go-source\" content=\"admini.dev/admini %s %s/tree/master{/dir} %s/blob/master{/dir}/{file}#L{line}\">"
		ps.HeaderContent = fmt.Sprintf(msg, util.AppSource, util.AppSource, util.AppSource, util.AppSource)
		return "", &vsite.GoSource{}, path, nil
	case keyDownload:
		dls := download.GetLinks(as.BuildInfo.Version)
		ps.Data = util.ValueMap{"base": "https://admini.dev/admini/releases/download/v" + as.BuildInfo.Version, "links": dls}
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

func siteData(result string, kvs ...string) util.ValueMap {
	ret := util.ValueMap{"app": util.AppName, "url": util.AppURL, "result": result}
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
