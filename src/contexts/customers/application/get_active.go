package customersusecases

import customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"

type GetActive struct {
	customersRepo customers.Repository
}

func NewGetActive(customersRepo customers.Repository) GetActive {
	return GetActive{
		customersRepo: customersRepo,
	}
}

func (c GetActive) Execute() ([]customers.Customer, error) {
	customersDB, err := c.customersRepo.GetActive()
	if err != nil {
		return nil, err
	}
	customers := make([]customers.Customer, len(customersDB))
	for i, customerDB := range customersDB {
		customers[i] = customerDB.ToCustomer()
	}
	return customers, nil
}
