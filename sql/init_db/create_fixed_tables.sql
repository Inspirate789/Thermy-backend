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
