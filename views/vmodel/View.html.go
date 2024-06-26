// Code generated by qtc from "View.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vmodel/View.html:1
package vmodel

//line views/vmodel/View.html:1
import (
	"strings"

	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/qualify"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/components/view"
	"admini.dev/admini/views/layout"
)

//line views/vmodel/View.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vmodel/View.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vmodel/View.html:14
type View struct {
	layout.Basic
	Req    *cutil.WorkspaceRequest
	Act    *action.Action
	Model  *model.Model
	Result []any
}

//line views/vmodel/View.html:22
func (p *View) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodel/View.html:22
	qw422016.N().S(`
  <div class="card">
`)
//line views/vmodel/View.html:25
	rowPK, err := model.GetStrings(p.Model.Fields, p.Model.GetPK(ps.Logger), p.Result)
	if err != nil {
		panic(err)
	}

//line views/vmodel/View.html:29
	qw422016.N().S(`    <div class="right"><a href="`)
//line views/vmodel/View.html:30
	qw422016.E().S(p.Req.RouteAct(p.Act, 1+len(rowPK), append([]string{`x`}, rowPK...)...))
//line views/vmodel/View.html:30
	qw422016.N().S(`"><button>Edit</button></a></div>
    <h3>`)
//line views/vmodel/View.html:31
	components.StreamSVGIcon(qw422016, p.Act.IconWithFallback(), ps)
//line views/vmodel/View.html:31
	qw422016.N().S(` `)
//line views/vmodel/View.html:31
	qw422016.E().S(strings.Join(rowPK, "/"))
//line views/vmodel/View.html:31
	qw422016.N().S(`</h3>
    <a href="`)
//line views/vmodel/View.html:32
	qw422016.E().S(p.Req.RouteAct(p.Act, len(rowPK)+1))
//line views/vmodel/View.html:32
	qw422016.N().S(`"><em>`)
//line views/vmodel/View.html:32
	qw422016.E().S(p.Model.Name())
//line views/vmodel/View.html:32
	qw422016.N().S(`</em></a>
  </div>
  <div class="card">
    <table>
      <tbody>
`)
//line views/vmodel/View.html:37
	for idx, f := range p.Model.Fields {
//line views/vmodel/View.html:37
		qw422016.N().S(`        <tr>
          <th class="shrink">`)
//line views/vmodel/View.html:39
		qw422016.E().S(f.Name())
//line views/vmodel/View.html:39
		qw422016.N().S(`</th>
          <td>
            `)
//line views/vmodel/View.html:41
		view.StreamAnyByType(qw422016, p.Result[idx], f.Type, ps)
//line views/vmodel/View.html:41
		qw422016.N().S(`
`)
//line views/vmodel/View.html:42
		for _, rel := range p.Model.ApplicableRelations(f.Key) {
//line views/vmodel/View.html:44
			quals, err := qualify.Handle(rel, p.Act, p.Req, p.Model, p.Result)
			if err != nil {
				panic(err)
			}

//line views/vmodel/View.html:49
			for _, q := range quals {
//line views/vmodel/View.html:49
				qw422016.N().S(`            <a href="`)
//line views/vmodel/View.html:50
				qw422016.E().S(p.Req.Route(q.String()))
//line views/vmodel/View.html:50
				qw422016.N().S(`" title="`)
//line views/vmodel/View.html:50
				qw422016.E().S(q.Help())
//line views/vmodel/View.html:50
				qw422016.N().S(`">`)
//line views/vmodel/View.html:50
				components.StreamSVGRef(qw422016, q.Icon, 16, 16, ``, ps)
//line views/vmodel/View.html:50
				qw422016.N().S(`</a>
`)
//line views/vmodel/View.html:51
			}
//line views/vmodel/View.html:52
		}
//line views/vmodel/View.html:52
		qw422016.N().S(`          </td>
        </tr>
`)
//line views/vmodel/View.html:55
	}
//line views/vmodel/View.html:55
	qw422016.N().S(`      </tbody>
    </table>
  </div>
  `)
//line views/vmodel/View.html:59
	streamviewReferences(qw422016, p.Model.References, ps)
//line views/vmodel/View.html:59
	qw422016.N().S(`
`)
//line views/vmodel/View.html:60
}

//line views/vmodel/View.html:60
func (p *View) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vmodel/View.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodel/View.html:60
	p.StreamBody(qw422016, as, ps)
//line views/vmodel/View.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/vmodel/View.html:60
}

//line views/vmodel/View.html:60
func (p *View) Body(as *app.State, ps *cutil.PageState) string {
//line views/vmodel/View.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodel/View.html:60
	p.WriteBody(qb422016, as, ps)
//line views/vmodel/View.html:60
	qs422016 := string(qb422016.B)
//line views/vmodel/View.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodel/View.html:60
	return qs422016
//line views/vmodel/View.html:60
}

//line views/vmodel/View.html:62
func streamviewReferences(qw422016 *qt422016.Writer, refs model.References, ps *cutil.PageState) {
//line views/vmodel/View.html:62
	qw422016.N().S(`
`)
//line views/vmodel/View.html:63
	if len(refs) > 0 {
//line views/vmodel/View.html:63
		qw422016.N().S(`  <div class="card">
    <h3>References</h3>
    <ul class="accordion">
`)
//line views/vmodel/View.html:67
		for _, r := range refs {
//line views/vmodel/View.html:67
			qw422016.N().S(`      <li>
        <input id="`)
//line views/vmodel/View.html:69
			qw422016.E().S(r.Key)
//line views/vmodel/View.html:69
			qw422016.N().S(`-count" type="checkbox" hidden />
        <label for="`)
//line views/vmodel/View.html:70
			qw422016.E().S(r.Key)
//line views/vmodel/View.html:70
			qw422016.N().S(`-count">`)
//line views/vmodel/View.html:70
			components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vmodel/View.html:70
			qw422016.N().S(` `)
//line views/vmodel/View.html:70
			qw422016.E().S(r.String())
//line views/vmodel/View.html:70
			qw422016.N().S(`</label>
        <div class="bd"><div><div>
          <noscript>
            <style>.jsrequired { display: none } </style>
            <a href="">View `)
//line views/vmodel/View.html:74
			qw422016.E().S(r.String())
//line views/vmodel/View.html:74
			qw422016.N().S(`</a> (or enable JavaScript)
          </noscript>
          <div class="jsrequired">
            <div class="relationship" data-key="`)
//line views/vmodel/View.html:77
			qw422016.E().S(r.Key)
//line views/vmodel/View.html:77
			qw422016.N().S(`" data-pkg="`)
//line views/vmodel/View.html:77
			qw422016.E().S(r.SourcePkg.String())
//line views/vmodel/View.html:77
			qw422016.N().S(`" data-model="`)
//line views/vmodel/View.html:77
			qw422016.E().S(r.SourceModel)
//line views/vmodel/View.html:77
			qw422016.N().S(`" data-fields="`)
//line views/vmodel/View.html:77
			qw422016.E().S(strings.Join(r.SourceFields, `//`))
//line views/vmodel/View.html:77
			qw422016.N().S(`">Loading...</div>
            <em>(nah, not really; soon though)</em>
          </div>
        </div></div></div>
      </li>
`)
//line views/vmodel/View.html:82
		}
//line views/vmodel/View.html:82
		qw422016.N().S(`    </ul>
  </div>
`)
//line views/vmodel/View.html:85
	}
//line views/vmodel/View.html:86
}

//line views/vmodel/View.html:86
func writeviewReferences(qq422016 qtio422016.Writer, refs model.References, ps *cutil.PageState) {
//line views/vmodel/View.html:86
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vmodel/View.html:86
	streamviewReferences(qw422016, refs, ps)
//line views/vmodel/View.html:86
	qt422016.ReleaseWriter(qw422016)
//line views/vmodel/View.html:86
}

//line views/vmodel/View.html:86
func viewReferences(refs model.References, ps *cutil.PageState) string {
//line views/vmodel/View.html:86
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vmodel/View.html:86
	writeviewReferences(qb422016, refs, ps)
//line views/vmodel/View.html:86
	qs422016 := string(qb422016.B)
//line views/vmodel/View.html:86
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vmodel/View.html:86
	return qs422016
//line views/vmodel/View.html:86
}
