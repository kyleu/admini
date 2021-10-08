-- {% func ListTables(schema string) %}
select
  table_name, table_schema, table_rows, table_comment
from
  information_schema.tables
where
  table_schema = 'Chinook';
-- {% endfunc %}
