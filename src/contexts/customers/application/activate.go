package customersusecases

import (
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
)

type Activate struct {
	customersRepo customers.Repository
}

func NewActivate(repos dependencies.Repositories) Activate {
	return Activate{
		customersRepo: repos.Customers,
	}
}

func (c Activate) Execute(customerID customers.ID) error {
	customerDB, err := c.customersRepo.GetByID(customerID)
	if err != nil {
		return err
	}
	customer := customerDB.ToCustomer()
	customer.Activate()
	return c.customersRepo.Update(customer.ID, customer.ToCustomerDB())
}
