// Code generated by qtc from "WSRow.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vresult/WSRow.html:1
package vresult

//line views/vresult/WSRow.html:1
import (
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema/field"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/qualify"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/components/view"
)

//line views/vresult/WSRow.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vresult/WSRow.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vresult/WSRow.html:11
func StreamWSRow(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []any, fields field.Fields, m *model.Model, indent int, showNum bool) {
//line views/vresult/WSRow.html:12
	components.StreamIndent(qw422016, true, indent)
//line views/vresult/WSRow.html:12
	qw422016.N().S(`<tr>`)
//line views/vresult/WSRow.html:14
	if showNum {
//line views/vresult/WSRow.html:15
		components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSRow.html:15
		qw422016.N().S(`<th class="shrink"><em>`)
//line views/vresult/WSRow.html:16
		qw422016.N().D(idx + 1)
//line views/vresult/WSRow.html:16
		qw422016.N().S(`</em></th>`)
//line views/vresult/WSRow.html:17
	}
//line views/vresult/WSRow.html:19
	for fIdx, f := range fields {
//line views/vresult/WSRow.html:20
		if m != nil && m.IsPK(f.Key, ws.PS.Logger) {
//line views/vresult/WSRow.html:21
			streamrow(qw422016, ws, act, idx, row, fields, m, fIdx, f, indent, showNum, true)
//line views/vresult/WSRow.html:22
		}
//line views/vresult/WSRow.html:23
	}
//line views/vresult/WSRow.html:24
	for fIdx, f := range fields {
//line views/vresult/WSRow.html:25
		if m == nil || !m.IsPK(f.Key, ws.PS.Logger) {
//line views/vresult/WSRow.html:26
			streamrow(qw422016, ws, act, idx, row, fields, m, fIdx, f, indent, showNum, false)
//line views/vresult/WSRow.html:27
		}
//line views/vresult/WSRow.html:28
	}
//line views/vresult/WSRow.html:29
	components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSRow.html:29
	qw422016.N().S(`<td class="tfill"></td>`)
//line views/vresult/WSRow.html:31
	components.StreamIndent(qw422016, true, indent)
//line views/vresult/WSRow.html:31
	qw422016.N().S(`</tr>`)
//line views/vresult/WSRow.html:33
}

//line views/vresult/WSRow.html:33
func WriteWSRow(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []any, fields field.Fields, m *model.Model, indent int, showNum bool) {
//line views/vresult/WSRow.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSRow.html:33
	StreamWSRow(qw422016, ws, act, idx, row, fields, m, indent, showNum)
//line views/vresult/WSRow.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSRow.html:33
}

//line views/vresult/WSRow.html:33
func WSRow(ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []any, fields field.Fields, m *model.Model, indent int, showNum bool) string {
//line views/vresult/WSRow.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSRow.html:33
	WriteWSRow(qb422016, ws, act, idx, row, fields, m, indent, showNum)
//line views/vresult/WSRow.html:33
	qs422016 := string(qb422016.B)
//line views/vresult/WSRow.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSRow.html:33
	return qs422016
//line views/vresult/WSRow.html:33
}

//line views/vresult/WSRow.html:35
func streamrow(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []any, fields field.Fields, m *model.Model, fIdx int, f *field.Field, indent int, showNum bool, header bool) {
//line views/vresult/WSRow.html:36
	col := row[fIdx]

//line views/vresult/WSRow.html:37
	components.StreamIndent(qw422016, true, indent+1)
//line views/vresult/WSRow.html:38
	if header {
//line views/vresult/WSRow.html:38
		qw422016.N().S(`<th>`)
//line views/vresult/WSRow.html:38
	} else {
//line views/vresult/WSRow.html:38
		qw422016.N().S(`<td>`)
//line views/vresult/WSRow.html:38
	}
//line views/vresult/WSRow.html:39
	if m == nil {
//line views/vresult/WSRow.html:40
		view.StreamAnyByType(qw422016, col, f.Type)
//line views/vresult/WSRow.html:41
	} else {
//line views/vresult/WSRow.html:42
		rels := m.ApplicableRelations(f.Key)

//line views/vresult/WSRow.html:43
		if len(rels) == 0 {
//line views/vresult/WSRow.html:44
			streamcell(qw422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:45
		} else {
//line views/vresult/WSRow.html:46
			for _, rel := range rels {
//line views/vresult/WSRow.html:48
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

//line views/vresult/WSRow.html:62
				if len(quals) == 0 {
//line views/vresult/WSRow.html:63
					streamcell(qw422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:64
				} else {
//line views/vresult/WSRow.html:64
					qw422016.N().S(`<div class="two-pane"><div class="l">`)
//line views/vresult/WSRow.html:66
					streamcell(qw422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:66
					qw422016.N().S(`</div><div class="r">`)
//line views/vresult/WSRow.html:68
					for _, q := range quals {
//line views/vresult/WSRow.html:69
						qw422016.N().S(` `)
//line views/vresult/WSRow.html:69
						qw422016.N().S(`<a href="`)
//line views/vresult/WSRow.html:70
						qw422016.E().S(ws.Route(q.Link()...))
//line views/vresult/WSRow.html:70
						qw422016.N().S(`" title="`)
//line views/vresult/WSRow.html:70
						qw422016.E().S(q.String())
//line views/vresult/WSRow.html:70
						qw422016.N().S(`" class="rel">`)
//line views/vresult/WSRow.html:70
						components.StreamSVGRef(qw422016, q.Icon, 16, 16, "", ws.PS)
//line views/vresult/WSRow.html:70
						qw422016.N().S(`</a>`)
//line views/vresult/WSRow.html:71
					}
//line views/vresult/WSRow.html:71
					qw422016.N().S(`</div></div>`)
//line views/vresult/WSRow.html:74
				}
//line views/vresult/WSRow.html:75
			}
//line views/vresult/WSRow.html:76
		}
//line views/vresult/WSRow.html:77
	}
//line views/vresult/WSRow.html:78
	if !header {
//line views/vresult/WSRow.html:78
		qw422016.N().S(`</td>`)
//line views/vresult/WSRow.html:78
	} else {
//line views/vresult/WSRow.html:78
		qw422016.N().S(`</th>`)
//line views/vresult/WSRow.html:78
	}
//line views/vresult/WSRow.html:79
}

//line views/vresult/WSRow.html:79
func writerow(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []any, fields field.Fields, m *model.Model, fIdx int, f *field.Field, indent int, showNum bool, header bool) {
//line views/vresult/WSRow.html:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSRow.html:79
	streamrow(qw422016, ws, act, idx, row, fields, m, fIdx, f, indent, showNum, header)
//line views/vresult/WSRow.html:79
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSRow.html:79
}

//line views/vresult/WSRow.html:79
func row(ws *cutil.WorkspaceRequest, act *action.Action, idx int, row []any, fields field.Fields, m *model.Model, fIdx int, f *field.Field, indent int, showNum bool, header bool) string {
//line views/vresult/WSRow.html:79
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSRow.html:79
	writerow(qb422016, ws, act, idx, row, fields, m, fIdx, f, indent, showNum, header)
//line views/vresult/WSRow.html:79
	qs422016 := string(qb422016.B)
//line views/vresult/WSRow.html:79
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSRow.html:79
	return qs422016
//line views/vresult/WSRow.html:79
}

//line views/vresult/WSRow.html:81
func streamcell(qw422016 *qt422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, row []any, fields field.Fields, m *model.Model, f *field.Field, col any) {
//line views/vresult/WSRow.html:82
	if m.IsPK(f.Key, ws.PS.Logger) {
//line views/vresult/WSRow.html:84
		rowPK, err := model.GetStrings(fields, m.GetPK(ws.PS.Logger), row)
		if err != nil {
			panic(err)
		}
		link := append([]string{`v`}, rowPK...)

//line views/vresult/WSRow.html:89
		qw422016.N().S(`<a href="`)
//line views/vresult/WSRow.html:90
		qw422016.E().S(ws.RouteAct(act, 0, link...))
//line views/vresult/WSRow.html:90
		qw422016.N().S(`" class="pklink">`)
//line views/vresult/WSRow.html:91
		view.StreamAnyByType(qw422016, col, f.Type)
//line views/vresult/WSRow.html:91
		qw422016.N().S(`</a>`)
//line views/vresult/WSRow.html:93
	} else {
//line views/vresult/WSRow.html:94
		view.StreamAnyByType(qw422016, col, f.Type)
//line views/vresult/WSRow.html:95
	}
//line views/vresult/WSRow.html:96
}

//line views/vresult/WSRow.html:96
func writecell(qq422016 qtio422016.Writer, ws *cutil.WorkspaceRequest, act *action.Action, row []any, fields field.Fields, m *model.Model, f *field.Field, col any) {
//line views/vresult/WSRow.html:96
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vresult/WSRow.html:96
	streamcell(qw422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:96
	qt422016.ReleaseWriter(qw422016)
//line views/vresult/WSRow.html:96
}

//line views/vresult/WSRow.html:96
func cell(ws *cutil.WorkspaceRequest, act *action.Action, row []any, fields field.Fields, m *model.Model, f *field.Field, col any) string {
//line views/vresult/WSRow.html:96
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vresult/WSRow.html:96
	writecell(qb422016, ws, act, row, fields, m, f, col)
//line views/vresult/WSRow.html:96
	qs422016 := string(qb422016.B)
//line views/vresult/WSRow.html:96
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vresult/WSRow.html:96
	return qs422016
//line views/vresult/WSRow.html:96
}
