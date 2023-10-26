package persistence

type Repository[ID any, T any] interface {
	GetAll() ([]T, error)
	GetByID(ID) (T, error)
	Create(T) error
	Update(ID, T) error
	Delete(ID) error
}
