begin;
create table film(
                      id serial primary key ,
                      title varchar(150),
                      description text,
                      release_date  date,
                      rating integer
--                       actors
);

create table actor_films(
                      actor_id integer references  actor(id),
                      film_id integer references  film(id)

);
commit;