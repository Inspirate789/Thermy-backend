select users.role
from public.users
where users.name = (select session_user);
