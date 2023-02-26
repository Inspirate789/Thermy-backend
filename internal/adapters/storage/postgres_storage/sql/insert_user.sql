insert into public.users overriding user value -- or overriding system value
select null, $1, $2, $3, now()::timestamp;
-- values (null, $1, $2, $3, (SELECT now()::timestamp))

create user $1 with password $2;
grant $3 to $1;
