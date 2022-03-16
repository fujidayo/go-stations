package model

import (
	"fmt"
	"time"
)

type ErrNotFound struct {
	When time.Time
	What string
}

func (e *ErrNotFound) Error() error {
	return fmt.Errorf("%s, %s", e.When, e.What)
}

func Run() *ErrNotFound {
	return &ErrNotFound{
		time.Now(),
		"not updated",
	}
}
