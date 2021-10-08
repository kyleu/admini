-- {% func ListColumns(schema string) %}
select
  c.table_schema,
  c.table_name,
  c.column_name,
  c.ordinal_position,
  c.column_default,
  c.is_nullable,
  c.data_type,
  c.character_maximum_length,
  c.character_octet_length,
  c.numeric_precision,
  c.numeric_scale,
  c.datetime_precision
from
  information_schema.columns c
where
  c.table_schema not in ('information_schema', 'pg_catalog')
  {% if schema != "" %} and c.table_schema = '{%s schema %}'{% endif %}
order by
  c.table_schema, c.table_name, c.ordinal_position;
-- {% endfunc %}
