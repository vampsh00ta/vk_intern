package repository

import (
	"context"
	"vk/internal/repository/models"
)

type Actor interface {
	AddActor(ctx context.Context, actor models.Actor) (int, error)
	DeleteActorByID(ctx context.Context, actorId int) error
	DeleteActorByFullName(ctx context.Context, name, middlename, surname string) error
	ChangeActorByID(ctx context.Context, actor models.Actor) error
	ChangeActorByFullName(ctx context.Context, name, middlename, surname string) error
	AddFilmsToActor(ctx context.Context, actorId int, films ...models.Film) error

	GetActorById(ctx context.Context, actorId string) (models.Actor, error)
	GetActorByFullName(ctx context.Context, name, middlename, surname string) (models.Actor, error)
	GetActors(ctx context.Context) ([]models.Actor, error)
}

func (p Pg) AddActor(ctx context.Context, actor models.Actor) (int, error) {
	var id int
	q := `insert into actor (  name ,surname , middlename  ,birth_date ) values($1,$2,$3,$4) returning id`
	tx := p.getTx(ctx)
	if err := tx.Raw(q, actor.Name, actor.Surname, actor.Middlename, actor.BirthDate).Scan(&id).Error; err != nil {

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

func (p Pg) DeleteActorByFullName(ctx context.Context, name, middlename, surname string) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) ChangeActorByID(ctx context.Context, actor models.Actor) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) ChangeActorByFullName(ctx context.Context, name, middlename, surname string) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) AddFilmsToActor(ctx context.Context, actorId int, films ...models.Film) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) GetActorById(ctx context.Context, actorId string) (models.Actor, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pg) GetActorByFullName(ctx context.Context, name, middlename, surname string) (models.Actor, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pg) GetActors(ctx context.Context) ([]models.Actor, error) {
	//TODO implement me
	panic("implement me")
}
