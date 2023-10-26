package workordersusecases

import (
	"fmt"
	"time"

	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
)

type Search struct {
	workOrdersRepo workorders.Repository
}

func NewSearch(workOrdersRepo workorders.Repository) *Search {
	return &Search{
		workOrdersRepo: workOrdersRepo,
	}
}

func (c *Search) Execute(since string, until string, status string) ([]workorders.WorkOrder, error) {
	ordersDB, err := c.workOrdersRepo.GetAll()
	if err != nil {
		return nil, err
	}
	matchStatus := func(order workorders.WorkOrder) bool {
		return status == "" || order.MatchStatus(workorders.Status(status))
	}
	matchPlannedDate := func(order workorders.WorkOrder) bool {
		if since == "" && until == "" {
			return true
		}
		var sinceDate time.Time
		var untilDate time.Time
		var err error
		if since != "" {
			sinceDate, err = time.Parse(time.Layout, since)
		}
		if err != nil {
			fmt.Println(err)
			return false
		}
		if until != "" {
			untilDate, err = time.Parse(time.Layout, until)
		}
		if err != nil {
			fmt.Println(err)
			return false
		}
		return (since == "" && sinceDate.Before(order.PlannedDateBegin)) && (until == "" && untilDate.After(order.PlannedDateEnd))
	}
	orders := make([]workorders.WorkOrder, 0)
	for _, orderDB := range ordersDB {
		order := orderDB.ToWorkOrder()
		if matchStatus(order) && matchPlannedDate(order) {
			orders = append(orders, order)
		}
	}

	return orders, nil
}
