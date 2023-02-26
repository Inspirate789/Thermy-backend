select password
from public.users
where name = $1
limit 1;