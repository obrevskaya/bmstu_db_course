create role guest with login;
grant select on wines to guest;
grant select, insert on users to guest;

create role user with login password 'user_password';
grant guest to user;
grant select, insert, update, delete on orders to user;
grant select, insert, update, delete on order_elements to user;
grant insert, select on bills to user;
grant insert, delete, select on user_wines to user;

create role administrator login superuser;