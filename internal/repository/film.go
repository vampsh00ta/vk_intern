package repository

import (
	"context"
	"fmt"
	"vk/internal/repository/models"
)

type Film interface {
	AddFilm(ctx context.Context, film models.Film) (int, error)

	DeleteFilmByID(ctx context.Context, filmId int) error
	DeleteFilmByTitle(ctx context.Context, filmTitle string) error

	ChangeFilmByID(ctx context.Context, film models.Film) error
	ChangeFilmByFullName(ctx context.Context, filmId int) error

	AddActorsToFilms(ctx context.Context, filmId int, actor ...models.Actor) error

	GetFilmById(ctx context.Context, filmId int) (models.Film, error)
	GetFilmByTitle(ctx context.Context, filmTitle string) (models.Film, error)
	GetFilms(ctx context.Context) ([]models.Film, error)
}

func (p Pg) AddFilm(ctx context.Context, film models.Film) (int, error) {
	var id int
	q := `insert into film ( title,description,release_date,rating ) values($1,$2,$3,$4) returning id`
	tx := p.getTx(ctx)
	if err := tx.Raw(q, film.Title, film.Description, film.ReleaseDate, film.Rating).Scan(&id).Error; err != nil {

		return -1, err
	}

	//if len(film.Actors) != 0 {
	//	if err := p.AddActorsToFilms(ctx, id, film.Actors...); err != nil {
	//		return -1, err
	//	}
	//}
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
	q := `update film where id = $1 set title = $2 ,  description = $3, release_date = $4,rating  = $5`
	tx := p.getTx(ctx)

	if err := tx.Raw(q, film.Id, film.Title, film.Description, film.ReleaseDate, film.Rating).
		Scan(&film.Id).Error; err != nil {

		return err
	}

	return nil
}

func (p Pg) ChangeFilmByFullName(ctx context.Context, filmId int) error {
	//TODO implement me
	panic("implement me")
}

func (p Pg) AddActorsToFilms(ctx context.Context, filmId int, actors ...models.Actor) error {
	q := `insert into actor_films ( actor_films,film_id) values `
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

	//film, err := pgx.CollectOneRow(row, pgx.RowToStructByNameLax[models.Film])
	//if err := pgxscan.ScanRow(&film, row); err != nil {
	//	return models.Film{}, err
	//}
	return film, nil
}

func (p Pg) GetFilmByTitle(ctx context.Context, filmTitle string) (models.Film, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pg) GetFilms(ctx context.Context) ([]models.Film, error) {
	//TODO implement me
	panic("implement me")
}
