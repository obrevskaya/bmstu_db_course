create table IF NOT EXISTS wines (
        id UUID default uuid_generate_v4() primary key,
        name text not null,
        year int not null,
        strength int,
        price int not null,
        type text,
        count int default 10
);