// Code generated by qtc from "ListTables.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/qsqlserver/ListTables.sql:1
package qsqlserver

//line queries/qsqlserver/ListTables.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/qsqlserver/ListTables.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/qsqlserver/ListTables.sql:1
func StreamListTables(qw422016 *qt422016.Writer, schema string) {
//line queries/qsqlserver/ListTables.sql:1
	qw422016.N().S(`
select
  schema_name(schema_id) as schema_name,
  name as table_name
from
  sys.tables
order by schema_name, table_name;
-- `)
//line queries/qsqlserver/ListTables.sql:8
}

//line queries/qsqlserver/ListTables.sql:8
func WriteListTables(qq422016 qtio422016.Writer, schema string) {
//line queries/qsqlserver/ListTables.sql:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/qsqlserver/ListTables.sql:8
	StreamListTables(qw422016, schema)
//line queries/qsqlserver/ListTables.sql:8
	qt422016.ReleaseWriter(qw422016)
//line queries/qsqlserver/ListTables.sql:8
}

//line queries/qsqlserver/ListTables.sql:8
func ListTables(schema string) string {
//line queries/qsqlserver/ListTables.sql:8
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/qsqlserver/ListTables.sql:8
	WriteListTables(qb422016, schema)
//line queries/qsqlserver/ListTables.sql:8
	qs422016 := string(qb422016.B)
//line queries/qsqlserver/ListTables.sql:8
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/qsqlserver/ListTables.sql:8
	return qs422016
//line queries/qsqlserver/ListTables.sql:8
}
