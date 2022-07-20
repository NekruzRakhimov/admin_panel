create table brands
(
    id               bigserial primary key,
    --  contract_number  text not null,
    brand            text    not null,
    brand_code       text    not null,
    discount_percent varchar not null,
    contract_id      integer references contracts (id)
);


SELECT id, bin, contract_parameters ->> 'contract_number' AS contract_number
FROM contracts
WHERE status = 'в работе'
  AND requisites ->> 'bin' = '160140011654';

create table brands
(
    id               bigserial primary key,
    brand            text    not null,
    discount_percent varchar not null,
    contract_id      integer references contracts (id)
);


drop table brands;

SELECT c.id,
       b.contract_id,
       c.contract_parameters ->> 'contract_number' AS contract_number,
       b.discount_percent,
       b.brand
FROM brands b
         JOIN contracts c ON b.contract_id = c.id
WHERE contract_id = '390';
--WHERE requisites ->> 'bin' = '080240011774';


SELECT id
FROM contracts
WHERE requisites ->> 'bin' = '960340000029';

SELECT discounts
FROM contracts;

SELECT brand, discount_percent
FROM brands
WHERE contract_id = 207;



SELECT c.id,
       b.contract_id,
       c.contract_parameters ->> 'contract_number' AS contract_number,
       b.discount_percent,
       b.brand
FROM contracts c
         JOIN brands b ON b.contract_id = c.id
WHERE c.requisites ->> 'bin' = '160140011654';

--SELECT id, brand, discount_percent FROM  brands =  where   contract_id

DELETE
FROM contracts
WHERE id = 248;


SELECT *
FROM contracts
WHERE id = 294;



SELECT *
from contracts
WHERE req



--WHERE cars_info -> 'sold' = 'true';


---Номер/Наименования договора
---Период
---Тип скидки/Бренд
---Процент Скидки если это фиксированная скидка
---Сумма скидки
---Итог


INSERT INTO brands(brand, discount_percent, contract_id)
VALUES ($1, $2, $3);


SELECT id,
       contract_parameters ->> 'contract_number' AS contract_number,
       discounts ->> 'discount_amount'           AS discount_amount,
       requisites ->> 'bin'                      AS bin
FROM contracts
WHERE requisites ->> 'bin' = '860418401075';


SELECT discounts ->> 'discount_amount' AS discount_amount
FROM contracts
where id = 260;


SELECT discounts
FROM contracts
WHERE id = 260;

SELECT id, discounts
FROM contracts
WHERE requisites ->> 'bin' = '070340005201';

SELECT arr.position, arr.item_object
FROM purchases,
     jsonb_array_elements(items_purchased) with ordinality arr(item_object, position)
WHERE id = 2;


SELECT discounts
FROM contracts
WHERE requisites ->> 'bin' = '090909098988';
SELECT id, discounts::json as discount
FROM contracts
WHERE requisites ->> 'bin' = '070340005201';
SELECT json_array_elements(discounts -> 'periods')
FROM contracts;

WHERE requisites ->> 'bin' = '070340005201';


SELECT discounts ->> 'periods' an
FROM contracts
WHERE requisites ->> 'bin' = '070340005201';



SELECT id, manager, ext_contract_code
FROM contracts
WHERE requisites ->> 'bin' = '100840008133';


SELECT id,
       requisites ->> 'bin'                 as BIN,
       contract_parameters ->> 'start_date' as start,
       contract_parameters ->> 'end_date'   as end_date
from contracts


SELECT id,
       requisites ->> 'bin'                 as BIN,
       contract_parameters ->> 'start_date' as start,
       contract_parameters ->> 'end_date'   as end_date
from contracts
WHERE requisites ->> 'bin' = '860418401075';

Select contract_parameters ->> 'end_date' as end_date
FROM contracts;


insert into brands (brand, discount_percent, contract_id)
VALUES ('Старт', 5, 285);


SELECT *
From brands;



Update contracts
SET ext_contract_code = 'K0054437'
WHERE requisites ->> 'bin' = '100840008133';

Update contracts
SET status = 'заверщённый'
where id = ?


SELECT *
FROM brands
WHERE contract_id = 285;

SELECT *
FROM contracts
where id = 285;



DELETE
FROM contracts
WHERE id = 264;

ALTER TABLE graphics
    ADD COLUMN is_removed boolean default false;

ALTER TABLE dictionary_values
    DROP COLUMN has_file;



SELECT *
FROM suppliers
WHERE client_name ILIKE '%рад%';



SELECT id
     , status
     , requisites ->> 'beneficiary'              AS beneficiary
     , contract_parameters ->> 'contract_number' AS contract_number,type AS contract_type
     , created_at
     , updated_at
     , manager                                   AS author
     , contract_parameters ->> 'contract_amount' AS amount
FROM contracts
WHERE contract_number like  '%9595_2%' AND status =  'на согласовании'


SELECT *FROM segment WHERE id = 1;



UPDATE formed_graphics SET is_letter = true WHERE id = 29;


SELECT *FROM formed_graphics;



SELECT formulas_id  FROM formed_graphics;

SELECT fg.id,
       fg.is_letter,
       fg.formula_id,
       g.number          as graphic_name,
       fg.graphic_id     as graphic_id,
       g.supplier_name   as supplier,
       g.store_name      as store,
       fg.by_matrix,
       g.application_day as schedule,
       fg.product_availability_days,
       fg.dister_days,
       fg.store_days,
       fg.status         as status,
       to_char(fg.created_at::date, 'DD.MM.YYYY')
FROM formed_graphics fg
         JOIN graphics g ON fg.graphic_id = g.id
WHERE fg.id = 25;

SELECT segment_code, name_segment FROM  segment WHERE  beneficiary = 'Прима Дистрибьюшн ТОО';


select formula_id from formed_graphics
