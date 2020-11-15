package apierrors

import "fmt"

// NotFountErr error to handle record not found
type NotFountErr struct {
	ID string
}

func (e *NotFountErr) Error() string {
	return fmt.Sprintf(notFount, e.ID)
}
