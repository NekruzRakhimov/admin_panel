SELECT contract_parameters ->> 'extend_date' AS extend_date
FROM contracts
WHERE id = 163
  AND status = 'в работе';


CREATE TABLE cars
(
    id        SERIAL PRIMARY KEY,
    cars_info JSONB NOT NULL
);
SELECT requisites ->> 'beneficiary'              AS beneficiary,
       contract_parameters ->> 'contract_number' AS contract_number,
       type,
       created_at,
       updated_at,
       manager,
       contract_parameters ->> 'contract_amount' AS price
FROM contracts
WHERE contract_parameters ->> 'contract_number' like '00000000000004%'
SELECT requisites ->> 'beneficiary'              AS beneficiary,
       contract_parameters ->> 'contract_number' AS contract_number,
       type,
       created_at,
       updated_at,
       manager,
       contract_parameters ->> 'contract_amount' AS price
FROM contracts
WHERE manager like '%Азиз%';

SELECT requisites ->> 'beneficiary'              AS beneficiary,
       contract_parameters ->> 'contract_number'
                                                 AS contract_number,
       type,
       created_at,
       updated_at,
       manager,
       contract_parameters ->> 'contract_amount' AS price
FROM contracts
WHERE contract_parameters ->> 'contract_number' like '%00000000000004%';



SELECT *
FROM contracts;

SELECT contract_parameters ->> 'is_extend_contract' AS is_extend_contract
FROM contracts
WHERE id = 159;


SELECT *
FROM contracts
WHERE id = 163
  AND status = 'в работе';
SELECT *
FROM contracts
WHERE id = 172
  AND status = 'в работе';
SELECT *
FROM cars
WHERE cars_info -> 'sold' = 'true';


SELECT *
FROM contracts
WHERE contract_parameters ->> 'contract_number' = '1643882282349';
UPDATE test
SET data = data #- '{tags,-1}';
Update contracts
SET contract_parameters = contract_parameters #- 'contract_number,-1';


UPDATE test
SET json = json || '{
  "c": 6
}'::jsonb
WHERE id = 5;

UPDATE contracts
SET contract_parameters = contract_parameters || '{
  "end_date": "11.10.2023"
}'::jsonb
WHERE id = 166
  AND status = 'в работе';



UPDATE contracts
SET contract_parameters = jsonb_set("contract_parameters", '{"end_date"}', to_jsonb('20.10.2022'::text), true)
WHERE id = 166;

UPDATE contracts
SET contract_parameters = jsonb_set("contract_parameters", '{"end_date"}', to_jsonb('11.11.2039'::text), true),
    is_individ          = true
WHERE id = 163
  AND status = 'в работе';
UPDATE contracts
SET contract_parameters = jsonb_set("contract_parameters", '{"is_extend_contract"}', to_jsonb(true::bool), true),
    is_individ          = true
WHERE id = 172
  AND status = 'в работе';



update contracts
SET contract_parameters -> 'end_date' = '03.02.2023'
WHERE contract_parameters ->> 'contract_number' = '1643882282349';

SELECT cars_info ->> 'brand' AS car_name
FROM cars;
INSERT INTO cars(cars_info)
VALUES ('{
  "brand": "Toyota",
  "color": [
    "red",
    "black"
  ],
  "price": 285000,
  "sold": true
}'),
       ('{
         "brand": "Honda",
         "color": [
           "blue",
           "pink"
         ],
         "price": 25000,
         "sold": false
       }'),
       ('{
         "brand": "Mitsubishi",
         "color": [
           "black",
           "gray"
         ],
         "price": 604520,
         "sold": true
       }');



SELECT cars_info value ->> 'brand' AS brand
FROM cars, jsonb_array_elements(cars.cars_info);

SELECT DISTINCT value -> 'Tag' AS tag
FROM Documents, jsonb_array_elements(Documents.Tags);



SELECT contract_parameters -> 'contract_date' AS end_date
FROM contracts;
SELECT contract_parameters -> 'prepayment' AS pre, contract_parameters -> 'contract_date' AS date
FROM contracts;


SELECT contract_parameters -> 'prepayment' AS prepayment
FROM contracts
WHERE id = 108;
SELECT contract_parameters -> 'prepayment' AS prepayment
FROM contracts
WHERE requisites -> 'bin' = '0909090989889'; -- не работает


SELECT id,
       status,
       requisites ->> 'beneficiary'              AS beneficiary,
       contract_parameters ->> 'contract_number' AS contract_number,
       type                                      AS contract_type,
       created_at,
       updated_at,
       manager                                   AS author,
       contract_parameters ->> 'contract_amount' AS price
FROM contracts
WHERE manager like '%Иван%'
  AND status = '';


SELECT id, manager
FROM contracts
WHERE manager like '%Иван%';



SELECT supplier_company_manager -> 'email' AS email
From contracts;
SELECT *
FROM contracts
where supplier_company_manager ->> 'email' = 'aziz.rahimov0001@gmail.com';

SELECT *
FROM contracts
WHERE data = '{"a":1}';


SELECT contract_date
FROM contracts;

SELECT contract_parameters -> 'contract_date' AS data
FROM contracts;

SELECT cars_info -> 'brand' AS brand, cars_info -> 'color' AS color
FROM cars;


SELECT requisites -> 'bin'                      AS bin,
       contract_parameters -> 'contract_date'   AS end_date,
       contract_parameters -> 'contract_number' AS contract_number,
       type,
       supplier_company_manager -> 'email'      AS email,
       status
FROM contracts
Where status = 'в работе';

SELECT contract_parameters -> ''


SELECT *
FROM contracts;

SELECT requisites -> 'contractor_name' AS contractor_name
FROM contracts;



SELECT *
FROM contracts
WHERE id not in (select prev_contract_id from contracts)
  AND is_active = true
  AND status = 'в работе';

SELECT id,
       requisites ->> 'beneficiary'              AS beneficiary,
       contract_parameters ->> 'contract_number' AS contract_number,
       type                                      AS contract_type,
       created_at,
       updated_at,
       manager                                   AS author,
       contract_parameters ->> 'contract_amount' AS amount
FROM contracts
WHERE manager like '%Алиса%';


SELECT contract_parameters ->> 'end_date' AS end_date
FROM contracts;

create table brands
(
    id               bigserial primary key,
    brand            text    not null,
    discount_percent varchar not null,
    contract_id      integer references contracts (id)
);

drop table brands;



SELECT *
FROM "contracts"
WHERE (requisites ->> 'client_code' = '000002149'
    AND contract_parameters ->> 'start_date' <= '01.04.2021' AND contract_parameters ->> 'end_date' >= '30.12.2022' AND
       status = 'в работе');


SELECT *
FROM "contracts"
WHERE (requisites ->> 'client_code' = '000002149'
    AND (contract_parameters ->> 'start_date' <= '01.04.2021'::date OR
         contract_parameters ->> 'end_date' >= '30.12.2023'::date) AND
       status = 'в работе');


CREATE TABLE segment
(
    id           bigserial primary key,
    segment_code varchar unique,
    name_segment varchar,
    beneficiary  varchar,
    bin          varchar,
    client_code  varchar,
    email        varchar,
    for_market   bool default false,
    product      jsonb,
    region       jsonb

);

drop table segment;





UPDATE segment set