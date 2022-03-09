// Code generated by qtc from "Sources.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Sources.html:1
package vaction

//line views/vaction/Sources.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/layout"
)

//line views/vaction/Sources.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Sources.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Sources.html:8
type Sources struct {
	layout.Basic
	Req *cutil.WorkspaceRequest
	Act *action.Action
}

//line views/vaction/Sources.html:14
func (p *Sources) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Sources.html:14
	qw422016.N().S(`
  <div class="card">
    `)
//line views/vaction/Sources.html:16
	StreamActionHeader(qw422016, p.Act, "Source", ps)
//line views/vaction/Sources.html:16
	qw422016.N().S(`
    <ul>
`)
//line views/vaction/Sources.html:18
	for _, src := range p.Req.Sources {
//line views/vaction/Sources.html:18
		qw422016.N().S(`      <li><a href="`)
//line views/vaction/Sources.html:19
		qw422016.E().S(p.Req.RouteAct(p.Act, 0, src.Key))
//line views/vaction/Sources.html:19
		qw422016.N().S(`">`)
//line views/vaction/Sources.html:19
		qw422016.E().S(src.Name())
//line views/vaction/Sources.html:19
		qw422016.N().S(`</a></li>
`)
//line views/vaction/Sources.html:20
	}
//line views/vaction/Sources.html:20
	qw422016.N().S(`    </ul>
  </div>
`)
//line views/vaction/Sources.html:23
	StreamResultChildren(qw422016, p.Req, p.Act, 1)
//line views/vaction/Sources.html:24
	StreamResultDebug(qw422016, p.Req, p.Act)
//line views/vaction/Sources.html:25
}

//line views/vaction/Sources.html:25
func (p *Sources) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Sources.html:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Sources.html:25
	p.StreamBody(qw422016, as, ps)
//line views/vaction/Sources.html:25
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Sources.html:25
}

//line views/vaction/Sources.html:25
func (p *Sources) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/Sources.html:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Sources.html:25
	p.WriteBody(qb422016, as, ps)
//line views/vaction/Sources.html:25
	qs422016 := string(qb422016.B)
//line views/vaction/Sources.html:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Sources.html:25
	return qs422016
//line views/vaction/Sources.html:25
}
