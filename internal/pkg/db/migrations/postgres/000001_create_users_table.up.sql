create table if not exists users(
    id serial,
    username varchar (127) not null unique,
    password varchar (127) not null,
    primary key (id)
)