package service

import (
	"context"
	"vk/internal/repository/models"
)

type Film interface {
	AddFilm(ctx context.Context, film models.Film) (int, error)

	DeleteFilm(ctx context.Context, id int) error

	ChangeFilm(ctx context.Context, film models.Film) error
	ChangeFilmPartly(ctx context.Context, film models.Film) error

	GetFilmsByParams(ctx context.Context, params models.SortParams) ([]models.Film, error)
	GetFilms(ctx context.Context, sortBy, orderBy string) ([]models.Film, error)
	GetFilmsByTitle(ctx context.Context, title string, sortBy, orderBy string) ([]models.Film, error)
	GetFilmsByActorName(ctx context.Context, name string, sortBy, orderBy string) ([]models.Film, error)
}

func (s service) ChangeFilm(ctx context.Context, film models.Film) error {
	ctx = s.repo.Begin(ctx)
	defer s.repo.Commit(ctx)
	if err := s.repo.ChangeFilmByID(ctx, film); err != nil {
		s.repo.Rollback(ctx)

		return err
	}
	return nil
}
func (s service) ChangeFilmPartly(ctx context.Context, film models.Film) error {
	ctx = s.repo.Begin(ctx)
	defer s.repo.Commit(ctx)
	if err := s.repo.ChangeFilmByIDPartly(ctx, film); err != nil {
		s.repo.Rollback(ctx)

		return err
	}
	return nil
}

func (s service) DeleteFilm(ctx context.Context, id int) error {

	if err := s.repo.DeleteFilmByID(ctx, id); err != nil {
		return err
	}
	return nil
}
func (s service) GetFilmsByTitle(ctx context.Context, title string, sortBy, orderBy string) ([]models.Film, error) {
	films, err := s.repo.GetFilmsByTitle(ctx, title, sortBy, orderBy)
	if err != nil {
		return nil, err
	}
	return films, nil
}
func (s service) GetFilmsByActorName(ctx context.Context, name string, sortBy, orderBy string) ([]models.Film, error) {
	films, err := s.repo.GetFilmsByActorName(ctx, name, sortBy, orderBy)
	if err != nil {
		return nil, err
	}
	return films, nil
}

func (s service) GetFilmsByParams(ctx context.Context, params models.SortParams) ([]models.Film, error) {
	films, err := s.repo.GetFilmsByParams(ctx, params)
	if err != nil {
		return nil, err
	}
	return films, nil
}
func (s service) AddFilm(ctx context.Context, film models.Film) (int, error) {
	ctx = s.repo.Begin(ctx)
	defer s.repo.Commit(ctx)
	id, err := s.repo.AddFilm(ctx, film)
	if err != nil {
		s.repo.Rollback(ctx)
		return -1, err
	}
	return id, nil
}

func (s service) GetFilms(ctx context.Context, sortBy, orderBy string) ([]models.Film, error) {
	films, err := s.repo.GetFilms(ctx, sortBy, orderBy)
	if err != nil {
		return nil, err
	}
	return films, nil
}
