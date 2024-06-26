// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vproject/Detail.html:1
package vproject

//line views/vproject/Detail.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/project"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vproject/Detail.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vproject/Detail.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vproject/Detail.html:10
type Detail struct {
	layout.Basic
	View *project.View
}

//line views/vproject/Detail.html:15
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Detail.html:15
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/project/`)
//line views/vproject/Detail.html:17
	qw422016.E().S(p.View.Project.Key)
//line views/vproject/Detail.html:17
	qw422016.N().S(`/edit"><button type="button">Edit</button></a></div>
    <h3>`)
//line views/vproject/Detail.html:18
	components.StreamSVGIcon(qw422016, p.View.Project.IconWithFallback(), ps)
//line views/vproject/Detail.html:18
	qw422016.N().S(` `)
//line views/vproject/Detail.html:18
	qw422016.E().S(p.View.Project.Name())
//line views/vproject/Detail.html:18
	qw422016.N().S(`</h3>
    `)
//line views/vproject/Detail.html:19
	if p.View.Project.Description != "" {
//line views/vproject/Detail.html:19
		qw422016.N().S(`<em>`)
//line views/vproject/Detail.html:19
		qw422016.E().S(p.View.Project.Description)
//line views/vproject/Detail.html:19
		qw422016.N().S(`</em>`)
//line views/vproject/Detail.html:19
	}
//line views/vproject/Detail.html:19
	qw422016.N().S(`
    <p>
      <a href="/x/`)
//line views/vproject/Detail.html:21
	qw422016.E().S(p.View.Project.Key)
//line views/vproject/Detail.html:21
	qw422016.N().S(`"><button type="button">Workspace</button></a>
      <a href="/project/`)
//line views/vproject/Detail.html:22
	qw422016.E().S(p.View.Project.Key)
//line views/vproject/Detail.html:22
	qw422016.N().S(`/test"><button type="button">Test</button></a>
    </p>
  </div>

  <div class="card">
`)
//line views/vproject/Detail.html:27
	if len(p.View.Sources) == 0 {
//line views/vproject/Detail.html:27
		qw422016.N().S(`    <h3>No sources in project</h3>
`)
//line views/vproject/Detail.html:29
	} else {
//line views/vproject/Detail.html:29
		qw422016.N().S(`    <h3>`)
//line views/vproject/Detail.html:30
		qw422016.E().S(util.StringPlural(len(p.View.Sources), `source`))
//line views/vproject/Detail.html:30
		qw422016.N().S(` in project</h3>
    <ul>
`)
//line views/vproject/Detail.html:32
		for _, s := range p.View.Sources {
//line views/vproject/Detail.html:32
			qw422016.N().S(`      <li><a href="/source/`)
//line views/vproject/Detail.html:33
			qw422016.E().S(s.Key)
//line views/vproject/Detail.html:33
			qw422016.N().S(`">`)
//line views/vproject/Detail.html:33
			components.StreamSVGRef(qw422016, s.IconWithFallback(), 16, 16, "icon", ps)
//line views/vproject/Detail.html:33
			qw422016.N().S(` `)
//line views/vproject/Detail.html:33
			qw422016.E().S(s.Name())
//line views/vproject/Detail.html:33
			qw422016.N().S(`</a></li>
`)
//line views/vproject/Detail.html:34
		}
//line views/vproject/Detail.html:34
		qw422016.N().S(`    </ul>
`)
//line views/vproject/Detail.html:36
	}
//line views/vproject/Detail.html:36
	qw422016.N().S(`  </div>

  <div class="drag-container readonly card">
    <div class="action-workbench">
      <div class="l">
        <div class="drag-edit right">
          <button type="button" onclick="admini.sortableEdit(this);">Edit</button>
        </div>
        <div class="drag-actions right no-changes">
          <div class="message"><em>no changes</em></div>
          <div class="form">
            <form action="/project/`)
//line views/vproject/Detail.html:48
	qw422016.E().S(p.View.Project.Key)
//line views/vproject/Detail.html:48
	qw422016.N().S(`/actions" method="post">
              <input type="hidden" class="drag-state-original" value=""/>
              <input type="hidden" class="drag-state" value="" name="ordering"/>
              <button type="submit">Save</button>
            </form>
          </div>
        </div>
        <h3 class="no-padding"><span class="drag-tracked-size" data-sing="action" data-plur="actions">`)
//line views/vproject/Detail.html:55
	qw422016.E().S(util.StringPlural(p.View.Project.Actions.Size(), "action"))
//line views/vproject/Detail.html:55
	qw422016.N().S(`</span></h3>
        <div class="clear"></div>
        <div class="container tracked">
`)
//line views/vproject/Detail.html:58
	for _, act := range p.View.Project.Actions {
//line views/vproject/Detail.html:58
		qw422016.N().S(`          `)
//line views/vproject/Detail.html:59
		StreamActionList(qw422016, p.View.Project.Key, act, as, ps, 5)
//line views/vproject/Detail.html:59
		qw422016.N().S(`
`)
//line views/vproject/Detail.html:60
	}
//line views/vproject/Detail.html:60
	qw422016.N().S(`        </div>
      </div>
      <div class="r">
        <h3 class="no-padding">Available</h3>
        <div class="container">
          `)
//line views/vproject/Detail.html:66
	StreamActionAvailable(qw422016, p.View, as, ps, 5)
//line views/vproject/Detail.html:66
	qw422016.N().S(`
        </div>
      </div>
    </div>
  </div>
  <script type="text/javascript" src="/assets/sortable.js"></script>
`)
//line views/vproject/Detail.html:72
}

//line views/vproject/Detail.html:72
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vproject/Detail.html:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vproject/Detail.html:72
	p.StreamBody(qw422016, as, ps)
//line views/vproject/Detail.html:72
	qt422016.ReleaseWriter(qw422016)
//line views/vproject/Detail.html:72
}

//line views/vproject/Detail.html:72
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vproject/Detail.html:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vproject/Detail.html:72
	p.WriteBody(qb422016, as, ps)
//line views/vproject/Detail.html:72
	qs422016 := string(qb422016.B)
//line views/vproject/Detail.html:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vproject/Detail.html:72
	return qs422016
//line views/vproject/Detail.html:72
}
