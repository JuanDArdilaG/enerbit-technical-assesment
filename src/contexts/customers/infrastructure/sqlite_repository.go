package customersinfra

import (
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/persistence"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	*persistence.SQLiteDB[customers.ID, customers.CustomerDB]
}

func NewSQLiteRepository(db *gorm.DB) (customers.Repository, error) {
	sqlitedb, err := persistence.NewSQLiteDB[customers.ID, customers.CustomerDB](db)
	if err != nil {
		return nil, err
	}
	return &SQLiteRepository{
		SQLiteDB: sqlitedb.(*persistence.SQLiteDB[customers.ID, customers.CustomerDB]),
	}, nil
}

func (r *SQLiteRepository) GetByCustomerID(customerID string) ([]customers.CustomerDB, error) {
	var items []customers.CustomerDB
	result := r.SQLiteDB.DB.Where("customer_id = ?", customerID).Find(&items)
	return items, result.Error
}

func (r *SQLiteRepository) GetActive() ([]customers.CustomerDB, error) {
	var items []customers.CustomerDB
	result := r.SQLiteDB.DB.Where("is_active = ?", true).Find(&items)
	return items, result.Error
}
