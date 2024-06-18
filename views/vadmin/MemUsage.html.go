// Code generated by qtc from "MemUsage.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/MemUsage.html:1
package vadmin

//line views/vadmin/MemUsage.html:1
import (
	"runtime"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vadmin/MemUsage.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/MemUsage.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/MemUsage.html:10
type MemUsage struct {
	layout.Basic
	Mem *runtime.MemStats
}

//line views/vadmin/MemUsage.html:15
func (p *MemUsage) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/MemUsage.html:15
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/MemUsage.html:17
	components.StreamSVGIcon(qw422016, `desktop`, ps)
//line views/vadmin/MemUsage.html:17
	qw422016.N().S(` Memory Usage</h3>
    <em>Better formatting is coming soon</em>
    `)
//line views/vadmin/MemUsage.html:19
	qw422016.N().S(components.JSON(p.Mem))
//line views/vadmin/MemUsage.html:19
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/MemUsage.html:21
}

//line views/vadmin/MemUsage.html:21
func (p *MemUsage) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/MemUsage.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/MemUsage.html:21
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/MemUsage.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/MemUsage.html:21
}

//line views/vadmin/MemUsage.html:21
func (p *MemUsage) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/MemUsage.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/MemUsage.html:21
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/MemUsage.html:21
	qs422016 := string(qb422016.B)
//line views/vadmin/MemUsage.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/MemUsage.html:21
	return qs422016
//line views/vadmin/MemUsage.html:21
}
