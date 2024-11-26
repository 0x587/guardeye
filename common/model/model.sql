create table node
(
    id         uuid   not null
        constraint node_pk
            primary key,
    client_id  uuid   not null,
    created_at time,
    updated_at time,
    alias      text,
    ips        text[] not null,
    macs       text[] not null,
    os         text   not null,
    os_version text   not null,
    hostname   text   not null
);

alter table node
    owner to postgres;

