create table IF NOT EXISTS users (
    id UUID default uuid_generate_v4() primary key,
    login text not null,
    password text not null,
    fio text,
    email text not null check ( email like '%@%.%' ),
    points int default 0,
    status int default 0
);