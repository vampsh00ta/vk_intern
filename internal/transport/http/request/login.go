package request

type Customer struct {
	Username string `json:"username" example:"admin" validate:"required"`
}
