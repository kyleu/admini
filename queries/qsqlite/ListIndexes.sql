-- {% func ListIndexes(schema string) %}
select
  m.name as "n",
  m.tbl_name as "xn",
  p.name as "cn",
  p.seqno as "ci"
from
  sqlite_master m
  left outer join pragma_index_info(m.name) p on m.name <> p.name
where
  m.type='index'
order by
  m.name, p.seqno
;
-- {% endfunc %}
