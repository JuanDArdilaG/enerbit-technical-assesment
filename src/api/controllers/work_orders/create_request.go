package workorderscontrollers

import (
	"time"

	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
)

type CreateRequest struct {
	CustomerID       customers.ID `json:"customer_id" validate:"required,uuid"`
	Title            string       `json:"title" validate:"min=5,max=255"`
	PlannedDateBegin time.Time    `json:"planned_date_begin" validate:"required"`
	PlannedDateEnd   time.Time    `json:"planned_date_end" validate:"required"`
}
