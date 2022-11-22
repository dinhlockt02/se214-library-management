package utils

import (
	"fmt"
	"time"
)

const (
	TimeLayout = "2006-01-02T15:04:05.000Z"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(TimeLayout))
	return []byte(stamp), nil
}
