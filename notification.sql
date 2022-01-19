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

SELECT bin, contract_number, contract_date, type, email FROM notification;


SELECT *
FROM notification;


INSERT Into notification(bin, contract_date, contract_number, type, email)
VALUES ('"0909090989889"', '"21.12.2021"', '96665556565', 'supply', '"bigboss@gmail.com"');


SELECT id
FROM notification
WHERE contract_number = '96665556565';

SELECT id
FROM notification
where contract_number = '96665556565';

drop table notification;

INSERT INTO notification(bin, contract_date, contract_type, email, status)
VALUES

SELECT id
FROM notification
where contract_number = '312312';