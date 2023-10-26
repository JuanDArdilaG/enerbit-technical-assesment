package workorders

import "github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/domain"

type ID domain.UUID

func NewID(val string) (ID, error) {
	uuid, err := domain.NewUUID(val)
	return ID(uuid), err
}

func NewRandomID() (ID, error) {
	uuid, err := domain.NewRandomUUID()
	return ID(uuid), err
}
