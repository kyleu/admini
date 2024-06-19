// Code generated by qtc from "Render.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/Render.html:1
package views

//line views/Render.html:1
import (
	"fmt"
	"slices"
	"strings"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/Render.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Render.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Render.html:14
func StreamRender(qw422016 *qt422016.Writer, page layout.Page, as *app.State, ps *cutil.PageState) {
//line views/Render.html:15
	ctx, span, _ := telemetry.StartSpan(ps.Context, "html:"+strings.TrimPrefix(fmt.Sprintf("%T", page), "*"), ps.Logger)
	ps.Context = ctx
	defer func() {
		span.Complete()
		x := recover()
		if x != nil {
			ps.LogError("error processing template [%T]: %+v", x, x)
			panic(x)
		}
	}()

//line views/Render.html:25
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">
<!-- `)
//line views/Render.html:27
	qw422016.E().S(cutil.PageComment)
//line views/Render.html:27
	qw422016.N().S(` -->
<head>`)
//line views/Render.html:28
	page.StreamHead(qw422016, as, ps)
//line views/Render.html:28
	qw422016.N().S(`</head>
<body`)
//line views/Render.html:29
	if clsDecl := ps.ClassDecl(); clsDecl != `-` {
//line views/Render.html:29
		qw422016.N().S(clsDecl)
//line views/Render.html:29
		if ps.Action != "" {
//line views/Render.html:29
			qw422016.N().S(` `)
//line views/Render.html:29
			qw422016.N().S(`data-action="`)
//line views/Render.html:29
			qw422016.E().S(ps.Action)
//line views/Render.html:29
			qw422016.N().S(`"`)
//line views/Render.html:29
		}
//line views/Render.html:29
		qw422016.N().S(` `)
//line views/Render.html:29
		qw422016.N().S(`data-version="`)
//line views/Render.html:29
		qw422016.E().S(as.BuildInfo.Version)
//line views/Render.html:29
		qw422016.N().S(`"`)
//line views/Render.html:29
	}
//line views/Render.html:29
	qw422016.N().S(`>`)
//line views/Render.html:29
	if len(ps.Flashes) > 0 {
//line views/Render.html:29
		streamrenderFlashes(qw422016, ps.Flashes)
//line views/Render.html:29
	}
//line views/Render.html:29
	page.StreamNav(qw422016, as, ps)
//line views/Render.html:29
	qw422016.N().S(`
<main id="content"`)
//line views/Render.html:30
	if ps.HideMenu {
//line views/Render.html:30
		qw422016.N().S(` class="nomenu"`)
//line views/Render.html:30
	}
//line views/Render.html:30
	qw422016.N().S(`>`)
//line views/Render.html:30
	page.StreamBody(qw422016, as, ps)
//line views/Render.html:30
	qw422016.N().S(`</main>
`)
//line views/Render.html:31
	streamrenderIcons(qw422016, ps.Icons)
//line views/Render.html:31
	qw422016.N().S(`</body>
</html>
`)
//line views/Render.html:33
}

//line views/Render.html:33
func WriteRender(qq422016 qtio422016.Writer, page layout.Page, as *app.State, ps *cutil.PageState) {
//line views/Render.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Render.html:33
	StreamRender(qw422016, page, as, ps)
//line views/Render.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/Render.html:33
}

//line views/Render.html:33
func Render(page layout.Page, as *app.State, ps *cutil.PageState) string {
//line views/Render.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Render.html:33
	WriteRender(qb422016, page, as, ps)
//line views/Render.html:33
	qs422016 := string(qb422016.B)
//line views/Render.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Render.html:33
	return qs422016
//line views/Render.html:33
}

//line views/Render.html:35
func streamrenderFlashes(qw422016 *qt422016.Writer, flashes []string) {
//line views/Render.html:36
	if len(flashes) > 0 {
//line views/Render.html:36
		qw422016.N().S(`<div id="flash-container">`)
//line views/Render.html:38
		for idx, f := range flashes {
//line views/Render.html:39
			level, msg := util.StringSplit(f, ':', true)

//line views/Render.html:39
			qw422016.N().S(`<div class="flash"><input type="radio" style="display:none;" id="hide-flash-`)
//line views/Render.html:41
			qw422016.N().D(idx)
//line views/Render.html:41
			qw422016.N().S(`"><label for="hide-flash-`)
//line views/Render.html:42
			qw422016.N().D(idx)
//line views/Render.html:42
			qw422016.N().S(`"><span>×</span></label><div class="content flash-`)
//line views/Render.html:43
			qw422016.E().S(level)
//line views/Render.html:43
			qw422016.N().S(`">`)
//line views/Render.html:44
			qw422016.E().S(msg)
//line views/Render.html:44
			qw422016.N().S(`</div></div>`)
//line views/Render.html:47
		}
//line views/Render.html:47
		qw422016.N().S(`</div>`)
//line views/Render.html:49
	}
//line views/Render.html:50
}

//line views/Render.html:50
func writerenderFlashes(qq422016 qtio422016.Writer, flashes []string) {
//line views/Render.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Render.html:50
	streamrenderFlashes(qw422016, flashes)
//line views/Render.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/Render.html:50
}

//line views/Render.html:50
func renderFlashes(flashes []string) string {
//line views/Render.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Render.html:50
	writerenderFlashes(qb422016, flashes)
//line views/Render.html:50
	qs422016 := string(qb422016.B)
//line views/Render.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Render.html:50
	return qs422016
//line views/Render.html:50
}

//line views/Render.html:52
func streamrenderIcons(qw422016 *qt422016.Writer, icons []string) {
//line views/Render.html:53
	slices.Sort(icons)

//line views/Render.html:54
	if len(icons) > 0 {
//line views/Render.html:54
		qw422016.N().S(`<div class="icon-list" style="display: none;">`)
//line views/Render.html:55
		qw422016.N().S(`
`)
//line views/Render.html:56
		for _, icon := range icons {
//line views/Render.html:57
			qw422016.N().S(` `)
//line views/Render.html:57
			qw422016.N().S(` `)
//line views/Render.html:57
			components.StreamSVG(qw422016, icon)
//line views/Render.html:57
			qw422016.N().S(`
`)
//line views/Render.html:58
		}
//line views/Render.html:58
		qw422016.N().S(`</div>`)
//line views/Render.html:59
		qw422016.N().S(`
`)
//line views/Render.html:60
	}
//line views/Render.html:61
}

//line views/Render.html:61
func writerenderIcons(qq422016 qtio422016.Writer, icons []string) {
//line views/Render.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Render.html:61
	streamrenderIcons(qw422016, icons)
//line views/Render.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/Render.html:61
}

//line views/Render.html:61
func renderIcons(icons []string) string {
//line views/Render.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Render.html:61
	writerenderIcons(qb422016, icons)
//line views/Render.html:61
	qs422016 := string(qb422016.B)
//line views/Render.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Render.html:61
	return qs422016
//line views/Render.html:61
}
