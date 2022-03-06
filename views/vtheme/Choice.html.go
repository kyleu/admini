// Code generated by qtc from "Choice.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtheme/Choice.html:2
package vtheme

//line views/vtheme/Choice.html:2
import (
	"admini.dev/app/controller/cutil"
	"admini.dev/app/lib/theme"
	"admini.dev/views/vutil"
)

//line views/vtheme/Choice.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/Choice.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/Choice.html:8
func StreamChoice(qw422016 *qt422016.Writer, themes theme.Themes, selected string, indent int, ps *cutil.PageState) {
//line views/vtheme/Choice.html:9
	vutil.StreamIndent(qw422016, true, indent)
//line views/vtheme/Choice.html:9
	qw422016.N().S(`<div class="choice">`)
//line views/vtheme/Choice.html:11
	for _, t := range themes {
//line views/vtheme/Choice.html:12
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/vtheme/Choice.html:12
		qw422016.N().S(`<label title="`)
//line views/vtheme/Choice.html:13
		qw422016.E().S(t.Key)
//line views/vtheme/Choice.html:13
		qw422016.N().S(`">`)
//line views/vtheme/Choice.html:14
		if t.Key == selected {
//line views/vtheme/Choice.html:14
			qw422016.N().S(`<input type="radio" name="theme" value="`)
//line views/vtheme/Choice.html:15
			qw422016.E().S(t.Key)
//line views/vtheme/Choice.html:15
			qw422016.N().S(`" checked="checked" />`)
//line views/vtheme/Choice.html:16
		} else {
//line views/vtheme/Choice.html:16
			qw422016.N().S(`<input type="radio" name="theme" value="`)
//line views/vtheme/Choice.html:17
			qw422016.E().S(t.Key)
//line views/vtheme/Choice.html:17
			qw422016.N().S(`" />`)
//line views/vtheme/Choice.html:18
		}
//line views/vtheme/Choice.html:19
		StreamMockupTheme(qw422016, t, true, indent+2, ps)
//line views/vtheme/Choice.html:19
		qw422016.N().S(`</label>`)
//line views/vtheme/Choice.html:21
	}
//line views/vtheme/Choice.html:22
	vutil.StreamIndent(qw422016, true, indent)
//line views/vtheme/Choice.html:22
	qw422016.N().S(`</div>`)
//line views/vtheme/Choice.html:24
}

//line views/vtheme/Choice.html:24
func WriteChoice(qq422016 qtio422016.Writer, themes theme.Themes, selected string, indent int, ps *cutil.PageState) {
//line views/vtheme/Choice.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Choice.html:24
	StreamChoice(qw422016, themes, selected, indent, ps)
//line views/vtheme/Choice.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Choice.html:24
}

//line views/vtheme/Choice.html:24
func Choice(themes theme.Themes, selected string, indent int, ps *cutil.PageState) string {
//line views/vtheme/Choice.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Choice.html:24
	WriteChoice(qb422016, themes, selected, indent, ps)
//line views/vtheme/Choice.html:24
	qs422016 := string(qb422016.B)
//line views/vtheme/Choice.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Choice.html:24
	return qs422016
//line views/vtheme/Choice.html:24
}
