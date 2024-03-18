package request

import (
	"time"
)

type AddActor struct {
	Name      string    `json:"name" example:"ivan"  validate:"required"`
	Gender    string    `json:"gender" example:"female" validate:"oneof=male female,required"`
	BirthDate time.Time `json:"birth_date" example:"2006-01-02T15:04:05Z"  validate:"required"`
	Films     []int     `json:"films,omitempty"example:"1,2"`
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
