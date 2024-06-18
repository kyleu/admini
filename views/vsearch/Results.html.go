// Code generated by qtc from "Results.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsearch/Results.html:1
package vsearch

//line views/vsearch/Results.html:1
import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/search"
	"admini.dev/admini/app/lib/search/result"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/layout"
)

//line views/vsearch/Results.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsearch/Results.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsearch/Results.html:11
type Results struct {
	layout.Basic
	Params  *search.Params
	Results result.Results
	Errors  []error
}

//line views/vsearch/Results.html:18
func (p *Results) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsearch/Results.html:18
	qw422016.N().S(`
  <div class="card">
    <h3 title="Search results for [`)
//line views/vsearch/Results.html:20
	qw422016.E().S(p.Params.Q)
//line views/vsearch/Results.html:20
	qw422016.N().S(`]">`)
//line views/vsearch/Results.html:20
	components.StreamSVGIcon(qw422016, "search", ps)
//line views/vsearch/Results.html:20
	qw422016.N().S(` Search Results</h3>
    <form class="mt expanded" action="`)
//line views/vsearch/Results.html:21
	qw422016.E().S(ps.SearchPath)
//line views/vsearch/Results.html:21
	qw422016.N().S(`">
      <input name="q" value="`)
//line views/vsearch/Results.html:22
	qw422016.E().S(p.Params.Q)
//line views/vsearch/Results.html:22
	qw422016.N().S(`" />
      <div class="mt"><button type="submit">Search Again</button></div>
    </form>
  </div>
`)
//line views/vsearch/Results.html:26
	if p.Params.Q != "" && len(p.Results) == 0 {
//line views/vsearch/Results.html:26
		qw422016.N().S(`  <div class="card">
    <h3>No results</h3>
  </div>
`)
//line views/vsearch/Results.html:30
	}
//line views/vsearch/Results.html:31
	for _, res := range p.Results {
//line views/vsearch/Results.html:31
		qw422016.N().S(`  `)
//line views/vsearch/Results.html:32
		StreamResult(qw422016, res, p.Params, as, ps)
//line views/vsearch/Results.html:32
		qw422016.N().S(`
`)
//line views/vsearch/Results.html:33
	}
//line views/vsearch/Results.html:33
	qw422016.N().S(`  `)
//line views/vsearch/Results.html:34
	if len(p.Errors) > 0 {
//line views/vsearch/Results.html:34
		qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vsearch/Results.html:36
		components.StreamSVGIcon(qw422016, "error", ps)
//line views/vsearch/Results.html:36
		qw422016.N().S(` `)
//line views/vsearch/Results.html:36
		qw422016.E().S(util.StringPlural(len(p.Errors), "Error"))
//line views/vsearch/Results.html:36
		qw422016.N().S(`</h3>
    <ul class="mt">
`)
//line views/vsearch/Results.html:38
		for _, e := range p.Errors {
//line views/vsearch/Results.html:38
			qw422016.N().S(`      <li>`)
//line views/vsearch/Results.html:39
			qw422016.E().S(e.Error())
//line views/vsearch/Results.html:39
			qw422016.N().S(`</li>
`)
//line views/vsearch/Results.html:40
		}
//line views/vsearch/Results.html:40
		qw422016.N().S(`    </ul>
  </div>
  `)
//line views/vsearch/Results.html:43
	}
//line views/vsearch/Results.html:43
	qw422016.N().S(`
`)
//line views/vsearch/Results.html:44
}

//line views/vsearch/Results.html:44
func (p *Results) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsearch/Results.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsearch/Results.html:44
	p.StreamBody(qw422016, as, ps)
//line views/vsearch/Results.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/vsearch/Results.html:44
}

//line views/vsearch/Results.html:44
func (p *Results) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsearch/Results.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsearch/Results.html:44
	p.WriteBody(qb422016, as, ps)
//line views/vsearch/Results.html:44
	qs422016 := string(qb422016.B)
//line views/vsearch/Results.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsearch/Results.html:44
	return qs422016
//line views/vsearch/Results.html:44
}
