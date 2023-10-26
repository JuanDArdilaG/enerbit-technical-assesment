package persistence

import (
	"fmt"

	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/apperrors"
	"gorm.io/gorm"
)

type SQLiteDB[ID ~string, T interface{}] struct {
	*gorm.DB
}

func NewSQLiteDB[ID ~string, T interface{}](db *gorm.DB) (Repository[ID, T], error) {
	return &SQLiteDB[ID, T]{DB: db}, nil
}

func (s SQLiteDB[ID, T]) GetAll() ([]T, error) {
	var items []T
	result := s.DB.Find(&items)
	return items, result.Error
}

func (s SQLiteDB[ID, T]) GetByID(id ID) (T, error) {
	var item T
	result := s.DB.First(&item, id)
	if result.Error == gorm.ErrRecordNotFound {
		return item, apperrors.NewErrorNotFound(fmt.Sprintf("%T", item), string(id))
	}
	return item, result.Error
}

func (s SQLiteDB[ID, T]) Create(item T) error {
	result := s.DB.Create(&item)
	return result.Error
}

func (s SQLiteDB[ID, T]) Update(id ID, item T) error {
	_, err := s.GetByID(id)
	if err != nil {
		return err
	}
	result := s.DB.Save(&item)
	return result.Error
}

func (s SQLiteDB[ID, T]) Delete(id ID) error {
	var item T
	result := s.DB.Delete(&item, id)
	if result.Error == gorm.ErrRecordNotFound {
		return apperrors.NewErrorNotFound(fmt.Sprintf("%T", item), string(id))
	}
	return result.Error
}
