insert into public.users(id, name, password, role, registration_date) overriding user value -- or overriding system value
values (null, :quoted_username, :quoted_password, 'admin', now()::timestamp);

create user :username with password :quoted_password;
grant admin to :username;
