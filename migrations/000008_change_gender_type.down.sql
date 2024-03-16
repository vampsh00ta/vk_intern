begin;


alter table actor drop column gender;
alter table actor add column gender varchar(10) check(gender in ('female','male'));

commit;