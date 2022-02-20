create table brands (
                        id bigserial primary key,
                        brand text not null,
                        discount_percent varchar not null,
                        contract_id integer references contracts(id)
);


SELECT c.id, b.contract_id, c.contract_parameters ->> 'contract_number' AS contract_number, b.discount_percent, b.brand
FROM brands b
JOIN contracts  c ON b.contract_id = c.id
WHERE requisites ->> 'bin' = '0909090989889';

--WHERE cars_info -> 'sold' = 'true';


---Номер/Наименования договора
---Период
---Тип скидки/Бренд
---Процент Скидки если это фиксированная скидка
---Сумма скидки
---Итог



SELECT *FROM  contracts WHERE  contract_parameters ->> 'contract_number' = '003';