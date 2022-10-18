package common

import (
	"errors"
	"github.com/sony/sonyflake"
)

type IDGenerator interface {
	NextID() (uint64, error)
}

func NewIDGenerator() (IDGenerator, error) {
	var st sonyflake.Settings
	sf := sonyflake.NewSonyflake(st)
	if sf == nil {
		return nil, errors.New("sonyflake not created")
	}
	return sf, nil
}
