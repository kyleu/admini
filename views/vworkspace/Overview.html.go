// Code generated by qtc from "Overview.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/Overview.html:1
package vworkspace

//line views/vworkspace/Overview.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
	"admini.dev/admini/views/vutil"
)

//line views/vworkspace/Overview.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/Overview.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/Overview.html:10
type WorkspaceOverview struct {
	layout.Basic
	Req *cutil.WorkspaceRequest
}

//line views/vworkspace/Overview.html:15
func (p *WorkspaceOverview) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/Overview.html:15
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vworkspace/Overview.html:17
	components.StreamSVGRef(qw422016, p.Req.Project.IconWithFallback(), 20, 20, "icon", ps)
//line views/vworkspace/Overview.html:17
	qw422016.E().S(p.Req.Project.Name())
//line views/vworkspace/Overview.html:17
	qw422016.N().S(`</h3>
    `)
//line views/vworkspace/Overview.html:18
	streamactionList(qw422016, p.Req, p.Req.Project.Actions, ps, 2)
//line views/vworkspace/Overview.html:18
	qw422016.N().S(`
  </div>
`)
//line views/vworkspace/Overview.html:20
}

//line views/vworkspace/Overview.html:20
func (p *WorkspaceOverview) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/Overview.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/Overview.html:20
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/Overview.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/Overview.html:20
}

//line views/vworkspace/Overview.html:20
func (p *WorkspaceOverview) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/Overview.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/Overview.html:20
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/Overview.html:20
	qs422016 := string(qb422016.B)
//line views/vworkspace/Overview.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/Overview.html:20
	return qs422016
//line views/vworkspace/Overview.html:20
}

//line views/vworkspace/Overview.html:22
func streamactionList(qw422016 *qt422016.Writer, req *cutil.WorkspaceRequest, acts action.Actions, ps *cutil.PageState, indent int) {
//line views/vworkspace/Overview.html:22
	qw422016.N().S(`<ul>`)
//line views/vworkspace/Overview.html:24
	for _, act := range acts {
//line views/vworkspace/Overview.html:25
		if act.TypeKey != action.TypeSeparator.Key {
//line views/vworkspace/Overview.html:26
			if len(act.Children) == 0 {
//line views/vworkspace/Overview.html:27
				vutil.StreamIndent(qw422016, true, indent+1)
//line views/vworkspace/Overview.html:27
				qw422016.N().S(`<li><a href="`)
//line views/vworkspace/Overview.html:28
				qw422016.E().S(req.RouteAct(act, 0))
//line views/vworkspace/Overview.html:28
				qw422016.N().S(`">`)
//line views/vworkspace/Overview.html:29
				components.StreamSVGRef(qw422016, act.IconWithFallback(), 16, 16, "icon", ps)
//line views/vworkspace/Overview.html:30
				qw422016.E().S(act.Name())
//line views/vworkspace/Overview.html:30
				qw422016.N().S(`</a></li>`)
//line views/vworkspace/Overview.html:32
			} else {
//line views/vworkspace/Overview.html:33
				vutil.StreamIndent(qw422016, true, indent+1)
//line views/vworkspace/Overview.html:33
				qw422016.N().S(`<li>`)
//line views/vworkspace/Overview.html:35
				vutil.StreamIndent(qw422016, true, indent+2)
//line views/vworkspace/Overview.html:35
				qw422016.N().S(`<a href="`)
//line views/vworkspace/Overview.html:36
				qw422016.E().S(req.RouteAct(act, 0))
//line views/vworkspace/Overview.html:36
				qw422016.N().S(`">`)
//line views/vworkspace/Overview.html:37
				components.StreamSVGRef(qw422016, act.IconWithFallback(), 16, 16, "icon", ps)
//line views/vworkspace/Overview.html:38
				qw422016.E().S(act.Name())
//line views/vworkspace/Overview.html:38
				qw422016.N().S(`</a>`)
//line views/vworkspace/Overview.html:40
				vutil.StreamIndent(qw422016, true, indent+2)
//line views/vworkspace/Overview.html:41
				streamactionList(qw422016, req, act.Children, ps, indent+2)
//line views/vworkspace/Overview.html:42
				vutil.StreamIndent(qw422016, true, indent+1)
//line views/vworkspace/Overview.html:42
				qw422016.N().S(`</li>`)
//line views/vworkspace/Overview.html:44
			}
//line views/vworkspace/Overview.html:45
		}
//line views/vworkspace/Overview.html:46
	}
//line views/vworkspace/Overview.html:47
	vutil.StreamIndent(qw422016, true, indent)
//line views/vworkspace/Overview.html:47
	qw422016.N().S(`</ul>`)
//line views/vworkspace/Overview.html:49
}

//line views/vworkspace/Overview.html:49
func writeactionList(qq422016 qtio422016.Writer, req *cutil.WorkspaceRequest, acts action.Actions, ps *cutil.PageState, indent int) {
//line views/vworkspace/Overview.html:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/Overview.html:49
	streamactionList(qw422016, req, acts, ps, indent)
//line views/vworkspace/Overview.html:49
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/Overview.html:49
}

//line views/vworkspace/Overview.html:49
func actionList(req *cutil.WorkspaceRequest, acts action.Actions, ps *cutil.PageState, indent int) string {
//line views/vworkspace/Overview.html:49
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/Overview.html:49
	writeactionList(qb422016, req, acts, ps, indent)
//line views/vworkspace/Overview.html:49
	qs422016 := string(qb422016.B)
//line views/vworkspace/Overview.html:49
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/Overview.html:49
	return qs422016
//line views/vworkspace/Overview.html:49
}
