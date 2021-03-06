// Code generated by qtc from "Result.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Result.html:1
package vaction

//line views/vaction/Result.html:1
import (
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/vutil"
)

//line views/vaction/Result.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Result.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Result.html:8
func StreamResultChildren(qw422016 *qt422016.Writer, req *cutil.WorkspaceRequest, act *action.Action, indent int) {
//line views/vaction/Result.html:9
	if len(act.Children) > 0 {
//line views/vaction/Result.html:9
		qw422016.N().S(`<div class="card">`)
//line views/vaction/Result.html:11
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/vaction/Result.html:11
		qw422016.N().S(`<h3>Children</h3>`)
//line views/vaction/Result.html:13
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/vaction/Result.html:13
		qw422016.N().S(`<ul>`)
//line views/vaction/Result.html:15
		for _, kid := range act.Children {
//line views/vaction/Result.html:16
			vutil.StreamIndent(qw422016, true, indent+2)
//line views/vaction/Result.html:16
			qw422016.N().S(`<li><a href="`)
//line views/vaction/Result.html:17
			qw422016.E().S(req.RouteAct(kid, 0))
//line views/vaction/Result.html:17
			qw422016.N().S(`">`)
//line views/vaction/Result.html:17
			qw422016.E().S(kid.Name())
//line views/vaction/Result.html:17
			qw422016.N().S(`</a></li>`)
//line views/vaction/Result.html:18
		}
//line views/vaction/Result.html:19
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/vaction/Result.html:19
		qw422016.N().S(`</ul>`)
//line views/vaction/Result.html:21
		vutil.StreamIndent(qw422016, true, indent)
//line views/vaction/Result.html:21
		qw422016.N().S(`</div>`)
//line views/vaction/Result.html:23
	}
//line views/vaction/Result.html:24
}

//line views/vaction/Result.html:24
func WriteResultChildren(qq422016 qtio422016.Writer, req *cutil.WorkspaceRequest, act *action.Action, indent int) {
//line views/vaction/Result.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Result.html:24
	StreamResultChildren(qw422016, req, act, indent)
//line views/vaction/Result.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Result.html:24
}

//line views/vaction/Result.html:24
func ResultChildren(req *cutil.WorkspaceRequest, act *action.Action, indent int) string {
//line views/vaction/Result.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Result.html:24
	WriteResultChildren(qb422016, req, act, indent)
//line views/vaction/Result.html:24
	qs422016 := string(qb422016.B)
//line views/vaction/Result.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Result.html:24
	return qs422016
//line views/vaction/Result.html:24
}

//line views/vaction/Result.html:26
func StreamResultDebug(qw422016 *qt422016.Writer, req *cutil.WorkspaceRequest, act *action.Action) {
//line views/vaction/Result.html:26
	qw422016.N().S(`
  `)
//line views/vaction/Result.html:27
	components.StreamJSONModal(qw422016, "action", "Action", act, 1)
//line views/vaction/Result.html:27
	qw422016.N().S(`
`)
//line views/vaction/Result.html:28
}

//line views/vaction/Result.html:28
func WriteResultDebug(qq422016 qtio422016.Writer, req *cutil.WorkspaceRequest, act *action.Action) {
//line views/vaction/Result.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Result.html:28
	StreamResultDebug(qw422016, req, act)
//line views/vaction/Result.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Result.html:28
}

//line views/vaction/Result.html:28
func ResultDebug(req *cutil.WorkspaceRequest, act *action.Action) string {
//line views/vaction/Result.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Result.html:28
	WriteResultDebug(qb422016, req, act)
//line views/vaction/Result.html:28
	qs422016 := string(qb422016.B)
//line views/vaction/Result.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Result.html:28
	return qs422016
//line views/vaction/Result.html:28
}
