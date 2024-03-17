package request

import (
	"time"
)

type AddActor struct {
	Name      string    `json:"name,omitempty" example:"ivan"`
	Gender    string    `json:"gender,omitempty" example:"female" `
	BirthDate time.Time `json:"birth_date,omitempty" example:"2006-01-02T15:04:05Z"`
	//Films     []models.Id `json:"films,omitempty""`
	Films []int `json:"films,omitempty"example:"1,2"`
}
type UpdateActor struct {
	Id        int       `json:"id"  validate:"required"`
	Name      string    `json:"name,omitempty"`
	Gender    string    `json:"gender,omitempty" `
	BirthDate time.Time `json:"birth_date,omitempty"`
	//Films     []models.Id `json:"films,omitempty""`
	Films []int `json:"films,omitempty""`
}
type DeleteActor struct {
	Id int `json:"id" validate:"required"`
}

type GetActors struct {
}
