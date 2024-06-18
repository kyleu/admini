// Code generated by qtc from "Error.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/verror/Error.html:1
package verror

//line views/verror/Error.html:1
import (
	"strings"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/user"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/layout"
)

//line views/verror/Error.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/verror/Error.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/verror/Error.html:11
func streamerrorStack(qw422016 *qt422016.Writer, ed *util.ErrorDetail) {
//line views/verror/Error.html:11
	qw422016.N().S(`    <div class="overflow full-width">
      <table>
        <tbody>
`)
//line views/verror/Error.html:15
	for _, f := range ed.Stack {
//line views/verror/Error.html:15
		qw422016.N().S(`          <tr>
            <td>
`)
//line views/verror/Error.html:18
		if strings.Contains(f.Key, util.AppKey) {
//line views/verror/Error.html:18
			qw422016.N().S(`              <div class="error-key error-owned">`)
//line views/verror/Error.html:19
			qw422016.E().S(f.Key)
//line views/verror/Error.html:19
			qw422016.N().S(`</div>
`)
//line views/verror/Error.html:20
		} else {
//line views/verror/Error.html:20
			qw422016.N().S(`              <div class="error-key">`)
//line views/verror/Error.html:21
			qw422016.E().S(f.Key)
//line views/verror/Error.html:21
			qw422016.N().S(`</div>
`)
//line views/verror/Error.html:22
		}
//line views/verror/Error.html:22
		qw422016.N().S(`              <div class="error-location">`)
//line views/verror/Error.html:23
		qw422016.E().S(f.Loc)
//line views/verror/Error.html:23
		qw422016.N().S(`</div>
            </td>
          </tr>
`)
//line views/verror/Error.html:26
	}
//line views/verror/Error.html:26
	qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/verror/Error.html:30
}

//line views/verror/Error.html:30
func writeerrorStack(qq422016 qtio422016.Writer, ed *util.ErrorDetail) {
//line views/verror/Error.html:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/verror/Error.html:30
	streamerrorStack(qw422016, ed)
//line views/verror/Error.html:30
	qt422016.ReleaseWriter(qw422016)
//line views/verror/Error.html:30
}

//line views/verror/Error.html:30
func errorStack(ed *util.ErrorDetail) string {
//line views/verror/Error.html:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/verror/Error.html:30
	writeerrorStack(qb422016, ed)
//line views/verror/Error.html:30
	qs422016 := string(qb422016.B)
//line views/verror/Error.html:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/verror/Error.html:30
	return qs422016
//line views/verror/Error.html:30
}

//line views/verror/Error.html:32
type Error struct {
	layout.Basic
	Err *util.ErrorDetail
}

//line views/verror/Error.html:37
func (p *Error) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/Error.html:37
	qw422016.N().S(`
  `)
//line views/verror/Error.html:38
	StreamDetail(qw422016, p.Err, as, ps)
//line views/verror/Error.html:38
	qw422016.N().S(`
`)
//line views/verror/Error.html:39
}

//line views/verror/Error.html:39
func (p *Error) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/Error.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/verror/Error.html:39
	p.StreamBody(qw422016, as, ps)
//line views/verror/Error.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/verror/Error.html:39
}

//line views/verror/Error.html:39
func (p *Error) Body(as *app.State, ps *cutil.PageState) string {
//line views/verror/Error.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/verror/Error.html:39
	p.WriteBody(qb422016, as, ps)
//line views/verror/Error.html:39
	qs422016 := string(qb422016.B)
//line views/verror/Error.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/verror/Error.html:39
	return qs422016
//line views/verror/Error.html:39
}

//line views/verror/Error.html:41
func StreamDetail(qw422016 *qt422016.Writer, ed *util.ErrorDetail, as *app.State, ps *cutil.PageState) {
//line views/verror/Error.html:41
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/verror/Error.html:43
	qw422016.E().S(ed.Message)
//line views/verror/Error.html:43
	qw422016.N().S(`</h3>
    <em>Internal Server Error</em>
`)
//line views/verror/Error.html:45
	if user.IsAdmin(ps.Accounts) {
//line views/verror/Error.html:45
		qw422016.N().S(`    `)
//line views/verror/Error.html:46
		streamerrorStack(qw422016, ed)
//line views/verror/Error.html:46
		qw422016.N().S(` `)
//line views/verror/Error.html:46
		cause := ed.Cause

//line views/verror/Error.html:46
		qw422016.N().S(`
`)
//line views/verror/Error.html:47
		for cause != nil {
//line views/verror/Error.html:47
			qw422016.N().S(`    <h3>Caused by</h3>
    <div>`)
//line views/verror/Error.html:49
			qw422016.E().S(cause.Message)
//line views/verror/Error.html:49
			qw422016.N().S(`</div>`)
//line views/verror/Error.html:49
			streamerrorStack(qw422016, cause)
//line views/verror/Error.html:49
			cause = cause.Cause

//line views/verror/Error.html:49
			qw422016.N().S(`
`)
//line views/verror/Error.html:50
		}
//line views/verror/Error.html:51
	}
//line views/verror/Error.html:51
	qw422016.N().S(`  </div>
`)
//line views/verror/Error.html:53
}

//line views/verror/Error.html:53
func WriteDetail(qq422016 qtio422016.Writer, ed *util.ErrorDetail, as *app.State, ps *cutil.PageState) {
//line views/verror/Error.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/verror/Error.html:53
	StreamDetail(qw422016, ed, as, ps)
//line views/verror/Error.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/verror/Error.html:53
}

//line views/verror/Error.html:53
func Detail(ed *util.ErrorDetail, as *app.State, ps *cutil.PageState) string {
//line views/verror/Error.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/verror/Error.html:53
	WriteDetail(qb422016, ed, as, ps)
//line views/verror/Error.html:53
	qs422016 := string(qb422016.B)
//line views/verror/Error.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/verror/Error.html:53
	return qs422016
//line views/verror/Error.html:53
}
