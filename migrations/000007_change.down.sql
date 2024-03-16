begin;
alter table actor add column middlename varchar(255);
alter table actor add column surname varchar(255);

commit ;