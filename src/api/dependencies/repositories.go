package dependencies

import (
	"os"

	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	customersinfra "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/infrastructure"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
	workordersinfra "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/infrastructure"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repositories struct {
	WorkOrders workorders.Repository
	Customers  customers.Repository
}

func BuildRepositories() (*Repositories, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB_NAME")))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&workorders.WorkOrderDB{}, &customers.CustomerDB{})
	woRepo, err := workordersinfra.NewSQLiteRepository(db)
	if err != nil {
		return nil, err
	}
	customersRepo, err := customersinfra.NewSQLiteRepository(db)
	if err != nil {
		return nil, err
	}
	return &Repositories{
		WorkOrders: woRepo,
		Customers:  customersRepo,
	}, nil
}
