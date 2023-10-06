-- +goose Up
-- +goose StatementBegin
create or replace function check_element()
returns trigger as $$
declare
elem_count int;
wine_count int;
begin
select count(*)
into elem_count
from order_elements
where id_wine = new.id_wine and id_order = new.id_order;

if elem_count > 0 then
        raise exception
       'wine already in order.';
end if;


select w.count
into wine_count
from wines w
where w.id = new.id_wine;

if wine_count < new.count then
        raise exception
       	'not enough count of wine.';
end if;

return new;
end;
$$ language plpgsql;

create trigger element_trigger
    before insert on order_elements
    for each row
    execute function check_element();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop trigger if exists element_trigger on order_elements;
-- +goose StatementEnd
