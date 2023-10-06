-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users (
                       id UUID default uuid_generate_v4() primary key,
                       login text not null,
                       password text,
                       fio text,
                       email text not null check ( email like '%@%.%' ),
                       points int default 0,
                       status int default 0
);

create table wines (
                       id UUID default uuid_generate_v4() primary key,
                       name text not null,
                       year int not null,
                       strength int,
                       price int not null,
                       type text,
                       count int default 10
);

create table orders (
                        id UUID default uuid_generate_v4() primary key,
                        id_user UUID references users(id) on delete cascade,
                        total_price int,
                        status text not null ,
                        is_points bool
);

create table order_elements (
                                id UUID default uuid_generate_v4() primary key,
                                id_order uuid references orders(id) on delete cascade,
                                id_wine uuid references wines(id) on delete cascade,
                                count int
);

create table bills (
                       id UUID default uuid_generate_v4() primary key,
                       id_order uuid references orders(id) on delete cascade,
                       price int,
                       status text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists bills;
drop table if exists order_elements;
drop table if exists orders;
drop table if exists wines;
drop table if exists users;

drop extension if exists "uuid-ossp";
-- +goose StatementEnd
