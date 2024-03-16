package request

type Customer struct {
	Username string `json:"username"  validate:"required"`
}
