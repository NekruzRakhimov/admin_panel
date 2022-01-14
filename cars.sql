CREATE TABLE cars(
                     id SERIAL PRIMARY KEY,
                     cars_info JSONB NOT NULL);





INSERT INTO cars(cars_info)
VALUES('{"brand": "Toyota", "color": ["red", "black"], "price": 285000, "sold": true}'),
      ('{"brand": "Honda", "color": ["blue", "pink"], "price": 25000, "sold": false}'),
      ('{"brand": "Mitsubishi", "color": ["black", "gray"], "price": 604520, "sold": true}');


SELECT cars_info -> 'brand' AS car_name FROM cars;


SELECT *FROM contracts;


SELECT contract_parameters -> 'contract_date' AS end_date FROM contracts;
SELECT contract_parameters -> 'prepayment' AS pre,   contract_parameters -> 'contract_date' AS date   FROM contracts;


SELECT contract_parameters -> 'prepayment' AS prepayment FROM contracts WHERE id =108;
SELECT contract_parameters -> 'prepayment' AS prepayment FROM contracts WHERE requisites -> 'bin' = '0909090989889'; -- не работает

SELECT * FROM cars WHERE cars_info -> 'sold' = 'true';

SELECT * FROM contracts WHERE data = '{"a":1}';


SELECT contract_date FROM contracts;

SELECT contract_parameters -> 'contract_date' AS data FROM contracts;

SELECT  cars_info -> 'brand' AS brand FROM cars;