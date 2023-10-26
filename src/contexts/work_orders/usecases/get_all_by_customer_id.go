package workordersusecases

import (
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
)

type GetAllByCustomerID struct {
	workOrdersRepo workorders.Repository
}

func NewGetAllByCustomerID(workOrdersRepo workorders.Repository) *GetAllByCustomerID {
	return &GetAllByCustomerID{
		workOrdersRepo: workOrdersRepo,
	}
}

func (c *GetAllByCustomerID) Execute(customerID customers.ID) ([]workorders.WorkOrder, error) {
	ordersDB, err := c.workOrdersRepo.GetAllByCustomerID(customerID)
	if err != nil {
		return nil, err
	}
	orders := make([]workorders.WorkOrder, len(ordersDB))
	for i, orderDB := range ordersDB {
		orders[i] = orderDB.ToWorkOrder()
	}
	return orders, nil
}
