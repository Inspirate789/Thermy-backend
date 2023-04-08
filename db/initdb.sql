create user postgres superuser;

-- Фиксированные таблицы БД
drop schema if exists public cascade;
create schema if not exists public;

-- Пользователи
-- drop table if exists public.users;
create table if not exists public.users(
    id int generated always as identity primary key,
    name text unique not null,
    password text not null,
    role text not null,
    registration_date timestamp not null
);

-- Контексты употребления
-- drop table if exists public.contexts;
create table if not exists public.contexts(
    id int generated always as identity primary key,
    registration_date timestamp not null,
    text text unique not null
);

-- Характеристики единиц языка
-- drop table if exists public.properties;
create table if not exists public.properties(
    id int generated always as identity primary key,
    property text unique not null
);



-- Роли БД
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
-- grant all on schema public to admin;
-- grant all on all tables in schema public to admin;



-- Хранимые процедуры, функции и триггеры

CREATE OR REPLACE PROCEDURE public.create_layer_tables(layer text)
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    EXECUTE format(
            '-- (Структурные) Модели слоя
            create table if not exists %I.models(
                id int generated always as identity primary key,
                name text unique not null
            );

            -- Элементы слоя (справочная таблица)
            -- drop table if exists %I.elements;
            create table if not exists %I.elements(
                id int generated always as identity primary key,
                name text unique not null
            );

            -- Единицы русского языка
            create table if not exists %I.units_ru(
                id int generated always as identity primary key,
                model_id int,
                foreign key (model_id) references %I.models(id),
                registration_date timestamp not null,
                text text unique not null
            );

            -- Единицы иностранного языка
            create table if not exists %I.units_en(
                id int generated always as identity primary key,
                model_id int,
                foreign key (model_id) references %I.models(id),
                registration_date timestamp not null,
                text text unique not null
            );

            -- Таблица-связка (модели слоя и элементы слоя)
            -- drop table if exists %I.models_and_elem;
            create table if not exists %I.models_and_elems(
                model_id int,
                foreign key (model_id) references %I.models(id),
                elem_id int,
                foreign key (elem_id) references %I.elements(id)
            );

            -- Таблица-связка (единицы русского языка и единицы иностранного языка)
            -- drop table if exists %I.units_ru_and_en;
            create table if not exists %I.units_ru_and_en(
                unit_ru_id int,
                foreign key (unit_ru_id) references %I.units_ru(id),
                unit_en_id int,
                foreign key (unit_en_id) references %I.units_en(id)
            );

            -- Таблица-связка (характеристики и единицы русского языка)
            -- drop table if exists %I.properties_and_units_ru;
            create table if not exists %I.properties_and_units_ru(
                property_id int,
                foreign key (property_id) references public.properties(id),
                unit_id int,
                foreign key (unit_id) references %I.units_ru(id)
            );

            -- Таблица-связка (характеристики и единицы иностранного языка)
            -- drop table if exists %I.properties_and_units_en;
            create table if not exists %I.properties_and_units_en(
                property_id int,
                foreign key (property_id) references public.properties(id),
                unit_id int,
                foreign key (unit_id) references %I.units_en(id)
            );

            -- Таблица-связка (контексты и единицы русского языка)
            -- drop table if exists %I.contexts_and_units_ru;
            create table if not exists %I.contexts_and_units_ru(
                context_id int,
                foreign key (context_id) references public.contexts(id),
                unit_id int,
                foreign key (unit_id) references %I.units_ru(id)
            );

            -- Таблица-связка (контексты и единицы иностранного языка)
            -- drop table if exists %I.contexts_and_units_en;
            create table if not exists %I.contexts_and_units_en(
                ctx_id int,
                foreign key (ctx_id) references public.contexts(id),
                unit_id int,
                foreign key (unit_id) references %I.units_en(id)
            );

            -- Таблица-связка (пользователи и единицы русского языка)
            -- drop table if exists %I.users_and_units_ru;
            create table if not exists %I.users_and_units_ru(
                user_id int,
                foreign key (user_id) references public.users(id),
                unit_id int,
                foreign key (unit_id) references %I.units_ru(id)
            );

            -- Таблица-связка (пользователи и единицы иностранного языка)
            -- drop table if exists %I.users_and_units_en;
            create table if not exists %I.users_and_units_en(
                user_id int,
                foreign key (user_id) references public.users(id),
                unit_id int,
                foreign key (unit_id) references %I.units_en(id)
            );',
            layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE public.grant_student_rights_to_layer_tables(layer text)
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    EXECUTE format(
            'grant select, insert, update on %I.units_ru to student;
            grant select, insert, update on %I.units_en to student;
            grant select, insert, update on %I.units_ru_and_en to student;
            grant select, insert, update, delete on %I.properties_and_units_ru to student;
            grant select, insert, update, delete on %I.properties_and_units_en to student;
            grant select, insert, update on %I.contexts_and_units_ru to student;
            grant select, insert, update on %I.contexts_and_units_en to student;
            grant select, insert, update on %I.users_and_units_ru to student;
            grant select, insert, update on %I.users_and_units_en to student;
            grant select on %I.models to student;
            grant select on %I.elements to student;
            grant select on %I.models_and_elems to student;',
            layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE public.grant_educator_rights_to_layer_tables(layer text)
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    EXECUTE format(
            'grant insert, update, delete on %I.models to educator;
            grant insert, update, delete on %I.elements to educator;
            grant insert, update, delete on %I.models_and_elems to educator;',
            layer_name, layer_name, layer_name
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE public.grant_admin_rights_to_layer_tables(layer text)
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    EXECUTE format(
            'grant usage, create on schema %I to admin;
            grant select, insert, update, delete on all tables in schema %I to admin;',
            layer_name, layer_name
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE public.grant_rights_to_layer_tables(layer text)
AS
$func$
BEGIN
    EXECUTE format(
            E'call public.grant_student_rights_to_layer_tables(\'%s\');
            call public.grant_educator_rights_to_layer_tables(\'%s\');
            call public.grant_admin_rights_to_layer_tables(\'%s\');',
            layer, layer, layer
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE public.create_layer(layer text)
AS
$func$
BEGIN
    EXECUTE format(
            E'set role educator;
            create schema if not exists %I;
            call public.create_layer_tables(\'%s\');
            call public.grant_rights_to_layer_tables(\'%s\');',
            (layer || '_layer'), layer, layer
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.insert_user(username text, password text, role text)
    RETURNS int
AS
$func$
DECLARE result int;
BEGIN
    EXECUTE format(
            E'set role admin;
            create user %I with password \'%s\';
            grant %I to %I;',
            username,
            password,
            role,
            username
        );
    insert into public.users(id, name, password, role, registration_date) overriding user value -- or overriding system value
    values (null, username, password, role, now()::timestamp)
    returning public.users.id into result;
    RETURN result;
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_all_models(layer text)
    RETURNS table (
                      id int,
                      name text
                  )
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    RETURN QUERY
    EXECUTE format(
            'select *
            from %I.models;',
            layer_name
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_all_model_elements(layer text)
    RETURNS table (
                      id int,
                      name text
                  )
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    RETURN QUERY
    EXECUTE format(
            'select * from %I.elements;',
            layer_name
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_contexts_id_by_unit(layer text, lang text, unit_id int)
    RETURNS table (
        id int
                  )
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    RETURN QUERY
    EXECUTE format(
            'select context_id
            from %I.%I
            where unit_id = $1;',
            layer_name,
            ('contexts_and_units_' || lang)
        )
        USING unit_id;
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_all_linked_units(layer text)
    RETURNS table (
                      unit_ru_id int,
                      unit_ru_model_id int,
                      unit_ru_registration_date timestamp,
                      unit_ru_text text,
                      unit_en_id int,
                      unit_en_model_id int,
                      unit_en_registration_date timestamp,
                      unit_en_text text
                  )
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    RETURN QUERY
    EXECUTE format(
            'select %I.units_ru.id as unit_ru_id,
               %I.units_ru.model_id as unit_ru_model_id,
               %I.units_ru.registration_date as unit_ru_registration_date,
               %I.units_ru.text as unit_ru_text,
               %I.units_en.id as unit_en_id,
               %I.units_en.model_id as unit_en_model_id,
               %I.units_en.registration_date as unit_en_registration_date,
               %I.units_en.text as unit_en_text
            from  %I.units_ru
                inner join %I.units_ru_and_en on %I.units_ru.id = unit_ru_id
                inner join %I.units_en on %I.units_en.id = unit_en_id;',
            layer_name, layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_linked_units_by_models_id(layer text, models_id int[])
    RETURNS table (
                      unit_ru_id int,
                      unit_ru_model_id int,
                      unit_ru_registration_date timestamp,
                      unit_ru_text text,
                      unit_en_id int,
                      unit_en_model_id int,
                      unit_en_registration_date timestamp,
                      unit_en_text text
                  )
AS
$func$
DECLARE layer_name text;
        id_string text;
BEGIN
    select (layer || '_layer') into layer_name;
    select format('%s', array_to_string(models_id, ',')) into id_string;
    RETURN QUERY
    EXECUTE format(
        'select %I.units_ru.id as unit_ru_id,
           %I.units_ru.model_id as unit_ru_model_id,
           %I.units_ru.registration_date as unit_ru_registration_date,
           %I.units_ru.text as unit_ru_text,
           %I.units_en.id as unit_en_id,
           %I.units_en.model_id as unit_en_model_id,
           %I.units_en.registration_date as unit_en_registration_date,
           %I.units_en.text as unit_en_text
        from  %I.units_ru
            inner join %I.units_ru_and_en on %I.units_ru.id = unit_ru_id and %I.units_ru.model_id = any(array[%s])
            inner join %I.units_en on %I.units_en.id = unit_en_id and %I.units_en.model_id = any(array[%s]);',
        layer_name, layer_name, layer_name, layer_name, layer_name, layer_name,
        layer_name, layer_name, layer_name, layer_name, layer_name, layer_name,
        id_string,
        layer_name, layer_name, layer_name,
        id_string
    );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_linked_units_by_properties_id(layer text, properties_id int[])
    RETURNS table (
                      unit_ru_id int,
                      unit_ru_model_id int,
                      unit_ru_registration_date timestamp,
                      unit_ru_text text,
                      unit_en_id int,
                      unit_en_model_id int,
                      unit_en_registration_date timestamp,
                      unit_en_text text
                  )
AS
$func$
DECLARE layer_name text;
        id_string text;
BEGIN
    select (layer || '_layer') into layer_name;
    select format('%s', array_to_string(properties_id, ',')) into id_string;
    RETURN QUERY
        EXECUTE format(
            E'select %I.units_ru.id as unit_ru_id,
               %I.units_ru.model_id as unit_ru_model_id,
               %I.units_ru.registration_date as unit_ru_registration_date,
               %I.units_ru.text as unit_ru_text,
               %I.units_en.id as unit_en_id,
               %I.units_en.model_id as unit_en_model_id,
               %I.units_en.registration_date as unit_en_registration_date,
               %I.units_en.text as unit_en_text
            from  %I.units_ru
                inner join %I.units_ru_and_en on %I.units_ru.id = unit_ru_id
                inner join %I.units_en on %I.units_en.id = unit_en_id
            where array[%s] <@ array(select * from public.select_properties_id_by_unit_id(\'%s\', \'ru\', %I.units_ru.id))
                or array[%s] <@ array(select * from public.select_properties_id_by_unit_id(\'%s\', \'en\', %I.units_en.id));',
            layer_name, layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name, layer_name, layer_name,
            layer_name,
            id_string, layer_name, layer_name,
            id_string, layer_name, layer_name
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_unlinked_units_by_lang(layer text, lang text)
    RETURNS table (
                      id int,
                      model_id int,
                      registration_date timestamp,
                      text text
                  )
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    RETURN QUERY
    EXECUTE format(
            'select id, model_id ,registration_date, text
            from %I.%I left join %I.units_ru_and_en on id <> %I;',
            layer_name,
            ('units_' || lang),
            layer_name,
            ('unit_' || lang || '_id')
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_unlinked_units_by_lang_and_models_id(layer text, lang text, models_id int[])
    RETURNS table (
                      id int,
                      model_id int,
                      registration_date timestamp,
                      text text
                  )
AS
$func$
DECLARE layer_name text;
        id_string text;
BEGIN
    select (layer || '_layer') into layer_name;
    select format('%s', array_to_string(models_id, ',')) into id_string;
    RETURN QUERY
    EXECUTE format(
            'select id, model_id ,registration_date, text
            from %I.%I left join %I.units_ru_and_en on id <> %I and %I.%I.model_id = any(array[%s]);',
            layer_name, ('units_' || lang),
            layer_name, ('unit_' || lang || '_id'),
            layer_name, ('units_' || lang),
            id_string
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_unlinked_units_by_lang_and_properties_id(layer text, lang text, properties_id int[])
    RETURNS table (
                      id int,
                      model_id int,
                      registration_date timestamp,
                      text text
                  )
AS
$func$
DECLARE layer_name text;
        id_string text;
BEGIN
    select (layer || '_layer') into layer_name;
    select format('%s', array_to_string(properties_id, ',')) into id_string;
    RETURN QUERY
        EXECUTE format(
            E'select id, model_id ,registration_date, text
            from %I.%I left join %I.units_ru_and_en on id <> %I
                and array[%s] <@ array(select * from public.select_properties_id_by_unit_id(\'%s\', \'%s\', %I.%I.id));',
            layer_name, ('units_' || lang),
            layer_name, ('unit_' || lang || '_id'),
            id_string, layer_name, lang,
            layer_name, ('units_' || lang)
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_properties_by_unit(layer text, lang text, unit_text text)
    RETURNS table (
                      id int,
                      property text
                  )
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    RETURN QUERY
    EXECUTE format(
            E'select public.properties.id, public.properties.property
             from public.properties
                inner join %I.%I on public.properties.id = property_id
                inner join %I.%I on %I.%I.id = unit_id and %I.%I.text = \'%s\';',
            layer_name,
            ('properties_and_units_' || lang),
            layer_name,
            ('units_' || lang),
            layer_name,
            ('units_' || lang),
            layer_name,
            ('units_' || lang),
            unit_text
        );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.select_properties_id_by_unit_id(layer text, lang text, unit_id int)
    RETURNS table (
        id int
                  )
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    RETURN QUERY
    EXECUTE format(
            'select property_id
            from %I.%I
            where unit_id = $1;',
            layer_name,
            ('properties_and_units_' || lang)
        )
        USING unit_id;
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.insert_properties(property_texts text[])
    RETURNS table (
        id int
                  )
AS
$func$
BEGIN
    RETURN QUERY
    insert into public.properties(id, property) overriding user value -- or overriding system value
    values (null, unnest(property_texts))
    returning public.properties.id;
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.insert_models(layer text, model_texts text[])
    RETURNS table (
        id int
                  )
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    RETURN QUERY
    EXECUTE format(
        E'insert into %I.models(id, name) overriding user value -- or overriding system value
        values (null, unnest(array[%s]))
        returning %I.models.id;',
        layer_name,
        format(
                E'\'%s\'',
                array_to_string(model_texts, E'\',\'')
            ),
        layer_name
    );
END
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.insert_model_elements(layer text, element_texts text[])
    RETURNS table (
        id int
                  )
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    RETURN QUERY
    EXECUTE format(
            E'insert into %I.elements(id, name) overriding user value -- or overriding system value
            values (null, unnest(array[%s]))
            returning %I.elements.id;',
            layer_name,
            format(
                E'\'%s\'',
                array_to_string(element_texts, E'\',\'')
                ),
            layer_name
        );
END
$func$ LANGUAGE plpgsql;
