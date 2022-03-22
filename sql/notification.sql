CREATE TABLE notification
(
    id              bigserial primary key,
    bin             varchar,
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    contract_date   varchar,
    contract_number varchar,
    type            varchar,
    email           varchar,
    status          bool
);

SELECT bin,  created_at, contract_number, contract_date, type, email FROM notification;

SELECT  *FROM  contracts;


UPDATE notification set status = true WHERE  contract_number = '00000000001';

SELECT *
FROM notification;



SELECT *FROM  contracts where  supplier_company_manager -> 'email' = 'aziz.rahimov0001@gmail.com';

INSERT Into notification(bin, contract_date, contract_number, type, email)
VALUES ('"21312312312"', '"21.12.2021"', '96665556565', 'supply', '"tdsadasdsadas@gmail.com"');


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




SELECT jsonb_object_keys( '{"brand": "Mitsubishi", "sold": true}'::jsonb );



SELECT jsonb_extract_path('{"brand": "Honda", "sold": false}'::jsonb, 'brand');