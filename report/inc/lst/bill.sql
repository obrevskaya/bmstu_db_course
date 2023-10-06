create table IF NOT EXISTS bills ( 
    id UUID default uuid_generate_v4() primary key,
    id_order uuid references orders(id) on delete cascade,
    price int,
    status text not null
);