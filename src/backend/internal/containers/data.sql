insert into users (id, login, password, fio, email, points, status) values
                ('4ec2b1fb-1102-4ce1-b0d5-2b7da62d0023', 'a', 'qwerty1234', 'f i o', '1@mail.ru', 100, 1),
                ('06100e99-e631-474f-85e4-e5e794925f67', 'b', 'qwerty1234', 'f i o', '2@mail.ru', 100, 0),
                ('fc268c8c-fba8-4bd8-ab51-5432cd0416bb', 'c', 'qwerty1234', 'f i o', '3@mail.ru', 100, 0),
                ('b871fb21-0572-4bec-b44e-16dac8b7ea22', 'd', 'qwerty1234', 'f i o', '4@mail.ru', 100, 0);


insert into wines (id, name, year, strength, price, type) values ('a8d6569f-4059-4161-8365-0041b3ac7b9c', '1', 1, 1, 1, '1');



insert into orders (id, id_user, total_price, status, is_points) values
               ('77dcd288-79b2-4655-9584-cc9b5329665d', '4ec2b1fb-1102-4ce1-b0d5-2b7da62d0023', 28519, 'placed', false),
               ('b8b87f28-3cbb-425c-ac4d-52015710d61b', '06100e99-e631-474f-85e4-e5e794925f67', 28519, 'placed', false),
               ('6cd1dcbc-02db-4980-8fba-eb56bb229fbe', 'fc268c8c-fba8-4bd8-ab51-5432cd0416bb', 28519, 'placed', false),
               ('1557a6be-1008-412a-88d9-1f06630d028c', 'b871fb21-0572-4bec-b44e-16dac8b7ea22', 28519, 'placed', false);

insert into order_elements (id, id_order, id_wine, count) values
       ('21bf7ace-965b-4679-86b8-93a89cba0094', '77dcd288-79b2-4655-9584-cc9b5329665d', 'a8d6569f-4059-4161-8365-0041b3ac7b9c', 1),
       ('eea07263-3444-40d0-adc4-345ad7728298', 'b8b87f28-3cbb-425c-ac4d-52015710d61b', 'a8d6569f-4059-4161-8365-0041b3ac7b9c', 1),
       ('c11add9b-207e-4b41-a964-2662ed3cae27', '6cd1dcbc-02db-4980-8fba-eb56bb229fbe', 'a8d6569f-4059-4161-8365-0041b3ac7b9c', 1),
       ('7e1053d5-f31f-4841-b3f9-d0e47849cfb3', '1557a6be-1008-412a-88d9-1f06630d028c', 'a8d6569f-4059-4161-8365-0041b3ac7b9c', 1);

insert into bills (id, id_order, price, status) values
       ('a52b8aea-d751-4933-91bb-691132e3b760', '77dcd288-79b2-4655-9584-cc9b5329665d', 28519, 'paid'),
       ('3f010aca-5008-4aa5-a1a3-a061a876783f', 'b8b87f28-3cbb-425c-ac4d-52015710d61b', 28519, 'paid'),
       ('9a9a818f-f8c8-4e24-8648-60c84c4fdeaa', '6cd1dcbc-02db-4980-8fba-eb56bb229fbe', 28519, 'paid'),
       ('a15c02f3-0fb8-4380-896e-b835f9542668', '1557a6be-1008-412a-88d9-1f06630d028c', 28519, 'paid');
