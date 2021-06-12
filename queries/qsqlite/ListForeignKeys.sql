-- {% func ListForeignKeys(schema string) %}
select
  p.seq as "idx",
  m.tbl_name as "src",
  p."from" as "src_col",
  p."table" as "tgt",
  p."to" as "tgt_col"
from
  sqlite_master m
  left outer join pragma_foreign_key_list(m.name) p on m.name <> p.[table]
where
  m.type in ('table', 'view')
  and "src_col" is not null
;
-- {% endfunc %}
