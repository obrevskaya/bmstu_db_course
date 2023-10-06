-- +goose Up
-- +goose StatementBegin
copy wines (name,year,strength,price,type,count) from '/var/lib/postgresql/data/data/wine.csv' delimiter ',' csv;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table wines cascade;
-- +goose StatementEnd
