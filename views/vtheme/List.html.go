// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtheme/List.html:2
package vtheme

//line views/vtheme/List.html:2
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/theme"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vtheme/List.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/List.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/List.html:10
type List struct {
	layout.Basic
	Themes theme.Themes
}

//line views/vtheme/List.html:15
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/List.html:15
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/theme/new" title="add new theme">`)
//line views/vtheme/List.html:17
	components.StreamSVGRefIcon(qw422016, `plus`, ps)
//line views/vtheme/List.html:17
	qw422016.N().S(`</a></div>
    <h3>Themes</h3>
    <div class="theme-container mt">
`)
//line views/vtheme/List.html:20
	for _, t := range p.Themes {
//line views/vtheme/List.html:20
		qw422016.N().S(`      <div class="theme-item">
        <a href="/theme/`)
//line views/vtheme/List.html:22
		qw422016.N().U(t.Key)
//line views/vtheme/List.html:22
		qw422016.N().S(`">
          `)
//line views/vtheme/List.html:23
		StreamMockupTheme(qw422016, t, true, 5, ps)
//line views/vtheme/List.html:23
		qw422016.N().S(`
        </a>
      </div>
`)
//line views/vtheme/List.html:26
	}
//line views/vtheme/List.html:26
	qw422016.N().S(`    </div>
  </div>
`)
//line views/vtheme/List.html:29
}

//line views/vtheme/List.html:29
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/List.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/List.html:29
	p.StreamBody(qw422016, as, ps)
//line views/vtheme/List.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/List.html:29
}

//line views/vtheme/List.html:29
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtheme/List.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/List.html:29
	p.WriteBody(qb422016, as, ps)
//line views/vtheme/List.html:29
	qs422016 := string(qb422016.B)
//line views/vtheme/List.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/List.html:29
	return qs422016
//line views/vtheme/List.html:29
}
