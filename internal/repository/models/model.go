package models

type Customer struct {
	Id       int    `json:"id,omitempty" db:"id"`
	Username string `json:"username,omitempty" db:"username"`
	Admin    bool   `json:"admin,omitempty" db:"admin"`
}

type Id struct {
	Id string `json:"id,omitempty"`
}
