select id, text
from public.contexts
where id = any($1);