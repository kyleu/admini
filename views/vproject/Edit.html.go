// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/Edit.html:1
package vproject

//line views/vproject/Edit.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/project"
	"admini.dev/admini/app/source"
	"admini.dev/admini/views/components/edit"
	"admini.dev/admini/views/layout"
)

//line views/vproject/Edit.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/Edit.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/Edit.html:10
type Edit struct {
	layout.Basic
	Project          *project.Project
	AvailableSources source.Sources
}

//line views/vproject/Edit.html:16
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Edit.html:16
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a class="link-confirm" data-message="Are you sure you want to delete project [`)
//line views/vproject/Edit.html:19
	qw422016.E().S(p.Project.Key)
//line views/vproject/Edit.html:19
	qw422016.N().S(`]?" href="/project/`)
//line views/vproject/Edit.html:19
	qw422016.E().S(p.Project.Key)
//line views/vproject/Edit.html:19
	qw422016.N().S(`/delete">Delete Project</a>
    </div>
    <h3>Edit [`)
//line views/vproject/Edit.html:21
	qw422016.E().S(p.Project.Key)
//line views/vproject/Edit.html:21
	qw422016.N().S(`]</h3>
    <form class="mt" action="/project/`)
//line views/vproject/Edit.html:22
	qw422016.E().S(p.Project.Key)
//line views/vproject/Edit.html:22
	qw422016.N().S(`" method="post" enctype="application/x-www-form-urlencoded">
      <table class="expanded">
        <tbody>
          `)
//line views/vproject/Edit.html:25
	edit.StreamStringTable(qw422016, "title", "", "Title", p.Project.Title, 5)
//line views/vproject/Edit.html:25
	qw422016.N().S(`
          `)
//line views/vproject/Edit.html:26
	edit.StreamIconsTable(qw422016, "icon", "Icon", p.Project.Icon, ps, 5)
//line views/vproject/Edit.html:26
	qw422016.N().S(`
          `)
//line views/vproject/Edit.html:27
	edit.StreamStringTable(qw422016, "description", "", "Description", p.Project.Description, 5)
//line views/vproject/Edit.html:27
	qw422016.N().S(`
          `)
//line views/vproject/Edit.html:28
	edit.StreamCheckboxTable(qw422016, "sources", "Sources", p.Project.Sources, p.AvailableSources.Keys(), p.AvailableSources.Names(), 5)
//line views/vproject/Edit.html:28
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
//line views/vproject/Edit.html:37
}

//line views/vproject/Edit.html:37
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Edit.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/Edit.html:37
	p.StreamBody(qw422016, as, ps)
//line views/vproject/Edit.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/Edit.html:37
}

//line views/vproject/Edit.html:37
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vproject/Edit.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/Edit.html:37
	p.WriteBody(qb422016, as, ps)
//line views/vproject/Edit.html:37
	qs422016 := string(qb422016.B)
//line views/vproject/Edit.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/Edit.html:37
	return qs422016
//line views/vproject/Edit.html:37
}
