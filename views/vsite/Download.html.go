// Code generated by qtc from "Download.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsite/Download.html:2
package vsite

//line views/vsite/Download.html:2
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/site/download"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vsite/Download.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsite/Download.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsite/Download.html:11
type Download struct {
	layout.Basic
	Links download.Links
}

//line views/vsite/Download.html:16
func (p *Download) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/Download.html:16
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vsite/Download.html:18
	components.StreamSVGRefIcon(qw422016, "download", ps)
//line views/vsite/Download.html:18
	qw422016.N().S(` Download `)
//line views/vsite/Download.html:18
	qw422016.E().S(util.AppName)
//line views/vsite/Download.html:18
	qw422016.N().S(` `)
//line views/vsite/Download.html:18
	qw422016.E().S(as.BuildInfo.Version)
//line views/vsite/Download.html:18
	qw422016.N().S(`</h3>
  </div>

  <div class="card">
    <h3>Desktop Version</h3>
    <em>Standalone application using your platform's native web viewer</em>
    <ul class="mt">
`)
//line views/vsite/Download.html:25
	for _, link := range p.Links.GetByModes("desktop") {
//line views/vsite/Download.html:25
		qw422016.N().S(`      <li>
        <a href="https://github.com/kyleu/projectforge/releases/download/v`)
//line views/vsite/Download.html:27
		qw422016.E().S(as.BuildInfo.Version)
//line views/vsite/Download.html:27
		qw422016.N().S(`/`)
//line views/vsite/Download.html:27
		qw422016.E().S(link.URL)
//line views/vsite/Download.html:27
		qw422016.N().S(`">
          `)
//line views/vsite/Download.html:28
		components.StreamSVGRef(qw422016, link.OSIcon(), 20, 20, "icon", ps)
//line views/vsite/Download.html:28
		qw422016.N().S(` `)
//line views/vsite/Download.html:28
		qw422016.E().S(link.OSString())
//line views/vsite/Download.html:28
		qw422016.N().S(`
        </a>
        <div class="clear"></div>
      </li>
`)
//line views/vsite/Download.html:32
	}
//line views/vsite/Download.html:32
	qw422016.N().S(`    </ul>
  </div>

  <div class="card">
    <h3>Server Version</h3>
    <em>A command line interface that can launch a web server</em>
    <table class="mt">
      <tbody>
`)
//line views/vsite/Download.html:41
	var currentOS string

//line views/vsite/Download.html:42
	for _, link := range p.Links.GetByModes("server", "mobile") {
//line views/vsite/Download.html:43
		if currentOS != link.OS {
//line views/vsite/Download.html:44
			if currentOS != "" {
//line views/vsite/Download.html:44
				qw422016.N().S(`          </td>
        </tr>
`)
//line views/vsite/Download.html:47
			}
//line views/vsite/Download.html:48
			currentOS = link.OS

//line views/vsite/Download.html:48
			qw422016.N().S(`        <tr>
          <td>`)
//line views/vsite/Download.html:50
			qw422016.E().S(link.OSString())
//line views/vsite/Download.html:50
			qw422016.N().S(`</td>
          <td>
`)
//line views/vsite/Download.html:52
		}
//line views/vsite/Download.html:53
		if link.OS == "linux" && (link.Arch == "ppc64" || link.Arch == "mips64_hardfloat" || link.Arch == "mips_hardfloat") {
//line views/vsite/Download.html:53
			qw422016.N().S(`            <br />
`)
//line views/vsite/Download.html:55
		}
//line views/vsite/Download.html:55
		qw422016.N().S(`            <a href="https://github.com/kyleu/admini/releases/download/v`)
//line views/vsite/Download.html:56
		qw422016.E().S(as.BuildInfo.Version)
//line views/vsite/Download.html:56
		qw422016.N().S(`/`)
//line views/vsite/Download.html:56
		qw422016.E().S(link.URL)
//line views/vsite/Download.html:56
		qw422016.N().S(`">`)
//line views/vsite/Download.html:56
		qw422016.E().S(link.Arch)
//line views/vsite/Download.html:56
		qw422016.N().S(`</a>
`)
//line views/vsite/Download.html:57
	}
//line views/vsite/Download.html:57
	qw422016.N().S(`          </td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vsite/Download.html:63
}

//line views/vsite/Download.html:63
func (p *Download) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/Download.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsite/Download.html:63
	p.StreamBody(qw422016, as, ps)
//line views/vsite/Download.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/vsite/Download.html:63
}

//line views/vsite/Download.html:63
func (p *Download) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsite/Download.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsite/Download.html:63
	p.WriteBody(qb422016, as, ps)
//line views/vsite/Download.html:63
	qs422016 := string(qb422016.B)
//line views/vsite/Download.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsite/Download.html:63
	return qs422016
//line views/vsite/Download.html:63
}
