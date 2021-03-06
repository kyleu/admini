// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsource/List.html:1
package vsource

//line views/vsource/List.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/source"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vsource/List.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsource/List.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsource/List.html:9
type List struct {
	layout.Basic
	Sources source.Sources
}

//line views/vsource/List.html:14
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/List.html:14
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/source/_new" title="add new source">`)
//line views/vsource/List.html:16
	components.StreamSVGRef(qw422016, `plus`, 20, 20, `icon`, ps)
//line views/vsource/List.html:16
	qw422016.N().S(`</a></div>
    <h3>Sources</h3>
    `)
//line views/vsource/List.html:18
	StreamTable(qw422016, p.Sources, as, ps)
//line views/vsource/List.html:18
	qw422016.N().S(`
  </div>
`)
//line views/vsource/List.html:20
}

//line views/vsource/List.html:20
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/List.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/List.html:20
	p.StreamBody(qw422016, as, ps)
//line views/vsource/List.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/List.html:20
}

//line views/vsource/List.html:20
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsource/List.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/List.html:20
	p.WriteBody(qb422016, as, ps)
//line views/vsource/List.html:20
	qs422016 := string(qb422016.B)
//line views/vsource/List.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/List.html:20
	return qs422016
//line views/vsource/List.html:20
}
