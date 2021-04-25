// Code generated by qtc from "listColumns.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/listColumns.sql:1
package queries

//line queries/listColumns.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/listColumns.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/listColumns.sql:1
func StreamListColumns(qw422016 *qt422016.Writer) {
//line queries/listColumns.sql:1
	qw422016.N().S(`
select
  c.table_schema,
  c.table_name,
  c.column_name,
  c.ordinal_position,
  c.column_default,
  c.is_nullable,
  c.data_type,
  e.data_type as array_type,
  c.character_maximum_length,
  c.character_octet_length,
  c.numeric_precision,
  c.numeric_precision_radix,
  c.numeric_scale,
  c.datetime_precision,
  c.interval_type,
  c.domain_schema,
  c.domain_name,
  c.udt_schema,
  c.udt_name,
  c.dtd_identifier,
  c.is_updatable
from
  information_schema.columns c
  left join information_schema.element_types e on (
    (c.table_catalog, c.table_schema, c.table_name, 'TABLE', c.dtd_identifier) = (e.object_catalog, e.object_schema, e.object_name, e.object_type, e.collection_type_identifier)
  )
where
  table_schema = 'public'
order by
  c.table_schema, c.table_name, c.ordinal_position;
-- `)
//line queries/listColumns.sql:33
}

//line queries/listColumns.sql:33
func WriteListColumns(qq422016 qtio422016.Writer) {
//line queries/listColumns.sql:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/listColumns.sql:33
	StreamListColumns(qw422016)
//line queries/listColumns.sql:33
	qt422016.ReleaseWriter(qw422016)
//line queries/listColumns.sql:33
}

//line queries/listColumns.sql:33
func ListColumns() string {
//line queries/listColumns.sql:33
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/listColumns.sql:33
	WriteListColumns(qb422016)
//line queries/listColumns.sql:33
	qs422016 := string(qb422016.B)
//line queries/listColumns.sql:33
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/listColumns.sql:33
	return qs422016
//line queries/listColumns.sql:33
}
