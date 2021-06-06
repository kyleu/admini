-- {% func ListForeignKeys(schema string) %}
select
  c.constraint_name,
  x.ordinal_position as ordinal,
  x.table_schema as schema_name,
  x.table_name,
  x.column_name,
  y.table_schema as foreign_schema_name,
  y.table_name as foreign_table_name,
  y.column_name as foreign_column_name
from
  information_schema.referential_constraints c
    join information_schema.key_column_usage x on x.constraint_name = c.constraint_name
    join information_schema.key_column_usage y on y.ordinal_position = x.position_in_unique_constraint and y.constraint_name = c.unique_constraint_name
{% if schema != "" %}where x.table_schema = '{%s schema %}'{% endif %}
order by
  c.constraint_name,
  x.ordinal_position;
-- {% endfunc %}
