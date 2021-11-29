create table marketing_services_contract
(
    id         serial primary key,
    manager    text      not null,
    kam        text      not null,
    status     bool,
    requisites_id references requisites (id),
    supplier_company_manager_id references supplier_company_manager (id),
    contract_parameters_id references contract_parameters (id),

    discountPercent_id references discountPercent (id),
    created_at timestamp not null DEFAULT current_timestamp


);

create table requisites
(
    id                  serial primary key,
    beneficiary         varchar not null,
    bank_of_beneficiary varchar not null,
    bin                 varchar not null,
    iic                 varchar not null,
    phone               varchar not null,
    account_number      varchar not null
);


create table if not exists supplier_company_manager
(
    id         serial primary key,
    work_phone varchar not null,
    email      varchar not null,
    skype      varchar not null,
    phone      varchar not null,
    position   varchar not null,
    base       varchar not null

);

create table contract_parameters
(
    id                          serial primary key,
    number_of_contract          varchar not null,
    amount_contract             integer not null,
    -- currency - не хватает
    prepayment                  int     not null,
    date_of_delivery            timestamptz,
    frequency_deferred_discount varchar not null,
    delivery_address            varchar not null
);

create table currency
(
    id               serial primary key
        constraint currencies_pk
            primary key,
    alpha3           varchar,
    symbol           varchar,
    name             varchar,
    image_name       varchar,
    created_at       timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at       timestamp with time zone,
    deleted_at       timestamp with time zone,
    is_removed       boolean                  default false not null,
    rate             double precision,
    base_currency_id integer
);

create table products
(
    id             serial primary key,
    product_number varchar not null,
    price          int     not null,
    currency_id    int references currency (id)
);

create table if not exists discountPercent
(
    id         serial primary key,
    type       varchar not null,
    name       varchar not null,
    amount     integer not null,
    is_active  bool default false,
    grace_days integer not null


)