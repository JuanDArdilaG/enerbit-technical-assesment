package customers

import (
	"time"

	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/apperrors"
	"github.com/go-playground/validator/v10"
)

type Customer struct {
	ID        ID         `json:"id"`
	FirstName string     `json:"first_name" validate:"required"`
	LastName  string     `json:"last_name" validate:"required"`
	Address   string     `json:"address" validate:"required"`
	IsActive  bool       `json:"is_active"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedAt time.Time  `json:"created_at"`
}

func New(firstName string, lastName string, address string) (*Customer, error) {
	id, err := NewRandomID()
	if err != nil {
		return nil, err
	}
	return &Customer{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		IsActive:  false,
		CreatedAt: time.Now(),
	}, nil
}

func (c *Customer) Activate() {
	c.IsActive = true
	now := time.Now()
	c.StartDate = &now
}

func (c *Customer) Deactivate() {
	c.IsActive = false
	now := time.Now()
	c.EndDate = &now
}

func (wo *Customer) Validate() error {
	validate := validator.New()
	err := validate.Struct(wo)
	if err != nil {
		return apperrors.NewBadRequest("workorder", err.Error())
	}
	return nil
}

func (c *Customer) ToCustomerDB() CustomerDB {
	return CustomerDB{
		ID:        string(c.ID),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address:   c.Address,
		IsActive:  c.IsActive,
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
		CreatedAt: c.CreatedAt,
	}
}
