package request

import (
	"time"
)

type AddFilm struct {
	Title       *string   `json:"title,omitempty"  example:"test" validate:"min=1,max=150,required"`
	Description string    `json:"description,omitempty" example:"test" validate:"max=1000,required"`
	Rating      int       `json:"rating,omitempty" example:"10" validate:"gte=0,lte=10"`
	ReleaseDate time.Time `json:"release_date,omitempty" example:"2006-01-02T15:04:05Z"`
	Actors      []int     `json:"actors,omitempty" example:"1,2"`
}
type GetFilm struct {
	OrderBy string `json:"order_by,omitempty" `
	SortBy  string `json:"sort_by,omitempty" `
	Name    string `json:"name,omitempty" `
	Title   string `json:"title,omitempty" `
}
type UpdateFilm struct {
	Id          int       `json:"id"  example:"1" validate:"required"`
	Title       string    `json:"title,omitempty" example:"test" validate:"min=1,max=150"`
	Description string    `json:"description,omitempty" example:"test" validate:"max=1000"`
	Rating      int       `json:"rating,omitempty" example:"10" validate:"gte=0,lte=10"`
	ReleaseDate time.Time `json:"release_date,omitempty" example:"2006-01-02T15:04:05Z"`
	Actors      []int     `json:"actors,omitempty" example:"1,2"`
}
type UpdateFilmPartly struct {
	Id          int       `json:"id"  example:"1" validate:"required"`
	Title       string    `json:"title,omitempty" example:"test" validate:"omitempty,min=1,max=150"`
	Description string    `json:"description,omitempty" example:"test" validate:"max=1000"`
	Rating      int       `json:"rating,omitempty" example:"10" validate:"omitempty,gte=0,lte=10"`
	ReleaseDate time.Time `json:"release_date,omitempty" example:"2006-01-02T15:04:05Z"`
	Actors      []int     `json:"actors,omitempty" example:"1,2"`
}
type DeleteFilm struct {
	Id int `json:"id" example:"1"  validate:"required"`
}
