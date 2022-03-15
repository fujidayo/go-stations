package model

import "fmt"

type ErrNotFound struct {
	When string
	What string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s, %s", e.When, e.What)
}
