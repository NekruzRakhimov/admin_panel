create table brands
(
    id               bigserial primary key,
    brand            text    not null,
    discount_percent varchar not null,
    contract_id      integer references contracts (id)
);


INSERT INTO brands(brand, discount_percent, contract_id) VALUES ($1, $2, $3);