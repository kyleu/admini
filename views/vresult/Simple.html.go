// Code generated by qtc from "Simple.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vresult/Simple.html:1
package vresult

//line views/vresult/Simple.html:1
import (
	"fmt"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/filter"
	"admini.dev/admini/app/lib/schema/field"
	"admini.dev/admini/app/result"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/components/view"
)

//line views/vresult/Simple.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vresult/Simple.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vresult/Simple.html:14
func StreamSimple(qw422016 *qt422016.Writer, r *result.Result, indent int, as *app.State, ps *cutil.PageState) {
//line views/vresult/Simple.html:14
	qw422016.N().S(`<div class="right">`)
//line views/vresult/Simple.html:15
	qw422016.E().S(util.StringPlural(len(r.Data), `row`))
//line views/vresult/Simple.html:15
	qw422016.N().S(`</div>`)
//line views/vresult/Simple.html:16
	components.StreamIndent(qw422016, true, indent)
//line views/vresult/Simple.html:16
	qw422016.N().S(`<h3>`)
//line views/vresult/Simple.html:17
	qw422016.E().S(r.Title)
//line views/vresult/Simple.html:17
	qw422016.N().S(`</h3>`)
//line views/vresult/Simple.html:18
	streamsimpleTable(qw422016, r, indent, false, nil, as, ps)
//line views/vresult/Simple.html:19
}

//line views/vresult/Simple.html:19
func WriteSimple(qq422016 qtio422016.Writer, r *result.Result, indent int, as *app.State, ps *cutil.PageState) {
//line views/vresult/Simple.html:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/Simple.html:19
	StreamSimple(qw422016, r, indent, as, ps)
//line views/vresult/Simple.html:19
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/Simple.html:19
}

//line views/vresult/Simple.html:19
func Simple(r *result.Result, indent int, as *app.State, ps *cutil.PageState) string {
//line views/vresult/Simple.html:19
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/Simple.html:19
	WriteSimple(qb422016, r, indent, as, ps)
//line views/vresult/Simple.html:19
	qs422016 := string(qb422016.B)
//line views/vresult/Simple.html:19
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/Simple.html:19
	return qs422016
//line views/vresult/Simple.html:19
}

//line views/vresult/Simple.html:21
func streamsimpleTable(qw422016 *qt422016.Writer, r *result.Result, indent int, showNum bool, params *filter.Params, as *app.State, ps *cutil.PageState) {
//line views/vresult/Simple.html:22
	components.StreamIndent(qw422016, true, indent)
//line views/vresult/Simple.html:22
	qw422016.N().S(`<table class="result-table">`)
//line views/vresult/Simple.html:24
	components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/Simple.html:24
	qw422016.N().S(`<thead>`)
//line views/vresult/Simple.html:26
	components.StreamIndent(qw422016, true, indent+2)
//line views/vresult/Simple.html:26
	qw422016.N().S(`<tr>`)
//line views/vresult/Simple.html:28
	if showNum {
//line views/vresult/Simple.html:29
		components.StreamIndent(qw422016, true, indent+3)
//line views/vresult/Simple.html:29
		qw422016.N().S(`<th class="no-padding"><div class="resize"></div></th>`)
//line views/vresult/Simple.html:31
	}
//line views/vresult/Simple.html:32
	for fIdx, field := range r.Fields {
//line views/vresult/Simple.html:33
		components.StreamIndent(qw422016, true, indent+3)
//line views/vresult/Simple.html:34
		tooltip := fmt.Sprintf(`%s: ordinal %d (%s)`, field.Key, fIdx, field.Type)

//line views/vresult/Simple.html:35
		components.StreamTableHeader(qw422016, "x", field.Key, field.Key, params, "", ps.URI, tooltip, false, "", true, ps)
//line views/vresult/Simple.html:36
	}
//line views/vresult/Simple.html:37
	components.StreamIndent(qw422016, true, indent+3)
//line views/vresult/Simple.html:37
	qw422016.N().S(`<th class="tfill"></th>`)
//line views/vresult/Simple.html:39
	components.StreamIndent(qw422016, true, indent+2)
//line views/vresult/Simple.html:39
	qw422016.N().S(`</tr>`)
//line views/vresult/Simple.html:41
	components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/Simple.html:41
	qw422016.N().S(`</thead>`)
//line views/vresult/Simple.html:43
	components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/Simple.html:43
	qw422016.N().S(`<tbody>`)
//line views/vresult/Simple.html:45
	for rIdx, row := range r.Data {
//line views/vresult/Simple.html:46
		streamsimpleRow(qw422016, rIdx, row, r.Fields, indent+2, showNum, params, as, ps)
//line views/vresult/Simple.html:47
	}
//line views/vresult/Simple.html:49
	if params.HasNextPage(r.Count) || params.HasPreviousPage() {
//line views/vresult/Simple.html:50
		components.StreamIndent(qw422016, true, indent+2)
//line views/vresult/Simple.html:50
		qw422016.N().S(`<tr><td colspan="`)
//line views/vresult/Simple.html:51
		qw422016.N().D(len(r.Fields) + 1)
//line views/vresult/Simple.html:51
		qw422016.N().S(`">`)
//line views/vresult/Simple.html:51
		components.StreamPagination(qw422016, r.Count, params, ps.URI)
//line views/vresult/Simple.html:51
		qw422016.N().S(`</td></tr>`)
//line views/vresult/Simple.html:52
	}
//line views/vresult/Simple.html:53
	components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/Simple.html:53
	qw422016.N().S(`</tbody>`)
//line views/vresult/Simple.html:55
	components.StreamIndent(qw422016, true, indent)
//line views/vresult/Simple.html:55
	qw422016.N().S(`</table>`)
//line views/vresult/Simple.html:57
}

//line views/vresult/Simple.html:57
func writesimpleTable(qq422016 qtio422016.Writer, r *result.Result, indent int, showNum bool, params *filter.Params, as *app.State, ps *cutil.PageState) {
//line views/vresult/Simple.html:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/Simple.html:57
	streamsimpleTable(qw422016, r, indent, showNum, params, as, ps)
//line views/vresult/Simple.html:57
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/Simple.html:57
}

//line views/vresult/Simple.html:57
func simpleTable(r *result.Result, indent int, showNum bool, params *filter.Params, as *app.State, ps *cutil.PageState) string {
//line views/vresult/Simple.html:57
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/Simple.html:57
	writesimpleTable(qb422016, r, indent, showNum, params, as, ps)
//line views/vresult/Simple.html:57
	qs422016 := string(qb422016.B)
//line views/vresult/Simple.html:57
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/Simple.html:57
	return qs422016
//line views/vresult/Simple.html:57
}

//line views/vresult/Simple.html:59
func streamsimpleRow(qw422016 *qt422016.Writer, idx int, row []any, fields field.Fields, indent int, showNum bool, params *filter.Params, as *app.State, ps *cutil.PageState) {
//line views/vresult/Simple.html:60
	components.StreamIndent(qw422016, true, indent)
//line views/vresult/Simple.html:60
	qw422016.N().S(`<tr>`)
//line views/vresult/Simple.html:62
	if showNum {
//line views/vresult/Simple.html:63
		components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/Simple.html:63
		qw422016.N().S(`<th><em>`)
//line views/vresult/Simple.html:64
		qw422016.N().D(idx + 1)
//line views/vresult/Simple.html:64
		qw422016.N().S(`</em></th>`)
//line views/vresult/Simple.html:65
	}
//line views/vresult/Simple.html:67
	for fIdx, f := range fields {
//line views/vresult/Simple.html:68
		col := row[fIdx]

//line views/vresult/Simple.html:69
		components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/Simple.html:69
		qw422016.N().S(`<td>`)
//line views/vresult/Simple.html:70
		view.StreamAnyByType(qw422016, col, f.Type, ps)
//line views/vresult/Simple.html:70
		qw422016.N().S(`</td>`)
//line views/vresult/Simple.html:71
	}
//line views/vresult/Simple.html:72
	components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/Simple.html:72
	qw422016.N().S(`<td class="tfill"></td>`)
//line views/vresult/Simple.html:74
	components.StreamIndent(qw422016, true, indent)
//line views/vresult/Simple.html:74
	qw422016.N().S(`</tr>`)
//line views/vresult/Simple.html:76
}

//line views/vresult/Simple.html:76
func writesimpleRow(qq422016 qtio422016.Writer, idx int, row []any, fields field.Fields, indent int, showNum bool, params *filter.Params, as *app.State, ps *cutil.PageState) {
//line views/vresult/Simple.html:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/Simple.html:76
	streamsimpleRow(qw422016, idx, row, fields, indent, showNum, params, as, ps)
//line views/vresult/Simple.html:76
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/Simple.html:76
}

//line views/vresult/Simple.html:76
func simpleRow(idx int, row []any, fields field.Fields, indent int, showNum bool, params *filter.Params, as *app.State, ps *cutil.PageState) string {
//line views/vresult/Simple.html:76
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/Simple.html:76
	writesimpleRow(qb422016, idx, row, fields, indent, showNum, params, as, ps)
//line views/vresult/Simple.html:76
	qs422016 := string(qb422016.B)
//line views/vresult/Simple.html:76
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/Simple.html:76
	return qs422016
//line views/vresult/Simple.html:76
}
