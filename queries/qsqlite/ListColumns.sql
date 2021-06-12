-- {% func ListColumns(schema string) %}
select
  m.name as "xn",
  p.cid as "i",
  p.name as "n",
  p.type as "t",
  p.pk as "pk",
  p.dflt_value as "dv",
  p."notnull" as "nn"
from
  sqlite_master m
  left outer join pragma_table_info(m.name) p on m.name <> p.name
where
  m.type in ('table', 'view')
order by
  xn, i
;

-- {% endfunc %}
