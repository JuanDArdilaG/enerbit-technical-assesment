package customerscontrollers

type CreateRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Address   string `json:"address" validate:"required"`
}
