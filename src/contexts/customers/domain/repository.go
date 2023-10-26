package customers

import "github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/persistence"

type Repository interface {
	persistence.Repository[ID, CustomerDB]
	GetActive() ([]CustomerDB, error)
}
