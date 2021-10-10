-- {% func ListIndexes(schema string) %}
select
  table_schema as schema_name,
  table_name as table_name,
  index_name as index_name,
  non_unique as non_unique,
  group_concat(column_name order by seq_in_index) as column_names
from
  information_schema.statistics{% if schema != "" %}
where
  table_schema = '{%s schema %}'{% endif %}
group by
  1, 2, 3, 4
order by
  1, 2, 3
;
-- {% endfunc %}
