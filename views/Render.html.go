// Code generated by qtc from "Render.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/Render.html:2
package views

//line views/Render.html:2
import (
	"fmt"
	"sort"
	"strings"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/Render.html:15
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Render.html:15
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Render.html:15
func StreamRender(qw422016 *qt422016.Writer, page layout.Page, as *app.State, ps *cutil.PageState) {
//line views/Render.html:16
	ctx, span, _ := telemetry.StartSpan(ps.Context, "html:"+strings.TrimPrefix(fmt.Sprintf("%T", page), "*"), ps.Logger)
	ps.Context = ctx
	defer func() {
		span.Complete()
		x := recover()
		if x != nil {
			ps.Logger.Errorf("error processing template [%T]: %+v", x, x)
			panic(x)
		}
	}()

//line views/Render.html:26
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">
<!-- `)
//line views/Render.html:28
	qw422016.E().S(cutil.PageComment)
//line views/Render.html:28
	qw422016.N().S(` -->
<head>`)
//line views/Render.html:29
	page.StreamHead(qw422016, as, ps)
//line views/Render.html:29
	qw422016.N().S(`</head>
<body`)
//line views/Render.html:30
	if ps.Profile.Mode != `` {
//line views/Render.html:30
		qw422016.N().S(` class="`)
//line views/Render.html:30
		qw422016.E().S(ps.Profile.ModeClass())
//line views/Render.html:30
		qw422016.N().S(`"`)
//line views/Render.html:30
	}
//line views/Render.html:30
	qw422016.N().S(`>
`)
//line views/Render.html:31
	if len(ps.Flashes) > 0 {
//line views/Render.html:31
		streamrenderFlashes(qw422016, ps.Flashes)
//line views/Render.html:31
	}
//line views/Render.html:32
	page.StreamNav(qw422016, as, ps)
//line views/Render.html:32
	qw422016.N().S(`
<main id="content"`)
//line views/Render.html:33
	if ps.HideMenu {
//line views/Render.html:33
		qw422016.N().S(` class="nomenu"`)
//line views/Render.html:33
	}
//line views/Render.html:33
	qw422016.N().S(`>`)
//line views/Render.html:33
	page.StreamBody(qw422016, as, ps)
//line views/Render.html:33
	qw422016.N().S(`</main>
`)
//line views/Render.html:34
	sort.Strings(ps.Icons)

//line views/Render.html:35
	if len(ps.Icons) > 0 {
//line views/Render.html:35
		qw422016.N().S(`<div class="icon-list" style="display: none;">`)
//line views/Render.html:35
		for _, icon := range ps.Icons {
//line views/Render.html:35
			qw422016.N().S(`
  `)
//line views/Render.html:36
			components.StreamSVG(qw422016, icon)
//line views/Render.html:36
		}
//line views/Render.html:36
		qw422016.N().S(`
</div>
`)
//line views/Render.html:38
	}
//line views/Render.html:38
	qw422016.N().S(`</body>
</html>
`)
//line views/Render.html:40
}

//line views/Render.html:40
func WriteRender(qq422016 qtio422016.Writer, page layout.Page, as *app.State, ps *cutil.PageState) {
//line views/Render.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Render.html:40
	StreamRender(qw422016, page, as, ps)
//line views/Render.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/Render.html:40
}

//line views/Render.html:40
func Render(page layout.Page, as *app.State, ps *cutil.PageState) string {
//line views/Render.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Render.html:40
	WriteRender(qb422016, page, as, ps)
//line views/Render.html:40
	qs422016 := string(qb422016.B)
//line views/Render.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Render.html:40
	return qs422016
//line views/Render.html:40
}

//line views/Render.html:42
func streamrenderFlashes(qw422016 *qt422016.Writer, flashes []string) {
//line views/Render.html:43
	if len(flashes) > 0 {
//line views/Render.html:43
		qw422016.N().S(`<div id="flash-container">`)
//line views/Render.html:45
		for idx, f := range flashes {
//line views/Render.html:46
			level, msg := util.StringSplit(f, ':', true)

//line views/Render.html:46
			qw422016.N().S(`<div class="flash"><input type="radio" style="display:none;" id="hide-flash-`)
//line views/Render.html:48
			qw422016.N().D(idx)
//line views/Render.html:48
			qw422016.N().S(`"><label for="hide-flash-`)
//line views/Render.html:49
			qw422016.N().D(idx)
//line views/Render.html:49
			qw422016.N().S(`"><span>×</span></label><div class="content flash-`)
//line views/Render.html:50
			qw422016.E().S(level)
//line views/Render.html:50
			qw422016.N().S(`">`)
//line views/Render.html:51
			qw422016.E().S(msg)
//line views/Render.html:51
			qw422016.N().S(`</div></div>`)
//line views/Render.html:54
		}
//line views/Render.html:54
		qw422016.N().S(`</div>`)
//line views/Render.html:56
	}
//line views/Render.html:57
}

//line views/Render.html:57
func writerenderFlashes(qq422016 qtio422016.Writer, flashes []string) {
//line views/Render.html:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Render.html:57
	streamrenderFlashes(qw422016, flashes)
//line views/Render.html:57
	qt422016.ReleaseWriter(qw422016)
//line views/Render.html:57
}

//line views/Render.html:57
func renderFlashes(flashes []string) string {
//line views/Render.html:57
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Render.html:57
	writerenderFlashes(qb422016, flashes)
//line views/Render.html:57
	qs422016 := string(qb422016.B)
//line views/Render.html:57
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Render.html:57
	return qs422016
//line views/Render.html:57
}
