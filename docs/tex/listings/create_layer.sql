CREATE OR REPLACE PROCEDURE public.create_layer_tables(layer text)
AS
$func$
DECLARE layer_name text;
BEGIN
    select (layer || '_layer') into layer_name;
    EXECUTE format(
            'create table if not exists %I.models(
                id int generated always as identity primary key,
                name text unique not null
            );

            create table if not exists %I.elements(
                id int generated always as identity primary key,
                name text unique not null
            );

            create table if not exists %I.units_ru(
                id int generated always as identity primary key,
                model_id int,
                foreign key (model_id) references %I.models(id),
                registration_date timestamp not null,
                text text unique not null
            );

            create table if not exists %I.units_en(
                id int generated always as identity primary key,
                model_id int,
                foreign key (model_id) references %I.models(id),
                registration_date timestamp not null,
                text text unique not null
            );

            create table if not exists %I.models_and_elems(
                model_id int,
                foreign key (model_id) references %I.models(id),
                elem_id int,
                foreign key (elem_id) references %I.elements(id),
                unique(model_id, elem_id)
            );

            create table if not exists %I.units_ru_and_en(
                unit_ru_id int,
                foreign key (unit_ru_id) references %I.units_ru(id),
                unit_en_id int,
                foreign key (unit_en_id) references %I.units_en(id),
                unique(unit_ru_id, unit_en_id)
            );

            create table if not exists %I.properties_and_units_ru(
                property_id int,
                foreign key (property_id) references public.properties(id),
                unit_id int,
                foreign key (unit_id) references %I.units_ru(id),
                unique(property_id, unit_id)
            );

            create table if not exists %I.properties_and_units_en(
                property_id int,
                foreign key (property_id) references public.properties(id),
                unit_id int,
                foreign key (unit_id) references %I.units_en(id),
                unique(property_id, unit_id)
            );

            create table if not exists %I.contexts_and_units_ru(
                context_id int,
                foreign key (context_id) references public.contexts(id),
                unit_id int,
                foreign key (unit_id) references %I.units_ru(id),
                unique(context_id, unit_id)
            );

            create table if not exists %I.contexts_and_units_en(
                context_id int,
                foreign key (context_id) references public.contexts(id),
                unit_id int,
                foreign key (unit_id) references %I.units_en(id),
                unique(context_id, unit_id)
            );

            create table if not exists %I.users_and_units_ru(
                user_id int,
                foreign key (user_id) references public.users(id),
                unit_id int,
                foreign key (unit_id) references %I.units_ru(id),
                unique(user_id, unit_id)
            );

            create table if not exists %I.users_and_units_en(
                user_id int,
                foreign key (user_id) references public.users(id),
                unit_id int,
                foreign key (unit_id) references %I.units_en(id),
                unique(user_id, unit_id)
            );',
            layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name,
            layer_name
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
            layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name,
            layer_name, layer_name, layer_name, layer_name
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
