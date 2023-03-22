-- Переменные таблицы БД (под каждый слой своя схема)
drop schema if exists $1_layer cascade;
create schema if not exists $1_layer;

-- (Структурные) Модели слоя
create table if not exists $1_layer.models(
    id int generated always as identity primary key,
    name text unique not null
);

-- Элементы слоя (справочная таблица)
-- drop table if exists $1_layer.elements;
create table if not exists $1_layer.elements(
    id int generated always as identity primary key,
    name text unique not null
);

-- Единицы русского языка
create table if not exists $1_layer.units_ru(
    id int generated always as identity primary key,
    model_id int,
    foreign key (model_id) references $1_layer.models(id),
    registration_date timestamp not null,
    text text unique not null
);

-- Единицы иностранного языка
create table if not exists $1_layer.units_en(
    id int generated always as identity primary key,
    model_id int,
    foreign key (model_id) references $1_layer.models(id),
    registration_date timestamp not null,
    text text unique not null
);

-- Таблица-связка (модели слоя и элементы слоя)
-- drop table if exists $1_layer.models_and_elem;
create table if not exists $1_layer.models_and_elems(
    model_id int,
    foreign key (model_id) references $1_layer.models(id),
    elem_id int,
    foreign key (elem_id) references $1_layer.elements(id)
);

-- Таблица-связка (единицы русского языка и единицы иностранного языка)
-- drop table if exists $1_layer.units_ru_and_en;
create table if not exists $1_layer.units_ru_and_en(
    unit_ru_id int,
    foreign key (unit_ru_id) references $1_layer.units_ru(id),
    unit_en_id int,
    foreign key (unit_en_id) references $1_layer.units_en(id)
);

-- Таблица-связка (характеристики и единицы русского языка)
-- drop table if exists $1_layer.properties_and_units_ru;
create table if not exists $1_layer.properties_and_units_ru(
    property_id int,
    foreign key (property_id) references public.properties(id),
    unit_id int,
    foreign key (unit_id) references $1_layer.units_ru(id)
);

-- Таблица-связка (характеристики и единицы иностранного языка)
-- drop table if exists $1_layer.properties_and_units_en;
create table if not exists $1_layer.properties_and_units_en(
    property_id int,
    foreign key (property_id) references public.properties(id),
    unit_id int,
    foreign key (unit_id) references $1_layer.units_en(id)
);

-- Таблица-связка (контексты и единицы русского языка)
-- drop table if exists $1_layer.contexts_and_units_ru;
create table if not exists $1_layer.contexts_and_units_ru(
    context_id int,
    foreign key (context_id) references public.contexts(id),
    unit_id int,
    foreign key (unit_id) references $1_layer.units_ru(id)
);

-- Таблица-связка (контексты и единицы иностранного языка)
-- drop table if exists $1_layer.contexts_and_units_en;
create table if not exists $1_layer.contexts_and_units_en(
    ctx_id int,
    foreign key (ctx_id) references public.contexts(id),
    unit_id int,
    foreign key (unit_id) references $1_layer.units_en(id)
);

-- Таблица-связка (пользователи и единицы русского языка)
-- drop table if exists $1_layer.users_and_units_ru;
create table if not exists $1_layer.users_and_units_ru(
    user_id int,
    foreign key (user_id) references public.users(id),
    unit_id int,
    foreign key (unit_id) references $1_layer.units_ru(id)
);

-- Таблица-связка (пользователи и единицы иностранного языка)
-- drop table if exists $1_layer.users_and_units_en;
create table if not exists $1_layer.users_and_units_en(
    user_id int,
    foreign key (user_id) references public.users(id),
    unit_id int,
    foreign key (unit_id) references $1_layer.units_en(id)
);

grant select on $1_layer.units_ru to student;
grant insert on $1_layer.units_ru to student;
grant update on $1_layer.units_ru to student;

grant select on $1_layer.units_en to student;
grant insert on $1_layer.units_en to student;
grant update on $1_layer.units_en to student;

grant select on $1_layer.units_ru_and_en to student;
grant insert on $1_layer.units_ru_and_en to student;
grant update on $1_layer.units_ru_and_en to student;

grant select on $1_layer.properties_and_units_ru to student;
grant insert on $1_layer.properties_and_units_ru to student;
grant update on $1_layer.properties_and_units_ru to student;

grant select on $1_layer.properties_and_units_en to student;
grant insert on $1_layer.properties_and_units_en to student;
grant update on $1_layer.properties_and_units_en to student;

grant select on $1_layer.contexts_and_units_ru to student;
grant insert on $1_layer.contexts_and_units_ru to student;
grant update on $1_layer.contexts_and_units_ru to student;

grant select on $1_layer.contexts_and_units_en to student;
grant insert on $1_layer.contexts_and_units_en to student;
grant update on $1_layer.contexts_and_units_en to student;

grant select on $1_layer.users_and_units_ru to student;
grant insert on $1_layer.users_and_units_ru to student;
grant update on $1_layer.users_and_units_ru to student;

grant select on $1_layer.users_and_units_en to student;
grant insert on $1_layer.users_and_units_en to student;
grant update on $1_layer.users_and_units_en to student;

grant select on $1_layer.models to student;

grant select on $1_layer.elements to student;

grant select on $1_layer.models_and_elems to student;

grant insert on $1_layer.models to educator;
grant update on $1_layer.models to educator;

grant insert on $1_layer.elements to educator;
grant update on $1_layer.elements to educator;

grant insert on $1_layer.models_and_elems to educator;
grant update on $1_layer.models_and_elems to educator;



grant all on schema $1_layer to admin;
grant all on all tables in schema $1_layer to admin;
