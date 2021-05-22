-- {% func ListTypes(schema string) %}
select
  n.nspname as "schema",
  pg_catalog.format_type(t.oid, null) as "name",
  t.typname as "internal",
  case when t.typrelid != 0
         then CAST('tuple' as pg_catalog.text)
       when t.typlen < 0
         then CAST('var' as pg_catalog.text)
       else CAST(t.typlen as pg_catalog.text)
    end as "size",
  pg_catalog.array_to_string(
      ARRAY(
          select e.enumlabel
          from pg_catalog.pg_enum e
          where e.enumtypid = t.oid
          order by e.enumsortorder
        ),
      E'\n'
    ) as "elements",
  pg_catalog.pg_get_userbyid(t.typowner) as "owner",
  pg_catalog.array_to_string(t.typacl, E'\n') as "privileges",
  pg_catalog.obj_description(t.oid, 'pg_type') as "description"
from
  pg_catalog.pg_type t
  left join pg_catalog.pg_namespace n on n.oid = t.typnamespace
where
  (t.typrelid = 0 or (select c.relkind = 'c' from pg_catalog.pg_class c where c.oid = t.typrelid))
  and not EXISTS(select 1 from pg_catalog.pg_type el where el.oid = t.typelem and el.typarray = t.oid)
  and n.nspname <> 'pg_catalog'
  and n.nspname <> 'information_schema'
  and pg_catalog.pg_type_is_visible(t.oid)
order by
  1,
  2
;
-- {% endfunc %}
