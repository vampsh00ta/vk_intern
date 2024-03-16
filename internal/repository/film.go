package repository

import (
	"context"
	"fmt"
	"reflect"
	"vk/internal/repository/models"
)

type Film interface {
	AddFilm(ctx context.Context, film models.Film) (int, error)

	DeleteFilmByID(ctx context.Context, filmId int) error
	DeleteFilmByTitle(ctx context.Context, filmTitle string) error

	ChangeFilmsActors(ctx context.Context, filmId int, actorsFilms ...int) error

	ChangeFilmByID(ctx context.Context, film models.Film) error
	ChangeFilmByIDPartly(ctx context.Context, film models.Film) error

	ChangeFilmByFullName(ctx context.Context, filmId int) error

	GetFilmById(ctx context.Context, filmId int) (models.Film, error)
	GetFilmsByActorName(ctx context.Context, name string, sortBy, orderBy string) ([]models.Film, error)
	GetFilmsByTitle(ctx context.Context, filmTitle string, sortBy, orderBy string) ([]models.Film, error)
	GetFilms(ctx context.Context, sortBy, orderBy string) ([]models.Film, error)
}

func (p Pg) AddFilm(ctx context.Context, film models.Film) (int, error) {
	var id int
	q := `insert into film ( title,description,release_date,rating ) values($1,$2,$3,$4) returning id`
	tx := p.getTx(ctx)
	if err := tx.Raw(q, film.Title, film.Description, film.ReleaseDate, film.Rating).Scan(&id).Error; err != nil {

		return -1, err
	}
	fmt.Println(q)

	if len(film.Actors) != 0 {
		if err := p.AddActorsToFilms(ctx, id, film.Actors...); err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (p Pg) DeleteFilmByID(ctx context.Context, filmId int) error {
	q := `delete from film where id = $1 `
	tx := p.getTx(ctx)

	if err := tx.Raw(q, filmId).Scan(&filmId).Error; err != nil {

		return err
	}

	return nil
}

func (p Pg) DeleteFilmByTitle(ctx context.Context, filmTitle string) error {
	q := `delete from film where title = $1 `
	tx := p.getTx(ctx)

	if err := tx.Raw(q, filmTitle).Scan(&filmTitle).Error; err != nil {
		return err
	}

	return nil
}

// // dfsdfsdf
func (p Pg) ChangeFilmByID(ctx context.Context, film models.Film) error {
	q := `update film set title = $2 ,  description = $3, release_date = $4,rating  = $5  where id = $1`

	tx := p.getTx(ctx)

	if err := tx.Raw(q, film.Id, film.Title, film.Description, film.ReleaseDate, film.Rating).
		Scan(&film.Id).Error; err != nil {

		return err
	}
	if len(film.ActorIds) > 0 {
		if err := p.ChangeFilmsActors(ctx, film.Id, film.ActorIds...); err != nil {
			return err
		}
	}

	return nil
}

// поправить рефлект
func (p Pg) ChangeFilmByIDPartly(ctx context.Context, film models.Film) error {
	q := `update film  set `

	tx := p.getTx(ctx)
	count := 2
	input := []any{film.Id}
	v := reflect.ValueOf(film)
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

	if len(film.ActorIds) > 0 {
		if err := p.ChangeFilmsActors(ctx, film.Id, film.ActorIds...); err != nil {
			return err
		}
	}

	return nil
}
func (p Pg) ChangeFilmsActors(ctx context.Context, filmId int, actorsFilms ...int) error {
	tx := p.getTx(ctx)

	q := `delete from  actor_films where film_id = $1 `
	if err := tx.Raw(q, filmId).Error; err != nil {
		return err
	}
	if err := tx.Raw(q, filmId).Scan(&filmId).Error; err != nil {
		return err
	}
	q = `insert into  actor_films (film_id,actor_id)  values `
	input := []any{filmId}
	for i, actor := range actorsFilms {
		q += fmt.Sprintf("($1,$%d),", i+2)
		input = append(input, actor)
	}
	q = q[:len(q)-1]

	if err := tx.Raw(q, input...).Scan(&filmId).Error; err != nil {
		return err
	}
	return nil
}
func (p Pg) ChangeFilmByFullName(ctx context.Context, filmId int) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) AddActorsToFilms(ctx context.Context, filmId int, actors ...models.Actor) error {
	q := `insert into actor_films ( film_id,actor_id) values `
	inputVals := []any{filmId}
	tx := p.getTx(ctx)
	for i, actor := range actors {
		q += fmt.Sprintf("($1,$%d),", i+2)
		inputVals = append(inputVals, actor.Id)
	}
	q = q[0 : len(q)-1]
	if err := tx.Raw(q, inputVals...).Scan(&filmId).Error; err != nil {
		return err
	}

	return nil
}

func (p Pg) GetFilmById(ctx context.Context, filmId int) (models.Film, error) {
	//q := `select  film.id as film_id,  title , description , release_date , rating ,actor.id as actor_id ,actor.name , surname , middlename , gender , birth_date from actor
	//join actor_films on actor.id =actor_films.actor_id
	//	join film on film.id = actor_films.film_id where film.id = $1
	//	`
	var film models.Film
	tx := p.getTx(ctx)

	err := tx.Model(&models.Film{}).Preload("Actors").
		Where("id = ?", filmId).
		First(&film).Error
	if err != nil {
		return models.Film{}, err
	}

	return film, nil
}
func (p Pg) GetFilmsByTitle(ctx context.Context, filmTitle string, sortBy, orderBy string) ([]models.Film, error) {
	var films []models.Film
	tx := p.getTx(ctx)
	sortVal := sortStatement(sortBy, orderBy)
	err := tx.Model(&models.Film{}).Preload("Actors").
		Where("title like $1 ", "%"+filmTitle+"%").
		Order(sortVal).
		Find(&films).Error
	if err != nil {
		return nil, err
	}
	return films, nil

}
func (p Pg) GetFilmsByActorName(ctx context.Context, name string, sortBy, orderBy string) ([]models.Film, error) {
	var films []models.Film
	tx := p.getTx(ctx)

	err := tx.Model(&models.Film{}).
		Joins("join actor_films on actor_films.film_id = film.id").
		Joins("join actor on actor_films.actor_id = actor.id").
		Where("actor.name like $1 ",
			"%"+name+"%").
		Preload("Actors").
		Find(&films).
		Error
	if err != nil {
		return nil, err
	}
	return films, nil
}

func (p Pg) GetFilms(ctx context.Context, sortBy, orderBy string) ([]models.Film, error) {
	var films []models.Film
	tx := p.getTx(ctx)
	sortVal := sortStatement(sortBy, orderBy)

	err := tx.Model(&models.Film{}).Preload("Actors").Order(sortVal).Find(&films).Error
	if err != nil {
		return nil, err
	}
	return films, nil
}
