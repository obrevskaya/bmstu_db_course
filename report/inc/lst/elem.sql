create table IF NOT EXISTS order_elements (
    id UUID default uuid_generate_v4() primary key,
    id_order uuid references orders(id) on delete cascade,
    id_wine uuid references wines(id) on delete cascade,
    count int
);