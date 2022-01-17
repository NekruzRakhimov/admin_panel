CREATE TABLE notification
(
    id              bigserial primary key,
    bin             varchar,
    created_at      timestamp default current_timestamp,
    contract_date   varchar,
    contract_number varchar,
    type            varchar,
    email           varchar,
    status          bool
);

SELECT *
FROM notification;


INSERT Into notification(bin, contract_date, contract_number, type, email)
VALUES ('"0909090989889"', '"21.12.2021"', '"2728"', 'supply', '"bigboss@gmail.com"');

drop table notification;

INSERT INTO notification(bin, contract_date, contract_type, email, status)
VALUES

SELECT id FROM notification where contract_number = '312312';