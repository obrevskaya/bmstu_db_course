create table IF NOT EXISTS user_wines(
    id_user uuid references users(id) on delete cascade,
    id_wine uuid references wines(id) on delete cascade,
    primary key (id_user, id_wine)
);