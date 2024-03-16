package request

import (
	"time"
	"vk/internal/repository/models"
)

type AddFilm struct {
	Title       string    `json:"title,omitempty" `
	Description string    `json:"description,omitempty" `
	Rating      int       `json:"rating,omitempty" `
	ReleaseDate time.Time `json:"release_date,omitempty"`
	//Actors      []models.Id `json:"actors,omitempty" `
	Actors []models.Actor `json:"actors,omitempty" `
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
	Title       string    `json:"title,omitempty" `
	Description string    `json:"description,omitempty" `
	Rating      int       `json:"rating,omitempty" `
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Actors      []int     `json:"actors,omitempty" `
}
type DeleteFilm struct {
	Id int `json:"id"  validate:"required"`
}
