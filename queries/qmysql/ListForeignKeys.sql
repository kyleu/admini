-- {% func ListForeignKeys(schema string) %}
select
  x.constraint_name as constraint_name,
  x.ordinal_position as ordinal,
  x.table_schema as schema_name,
  x.table_name as table_name,
  x.column_name as column_name,
  x.table_schema as foreign_schema_name,
  x.referenced_table_name as foreign_table_name,
  x.referenced_column_name as foreign_column_name
from
  information_schema.key_column_usage x
where
  x.referenced_table_name is not null{% if schema != "" %}
  and x.table_schema = '{%s schema %}'{% endif %}
order by
  x.constraint_name,
  x.ordinal_position;
-- {% endfunc %}
