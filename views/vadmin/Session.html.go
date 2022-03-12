// Code generated by qtc from "Session.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Session.html:2
package vadmin

//line views/vadmin/Session.html:2
import (
	"fmt"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/layout"
)

//line views/vadmin/Session.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Session.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Session.html:10
type Session struct{ layout.Basic }

//line views/vadmin/Session.html:12
func (p *Session) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Session.html:12
	qw422016.N().S(`
  <div class="card">
    <h3>Session</h3>
    <em>`)
//line views/vadmin/Session.html:15
	qw422016.N().D(len(ps.Session))
//line views/vadmin/Session.html:15
	qw422016.N().S(` values</em>
  </div>
`)
//line views/vadmin/Session.html:17
	if len(ps.Session) > 0 {
//line views/vadmin/Session.html:17
		qw422016.N().S(`  <div class="card">
    <h3>Values</h3>
    <ul class="mt">
`)
//line views/vadmin/Session.html:21
		for k, v := range ps.Session {
//line views/vadmin/Session.html:21
			qw422016.N().S(`      <li>`)
//line views/vadmin/Session.html:22
			qw422016.E().S(k)
//line views/vadmin/Session.html:22
			qw422016.N().S(`: `)
//line views/vadmin/Session.html:22
			qw422016.E().S(fmt.Sprint(v))
//line views/vadmin/Session.html:22
			qw422016.N().S(`</li>
`)
//line views/vadmin/Session.html:23
		}
//line views/vadmin/Session.html:23
		qw422016.N().S(`    </ul>
  </div>
`)
//line views/vadmin/Session.html:26
	} else {
//line views/vadmin/Session.html:26
		qw422016.N().S(`  <div class="card">
    <em>Empty session</em>
  </div>
`)
//line views/vadmin/Session.html:30
	}
//line views/vadmin/Session.html:31
}

//line views/vadmin/Session.html:31
func (p *Session) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Session.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Session.html:31
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Session.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Session.html:31
}

//line views/vadmin/Session.html:31
func (p *Session) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Session.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Session.html:31
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Session.html:31
	qs422016 := string(qb422016.B)
//line views/vadmin/Session.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Session.html:31
	return qs422016
//line views/vadmin/Session.html:31
}