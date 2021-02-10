create table users (
    id serial not null,
    name varchar(255) unique not null,
    score integer not null default 0,
    primary key (id)
);
