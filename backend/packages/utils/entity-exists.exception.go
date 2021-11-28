package exceptions

import "fmt"

type EntityExistsExcpetion struct {
	Entity   string
	Identity string
	Value    string
}

func (e *EntityExistsExcpetion) Error() string {
	return fmt.Sprintf("Requested entity %s with identity %s and value %s already exists", e.Entity, e.Identity, e.Value)
}
