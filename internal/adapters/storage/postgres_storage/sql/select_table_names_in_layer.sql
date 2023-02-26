select table_schema||'.'||table_name as full_rel_name -- select table_schema||'.'||table_name
from information_schema.tables
where table_schema = 'terms_layer';
