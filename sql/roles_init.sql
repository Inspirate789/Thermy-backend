create role student;

grant select on public.contexts to student;
grant insert on public.contexts to student;
grant update on public.contexts to student;

grant select on public.properties to student;
grant insert on public.properties to student;
grant update on public.properties to student;

grant select on information_schema.schemata to student;
grant select on information_schema.tables to student;



create role educator;

grant student to educator;

grant create on database postgres to educator;



create role admin with CREATEDB CREATEROLE;






