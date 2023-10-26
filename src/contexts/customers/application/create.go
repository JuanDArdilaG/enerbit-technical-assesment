package customersusecases

import (
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
)

type Create struct {
	customersRepo customers.Repository
}

func NewCreate(customersRepo customers.Repository) *Create {
	return &Create{
		customersRepo: customersRepo,
	}
}

func (c *Create) Execute(firstName string, lastName string, address string) error {
	customer, err := customers.New(firstName, lastName, address)
	if err != nil {
		return err
	}
	return c.customersRepo.Create(customer.ToCustomerDB())
}
