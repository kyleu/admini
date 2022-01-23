-- {% func ListIndexes(schema string) %}
select
  n.nspname as schema_name,
  t.relname as table_name,
  i.relname as index_name,
  case when idx.indisprimary then 1 else 0 end as pk,
  case when idx.indisunique then 1 else 0 end as u,
  array_to_string(array_agg(a.attname), ',') as column_names
from
  pg_class t,
  pg_class i,
  pg_index idx,
  pg_attribute a,
  pg_namespace n
where
  t.oid = idx.indrelid
  and i.oid = idx.indexrelid
  and a.attrelid = t.oid
  and n.oid = t.relnamespace
  and a.attnum = any(idx.indkey)
  and t.relkind = 'r'
  and n.nspname not in ('information_schema', 'pg_catalog')
  {% if schema != "" %}
  and n.nspname = '{%s schema %}'
  {% endif %}
group by
  n.nspname,
  t.relname,
  i.relname,
  idx.indisprimary,
  idx.indisunique
order by
  t.relname,
  i.relname;

-- {% endfunc %}
