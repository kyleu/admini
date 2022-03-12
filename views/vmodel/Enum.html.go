// Code generated by qtc from "Enum.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vmodel/Enum.html:1
package vmodel

//line views/vmodel/Enum.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
	"admini.dev/admini/views/vaction"
)

//line views/vmodel/Enum.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vmodel/Enum.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vmodel/Enum.html:11
type Enum struct {
	layout.Basic
	Req   *cutil.WorkspaceRequest
	Act   *action.Action
	Model *model.Model
	Refs  model.Relationships
}

//line views/vmodel/Enum.html:19
func (p *Enum) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodel/Enum.html:19
	qw422016.N().S(`
  <div class="card">
    `)
//line views/vmodel/Enum.html:21
	vaction.StreamActionHeader(qw422016, p.Act, "Enum ["+p.Model.Name()+"] ("+p.Model.Type.String()+")", ps)
//line views/vmodel/Enum.html:21
	qw422016.N().S(`
    <a href="#modal-model"><button type="button">Model</button></a>
  </div>

  <div class="card">
    <h3>Values</h3>
    <ul>
`)
//line views/vmodel/Enum.html:28
	for _, f := range p.Model.Fields {
//line views/vmodel/Enum.html:28
		qw422016.N().S(`      <li>`)
//line views/vmodel/Enum.html:29
		qw422016.E().S(f.Key)
//line views/vmodel/Enum.html:29
		qw422016.N().S(` <em>`)
//line views/vmodel/Enum.html:29
		qw422016.E().S(f.Description())
//line views/vmodel/Enum.html:29
		qw422016.N().S(`</em></li>
`)
//line views/vmodel/Enum.html:30
	}
//line views/vmodel/Enum.html:30
	qw422016.N().S(`    </ul>
  </div>

`)
//line views/vmodel/Enum.html:34
	if len(p.Refs) > 0 {
//line views/vmodel/Enum.html:34
		qw422016.N().S(`  <div class="card">
    <h3>References</h3>
    <ul>
`)
//line views/vmodel/Enum.html:38
		for _, ref := range p.Refs {
//line views/vmodel/Enum.html:38
			qw422016.N().S(`      <li>`)
//line views/vmodel/Enum.html:39
			qw422016.E().V(ref)
//line views/vmodel/Enum.html:39
			qw422016.N().S(`</li>
`)
//line views/vmodel/Enum.html:40
		}
//line views/vmodel/Enum.html:40
		qw422016.N().S(`    </ul>
  </div>
`)
//line views/vmodel/Enum.html:43
	}
//line views/vmodel/Enum.html:43
	qw422016.N().S(`
  `)
//line views/vmodel/Enum.html:45
	components.StreamJSONModal(qw422016, "model", "Model JSON", p.Model, 1)
//line views/vmodel/Enum.html:45
	qw422016.N().S(`
`)
//line views/vmodel/Enum.html:46
	vaction.StreamResultChildren(qw422016, p.Req, p.Act, 1)
//line views/vmodel/Enum.html:47
	vaction.StreamResultDebug(qw422016, p.Req, p.Act)
//line views/vmodel/Enum.html:48
}

//line views/vmodel/Enum.html:48
func (p *Enum) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodel/Enum.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodel/Enum.html:48
	p.StreamBody(qw422016, as, ps)
//line views/vmodel/Enum.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/vmodel/Enum.html:48
}

//line views/vmodel/Enum.html:48
func (p *Enum) Body(as *app.State, ps *cutil.PageState) string {
//line views/vmodel/Enum.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodel/Enum.html:48
	p.WriteBody(qb422016, as, ps)
//line views/vmodel/Enum.html:48
	qs422016 := string(qb422016.B)
//line views/vmodel/Enum.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodel/Enum.html:48
	return qs422016
//line views/vmodel/Enum.html:48
}