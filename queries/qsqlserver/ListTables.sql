-- {% func ListTables(schema string) %}
select
  schema_name(schema_id) as schema_name,
  name as table_name
from
  sys.tables
order by schema_name, table_name;
-- {% endfunc %}
