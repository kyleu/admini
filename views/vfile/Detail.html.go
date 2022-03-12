// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vfile/Detail.html:2
package vfile

//line views/vfile/Detail.html:2
import (
	"path/filepath"
	"unicode/utf8"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
)

//line views/vfile/Detail.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vfile/Detail.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vfile/Detail.html:10
func StreamDetail(qw422016 *qt422016.Writer, path []string, b []byte, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Detail.html:10
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vfile/Detail.html:12
	for idx, p := range path {
//line views/vfile/Detail.html:12
		qw422016.N().S(`/<a href="`)
//line views/vfile/Detail.html:12
		qw422016.E().S(urlPrefix)
//line views/vfile/Detail.html:12
		qw422016.N().S(`/`)
//line views/vfile/Detail.html:12
		qw422016.E().S(filepath.Join(path[:idx+1]...))
//line views/vfile/Detail.html:12
		qw422016.N().S(`">`)
//line views/vfile/Detail.html:12
		qw422016.E().S(p)
//line views/vfile/Detail.html:12
		qw422016.N().S(`</a>`)
//line views/vfile/Detail.html:12
	}
//line views/vfile/Detail.html:12
	qw422016.N().S(`</h3>
    <div class="mt">
`)
//line views/vfile/Detail.html:14
	if len(b) > (1024 * 128) {
//line views/vfile/Detail.html:14
		qw422016.N().S(`      <em>File is `)
//line views/vfile/Detail.html:15
		qw422016.N().D(len(b))
//line views/vfile/Detail.html:15
		qw422016.N().S(` bytes, which is too large for the file viewer</em>
`)
//line views/vfile/Detail.html:16
	} else if utf8.Valid(b) {
//line views/vfile/Detail.html:17
		out, _ := cutil.FormatFilename(string(b), path[len(path)-1])

//line views/vfile/Detail.html:17
		qw422016.N().S(`      `)
//line views/vfile/Detail.html:18
		qw422016.N().S(out)
//line views/vfile/Detail.html:18
		qw422016.N().S(`
`)
//line views/vfile/Detail.html:19
	} else {
//line views/vfile/Detail.html:19
		qw422016.N().S(`      <em>File is binary, which is unsupported for the file viewer</em>
`)
//line views/vfile/Detail.html:21
	}
//line views/vfile/Detail.html:21
	qw422016.N().S(`    </div>
  </div>
`)
//line views/vfile/Detail.html:24
}

//line views/vfile/Detail.html:24
func WriteDetail(qq422016 qtio422016.Writer, path []string, b []byte, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Detail.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Detail.html:24
	StreamDetail(qw422016, path, b, urlPrefix, as, ps)
//line views/vfile/Detail.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Detail.html:24
}

//line views/vfile/Detail.html:24
func Detail(path []string, b []byte, urlPrefix string, as *app.State, ps *cutil.PageState) string {
//line views/vfile/Detail.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Detail.html:24
	WriteDetail(qb422016, path, b, urlPrefix, as, ps)
//line views/vfile/Detail.html:24
	qs422016 := string(qb422016.B)
//line views/vfile/Detail.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Detail.html:24
	return qs422016
//line views/vfile/Detail.html:24
}