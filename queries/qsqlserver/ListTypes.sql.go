// Code generated by qtc from "ListTypes.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/qsqlserver/ListTypes.sql:1
package qsqlserver

//line queries/qsqlserver/ListTypes.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/qsqlserver/ListTypes.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/qsqlserver/ListTypes.sql:1
func StreamListTypes(qw422016 *qt422016.Writer, schema string) {
//line queries/qsqlserver/ListTypes.sql:1
	qw422016.N().S(`
-- `)
//line queries/qsqlserver/ListTypes.sql:2
}

//line queries/qsqlserver/ListTypes.sql:2
func WriteListTypes(qq422016 qtio422016.Writer, schema string) {
//line queries/qsqlserver/ListTypes.sql:2
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/qsqlserver/ListTypes.sql:2
	StreamListTypes(qw422016, schema)
//line queries/qsqlserver/ListTypes.sql:2
	qt422016.ReleaseWriter(qw422016)
//line queries/qsqlserver/ListTypes.sql:2
}

//line queries/qsqlserver/ListTypes.sql:2
func ListTypes(schema string) string {
//line queries/qsqlserver/ListTypes.sql:2
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/qsqlserver/ListTypes.sql:2
	WriteListTypes(qb422016, schema)
//line queries/qsqlserver/ListTypes.sql:2
	qs422016 := string(qb422016.B)
//line queries/qsqlserver/ListTypes.sql:2
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/qsqlserver/ListTypes.sql:2
	return qs422016
//line queries/qsqlserver/ListTypes.sql:2
}