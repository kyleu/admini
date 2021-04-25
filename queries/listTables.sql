-- {% func ListTables() %}
select
  n.nspname as "schema",
  c.relname as "name",
  case c.relkind
    when 'r' then 'table'
    when 'v' then 'view'
    when 'm' then 'materialized view'
    when 'i' then 'index'
    when 'S' then 'sequence'
    when 's' then 'special'
    when 'f' then 'foreign table'
    when 'p' then 'partitioned table'
    when 'I' then 'partitioned index'
  end as "type",
  pg_catalog.pg_get_userbyid(c.relowner) as "owner"
from
  pg_catalog.pg_class c
  left join pg_catalog.pg_namespace n ON n.oid = c.relnamespace
where
  c.relkind IN ('r','p','v','m','S','f','')
  and n.nspname <> 'pg_catalog'
  and n.nspname <> 'information_schema'
  and n.nspname !~ '^pg_toast'
  and pg_catalog.pg_table_is_visible(c.oid)
order by
  "schema",
  "name"
;
-- {% endfunc %}
