package workordersinfra

import (
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/persistence"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	*persistence.SQLiteDB[workorders.ID, workorders.WorkOrderDB]
}

func NewSQLiteRepository(db *gorm.DB) (workorders.Repository, error) {
	sqlitedb, err := persistence.NewSQLiteDB[workorders.ID, workorders.WorkOrderDB](db)
	if err != nil {
		return nil, err
	}
	return &SQLiteRepository{
		SQLiteDB: (sqlitedb).(*persistence.SQLiteDB[workorders.ID, workorders.WorkOrderDB]),
	}, nil
}

func (r *SQLiteRepository) GetAll() ([]workorders.WorkOrderDB, error) {
	var items []workorders.WorkOrderDB
	result := r.SQLiteDB.DB.Preload("Customer").Find(&items)
	return items, result.Error
}

func (r *SQLiteRepository) GetByID(id workorders.ID) (workorders.WorkOrderDB, error) {
	var item workorders.WorkOrderDB
	result := r.SQLiteDB.DB.Where("id = ?", id).Preload("Customer").First(&item)
	return item, result.Error
}

func (r *SQLiteRepository) GetAllByCustomerID(customerID customers.ID) ([]workorders.WorkOrderDB, error) {
	var items []workorders.WorkOrderDB
	result := r.SQLiteDB.DB.Where(
		"customer_id = ?", customerID,
	).Preload("Customer").Find(&items)
	return items, result.Error
}

func (r *SQLiteRepository) GetDoneByCustomerID(customerID customers.ID) ([]workorders.WorkOrderDB, error) {
	var items []workorders.WorkOrderDB
	result := r.SQLiteDB.DB.Where(
		"customer_id = ? AND status = ?", customerID, workorders.StatusDone,
	).Preload("Customer").Find(&items)
	return items, result.Error
}
