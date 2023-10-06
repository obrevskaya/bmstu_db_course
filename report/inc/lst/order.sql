create table IF NOT EXISTS orders (
    id UUID default uuid_generate_v4() primary key,
    id_user UUID references users(id) on delete cascade,
    total_price int,
    status text not null ,
    is_points bool
);