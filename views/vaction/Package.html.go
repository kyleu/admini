// Code generated by qtc from "Package.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Package.html:1
package vaction

//line views/vaction/Package.html:1
import (
	"admini.dev/app"
	"admini.dev/app/action"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/lib/schema/model"
	"admini.dev/views/layout"
)

//line views/vaction/Package.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Package.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Package.html:9
type Package struct {
	layout.Basic
	Req *cutil.WorkspaceRequest
	Act *action.Action
	Pkg *model.Package
}

//line views/vaction/Package.html:16
func (p *Package) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Package.html:16
	qw422016.N().S(`
  <div class="card">
    `)
//line views/vaction/Package.html:18
	StreamActionHeader(qw422016, p.Act, "Package ["+p.Pkg.Key+"]", ps)
//line views/vaction/Package.html:18
	qw422016.N().S(`
`)
//line views/vaction/Package.html:19
	if len(p.Pkg.ChildPackages) > 0 {
//line views/vaction/Package.html:19
		qw422016.N().S(`    <ul>
`)
//line views/vaction/Package.html:21
		for _, pkg := range p.Pkg.ChildPackages {
//line views/vaction/Package.html:21
			qw422016.N().S(`      <li><a href="`)
//line views/vaction/Package.html:22
			qw422016.E().S(p.Req.RouteAct(p.Act, 0, pkg.Key))
//line views/vaction/Package.html:22
			qw422016.N().S(`">`)
//line views/vaction/Package.html:22
			qw422016.E().S(pkg.Name())
//line views/vaction/Package.html:22
			qw422016.N().S(`</a></li>
`)
//line views/vaction/Package.html:23
		}
//line views/vaction/Package.html:23
		qw422016.N().S(`    </ul>
`)
//line views/vaction/Package.html:25
	}
//line views/vaction/Package.html:25
	qw422016.N().S(`
`)
//line views/vaction/Package.html:27
	if len(p.Pkg.ChildModels) > 0 {
//line views/vaction/Package.html:27
		qw422016.N().S(`    <ul>
`)
//line views/vaction/Package.html:29
		for _, m := range p.Pkg.ChildModels {
//line views/vaction/Package.html:29
			qw422016.N().S(`      <li><a href="`)
//line views/vaction/Package.html:30
			qw422016.E().S(p.Req.RouteAct(p.Act, 0, m.Key))
//line views/vaction/Package.html:30
			qw422016.N().S(`">`)
//line views/vaction/Package.html:30
			qw422016.E().S(m.Name())
//line views/vaction/Package.html:30
			qw422016.N().S(`</a></li>
`)
//line views/vaction/Package.html:31
		}
//line views/vaction/Package.html:31
		qw422016.N().S(`    </ul>
`)
//line views/vaction/Package.html:33
	}
//line views/vaction/Package.html:33
	qw422016.N().S(`  </div>
`)
//line views/vaction/Package.html:35
	StreamResultChildren(qw422016, p.Req, p.Act, 1)
//line views/vaction/Package.html:36
	StreamResultDebug(qw422016, p.Req, p.Act)
//line views/vaction/Package.html:37
}

//line views/vaction/Package.html:37
func (p *Package) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Package.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Package.html:37
	p.StreamBody(qw422016, as, ps)
//line views/vaction/Package.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Package.html:37
}

//line views/vaction/Package.html:37
func (p *Package) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/Package.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Package.html:37
	p.WriteBody(qb422016, as, ps)
//line views/vaction/Package.html:37
	qs422016 := string(qb422016.B)
//line views/vaction/Package.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Package.html:37
	return qs422016
//line views/vaction/Package.html:37
}
