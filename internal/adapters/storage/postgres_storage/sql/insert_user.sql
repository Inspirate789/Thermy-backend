insert into public.users(id, name, password, role, registration_date) overriding user value -- or overriding system value
values (null, '$1', '$2', '$3', now()::timestamp);

create user $1 with password '$2';
grant $3 to $1;
