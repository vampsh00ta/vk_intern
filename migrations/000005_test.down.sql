begin;
delete from customer where username = 'admin' or username = 'notadmin'
commit;