package entity

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(
		uuid.New(),
	)
}

func StringToID(s string) (*ID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return nil, coreerror.NewBadRequestError("Invalid id string", err)
	}
	return &id, nil
}
