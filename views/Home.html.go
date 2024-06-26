// Code generated by qtc from "Home.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/Home.html:1
package views

//line views/Home.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/project"
	"admini.dev/admini/app/source"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
	"admini.dev/admini/views/vproject"
	"admini.dev/admini/views/vsource"
)

//line views/Home.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Home.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Home.html:13
type Home struct {
	layout.Basic
	Sources  source.Sources
	Projects project.Projects
}

//line views/Home.html:19
func (p *Home) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Home.html:19
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/Home.html:21
	components.StreamSVGRef(qw422016, `app`, 20, 20, `icon`, ps)
//line views/Home.html:21
	qw422016.N().S(` `)
//line views/Home.html:21
	qw422016.E().S(util.AppName)
//line views/Home.html:21
	qw422016.N().S(`</h3>
  </div>

  <div class="card">
    <div class="right"><a href="/project/_new" title="add new project">`)
//line views/Home.html:25
	components.StreamSVGRef(qw422016, `plus`, 20, 20, `icon`, ps)
//line views/Home.html:25
	qw422016.N().S(`</a></div>
    <h3><a href="/project">Projects</a></h3>
`)
//line views/Home.html:27
	if len(p.Projects) == 0 {
//line views/Home.html:27
		qw422016.N().S(`    <p>no projects available, why not <a href="/project/_new" title="add new project">add one</a></p>
`)
//line views/Home.html:29
	} else {
//line views/Home.html:30
		vproject.StreamTable(qw422016, p.Projects, as, ps)
//line views/Home.html:31
	}
//line views/Home.html:31
	qw422016.N().S(`  </div>

  <div class="card">
    <div class="right"><a href="/source/_new" title="add new source">`)
//line views/Home.html:35
	components.StreamSVGRef(qw422016, `plus`, 20, 20, `icon`, ps)
//line views/Home.html:35
	qw422016.N().S(`</a></div>
    <h3><a href="/source">Sources</a></h3>
`)
//line views/Home.html:37
	if len(p.Sources) == 0 {
//line views/Home.html:37
		qw422016.N().S(`    <p>No sources configured. Would you like to <a href="/source/_new" title="add new source">add a new source</a> or <a href="/source/_example">load the example database</a>?</p>
`)
//line views/Home.html:39
	} else {
//line views/Home.html:40
		vsource.StreamTable(qw422016, p.Sources, as, ps)
//line views/Home.html:41
	}
//line views/Home.html:41
	qw422016.N().S(`  </div>
`)
//line views/Home.html:43
}

//line views/Home.html:43
func (p *Home) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Home.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Home.html:43
	p.StreamBody(qw422016, as, ps)
//line views/Home.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/Home.html:43
}

//line views/Home.html:43
func (p *Home) Body(as *app.State, ps *cutil.PageState) string {
//line views/Home.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Home.html:43
	p.WriteBody(qb422016, as, ps)
//line views/Home.html:43
	qs422016 := string(qb422016.B)
//line views/Home.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Home.html:43
	return qs422016
//line views/Home.html:43
}
