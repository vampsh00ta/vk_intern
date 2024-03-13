package repository

import (
	"context"
	"vk/internal/repository/models"
)

type Actor interface {
	AddActor(ctx context.Context, actor models.Actor) error
	DeleteActorByID(ctx context.Context, actorId int) error
	DeleteActorByFullName(ctx context.Context, name, middlename, surname string) error
	ChangeActorByID(ctx context.Context, actor models.Actor) error
	ChangeActorByFullName(ctx context.Context, name, middlename, surname string) error
	AddFilmsToActor(ctx context.Context, actorId int, films ...models.Film) error

	GetActorById(ctx context.Context, actorId string) (models.Actor, error)
	GetActorByFullName(ctx context.Context, name, middlename, surname string) (models.Actor, error)
	GetActors(ctx context.Context) ([]models.Actor, error)
}
