package entity

import (
	"github.com/rs/xid"
)

type ID = xid.ID

func NewID() ID {
	return xid.New()
}

func StringToID(s string) (*ID, error) {
	id, err := xid.FromString(s)
	return &id, err
}
