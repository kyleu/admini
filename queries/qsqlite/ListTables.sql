-- {% func ListTables(schema string) %}
select
  m.type as "t",
  m.name as "n"
from
  sqlite_master m
where
  m.type in ('table', 'view')
order by
  n
;
-- {% endfunc %}
