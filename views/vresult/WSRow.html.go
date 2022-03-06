// Code generated by qtc from "WSRow.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vresult/WSRow.html:1
package vresult

//line views/vresult/WSRow.html:1
import (
	"admini.dev/app/action"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/lib/schema/field"
	"admini.dev/app/lib/schema/model"
	"admini.dev/app/qualify"
	"admini.dev/views/components"
	"admini.dev/views/components/fieldview"
	"admini.dev/views/vutil"
)

//line views/vresult/WSRow.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vresult/WSRow.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vresult/WSRow.html:12
func StreamWSRow(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []interface{}, fields field.Fields, m *model.Model, indent int, showNum bool) {
//line views/vresult/WSRow.html:13
	vutil.StreamIndent(qw422016, true, indent)
//line views/vresult/WSRow.html:13
	qw422016.N().S(`<tr>`)
//line views/vresult/WSRow.html:15
	if showNum {
//line views/vresult/WSRow.html:16
		vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSRow.html:16
		qw422016.N().S(`<th class="shrink"><em>`)
//line views/vresult/WSRow.html:17
		qw422016.N().D(idx + 1)
//line views/vresult/WSRow.html:17
		qw422016.N().S(`</em></th>`)
//line views/vresult/WSRow.html:18
	}
//line views/vresult/WSRow.html:20
	for fIdx, f := range fields {
//line views/vresult/WSRow.html:21
		if m != nil && m.IsPK(f.Key, ws.PS.Logger) {
//line views/vresult/WSRow.html:22
			streamrow(qw422016, ws, act, idx, row, fields, m, fIdx, f, indent, showNum, true)
//line views/vresult/WSRow.html:23
		}
//line views/vresult/WSRow.html:24
	}
//line views/vresult/WSRow.html:25
	for fIdx, f := range fields {
//line views/vresult/WSRow.html:26
		if m == nil || !m.IsPK(f.Key, ws.PS.Logger) {
//line views/vresult/WSRow.html:27
			streamrow(qw422016, ws, act, idx, row, fields, m, fIdx, f, indent, showNum, false)
//line views/vresult/WSRow.html:28
		}
//line views/vresult/WSRow.html:29
	}
//line views/vresult/WSRow.html:30
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSRow.html:30
	qw422016.N().S(`<td class="tfill"></td>`)
//line views/vresult/WSRow.html:32
	vutil.StreamIndent(qw422016, true, indent)
//line views/vresult/WSRow.html:32
	qw422016.N().S(`</tr>`)
//line views/vresult/WSRow.html:34
}

//line views/vresult/WSRow.html:34
func WriteWSRow(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []interface{}, fields field.Fields, m *model.Model, indent int, showNum bool) {
//line views/vresult/WSRow.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSRow.html:34
	StreamWSRow(qw422016, ws, act, idx, row, fields, m, indent, showNum)
//line views/vresult/WSRow.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSRow.html:34
}

//line views/vresult/WSRow.html:34
func WSRow(ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []interface{}, fields field.Fields, m *model.Model, indent int, showNum bool) string {
//line views/vresult/WSRow.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSRow.html:34
	WriteWSRow(qb422016, ws, act, idx, row, fields, m, indent, showNum)
//line views/vresult/WSRow.html:34
	qs422016 := string(qb422016.B)
//line views/vresult/WSRow.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSRow.html:34
	return qs422016
//line views/vresult/WSRow.html:34
}

//line views/vresult/WSRow.html:36
func streamrow(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []interface{}, fields field.Fields, m *model.Model, fIdx int, f *field.Field, indent int, showNum bool, header bool) {
//line views/vresult/WSRow.html:37
	col := row[fIdx]

//line views/vresult/WSRow.html:38
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSRow.html:39
	if header {
//line views/vresult/WSRow.html:39
		qw422016.N().S(`<th>`)
//line views/vresult/WSRow.html:39
	} else {
//line views/vresult/WSRow.html:39
		qw422016.N().S(`<td>`)
//line views/vresult/WSRow.html:39
	}
//line views/vresult/WSRow.html:40
	if m == nil {
//line views/vresult/WSRow.html:41
		fieldview.StreamAny(qw422016, col, f.Type)
//line views/vresult/WSRow.html:42
	} else {
//line views/vresult/WSRow.html:43
		rels := m.ApplicableRelations(f.Key)

//line views/vresult/WSRow.html:44
		if len(rels) == 0 {
//line views/vresult/WSRow.html:45
			streamcell(qw422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:46
		} else {
//line views/vresult/WSRow.html:47
			for _, rel := range rels {
//line views/vresult/WSRow.html:49
				rowFK, err := model.GetStrings(fields, rel.SourceFields, row)
				if err != nil {
					panic(err)
				}
				src := act.Config["source"]
				if act.TypeKey == action.TypeAll.Key {
					src = ws.Path[0]
				}
				req := qualify.NewRequest("model", "view", "source", src, "model", rel.Path(), "keys", rowFK)
				quals, err := qualify.Qualify(req, ws.Project.Actions, ws.Schemata)
				if err != nil {
					panic(err)
				}

//line views/vresult/WSRow.html:63
				if len(quals) == 0 {
//line views/vresult/WSRow.html:64
					streamcell(qw422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:65
				} else {
//line views/vresult/WSRow.html:65
					qw422016.N().S(`<div class="two-pane"><div class="l">`)
//line views/vresult/WSRow.html:67
					streamcell(qw422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:67
					qw422016.N().S(`</div><div class="r">`)
//line views/vresult/WSRow.html:69
					for _, q := range quals {
//line views/vresult/WSRow.html:70
						qw422016.N().S(` `)
//line views/vresult/WSRow.html:70
						qw422016.N().S(`<a href="`)
//line views/vresult/WSRow.html:71
						qw422016.E().S(ws.Route(q.Link()...))
//line views/vresult/WSRow.html:71
						qw422016.N().S(`" title="`)
//line views/vresult/WSRow.html:71
						qw422016.E().S(q.String())
//line views/vresult/WSRow.html:71
						qw422016.N().S(`" class="rel">`)
//line views/vresult/WSRow.html:71
						components.StreamSVGRef(qw422016, q.Icon, 16, 16, "", ws.PS)
//line views/vresult/WSRow.html:71
						qw422016.N().S(`</a>`)
//line views/vresult/WSRow.html:72
					}
//line views/vresult/WSRow.html:72
					qw422016.N().S(`</div></div>`)
//line views/vresult/WSRow.html:75
				}
//line views/vresult/WSRow.html:76
			}
//line views/vresult/WSRow.html:77
		}
//line views/vresult/WSRow.html:78
	}
//line views/vresult/WSRow.html:79
	if !header {
//line views/vresult/WSRow.html:79
		qw422016.N().S(`</td>`)
//line views/vresult/WSRow.html:79
	} else {
//line views/vresult/WSRow.html:79
		qw422016.N().S(`</th>`)
//line views/vresult/WSRow.html:79
	}
//line views/vresult/WSRow.html:80
}

//line views/vresult/WSRow.html:80
func writerow(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []interface{}, fields field.Fields, m *model.Model, fIdx int, f *field.Field, indent int, showNum bool, header bool) {
//line views/vresult/WSRow.html:80
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSRow.html:80
	streamrow(qw422016, ws, act, idx, row, fields, m, fIdx, f, indent, showNum, header)
//line views/vresult/WSRow.html:80
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSRow.html:80
}

//line views/vresult/WSRow.html:80
func row(ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []interface{}, fields field.Fields, m *model.Model, fIdx int, f *field.Field, indent int, showNum bool, header bool) string {
//line views/vresult/WSRow.html:80
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSRow.html:80
	writerow(qb422016, ws, act, idx, row, fields, m, fIdx, f, indent, showNum, header)
//line views/vresult/WSRow.html:80
	qs422016 := string(qb422016.B)
//line views/vresult/WSRow.html:80
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSRow.html:80
	return qs422016
//line views/vresult/WSRow.html:80
}

//line views/vresult/WSRow.html:82
func streamcell(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, row []interface{}, fields field.Fields, m *model.Model, f *field.Field, col interface{}) {
//line views/vresult/WSRow.html:83
	if m.IsPK(f.Key, ws.PS.Logger) {
//line views/vresult/WSRow.html:85
		rowPK, err := model.GetStrings(fields, m.GetPK(ws.PS.Logger), row)
		if err != nil {
			panic(err)
		}
		link := append([]string{`v`}, rowPK...)

//line views/vresult/WSRow.html:90
		qw422016.N().S(`<a href="`)
//line views/vresult/WSRow.html:91
		qw422016.E().S(ws.RouteAct(act, 0, link...))
//line views/vresult/WSRow.html:91
		qw422016.N().S(`" class="pklink">`)
//line views/vresult/WSRow.html:92
		fieldview.StreamAny(qw422016, col, f.Type)
//line views/vresult/WSRow.html:92
		qw422016.N().S(`</a>`)
//line views/vresult/WSRow.html:94
	} else {
//line views/vresult/WSRow.html:95
		fieldview.StreamAny(qw422016, col, f.Type)
//line views/vresult/WSRow.html:96
	}
//line views/vresult/WSRow.html:97
}

//line views/vresult/WSRow.html:97
func writecell(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, row []interface{}, fields field.Fields, m *model.Model, f *field.Field, col interface{}) {
//line views/vresult/WSRow.html:97
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSRow.html:97
	streamcell(qw422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:97
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSRow.html:97
}

//line views/vresult/WSRow.html:97
func cell(ws *cutil.WorkspaceRequest, act *action.Action, row []interface{}, fields field.Fields, m *model.Model, f *field.Field, col interface{}) string {
//line views/vresult/WSRow.html:97
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSRow.html:97
	writecell(qb422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:97
	qs422016 := string(qb422016.B)
//line views/vresult/WSRow.html:97
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSRow.html:97
	return qs422016
//line views/vresult/WSRow.html:97
}
