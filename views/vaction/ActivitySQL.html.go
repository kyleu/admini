// Code generated by qtc from "ActivitySQL.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/ActivitySQL.html:1
package vaction

//line views/vaction/ActivitySQL.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/result"
	"admini.dev/admini/views/layout"
	"admini.dev/admini/views/vresult"
)

//line views/vaction/ActivitySQL.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/ActivitySQL.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/ActivitySQL.html:10
type ActivitySQL struct {
	layout.Basic
	Req *cutil.WorkspaceRequest
	Act *action.Action
	SQL string
	Res *result.Result
}

//line views/vaction/ActivitySQL.html:18
func (p *ActivitySQL) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/ActivitySQL.html:18
	qw422016.N().S(`
  <div class="card">
    `)
//line views/vaction/ActivitySQL.html:20
	StreamActionHeader(qw422016, p.Act, "SQL Playground", ps)
//line views/vaction/ActivitySQL.html:20
	qw422016.N().S(`
    <form class="sql" action="" method="post">
      <textarea class="pre mt mb" name="sql" placeholder="SQL" rows="10">`)
//line views/vaction/ActivitySQL.html:22
	qw422016.E().S(p.SQL)
//line views/vaction/ActivitySQL.html:22
	qw422016.N().S(`</textarea>
      <button type="submit">Run</button>
    </form>
  </div>
`)
//line views/vaction/ActivitySQL.html:26
	if p.Res != nil {
//line views/vaction/ActivitySQL.html:26
		qw422016.N().S(`  `)
//line views/vaction/ActivitySQL.html:27
		vresult.StreamSimple(qw422016, p.Res, 2, as, ps)
//line views/vaction/ActivitySQL.html:27
		qw422016.N().S(`
`)
//line views/vaction/ActivitySQL.html:28
	}
//line views/vaction/ActivitySQL.html:28
	qw422016.N().S(`
`)
//line views/vaction/ActivitySQL.html:30
	StreamResultChildren(qw422016, p.Req, p.Act, 2)
//line views/vaction/ActivitySQL.html:31
	StreamResultDebug(qw422016, p.Req, p.Act)
//line views/vaction/ActivitySQL.html:32
}

//line views/vaction/ActivitySQL.html:32
func (p *ActivitySQL) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/ActivitySQL.html:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/ActivitySQL.html:32
	p.StreamBody(qw422016, as, ps)
//line views/vaction/ActivitySQL.html:32
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/ActivitySQL.html:32
}

//line views/vaction/ActivitySQL.html:32
func (p *ActivitySQL) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/ActivitySQL.html:32
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/ActivitySQL.html:32
	p.WriteBody(qb422016, as, ps)
//line views/vaction/ActivitySQL.html:32
	qs422016 := string(qb422016.B)
//line views/vaction/ActivitySQL.html:32
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/ActivitySQL.html:32
	return qs422016
//line views/vaction/ActivitySQL.html:32
}
