// Code generated by qtc from "Modules.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Modules.html:2
package vadmin

//line views/vadmin/Modules.html:2
import (
	"runtime/debug"

	"admini.dev/app"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/util"
	"admini.dev/views/layout"
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
	Info *util.DebugInfo
}

//line views/vadmin/Modules.html:16
func (p *Modules) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Modules.html:16
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/Modules.html:18
	qw422016.E().S(util.AppName)
//line views/vadmin/Modules.html:18
	qw422016.N().S(` v`)
//line views/vadmin/Modules.html:18
	qw422016.E().S(as.BuildInfo.Version)
//line views/vadmin/Modules.html:18
	qw422016.N().S(`</h3>
    <ul class="mt">
`)
//line views/vadmin/Modules.html:20
	for _, k := range p.Info.Tags.Order {
//line views/vadmin/Modules.html:20
		qw422016.N().S(`      <li><strong>`)
//line views/vadmin/Modules.html:21
		qw422016.E().S(k)
//line views/vadmin/Modules.html:21
		qw422016.N().S(`</strong>: `)
//line views/vadmin/Modules.html:21
		qw422016.E().S(p.Info.Tags.GetSimple(k))
//line views/vadmin/Modules.html:21
		qw422016.N().S(`</li>
`)
//line views/vadmin/Modules.html:22
	}
//line views/vadmin/Modules.html:22
	qw422016.N().S(`    </ul>
  </div>
  `)
//line views/vadmin/Modules.html:25
	streammoduleList(qw422016, p.Info.Mods)
//line views/vadmin/Modules.html:25
	qw422016.N().S(`
`)
//line views/vadmin/Modules.html:26
}

//line views/vadmin/Modules.html:26
func (p *Modules) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Modules.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Modules.html:26
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Modules.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Modules.html:26
}

//line views/vadmin/Modules.html:26
func (p *Modules) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Modules.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Modules.html:26
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Modules.html:26
	qs422016 := string(qb422016.B)
//line views/vadmin/Modules.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Modules.html:26
	return qs422016
//line views/vadmin/Modules.html:26
}

//line views/vadmin/Modules.html:28
func streammoduleList(qw422016 *qt422016.Writer, Mods []*debug.Module) {
//line views/vadmin/Modules.html:28
	qw422016.N().S(`
  <div class="card">
    <h3>Go Modules</h3>
    <table class="mt">
      <thead>
        <tr>
          <th>Name</th>
          <th>Version</th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vadmin/Modules.html:39
	for _, m := range Mods {
//line views/vadmin/Modules.html:39
		qw422016.N().S(`        <tr>
          <td><a target="_blank" rel="noopener noreferrer" href="https://`)
//line views/vadmin/Modules.html:41
		qw422016.E().S(m.Path)
//line views/vadmin/Modules.html:41
		qw422016.N().S(`">`)
//line views/vadmin/Modules.html:41
		qw422016.E().S(m.Path)
//line views/vadmin/Modules.html:41
		qw422016.N().S(`</a></td>
          <td title="`)
//line views/vadmin/Modules.html:42
		qw422016.E().S(m.Sum)
//line views/vadmin/Modules.html:42
		qw422016.N().S(`">`)
//line views/vadmin/Modules.html:42
		qw422016.E().S(m.Version)
//line views/vadmin/Modules.html:42
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vadmin/Modules.html:44
	}
//line views/vadmin/Modules.html:44
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vadmin/Modules.html:48
}

//line views/vadmin/Modules.html:48
func writemoduleList(qq422016 qtio422016.Writer, Mods []*debug.Module) {
//line views/vadmin/Modules.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Modules.html:48
	streammoduleList(qw422016, Mods)
//line views/vadmin/Modules.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Modules.html:48
}

//line views/vadmin/Modules.html:48
func moduleList(Mods []*debug.Module) string {
//line views/vadmin/Modules.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Modules.html:48
	writemoduleList(qb422016, Mods)
//line views/vadmin/Modules.html:48
	qs422016 := string(qb422016.B)
//line views/vadmin/Modules.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Modules.html:48
	return qs422016
//line views/vadmin/Modules.html:48
}
