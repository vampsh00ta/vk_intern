begin;
drop  table actor_films;
create table actor_films(
                            actor_id integer references  actor(id) on delete cascade ,
                            film_id integer references  film(id) on delete cascade

);
commit ;