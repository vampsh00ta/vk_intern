package repository

import (
	"context"
	"fmt"
	"reflect"
	"vk/internal/repository/models"
)

type Actor interface {
	AddActor(ctx context.Context, actor models.Actor) (int, error)

	DeleteActorByID(ctx context.Context, actorId int) error
	DeleteActorByFullName(ctx context.Context, name string) error

	ChangeActorByID(ctx context.Context, actor models.Actor) error
	ChangeActorByFullName(ctx context.Context, name string) error
	//ChangeActorFilms(ctx context.Context, actor int, filmIds ...int) error
	ChangeActorByIDPartly(ctx context.Context, actor models.Actor) error

	GetActorById(ctx context.Context, actorId int) (models.Actor, error)
	GetActorByFullName(ctx context.Context, name string) (models.Actor, error)
	GetActors(ctx context.Context) ([]models.Actor, error)
}

// func (p Pg) AddFilmsToActorById(ctx context.Context, actorId int, filmIds ...models.Film) error
func (p Pg) AddActor(ctx context.Context, actor models.Actor) (int, error) {
	var id int
	//q := `insert into actor (  name ,surname , middlename  ,birth_date ) values($1,$2,$3,$4) returning id`
	q := `insert into actor (  name ,gender  ,birth_date ) values($1,$2,$3) returning id`
	tx := p.getTx(ctx)
	if err := tx.Raw(q, actor.Name, actor.Gender, actor.BirthDate).Scan(&id).Error; err != nil {

		return -1, err
	}

	if len(actor.FilmIds) != 0 {
		if err := p.AddFilmsToActorByIds(ctx, id, actor.FilmIds...); err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (p Pg) AddFilmsToActorByIds(ctx context.Context, actorId int, filmIds ...int) error {
	q := `insert into actor_films ( actor_id ,film_id) values `
	inputVals := []any{actorId}
	tx := p.getTx(ctx)
	for i, id := range filmIds {
		q += fmt.Sprintf("($1,$%d),", i+2)
		inputVals = append(inputVals, id)
	}
	q = q[0 : len(q)-1]
	if err := tx.Raw(q, inputVals...).Scan(&actorId).Error; err != nil {
		return err
	}

	return nil
}
func (p Pg) DeleteActorByID(ctx context.Context, actorId int) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) DeleteActorByFullName(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) ChangeActorByFullName(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}
func (p Pg) ChangeActorFilms(ctx context.Context, actorId int, filmIds ...int) error {
	tx := p.getTx(ctx)

	q := `delete from  actor_films where actor_id = $1 `

	if err := tx.Raw(q, actorId).Scan(&actorId).Error; err != nil {
		return err
	}

	input := []any{actorId}

	q = `insert into  actor_films ( actor_id ,film_id)  values `
	for i, film := range filmIds {
		q += fmt.Sprintf("($1,$%d),", i+2)
		input = append(input, film)
	}
	q = q[:len(q)-1]
	if err := tx.Raw(q, input...).Scan(&actorId).Error; err != nil {
		return err
	}
	return nil

}

func (p Pg) ChangeActorByID(ctx context.Context, actor models.Actor) error {
	q := `update actor set name = $2 ,  birth_date = $3, gender = $4 where id = $1`

	tx := p.getTx(ctx)

	if err := tx.Raw(q, actor.Id, actor.Name, actor.BirthDate, actor.Gender).
		Scan(&actor.Id).Error; err != nil {

		return err
	}
	if len(actor.FilmIds) > 0 {
		if err := p.ChangeActorFilms(ctx, actor.Id, actor.FilmIds...); err != nil {
			return err
		}
	}
	return nil
}
func (p Pg) ChangeActorByIDPartly(ctx context.Context, actor models.Actor) error {
	q := `update actor  set `

	tx := p.getTx(ctx)
	count := 2
	input := []any{actor.Id}
	v := reflect.ValueOf(actor)
	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)

		if tagv := fi.Tag.Get("db"); tagv != "" && tagv != "id" {
			if !v.Field(i).IsZero() {
				q += fmt.Sprintf("%s  = $%d,", tagv, count)
				input = append(input, v.Field(i).Interface())
				count += 1

			}

		}
	}
	if count > 2 {
		q = q[:len(q)-1] + " where id = $1"
		if err := tx.Raw(q, input...).Scan(nil).Error; err != nil {

			return err
		}
	}

	if len(actor.FilmIds) > 0 {
		if err := p.ChangeActorFilms(ctx, actor.Id, actor.FilmIds...); err != nil {
			return err
		}
	}

	return nil

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
