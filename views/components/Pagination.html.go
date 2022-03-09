// Code generated by qtc from "Pagination.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Pagination.html:2
package components

//line views/components/Pagination.html:2
import (
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app/lib/filter"
)

//line views/components/Pagination.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Pagination.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Pagination.html:8
func StreamPagination(qw422016 *qt422016.Writer, count int, params *filter.Params, u *fasthttp.URI) {
//line views/components/Pagination.html:9
	if params != nil {
//line views/components/Pagination.html:10
		if params.HasNextPage(count) {
//line views/components/Pagination.html:10
			qw422016.N().S(`<div class="right"><a href="?`)
//line views/components/Pagination.html:12
			qw422016.E().S(params.NextPage().ToQueryString(u))
//line views/components/Pagination.html:12
			qw422016.N().S(`">Next page</a></div>`)
//line views/components/Pagination.html:14
		}
//line views/components/Pagination.html:15
		if params.HasPreviousPage() {
//line views/components/Pagination.html:15
			qw422016.N().S(`<div class="left"><a href="?`)
//line views/components/Pagination.html:17
			qw422016.E().S(params.PreviousPage().ToQueryString(u))
//line views/components/Pagination.html:17
			qw422016.N().S(`">Previous page</a></div>`)
//line views/components/Pagination.html:19
		}
//line views/components/Pagination.html:20
	}
//line views/components/Pagination.html:21
}

//line views/components/Pagination.html:21
func WritePagination(qq422016 qtio422016.Writer, count int, params *filter.Params, u *fasthttp.URI) {
//line views/components/Pagination.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Pagination.html:21
	StreamPagination(qw422016, count, params, u)
//line views/components/Pagination.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/components/Pagination.html:21
}

//line views/components/Pagination.html:21
func Pagination(count int, params *filter.Params, u *fasthttp.URI) string {
//line views/components/Pagination.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Pagination.html:21
	WritePagination(qb422016, count, params, u)
//line views/components/Pagination.html:21
	qs422016 := string(qb422016.B)
//line views/components/Pagination.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Pagination.html:21
	return qs422016
//line views/components/Pagination.html:21
}
