CREATE TABLE notification
(
    id  bigserial primary key,
    bin           varchar,
    created_at    timestamp default current_timestamp,
    contract_date      varchar,
    contract_type varchar,
    email         varchar,
    status        bool
);

drop table notification;