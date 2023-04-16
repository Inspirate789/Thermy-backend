create role student;

grant select, insert, update on public.contexts to student;

grant select, insert, update, delete on public.properties to student;

grant select on information_schema.schemata to student;
grant select on information_schema.tables to student;

create role educator inherit;
grant student to educator;

grant delete on public.contexts to educator;

grant create on database :dbname to educator;
grant usage, create on schema public to educator;
grant references on all tables in schema public to educator;

create role admin with inherit CREATEDB CREATEROLE;
grant educator to admin;
alter database :dbname owner to admin;
grant usage, create on schema public to admin;
grant references, select, insert, update, delete on all tables in schema public to admin;
