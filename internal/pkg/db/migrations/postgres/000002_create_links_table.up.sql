create table if not exists links(
    id serial,
    title varchar (255) ,
    address varchar (255) ,
    userid integer ,
    foreign key (userid) references users(id) ,
    primary key (id)
)