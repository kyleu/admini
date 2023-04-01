-- {% func ListForeignKeys(schema string) %}
select
  obj.name as fk,
  sch.name as sch,
  tab1.name as tbl,
  col1.name as col,
  tab2.name as ref_tbl,
  col2.name as ref_col
from
  sys.foreign_key_columns fkc
    inner join sys.objects obj on obj.object_id = fkc.constraint_object_id
    inner join sys.tables tab1 on tab1.object_id = fkc.parent_object_id
    inner join sys.schemas sch on tab1.schema_id = sch.schema_id
    inner join sys.columns col1 on col1.column_id = parent_column_id and col1.object_id = tab1.object_id
    inner join sys.tables tab2 on tab2.object_id = fkc.referenced_object_id
    inner join sys.columns col2 on col2.column_id = referenced_column_id and col2.object_id = tab2.object_id
order by
  sch, tbl, col
-- {% endfunc %}
