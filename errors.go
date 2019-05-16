package proverbs

import "fmt"

// ErrNthProverbNotFound is returned when the requested nth proverb isn't found in the dictionary of proverbs.
type ErrNthProverbNotFound struct {
	N int
}

func (e *ErrNthProverbNotFound) Error() string {
	return fmt.Sprintf("proverb %d not found", e.N)
}
