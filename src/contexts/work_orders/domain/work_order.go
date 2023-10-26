package workorders

import (
	"errors"
	"time"

	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/apperrors"
	"github.com/go-playground/validator/v10"
)

var ErrInvalidPlannedDate = errors.New("invalid planned date. the difference between planned date begin and planned date end must be less than 2 hours")

type WorkOrder struct {
	ID               ID                 `json:"id"`
	CustomerID       customers.ID       `json:"customer_id" validate:"required,uuid"`
	Customer         customers.Customer `json:"customer" validate:"-"`
	Title            string             `json:"title" validate:"min=5,max=255"`
	PlannedDateBegin time.Time          `json:"planned_date_begin"`
	PlannedDateEnd   time.Time          `json:"planned_date_end"`
	Status           Status             `json:"status" validate:"oneof=new done cancelled"`
	CreatedAt        time.Time          `json:"created_at"`
}

func New(
	customer customers.Customer, title string, plannedDateBegin time.Time, plannedDateEnd time.Time,
) (*WorkOrder, error) {
	id, err := NewRandomID()
	if err != nil {
		return nil, err
	}
	order := &WorkOrder{
		ID:               id,
		CustomerID:       customer.ID,
		Customer:         customer,
		Title:            title,
		PlannedDateBegin: plannedDateBegin,
		PlannedDateEnd:   plannedDateEnd,
		Status:           StatusNew,
		CreatedAt:        time.Now(),
	}
	err = order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func validPlannedDate(plannedDateBegin time.Time, plannedDateEnd time.Time) error {
	datesDiff := plannedDateEnd.Sub(plannedDateBegin)
	if datesDiff.Seconds() > 0 && datesDiff.Hours() <= 2 {
		return nil
	}
	return ErrInvalidPlannedDate
}

func (wo *WorkOrder) ToWorkOrderDB() WorkOrderDB {
	return WorkOrderDB{
		ID:               string(wo.ID),
		CustomerID:       string(wo.Customer.ID),
		Customer:         wo.Customer,
		Title:            wo.Title,
		PlannedDateBegin: wo.PlannedDateBegin,
		PlannedDateEnd:   wo.PlannedDateEnd,
		Status:           string(wo.Status),
		CreatedAt:        wo.CreatedAt,
	}
}

func (wo *WorkOrder) Finish() {
	wo.Status = StatusDone
}

func (wo *WorkOrder) ChangeStatus(s Status) {
	wo.Status = s
}

func (wo *WorkOrder) MatchStatus(s Status) bool {
	return string(wo.Status) == string(s)
}

func (wo *WorkOrder) MatchPlannedDate(plannedDateBegin time.Time, plannedDateEnd time.Time) bool {
	return wo.PlannedDateBegin.Equal(plannedDateBegin) && wo.PlannedDateEnd.Equal(plannedDateEnd)
}

func (wo *WorkOrder) Validate() error {
	if err := validPlannedDate(wo.PlannedDateBegin, wo.PlannedDateEnd); err != nil {
		return apperrors.NewBadRequest("workorder", err.Error())
	}
	validate := validator.New()
	err := validate.Struct(wo)
	if err != nil {
		return apperrors.NewBadRequest("workorder", err.Error())
	}
	return nil
}
