package workorders

import (
	"time"

	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	"gorm.io/gorm"
)

type WorkOrderDB struct {
	gorm.Model
	ID               string             `json:"id" gorm:"size:36;unique;primary_key;not null"`
	CustomerID       string             `json:"-" gorm:"foreignKey:id;not null"`
	Customer         customers.Customer `json:"customer" gorm:"foreignKey:CustomerID"`
	Title            string             `json:"title" gorm:"not null;size:255"`
	PlannedDateBegin time.Time          `json:"planned_date_begin" gorm:"not null"`
	PlannedDateEnd   time.Time          `json:"planned_date_end" gorm:"not null"`
	Status           string             `json:"status" gorm:"not null"`
	CreatedAt        time.Time          `json:"created_at" gorm:"not null"`
}

func (w WorkOrderDB) ToWorkOrder() WorkOrder {
	return WorkOrder{
		ID:               ID(w.ID),
		CustomerID:       customers.ID(w.CustomerID),
		Customer:         w.Customer,
		Title:            w.Title,
		PlannedDateBegin: w.PlannedDateBegin,
		PlannedDateEnd:   w.PlannedDateEnd,
		Status:           Status(w.Status),
		CreatedAt:        w.CreatedAt,
	}
}
