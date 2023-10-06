-- +goose Up
-- +goose StatementBegin
copy wines (id,name,year,strength,price,type,count) from '/data/wineResearch.csv' delimiter ',' csv;
copy users (id,login,password,fio,email,points,status ) from '/data/user.csv' delimiter ',' csv;
copy orders (id,id_user,total_price,status,is_points ) from '/data/order.csv' delimiter ',' csv;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table orders cascade;
truncate table users cascade;
truncate table wines cascade;
-- +goose StatementEnd
