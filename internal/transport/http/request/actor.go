package request

import (
	"time"
	"vk/internal/repository/models"
)

type AddActor struct {
	Name      string    `json:"name,omitempty"`
	Gender    string    `json:"gender,omitempty" `
	BirthDate time.Time `json:"birth_date,omitempty"`
	//Films     []models.Id `json:"films,omitempty""`
	Films []models.Film `json:"films,omitempty""`
}
type DeleteActor struct {
	Id int `json:"id" validate:"required"`
}

type GetActors struct {
}
