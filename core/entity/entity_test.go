package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	newId := NewID()

	assert.NotNil(t, newId, "invalid id: id is nil")

}

func TestStringToID(t *testing.T) {
	id := NewID()
	idString := id.String()

	rs, err := StringToID(idString)

	if assert.Nil(t, err) {
		assert.Equal(t, rs.String(), id.String())
	}

}
