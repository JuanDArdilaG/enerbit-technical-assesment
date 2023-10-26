package domain

import "github.com/google/uuid"

type UUID string

func NewUUID(val string) (UUID, error) {
	uuid, err := uuid.Parse(val)
	if err != nil {
		return "", err
	}
	return UUID(uuid.String()), nil
}

func NewRandomUUID() (UUID, error) {
	uuid, err := uuid.NewRandom()
	return UUID(uuid.String()), err
}

func (u *UUID) String() string {
	if u == nil {
		return ""
	}
	return string(*u)
}
