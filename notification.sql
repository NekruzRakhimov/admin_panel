CREATE TABLE notification
    id primary big serial,
    bin varchar,
    created_at timestamp  default current_timestamp,
    end_date timestamp,
    contract_type varchar,
    email varchar,
    status bool;
