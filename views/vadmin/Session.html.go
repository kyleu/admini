// Code generated by qtc from "Session.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/Session.html:1
package vadmin

//line views/vadmin/Session.html:1
import (
	"fmt"

	"github.com/samber/lo"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vadmin/Session.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Session.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Session.html:13
type Session struct{ layout.Basic }

//line views/vadmin/Session.html:15
func (p *Session) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Session.html:15
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/Session.html:17
	components.StreamSVGIcon(qw422016, `desktop`, ps)
//line views/vadmin/Session.html:17
	qw422016.N().S(` Session</h3>
    <em>`)
//line views/vadmin/Session.html:18
	qw422016.N().D(len(ps.Session))
//line views/vadmin/Session.html:18
	qw422016.N().S(` values</em>
  </div>
`)
//line views/vadmin/Session.html:20
	if len(ps.Session) > 0 {
//line views/vadmin/Session.html:20
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vadmin/Session.html:22
		components.StreamSVGIcon(qw422016, `list`, ps)
//line views/vadmin/Session.html:22
		qw422016.N().S(` Values</h3>
    <div class="overflow full-width">
      <table class="mt expanded">
        <tbody>
`)
//line views/vadmin/Session.html:26
		for _, k := range util.ArraySorted(lo.Keys(ps.Session)) {
//line views/vadmin/Session.html:27
			v := ps.Session[k]

//line views/vadmin/Session.html:27
			qw422016.N().S(`          <tr>
            <th class="shrink">`)
//line views/vadmin/Session.html:29
			qw422016.E().S(k)
//line views/vadmin/Session.html:29
			qw422016.N().S(`</th>
            <td>`)
//line views/vadmin/Session.html:30
			qw422016.E().S(fmt.Sprint(v))
//line views/vadmin/Session.html:30
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vadmin/Session.html:32
		}
//line views/vadmin/Session.html:32
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Session.html:37
	} else {
//line views/vadmin/Session.html:37
		qw422016.N().S(`  <div class="card">
    <em>Empty session</em>
  </div>
`)
//line views/vadmin/Session.html:41
	}
//line views/vadmin/Session.html:41
	qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vadmin/Session.html:43
	components.StreamSVGIcon(qw422016, `profile`, ps)
//line views/vadmin/Session.html:43
	qw422016.N().S(` Profile</h3>
    <div class="mt">`)
//line views/vadmin/Session.html:44
	components.StreamJSON(qw422016, ps.Profile)
//line views/vadmin/Session.html:44
	qw422016.N().S(`</div>
  </div>
`)
//line views/vadmin/Session.html:46
	if len(ps.Accounts) > 0 {
//line views/vadmin/Session.html:46
		qw422016.N().S(`  <div class="card">
    <h3>Accounts</h3>
    <div class="overflow full-width">
      <table class="mt">
        <thead>
          <tr>
            <th>Provider</th>
            <th>Email</th>
            <th>Token</th>
            <th>Picture</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vadmin/Session.html:60
		for _, acct := range ps.Accounts {
//line views/vadmin/Session.html:60
			qw422016.N().S(`        <tr>
          <td>`)
//line views/vadmin/Session.html:62
			qw422016.E().S(acct.Provider)
//line views/vadmin/Session.html:62
			qw422016.N().S(`</td>
          <td>`)
//line views/vadmin/Session.html:63
			qw422016.E().S(acct.Email)
//line views/vadmin/Session.html:63
			qw422016.N().S(`</td>
          <td><div class="break-word">`)
//line views/vadmin/Session.html:64
			qw422016.E().S(acct.Token)
//line views/vadmin/Session.html:64
			qw422016.N().S(`</div></td>
          <td>`)
//line views/vadmin/Session.html:65
			qw422016.E().S(acct.Picture)
//line views/vadmin/Session.html:65
			qw422016.N().S(`</td>
        </tr>
`)
//line views/vadmin/Session.html:67
		}
//line views/vadmin/Session.html:67
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Session.html:72
	}
//line views/vadmin/Session.html:73
}

//line views/vadmin/Session.html:73
func (p *Session) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Session.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Session.html:73
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Session.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Session.html:73
}

//line views/vadmin/Session.html:73
func (p *Session) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Session.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Session.html:73
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Session.html:73
	qs422016 := string(qb422016.B)
//line views/vadmin/Session.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Session.html:73
	return qs422016
//line views/vadmin/Session.html:73
}
