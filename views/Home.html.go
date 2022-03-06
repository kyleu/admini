// Code generated by qtc from "Home.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- $PF_IGNORE$ -->

//line views/Home.html:2
package views

//line views/Home.html:2
import (
	"admini.dev/app"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/lib/sandbox"
	"admini.dev/app/project"
	"admini.dev/app/source"
	"admini.dev/app/util"
	"admini.dev/views/components"
	"admini.dev/views/layout"
	"admini.dev/views/vproject"
	"admini.dev/views/vsource"
)

//line views/Home.html:15
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Home.html:15
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Home.html:15
type Home struct {
	layout.Basic
	Sources  source.Sources
	Projects project.Projects
}

//line views/Home.html:21
func (p *Home) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Home.html:21
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/Home.html:23
	components.StreamSVGRef(qw422016, `app`, 20, 20, `icon`, ps)
//line views/Home.html:23
	qw422016.E().S(util.AppName)
//line views/Home.html:23
	qw422016.N().S(`</h3>
  </div>

  <div class="card">
    <div class="right"><a href="/project/_new" title="add new project">`)
//line views/Home.html:27
	components.StreamSVGRef(qw422016, `plus`, 20, 20, `icon`, ps)
//line views/Home.html:27
	qw422016.N().S(`</a></div>
    <h3><a href="/project">Projects</a></h3>
`)
//line views/Home.html:29
	if len(p.Projects) == 0 {
//line views/Home.html:29
		qw422016.N().S(`    <p>no projects available, why not <a href="/project/_new" title="add new project">add one</a></p>
`)
//line views/Home.html:31
	} else {
//line views/Home.html:32
		vproject.StreamTable(qw422016, p.Projects, as, ps)
//line views/Home.html:33
	}
//line views/Home.html:33
	qw422016.N().S(`  </div>

  <div class="card">
    <div class="right"><a href="/source/_new" title="add new source">`)
//line views/Home.html:37
	components.StreamSVGRef(qw422016, `plus`, 20, 20, `icon`, ps)
//line views/Home.html:37
	qw422016.N().S(`</a></div>
    <h3><a href="/source">Sources</a></h3>
`)
//line views/Home.html:39
	if len(p.Sources) == 0 {
//line views/Home.html:39
		qw422016.N().S(`    <p>No sources configured. Would you like to <a href="/source/_new" title="add new source">add a new source</a> or <a href="/source/_example">load the example database</a>?</p>
`)
//line views/Home.html:41
	} else {
//line views/Home.html:42
		vsource.StreamTable(qw422016, p.Sources, as, ps)
//line views/Home.html:43
	}
//line views/Home.html:43
	qw422016.N().S(`  </div>

  <div class="card">
    <h3><a href="/sandbox">Sandboxes</a></h3>
    <ul>
`)
//line views/Home.html:49
	for _, s := range sandbox.AllSandboxes {
//line views/Home.html:49
		qw422016.N().S(`      <li><a href="/sandbox/`)
//line views/Home.html:50
		qw422016.E().S(s.Key)
//line views/Home.html:50
		qw422016.N().S(`">`)
//line views/Home.html:50
		qw422016.E().S(s.Title)
//line views/Home.html:50
		qw422016.N().S(`</a></li>
`)
//line views/Home.html:51
	}
//line views/Home.html:51
	qw422016.N().S(`    </ul>
  </div>
`)
//line views/Home.html:54
}

//line views/Home.html:54
func (p *Home) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Home.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Home.html:54
	p.StreamBody(qw422016, as, ps)
//line views/Home.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/Home.html:54
}

//line views/Home.html:54
func (p *Home) Body(as *app.State, ps *cutil.PageState) string {
//line views/Home.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Home.html:54
	p.WriteBody(qb422016, as, ps)
//line views/Home.html:54
	qs422016 := string(qb422016.B)
//line views/Home.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Home.html:54
	return qs422016
//line views/Home.html:54
}
