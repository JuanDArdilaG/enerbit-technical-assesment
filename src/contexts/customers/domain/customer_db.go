package customers

import (
	"time"

	"gorm.io/gorm"
)

type CustomerDB struct {
	gorm.Model
	ID        string     `json:"id" gorm:"size:36;unique;primary_key;not null"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Address   string     `json:"address"`
	IsActive  bool       `json:"is_active"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedAt time.Time  `json:"created_at"`
}

func (c CustomerDB) ToCustomer() Customer {
	return Customer{
		ID:        ID(c.ID),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address:   c.Address,
		IsActive:  c.IsActive,
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
		CreatedAt: c.CreatedAt,
	}
}
