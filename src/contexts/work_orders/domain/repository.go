package workorders

import (
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/persistence"
)

type Repository interface {
	persistence.Repository[ID, WorkOrderDB]
	GetAllByCustomerID(customerID customers.ID) ([]WorkOrderDB, error)
	GetDoneByCustomerID(customerID customers.ID) ([]WorkOrderDB, error)
}
