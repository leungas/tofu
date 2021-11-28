package exceptions

import "fmt"

type EntityNotFoundException struct {
	Entity string
	Value  string
}

func (e *EntityNotFoundException) Error() string {
	return fmt.Sprintf("Request enitty %s with identity value %s is not available", e.Entity, e.Value)
}
