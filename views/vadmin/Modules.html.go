// Code generated by qtc from "Modules.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/Modules.html:1
package vadmin

//line views/vadmin/Modules.html:1
import (
	"runtime/debug"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vadmin/Modules.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Modules.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Modules.html:11
type Modules struct {
	layout.Basic
	Modules []*debug.Module
}

//line views/vadmin/Modules.html:16
func (p *Modules) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Modules.html:16
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vadmin/Modules.html:18
	qw422016.E().S(util.AppName)
//line views/vadmin/Modules.html:18
	qw422016.N().S(` v`)
//line views/vadmin/Modules.html:18
	qw422016.E().S(as.BuildInfo.Version)
//line views/vadmin/Modules.html:18
	qw422016.N().S(`</div>
    <h3>`)
//line views/vadmin/Modules.html:19
	components.StreamSVGIcon(qw422016, `archive`, ps)
//line views/vadmin/Modules.html:19
	qw422016.N().S(` Go Modules</h3>
    `)
//line views/vadmin/Modules.html:20
	streammoduleTable(qw422016, p.Modules)
//line views/vadmin/Modules.html:20
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/Modules.html:22
}

//line views/vadmin/Modules.html:22
func (p *Modules) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Modules.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Modules.html:22
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Modules.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Modules.html:22
}

//line views/vadmin/Modules.html:22
func (p *Modules) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Modules.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Modules.html:22
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Modules.html:22
	qs422016 := string(qb422016.B)
//line views/vadmin/Modules.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Modules.html:22
	return qs422016
//line views/vadmin/Modules.html:22
}

//line views/vadmin/Modules.html:24
func streammoduleTable(qw422016 *qt422016.Writer, mods []*debug.Module) {
//line views/vadmin/Modules.html:24
	qw422016.N().S(`
    <div class="overflow full-width">
      <table class="mt">
        <thead>
          <tr>
            <th>Name</th>
            <th>Version</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vadmin/Modules.html:34
	for _, m := range mods {
//line views/vadmin/Modules.html:34
		qw422016.N().S(`          <tr>
            <td><a target="_blank" rel="noopener noreferrer" href="https://`)
//line views/vadmin/Modules.html:36
		qw422016.E().S(m.Path)
//line views/vadmin/Modules.html:36
		qw422016.N().S(`">`)
//line views/vadmin/Modules.html:36
		qw422016.E().S(m.Path)
//line views/vadmin/Modules.html:36
		qw422016.N().S(`</a></td>
            <td title="`)
//line views/vadmin/Modules.html:37
		qw422016.E().S(m.Sum)
//line views/vadmin/Modules.html:37
		qw422016.N().S(`">`)
//line views/vadmin/Modules.html:37
		qw422016.E().S(m.Version)
//line views/vadmin/Modules.html:37
		qw422016.N().S(`</td>
          </tr>
`)
//line views/vadmin/Modules.html:39
	}
//line views/vadmin/Modules.html:39
	qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/vadmin/Modules.html:43
}

//line views/vadmin/Modules.html:43
func writemoduleTable(qq422016 qtio422016.Writer, mods []*debug.Module) {
//line views/vadmin/Modules.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Modules.html:43
	streammoduleTable(qw422016, mods)
//line views/vadmin/Modules.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Modules.html:43
}

//line views/vadmin/Modules.html:43
func moduleTable(mods []*debug.Module) string {
//line views/vadmin/Modules.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Modules.html:43
	writemoduleTable(qb422016, mods)
//line views/vadmin/Modules.html:43
	qs422016 := string(qb422016.B)
//line views/vadmin/Modules.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Modules.html:43
	return qs422016
//line views/vadmin/Modules.html:43
}
