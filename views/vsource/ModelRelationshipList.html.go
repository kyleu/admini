// Code generated by qtc from "ModelRelationshipList.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsource/ModelRelationshipList.html:1
package vsource

//line views/vsource/ModelRelationshipList.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components/fieldview"
)

//line views/vsource/ModelRelationshipList.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsource/ModelRelationshipList.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsource/ModelRelationshipList.html:9
func StreamModelRelationshipList(qw422016 *qt422016.Writer, m *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vsource/ModelRelationshipList.html:9
	if len(m.Relationships) > 0 {
//line views/vsource/ModelRelationshipList.html:9
		qw422016.N().S(`    <div class="card">
      <h3>`)
//line views/vsource/ModelRelationshipList.html:11
		qw422016.E().S(util.StringPlural(len(m.Relationships), `relationship`))
//line views/vsource/ModelRelationshipList.html:11
		qw422016.N().S(`</h3>
      <table>
        <thead>
          <tr>
            <th>Key</th>
            <th>Source Fields</th>
            <th>Target Package</th>
            <th>Target Model</th>
            <th>Target Fields</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vsource/ModelRelationshipList.html:23
		for _, rel := range m.Relationships {
//line views/vsource/ModelRelationshipList.html:23
			qw422016.N().S(`          <tr>
            <td>`)
//line views/vsource/ModelRelationshipList.html:25
			fieldview.StreamString(qw422016, rel.Key)
//line views/vsource/ModelRelationshipList.html:25
			qw422016.N().S(`</td>
            <td>`)
//line views/vsource/ModelRelationshipList.html:26
			fieldview.StreamArrayString(qw422016, rel.SourceFields)
//line views/vsource/ModelRelationshipList.html:26
			qw422016.N().S(`</td>
            <td>`)
//line views/vsource/ModelRelationshipList.html:27
			fieldview.StreamPackage(qw422016, rel.TargetPkg)
//line views/vsource/ModelRelationshipList.html:27
			qw422016.N().S(`</td>
            <td>`)
//line views/vsource/ModelRelationshipList.html:28
			fieldview.StreamString(qw422016, rel.TargetModel)
//line views/vsource/ModelRelationshipList.html:28
			qw422016.N().S(`</td>
            <td>`)
//line views/vsource/ModelRelationshipList.html:29
			fieldview.StreamArrayString(qw422016, rel.TargetFields)
//line views/vsource/ModelRelationshipList.html:29
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vsource/ModelRelationshipList.html:31
		}
//line views/vsource/ModelRelationshipList.html:31
		qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/vsource/ModelRelationshipList.html:35
	}
//line views/vsource/ModelRelationshipList.html:35
}

//line views/vsource/ModelRelationshipList.html:35
func WriteModelRelationshipList(qq422016 qtio422016.Writer, m *model.Model, as *app.State, ps *cutil.PageState) {
//line views/vsource/ModelRelationshipList.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/ModelRelationshipList.html:35
	StreamModelRelationshipList(qw422016, m, as, ps)
//line views/vsource/ModelRelationshipList.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/ModelRelationshipList.html:35
}

//line views/vsource/ModelRelationshipList.html:35
func ModelRelationshipList(m *model.Model, as *app.State, ps *cutil.PageState) string {
//line views/vsource/ModelRelationshipList.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/ModelRelationshipList.html:35
	WriteModelRelationshipList(qb422016, m, as, ps)
//line views/vsource/ModelRelationshipList.html:35
	qs422016 := string(qb422016.B)
//line views/vsource/ModelRelationshipList.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/ModelRelationshipList.html:35
	return qs422016
//line views/vsource/ModelRelationshipList.html:35
}