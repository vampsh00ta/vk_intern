package models

import "time"

type Film struct {
	Id          int       `json:"id,omitempty" db:"id"`
	Title       string    `json:"title,omitempty" db:"title" mapstructure:"title"`
	Description string    `json:"description,omitempty" db:"description" mapstructure:"description"`
	Rating      int       `json:"rating,omitempty" db:"rating" mapstructure:"rating"`
	ReleaseDate time.Time `json:"release_date,omitempty" db:"release_date" mapstructure:"release_date"`
	Actors      []Actor   `json:"actors,omitempty" gorm:"many2many:actor_films;"`
	ActorIds    []int     `gorm:"-"  json:"-"`
}

func (Film) TableName() string {
	return "film"
}
