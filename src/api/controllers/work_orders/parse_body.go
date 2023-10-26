package workorderscontrollers

import (
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/apperrors"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
	"github.com/gin-gonic/gin"
)

func ParseBody(c *gin.Context) (workorders.WorkOrder, error) {
	workOrder := workorders.WorkOrder{}
	err := c.ShouldBindJSON(&workOrder)
	if err != nil {
		return workorders.WorkOrder{}, apperrors.NewBadRequest("workorder", err.Error())
	}
	workOrder.Status = workorders.StatusNew
	err = workOrder.Validate()
	if err != nil {
		return workorders.WorkOrder{}, err
	}
	return workOrder, nil
}
