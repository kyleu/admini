// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/List.html:1
package vproject

//line views/vproject/List.html:1
import (
	"admini.dev/app"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/project"
	"admini.dev/views/components"
	"admini.dev/views/layout"
)

//line views/vproject/List.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/List.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/List.html:9
type List struct {
	layout.Basic
	Projects project.Projects
}

//line views/vproject/List.html:14
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/List.html:14
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/project/_new" title="add new project">`)
//line views/vproject/List.html:16
	components.StreamSVGRef(qw422016, `plus`, 20, 20, `icon`, ps)
//line views/vproject/List.html:16
	qw422016.N().S(`</a></div>
    <h3>Projects</h3>
    `)
//line views/vproject/List.html:18
	StreamTable(qw422016, p.Projects, as, ps)
//line views/vproject/List.html:18
	qw422016.N().S(`
  </div>
`)
//line views/vproject/List.html:20
}

//line views/vproject/List.html:20
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/List.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/List.html:20
	p.StreamBody(qw422016, as, ps)
//line views/vproject/List.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/List.html:20
}

//line views/vproject/List.html:20
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vproject/List.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/List.html:20
	p.WriteBody(qb422016, as, ps)
//line views/vproject/List.html:20
	qs422016 := string(qb422016.B)
//line views/vproject/List.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/List.html:20
	return qs422016
//line views/vproject/List.html:20
}
