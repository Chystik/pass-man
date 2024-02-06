create schema if not exists passman;

create table if not exists passman.user (
    login varchar(50) primary key not null unique,
    password bytea not null,
    vault_key bytea not null
);

create table if not exists passman.password (
    id uuid primary key default gen_random_uuid(),
    user_id varchar(50) references passman.user(login) on delete cascade not null,
    meta bytea not null,
    username bytea not null,
    password bytea not null
);

create table if not exists passman.card (
    id uuid primary key default gen_random_uuid(),
    user_id varchar(50) references passman.user(login) on delete cascade not null,
    meta bytea not null,
    number bytea not null,
    valid_thru bytea not null,
    holder bytea not null,
    cvv bytea not null
);

create table if not exists passman.note (
    id uuid primary key default gen_random_uuid(),
    user_id varchar(50) references passman.user(login) on delete cascade not null,
    meta bytea not null,
    note bytea not null
);

create table if not exists passman.file (
    id uuid primary key default gen_random_uuid(),
    user_id varchar(50) references passman.user(login) on delete cascade not null,
    meta bytea not null,
    full_name bytea not null,
    data oid not null
);