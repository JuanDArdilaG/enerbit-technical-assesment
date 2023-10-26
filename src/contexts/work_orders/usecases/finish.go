package workordersusecases

import (
	customersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/application"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
)

type Finish struct {
	woRepo           workorders.Repository
	stream           workorders.Publisher
	activateCustomer customersusecases.Activate
}

func NewFinish(
	woRepo workorders.Repository,
	stream workorders.Publisher,
	activateCustomer customersusecases.Activate,
) *Finish {
	return &Finish{
		woRepo:           woRepo,
		stream:           stream,
		activateCustomer: activateCustomer,
	}
}

func (c *Finish) Execute(id workorders.ID) error {
	workOrderDB, err := c.woRepo.GetByID(id)
	if err != nil {
		return err
	}
	workOrder := workOrderDB.ToWorkOrder()

	err = c.activateCustomer.Execute(workOrder.CustomerID)
	if err != nil {
		return err
	}

	workOrder.Finish()
	c.stream.Publish(workOrder.ID, workOrder.Status)
	return c.woRepo.Update(workOrder.ID, workOrder.ToWorkOrderDB())
}
