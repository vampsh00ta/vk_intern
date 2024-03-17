package service

import (
	"context"
	"vk/internal/repository/models"
)

type Actor interface {
	AddActor(ctx context.Context, actor models.Actor) (int, error)
	DeleteActorByID(ctx context.Context, actorId int) error
	//DeleteActorByFullName(ctx context.Context, name string) error
	ChangeActor(ctx context.Context, actor models.Actor) error
	ChangeActorPartly(ctx context.Context, actor models.Actor) error

	//ChangeActorByFullName(ctx context.Context, name string) error
	//AddFilmsToActor(ctx context.Context, actorId int, films ...models.Film) error
	//
	//GetActorById(ctx context.Context, actorId int) (models.Actor, error)
	//GetActorByFullName(ctx context.Context, name string) (models.Actor, error)
	GetActors(ctx context.Context) ([]models.Actor, error)
}

func (s service) ChangeActor(ctx context.Context, actor models.Actor) error {
	ctx = s.repo.Begin(ctx)
	defer s.repo.Commit(ctx)
	if err := s.repo.ChangeActorByID(ctx, actor); err != nil {
		s.repo.Rollback(ctx)

		return err
	}
	return nil
}
func (s service) ChangeActorPartly(ctx context.Context, actor models.Actor) error {
	ctx = s.repo.Begin(ctx)
	defer s.repo.Commit(ctx)
	if err := s.repo.ChangeActorByIDPartly(ctx, actor); err != nil {
		s.repo.Rollback(ctx)

		return err
	}
	return nil
}
func (s service) DeleteActorByID(ctx context.Context, actorId int) error {
	if err := s.repo.DeleteActorByID(ctx, actorId); err != nil {
		return err
	}
	return nil
}
func (s service) AddActor(ctx context.Context, actor models.Actor) (int, error) {
	id, err := s.repo.AddActor(ctx, actor)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (s service) GetActors(ctx context.Context) ([]models.Actor, error) {
	actors, err := s.repo.GetActors(ctx)
	if err != nil {
		return nil, err
	}
	return actors, nil
}
