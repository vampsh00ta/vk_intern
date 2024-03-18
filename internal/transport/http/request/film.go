package request

import (
	"time"
)

type AddFilm struct {
	Title       string    `json:"title,omitempty" validate:"min=1,max=150,required"`
	Description string    `json:"description,omitempty" validate:"max=1000,required"`
	Rating      int       `json:"rating,omitempty" validate:"gte=0,lte=10"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	//Actors      []models.Id `json:"actors,omitempty" `
	Actors []int `json:"actors,omitempty" `
	//
	//models.Film
}
type GetFilm struct {
	OrderBy string `json:"order_by,omitempty" schema:"order_by"`
	SortBy  string `json:"sort_by,omitempty" schema:"sort_by"`
	Name    string `json:"name,omitempty" schema:"name"`
	Title   string `json:"title,omitempty" schema:"title"`
}
type UpdateFilm struct {
	Id          int       `json:"id"  validate:"required"`
	Title       string    `json:"title,omitempty" validate:"min=1,max=150"`
	Description string    `json:"description,omitempty" validate:"max=1000"`
	Rating      int       `json:"rating,omitempty" validate:"gte=0,lte=10"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Actors      []int     `json:"actors,omitempty" `
}
type DeleteFilm struct {
	Id int `json:"id"  validate:"required"`
}
