// Code generated by qtc from "Args.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vpage/Args.html:1
package vpage

//line views/vpage/Args.html:1
import (
	"strconv"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components/edit"
	"admini.dev/admini/views/layout"
)

//line views/vpage/Args.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vpage/Args.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vpage/Args.html:11
type Args struct {
	layout.Basic
	URL        string
	Directions string
	ArgRes     *cutil.ArgResults
	Hidden     map[string]string
}

//line views/vpage/Args.html:19
func (p *Args) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Args.html:19
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vpage/Args.html:21
	if p.Directions == "" {
//line views/vpage/Args.html:21
		qw422016.N().S(`Enter Data`)
//line views/vpage/Args.html:21
	} else {
//line views/vpage/Args.html:21
		qw422016.E().S(p.Directions)
//line views/vpage/Args.html:21
	}
//line views/vpage/Args.html:21
	qw422016.N().S(`</h3>
    <form action="`)
//line views/vpage/Args.html:22
	qw422016.E().S(p.URL)
//line views/vpage/Args.html:22
	qw422016.N().S(`" method="get">
`)
//line views/vpage/Args.html:23
	for k, v := range p.Hidden {
//line views/vpage/Args.html:23
		qw422016.N().S(`      <input type="hidden" name="`)
//line views/vpage/Args.html:24
		qw422016.E().S(k)
//line views/vpage/Args.html:24
		qw422016.N().S(`" value="`)
//line views/vpage/Args.html:24
		qw422016.E().S(v)
//line views/vpage/Args.html:24
		qw422016.N().S(`" />
`)
//line views/vpage/Args.html:25
	}
//line views/vpage/Args.html:25
	qw422016.N().S(`      <div class="overflow full-width">
        <table class="mt min-200 expanded">
          <tbody>
`)
//line views/vpage/Args.html:29
	for _, arg := range p.ArgRes.Args {
//line views/vpage/Args.html:31
		v := util.OrDefault(p.ArgRes.Values.GetStringOpt(arg.Key), arg.Default)
		title := arg.Title
		if len(title) > 50 {
			title = title[:50] + "..."
		}

//line views/vpage/Args.html:37
		switch arg.Type {
//line views/vpage/Args.html:38
		case "bool":
//line views/vpage/Args.html:38
			qw422016.N().S(`            `)
//line views/vpage/Args.html:39
			edit.StreamBoolTable(qw422016, arg.Key, title, v == "true", 5, arg.Description)
//line views/vpage/Args.html:39
			qw422016.N().S(`
`)
//line views/vpage/Args.html:40
		case "textarea":
//line views/vpage/Args.html:40
			qw422016.N().S(`            `)
//line views/vpage/Args.html:41
			edit.StreamTextareaTable(qw422016, arg.Key, "", title, 12, v, 5, arg.Description)
//line views/vpage/Args.html:41
			qw422016.N().S(`
`)
//line views/vpage/Args.html:42
		case "number", "int":
//line views/vpage/Args.html:43
			i, _ := strconv.ParseInt(v, 10, 32)

//line views/vpage/Args.html:43
			qw422016.N().S(`            `)
//line views/vpage/Args.html:44
			edit.StreamIntTable(qw422016, arg.Key, "", title, int(i), 5, arg.Description)
//line views/vpage/Args.html:44
			qw422016.N().S(`
`)
//line views/vpage/Args.html:45
		default:
//line views/vpage/Args.html:45
			qw422016.N().S(`            `)
//line views/vpage/Args.html:46
			edit.StreamDatalistTable(qw422016, arg.Key, "", title, v, arg.Choices, nil, 5, arg.Description)
//line views/vpage/Args.html:46
			qw422016.N().S(`
`)
//line views/vpage/Args.html:47
		}
//line views/vpage/Args.html:48
	}
//line views/vpage/Args.html:48
	qw422016.N().S(`          </tbody>
        </table>
      </div>
      <button class="mt" type="submit">Submit</button>
    </form>
  </div>
`)
//line views/vpage/Args.html:55
}

//line views/vpage/Args.html:55
func (p *Args) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Args.html:55
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vpage/Args.html:55
	p.StreamBody(qw422016, as, ps)
//line views/vpage/Args.html:55
	qt422016.ReleaseWriter(qw422016)
//line views/vpage/Args.html:55
}

//line views/vpage/Args.html:55
func (p *Args) Body(as *app.State, ps *cutil.PageState) string {
//line views/vpage/Args.html:55
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vpage/Args.html:55
	p.WriteBody(qb422016, as, ps)
//line views/vpage/Args.html:55
	qs422016 := string(qb422016.B)
//line views/vpage/Args.html:55
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vpage/Args.html:55
	return qs422016
//line views/vpage/Args.html:55
}
