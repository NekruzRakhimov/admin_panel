
create table brands (
                        id bigserial primary key,
                        brand text not null,
                        discount_percent varchar not null,
                        contract_id integer references contracts(id)
);


SELECT c.id, b.contract_id, c.contract_parameters ->> 'contract_number' AS contract_number, b.discount_percent, b.brand
FROM brands b
JOIN contracts  c ON b.contract_id = c.id
WHERE requisites ->> 'bin' = '080240011774';



SELECT id FROM contracts WHERE requisites ->> 'bin' = '080240011774';

SELECT  discounts FROM contracts;

SELECT brand, discount_percent FROM  brands WHERE contract_id = 207;



SELECT c.id, b.contract_id, c.contract_parameters ->> 'contract_number' AS contract_number, b.discount_percent, b.brand FROM contracts c
		JOIN brands  b ON b.contract_id = c.id WHERE c.requisites ->> 'bin' = '0909090989889';

--SELECT id, brand, discount_percent FROM  brands =  where   contract_id

DELETE FROM contracts WHERE  id =248;



--WHERE cars_info -> 'sold' = 'true';


---Номер/Наименования договора
---Период
---Тип скидки/Бренд
---Процент Скидки если это фиксированная скидка
---Сумма скидки
---Итог




create table brands
(
    id               bigserial primary key,
    brand            text    not null,
    discount_percent varchar not null,
    contract_id      integer references contracts (id)
);


INSERT INTO brands(brand, discount_percent, contract_id) VALUES ($1, $2, $3);


SELECT id, contract_parameters ->> 'contract_number' AS contract_number, discounts ->> 'discount_amount' AS discount_amount, requisites ->> 'bin' AS bin FROM contracts
WHERE requisites ->> 'bin' = '860418401075';


SELECT discounts ->> 'discount_amount' AS discount_amount FROM contracts where  id = 260;


SELECT discounts FROM contracts WHERE id = 260;

SELECT id, discounts FROM contracts
WHERE requisites ->> 'bin' = '070340005201';

SELECT arr.position,arr.item_object
FROM purchases,
     jsonb_array_elements(items_purchased) with ordinality arr(item_object, position)
WHERE id=2;


SELECT discounts  FROM contracts WHERE requisites ->> 'bin' = '090909098988';
SELECT id, discounts::json as discount  FROM contracts WHERE requisites ->> 'bin' = '070340005201';
SELECT json_array_elements(discounts ->'periods') FROM  contracts;

WHERE requisites ->> 'bin' = '070340005201';


SELECT discounts ->> 'periods'  an  FROM contracts WHERE requisites ->> 'bin' = '070340005201';



SELECT id, manager, ext_contract_code FROM contracts WHERE requisites ->> 'bin' =  '100840008133';

Update contracts SET ext_contract_code = 'K0054437' WHERE requisites ->> 'bin' =  '100840008133';


DELETE FROM contracts WHERE  id = 264;