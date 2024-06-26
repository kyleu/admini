// Code generated by qtc from "New.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vmodel/New.html:1
package vmodel

//line views/vmodel/New.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/views/components/edit"
	"admini.dev/admini/views/layout"
)

//line views/vmodel/New.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vmodel/New.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vmodel/New.html:10
type New struct {
	layout.Basic
	Req      *cutil.WorkspaceRequest
	Act      *action.Action
	Model    *model.Model
	Defaults []any
}

//line views/vmodel/New.html:18
func (p *New) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodel/New.html:18
	qw422016.N().S(`
  <div class="card">
    <h3>New `)
//line views/vmodel/New.html:20
	qw422016.E().S(p.Model.Name())
//line views/vmodel/New.html:20
	qw422016.N().S(`</h3>
    <form action="`)
//line views/vmodel/New.html:21
	qw422016.E().S(p.Req.RouteAct(p.Act, 0))
//line views/vmodel/New.html:21
	qw422016.N().S(`" method="post" enctype="application/x-www-form-urlencoded">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vmodel/New.html:24
	for idx, f := range p.Model.Fields {
//line views/vmodel/New.html:24
		qw422016.N().S(`
            <tr>
              <th class="shrink"><label><input type="checkbox" value="true" name="`)
//line views/vmodel/New.html:26
		qw422016.E().S(f.Key)
//line views/vmodel/New.html:26
		qw422016.N().S(`--selected" /> `)
//line views/vmodel/New.html:26
		qw422016.E().S(f.Key)
//line views/vmodel/New.html:26
		qw422016.N().S(`</label></th>
              <td>`)
//line views/vmodel/New.html:27
		edit.StreamAnyByType(qw422016, f.Key, f.Key, p.Defaults[idx], f.Type)
//line views/vmodel/New.html:27
		qw422016.N().S(`</td>
            </tr>
          `)
//line views/vmodel/New.html:29
	}
//line views/vmodel/New.html:29
	qw422016.N().S(`
        </tbody>
      </table>
      <div class="mt">
        <button type="submit">Save Changes</button>
        <button type="reset">Reset</button>
      </div>
    </form>
  </div>
`)
//line views/vmodel/New.html:38
}

//line views/vmodel/New.html:38
func (p *New) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodel/New.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodel/New.html:38
	p.StreamBody(qw422016, as, ps)
//line views/vmodel/New.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vmodel/New.html:38
}

//line views/vmodel/New.html:38
func (p *New) Body(as *app.State, ps *cutil.PageState) string {
//line views/vmodel/New.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodel/New.html:38
	p.WriteBody(qb422016, as, ps)
//line views/vmodel/New.html:38
	qs422016 := string(qb422016.B)
//line views/vmodel/New.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodel/New.html:38
	return qs422016
//line views/vmodel/New.html:38
}
