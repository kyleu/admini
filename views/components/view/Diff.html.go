// Code generated by qtc from "Diff.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Diff.html:2
package view

//line views/components/view/Diff.html:2
import (
	"github.com/samber/lo"

	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
)

//line views/components/view/Diff.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Diff.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Diff.html:10
func StreamDiffs(qw422016 *qt422016.Writer, value util.Diffs) {
//line views/components/view/Diff.html:11
	if len(value) == 0 {
//line views/components/view/Diff.html:11
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/view/Diff.html:13
	} else {
//line views/components/view/Diff.html:13
		qw422016.N().S(`<div class="overflow full-width"><table class="expanded"><thead><tr><th>Path</th><th>Old</th><th></th><th>New</th></tr></thead><tbody>`)
//line views/components/view/Diff.html:25
		for _, d := range value {
//line views/components/view/Diff.html:25
			qw422016.N().S(`<tr><td style="width: 30%;"><code>`)
//line views/components/view/Diff.html:27
			qw422016.E().S(d.Path)
//line views/components/view/Diff.html:27
			qw422016.N().S(`</code></td><td style="width: 30%;"><code><em>`)
//line views/components/view/Diff.html:28
			qw422016.E().S(d.Old)
//line views/components/view/Diff.html:28
			qw422016.N().S(`</em></code></td><td style="width: 10%;">→</td><td style="width: 30%;"><code class="success">`)
//line views/components/view/Diff.html:30
			qw422016.E().S(d.New)
//line views/components/view/Diff.html:30
			qw422016.N().S(`</code></td></tr>`)
//line views/components/view/Diff.html:32
		}
//line views/components/view/Diff.html:32
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/view/Diff.html:36
	}
//line views/components/view/Diff.html:37
}

//line views/components/view/Diff.html:37
func WriteDiffs(qq422016 qtio422016.Writer, value util.Diffs) {
//line views/components/view/Diff.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Diff.html:37
	StreamDiffs(qw422016, value)
//line views/components/view/Diff.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Diff.html:37
}

//line views/components/view/Diff.html:37
func Diffs(value util.Diffs) string {
//line views/components/view/Diff.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Diff.html:37
	WriteDiffs(qb422016, value)
//line views/components/view/Diff.html:37
	qs422016 := string(qb422016.B)
//line views/components/view/Diff.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Diff.html:37
	return qs422016
//line views/components/view/Diff.html:37
}

//line views/components/view/Diff.html:39
func StreamDiffsSet(qw422016 *qt422016.Writer, key string, value util.DiffsSet, ps *cutil.PageState) {
//line views/components/view/Diff.html:40
	if len(value) == 0 {
//line views/components/view/Diff.html:40
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/view/Diff.html:42
	} else {
//line views/components/view/Diff.html:42
		qw422016.N().S(`<ul class="accordion">`)
//line views/components/view/Diff.html:44
		for idx, k := range util.ArraySorted[string](lo.Keys(value)) {
//line views/components/view/Diff.html:45
			dk, u := util.StringSplitLast(k, '^', true)

//line views/components/view/Diff.html:46
			v := value[k]

//line views/components/view/Diff.html:47
			if idx < 100 {
//line views/components/view/Diff.html:47
				qw422016.N().S(`<li><input id="accordion-`)
//line views/components/view/Diff.html:49
				qw422016.E().S(k)
//line views/components/view/Diff.html:49
				qw422016.N().S(`-`)
//line views/components/view/Diff.html:49
				qw422016.N().D(idx)
//line views/components/view/Diff.html:49
				qw422016.N().S(`" type="checkbox" hidden="hidden" /><label for="accordion-`)
//line views/components/view/Diff.html:50
				qw422016.E().S(k)
//line views/components/view/Diff.html:50
				qw422016.N().S(`-`)
//line views/components/view/Diff.html:50
				qw422016.N().D(idx)
//line views/components/view/Diff.html:50
				qw422016.N().S(`"><div class="right">`)
//line views/components/view/Diff.html:52
				if len(v) == 1 {
//line views/components/view/Diff.html:52
					qw422016.N().S(`<em>(`)
//line views/components/view/Diff.html:53
					qw422016.E().S(v[0].String())
//line views/components/view/Diff.html:53
					qw422016.N().S(`)</em>`)
//line views/components/view/Diff.html:53
					qw422016.N().S(` `)
//line views/components/view/Diff.html:54
				}
//line views/components/view/Diff.html:55
				qw422016.E().S(util.StringPlural(len(v), "diff"))
//line views/components/view/Diff.html:55
				qw422016.N().S(`</div>`)
//line views/components/view/Diff.html:57
				components.StreamExpandCollapse(qw422016, 3, ps)
//line views/components/view/Diff.html:58
				if u != "" {
//line views/components/view/Diff.html:58
					qw422016.N().S(`<a href="`)
//line views/components/view/Diff.html:58
					qw422016.E().S(u)
//line views/components/view/Diff.html:58
					qw422016.N().S(`">`)
//line views/components/view/Diff.html:58
					qw422016.E().S(dk)
//line views/components/view/Diff.html:58
					qw422016.N().S(`</a>`)
//line views/components/view/Diff.html:58
				} else {
//line views/components/view/Diff.html:58
					qw422016.E().S(dk)
//line views/components/view/Diff.html:58
				}
//line views/components/view/Diff.html:58
				qw422016.N().S(`</label><div class="bd"><div><div>`)
//line views/components/view/Diff.html:61
				StreamDiffs(qw422016, v)
//line views/components/view/Diff.html:61
				qw422016.N().S(`</div></div></div></li>`)
//line views/components/view/Diff.html:64
			}
//line views/components/view/Diff.html:65
			if idx == 100 {
//line views/components/view/Diff.html:65
				qw422016.N().S(`<li><input id="accordion-`)
//line views/components/view/Diff.html:67
				qw422016.E().S(k)
//line views/components/view/Diff.html:67
				qw422016.N().S(`-extras" type="checkbox" hidden="hidden" /><label for="accordion-`)
//line views/components/view/Diff.html:68
				qw422016.E().S(k)
//line views/components/view/Diff.html:68
				qw422016.N().S(`-extras">...and`)
//line views/components/view/Diff.html:68
				qw422016.N().S(` `)
//line views/components/view/Diff.html:68
				qw422016.N().D(len(value) - 100)
//line views/components/view/Diff.html:68
				qw422016.N().S(` `)
//line views/components/view/Diff.html:68
				qw422016.N().S(`extra</label></li>`)
//line views/components/view/Diff.html:70
			}
//line views/components/view/Diff.html:71
		}
//line views/components/view/Diff.html:71
		qw422016.N().S(`</ul>`)
//line views/components/view/Diff.html:73
	}
//line views/components/view/Diff.html:74
}

//line views/components/view/Diff.html:74
func WriteDiffsSet(qq422016 qtio422016.Writer, key string, value util.DiffsSet, ps *cutil.PageState) {
//line views/components/view/Diff.html:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Diff.html:74
	StreamDiffsSet(qw422016, key, value, ps)
//line views/components/view/Diff.html:74
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Diff.html:74
}

//line views/components/view/Diff.html:74
func DiffsSet(key string, value util.DiffsSet, ps *cutil.PageState) string {
//line views/components/view/Diff.html:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Diff.html:74
	WriteDiffsSet(qb422016, key, value, ps)
//line views/components/view/Diff.html:74
	qs422016 := string(qb422016.B)
//line views/components/view/Diff.html:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Diff.html:74
	return qs422016
//line views/components/view/Diff.html:74
}
