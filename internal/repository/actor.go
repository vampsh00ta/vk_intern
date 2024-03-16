package repository

import (
	"context"
	"vk/internal/repository/models"
)

type Actor interface {
	AddActor(ctx context.Context, actor models.Actor) (int, error)
	DeleteActorByID(ctx context.Context, actorId int) error
	DeleteActorByFullName(ctx context.Context, name string) error
	ChangeActorByID(ctx context.Context, actor models.Actor) error
	ChangeActorByFullName(ctx context.Context, name string) error
	AddFilmsToActor(ctx context.Context, actorId int, films ...models.Film) error

	GetActorById(ctx context.Context, actorId int) (models.Actor, error)
	GetActorByFullName(ctx context.Context, name string) (models.Actor, error)
	GetActors(ctx context.Context) ([]models.Actor, error)
}

func (p Pg) AddActor(ctx context.Context, actor models.Actor) (int, error) {
	var id int
	//q := `insert into actor (  name ,surname , middlename  ,birth_date ) values($1,$2,$3,$4) returning id`
	q := `insert into actor (  name ,gender  ,birth_date ) values($1,$2,$3,$4) returning id`
	tx := p.getTx(ctx)
	if err := tx.Raw(q, actor.Name, actor.Gender, actor.BirthDate).Scan(&id).Error; err != nil {

		return -1, err
	}

	//if len(film.Actors) != 0 {
	//	if err := p.AddActorsToFilms(ctx, id, film.Actors...); err != nil {
	//		return -1, err
	//	}
	//}
	return id, nil
}

func (p Pg) DeleteActorByID(ctx context.Context, actorId int) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) DeleteActorByFullName(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) ChangeActorByID(ctx context.Context, actor models.Actor) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) ChangeActorByFullName(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) AddFilmsToActor(ctx context.Context, actorId int, films ...models.Film) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) GetActorById(ctx context.Context, filmId int) (models.Actor, error) {
	//q := `select  film.id as film_id,  title , description , release_date , rating ,actor.id as actor_id ,actor.name , surname , middlename , gender , birth_date from actor
	//join actor_films on actor.id =actor_films.actor_id
	//	join film on film.id = actor_films.film_id where film.id = $1
	//	`
	var actor models.Actor
	tx := p.getTx(ctx)

	err := tx.Model(&models.Actor{}).Preload("Films").
		Where("id = ?", filmId).
		First(&actor).Error
	if err != nil {
		return models.Actor{}, err
	}

	return actor, nil
}
func (p Pg) GetActorByFullName(ctx context.Context, name string) (models.Actor, error) {
	var actor models.Actor
	tx := p.getTx(ctx)

	err := tx.Model(&models.Actor{}).Preload("Films").
		Where("name = ?   ", name).
		First(&actor).Error
	if err != nil {
		return models.Actor{}, err
	}
	return actor, nil

}

func (p Pg) GetActors(ctx context.Context) ([]models.Actor, error) {
	var actors []models.Actor
	tx := p.getTx(ctx)

	err := tx.Model(&models.Actor{}).Preload("Films").Find(&actors).Error
	if err != nil {
		return nil, err
	}
	return actors, nil
}
