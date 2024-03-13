package models

import "time"

type Actor struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Middlename string    `json:"middlename"`
	Surname    string    `json:"surname"`
	BirthDate  time.Time `json:"birth_date"`
	Films      []Film
}

type Film struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	ReleaseDate time.Time `json:"release_date"`
	Actors      Actor
}
