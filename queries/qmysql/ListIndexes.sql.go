// Code generated by qtc from "ListIndexes.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/qmysql/ListIndexes.sql:1
package qmysql

//line queries/qmysql/ListIndexes.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/qmysql/ListIndexes.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/qmysql/ListIndexes.sql:1
func StreamListIndexes(qw422016 *qt422016.Writer, schema string) {
//line queries/qmysql/ListIndexes.sql:1
	qw422016.N().S(`
select
  table_schema as schema_name,
  table_name as table_name,
  index_name as index_name,
  non_unique as non_unique,
  group_concat(column_name order by seq_in_index) as column_names
from
  information_schema.statistics`)
//line queries/qmysql/ListIndexes.sql:9
	if schema != "" {
//line queries/qmysql/ListIndexes.sql:9
		qw422016.N().S(`
where
  table_schema = '`)
//line queries/qmysql/ListIndexes.sql:11
		qw422016.E().S(schema)
//line queries/qmysql/ListIndexes.sql:11
		qw422016.N().S(`'`)
//line queries/qmysql/ListIndexes.sql:11
	}
//line queries/qmysql/ListIndexes.sql:11
	qw422016.N().S(`
group by
  1, 2, 3, 4
order by
  1, 2, 3
;
-- `)
//line queries/qmysql/ListIndexes.sql:17
}

//line queries/qmysql/ListIndexes.sql:17
func WriteListIndexes(qq422016 qtio422016.Writer, schema string) {
//line queries/qmysql/ListIndexes.sql:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/qmysql/ListIndexes.sql:17
	StreamListIndexes(qw422016, schema)
//line queries/qmysql/ListIndexes.sql:17
	qt422016.ReleaseWriter(qw422016)
//line queries/qmysql/ListIndexes.sql:17
}

//line queries/qmysql/ListIndexes.sql:17
func ListIndexes(schema string) string {
//line queries/qmysql/ListIndexes.sql:17
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/qmysql/ListIndexes.sql:17
	WriteListIndexes(qb422016, schema)
//line queries/qmysql/ListIndexes.sql:17
	qs422016 := string(qb422016.B)
//line queries/qmysql/ListIndexes.sql:17
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/qmysql/ListIndexes.sql:17
	return qs422016
//line queries/qmysql/ListIndexes.sql:17
}