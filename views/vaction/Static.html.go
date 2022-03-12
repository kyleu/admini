// Code generated by qtc from "Static.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Static.html:1
package vaction

//line views/vaction/Static.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/layout"
)

//line views/vaction/Static.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Static.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Static.html:8
type Static struct {
	layout.Basic
	Req *cutil.WorkspaceRequest
	Act *action.Action
}

//line views/vaction/Static.html:14
func (p *Static) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Static.html:14
	qw422016.N().S(`
  <div class="card">
    `)
//line views/vaction/Static.html:16
	StreamActionHeader(qw422016, p.Act, "Static", ps)
//line views/vaction/Static.html:16
	qw422016.N().S(`
  </div>
  <div class="card">
`)
//line views/vaction/Static.html:19
	ct := p.Act.Config.GetStringOpt("content")

//line views/vaction/Static.html:20
	switch p.Act.Config.GetStringOpt("format") {
//line views/vaction/Static.html:21
	case "html":
//line views/vaction/Static.html:21
		qw422016.N().S(`    `)
//line views/vaction/Static.html:22
		qw422016.N().S(ct)
//line views/vaction/Static.html:22
		qw422016.N().S(`
`)
//line views/vaction/Static.html:23
	case "text":
//line views/vaction/Static.html:23
		qw422016.N().S(`    `)
//line views/vaction/Static.html:24
		qw422016.E().S(ct)
//line views/vaction/Static.html:24
		qw422016.N().S(`
`)
//line views/vaction/Static.html:25
	case "code":
//line views/vaction/Static.html:25
		qw422016.N().S(`    <pre>`)
//line views/vaction/Static.html:26
		qw422016.E().S(ct)
//line views/vaction/Static.html:26
		qw422016.N().S(`</pre>
`)
//line views/vaction/Static.html:27
	default:
//line views/vaction/Static.html:27
		qw422016.N().S(`    `)
//line views/vaction/Static.html:28
		qw422016.E().S(ct)
//line views/vaction/Static.html:28
		qw422016.N().S(`
`)
//line views/vaction/Static.html:29
	}
//line views/vaction/Static.html:29
	qw422016.N().S(`  </div>
`)
//line views/vaction/Static.html:31
	StreamResultChildren(qw422016, p.Req, p.Act, 1)
//line views/vaction/Static.html:32
	StreamResultDebug(qw422016, p.Req, p.Act)
//line views/vaction/Static.html:33
}

//line views/vaction/Static.html:33
func (p *Static) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Static.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Static.html:33
	p.StreamBody(qw422016, as, ps)
//line views/vaction/Static.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Static.html:33
}

//line views/vaction/Static.html:33
func (p *Static) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/Static.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Static.html:33
	p.WriteBody(qb422016, as, ps)
//line views/vaction/Static.html:33
	qs422016 := string(qb422016.B)
//line views/vaction/Static.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Static.html:33
	return qs422016
//line views/vaction/Static.html:33
}