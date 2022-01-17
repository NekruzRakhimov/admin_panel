CREATE TABLE cars(
                     id SERIAL PRIMARY KEY,
                     cars_info JSONB NOT NULL);





INSERT INTO cars(cars_info)
VALUES('{"brand": "Toyota", "color": ["red", "black"], "price": 285000, "sold": true}'),
      ('{"brand": "Honda", "color": ["blue", "pink"], "price": 25000, "sold": false}'),
      ('{"brand": "Mitsubishi", "color": ["black", "gray"], "price": 604520, "sold": true}');


SELECT cars_info -> 'brand' AS car_name FROM cars;





SELECT contract_parameters -> 'contract_date' AS end_date FROM contracts;
SELECT contract_parameters -> 'prepayment' AS pre,   contract_parameters -> 'contract_date' AS date   FROM contracts;


SELECT contract_parameters -> 'prepayment' AS prepayment FROM contracts WHERE id =108;
SELECT contract_parameters -> 'prepayment' AS prepayment FROM contracts WHERE requisites -> 'bin' = '0909090989889'; -- не работает

SELECT * FROM cars WHERE cars_info -> 'sold' = 'true';

SELECT * FROM contracts WHERE data = '{"a":1}';


SELECT contract_date FROM contracts;

SELECT contract_parameters -> 'contract_date' AS data FROM contracts;

SELECT  cars_info -> 'brand' AS brand, cars_info -> 'color' AS color FROM cars;


SELECT requisites -> 'bin' AS bin, contract_parameters -> 'contract_date' AS end_date, contract_parameters -> 'contract_number'  AS   contract_number, type, supplier_company_manager -> 'email'  AS email
FROM contracts Where  status = 'в работе';

SELECT  contract_parameters -> ''