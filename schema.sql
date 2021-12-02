-- auto-generated definition
create table rights
(
    id          serial                                 not null
        constraint rights_pk
            primary key,
    code        varchar,
    section     varchar,
    description text,
    created_at  timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    is_removed  boolean                  default false not null
);


create unique index rights_id_uindex
    on rights (id);

create unique index rights_code_uindex
    on rights (code);


-- auto-generated definition
create table roles
(
    id          serial                                             not null
        constraint roles_pk
            primary key,
    code        varchar,
    name        varchar,
    description text,
    status      varchar,
    created_at  timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    is_removed  boolean                  default false             not null
);


create unique index roles_id_uindex
    on roles (id);





-- auto-generated definition
create table roles_rights
(
    id         serial not null
        constraint roles_rights_pk
            primary key,
    role_id    integer
        constraint roles_rights_roles_id_fk
            references roles
            on update cascade on delete cascade,
    right_id   integer
        constraint roles_rights_rights_id_fk
            references rights
            on update cascade on delete cascade,
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    is_removed boolean                  default false
);


create unique index roles_rights_id_uindex
    on roles_rights (id);




-- auto-generated definition
create table users
(
    id         serial                                 not null
        constraint users_pk
            primary key,
    name       varchar                                not null,
    surname    varchar,
    last_name  varchar,
    login      varchar,
    email      varchar,
    password   varchar,
    status     varchar,
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    is_removed boolean                  default false not null
);


create unique index users_id_uindex
    on users (id);



-- auto-generated definition
create table if not exists users_roles
(
    id         serial not null
        constraint users_roles_pk
            primary key,
    role_id    integer
        constraint users_roles_roles_id_fk
            references roles
            on update cascade on delete set null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    is_removed boolean                  default false,
    user_id    integer
        constraint users_roles_users_id_fk
            references users
            on update cascade on delete cascade
);

create unique index users_roles_id_uindex
    on users_roles (id);