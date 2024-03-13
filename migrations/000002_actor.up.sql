begin;
create table actor(
                      id serial primary key ,
                      name varchar(255),
                      surname varchar(255),
                      middlename  varchar(255),
                      gender boolean,
                      birth_date date
);
commit;