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



--SELECT contract_parameters ->> 'contract_number', id FROM contracts  WHERE requisites ->> 'bin' = '860418401075';

SELECT discounts,id FROM contracts  WHERE contract_parameters ->> 'contract_number' = '9090';



SELECT id, status,  contract_parameters ->> 'contract_number' AS contract_number FROM contracts WHERE status = 'в работе' AND requisites ->> 'bin' = '160140011654';
SELECT brand AS brand_name, discount_percent, contract_id FROM  brands WHERE contract_id = 381; -- у этого есть бренды, но они не совпадают
SELECT brand AS brand_name, discount_percent, contract_id FROM  brands WHERE contract_id = 383; -- вот что должно быть --
SELECT brand AS brand_name, discount_percent, contract_id FROM  brands WHERE contract_id = 386; -- его не должно быть

SELECT *FROM stored_reports WHERE bin LIKE '%86%';
SELECT *FROM stored_reports WHERE id LIKE '%25%';
