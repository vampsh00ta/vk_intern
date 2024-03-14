
begin;
create table customer(
    id serial primary key ,
    username varchar(50),
    admin boolean
);


commit;