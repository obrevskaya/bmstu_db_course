-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table IF NOT EXISTS users (
    id UUID default uuid_generate_v4() primary key,
    login text not null,
    password text not null,
    fio text,
    email text not null check ( email like '%@%.%' ),
    points int default 0,
    status int default 0
);

create table IF NOT EXISTS wines (
        id UUID default uuid_generate_v4() primary key,
        name text not null,
        year int not null,
        strength int,
        price int not null,
        type text,
        count int default 10
);

create table IF NOT EXISTS orders (
    id UUID default uuid_generate_v4() primary key,
    id_user UUID references users(id) on delete cascade,
    total_price int,
    status text not null ,
    is_points bool
);

create table IF NOT EXISTS order_elements (
    id UUID default uuid_generate_v4() primary key,
    id_order uuid references orders(id) on delete cascade,
    id_wine uuid references wines(id) on delete cascade,
    count int
);

create table IF NOT EXISTS bills ( 
    id UUID default uuid_generate_v4() primary key,
    id_order uuid references orders(id) on delete cascade,
    price int,
    status text not null
);

create table IF NOT EXISTS user_wines(
    id_user uuid references users(id) on delete cascade,
    id_wine uuid references wines(id) on delete cascade,
    primary key (id_user, id_wine)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists user_wines cascade;
drop table if exists bills;
drop table if exists order_elements;
drop table if exists orders;
drop table if exists wines;
drop table if exists users;


drop extension if exists "uuid-ossp";
-- +goose StatementEnd
