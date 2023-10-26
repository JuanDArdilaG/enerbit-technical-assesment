package workordersusecases

import (
	"time"

	customersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/application"
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
)

type Create struct {
	ordersRepo         workorders.Repository
	customersRepo      customers.Repository
	deactivateCustomer customersusecases.Deactivate
}

func NewCreate(ordersRepo workorders.Repository, customersRepo customers.Repository, deactivateCustomer customersusecases.Deactivate) *Create {
	return &Create{
		ordersRepo:         ordersRepo,
		customersRepo:      customersRepo,
		deactivateCustomer: deactivateCustomer,
	}
}

func (c *Create) Execute(customerID customers.ID, title string, plannedDateBegin time.Time, plannedDateEnd time.Time) error {
	customerDB, err := c.customersRepo.GetByID(customerID)
	if err != nil {
		return err
	}
	workOrder, err := workorders.New(customerDB.ToCustomer(), title, plannedDateBegin, plannedDateEnd)
	if err != nil {
		return err
	}
	err = c.deactivateCustomer.Execute(customerID)
	if err != nil {
		return err
	}
	return c.ordersRepo.Create(workOrder.ToWorkOrderDB())
}
