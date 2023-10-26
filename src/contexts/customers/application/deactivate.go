package customersusecases

import (
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
)

type Deactivate struct {
	customersRepo customers.Repository
}

func NewDeactivate(repos dependencies.Repositories) Deactivate {
	return Deactivate{
		customersRepo: repos.Customers,
	}
}

func (c Deactivate) Execute(customerID customers.ID) error {
	customerDB, err := c.customersRepo.GetByID(customerID)
	if err != nil {
		return err
	}
	customer := customerDB.ToCustomer()
	customer.Deactivate()
	return c.customersRepo.Update(customer.ID, customer.ToCustomerDB())
}
