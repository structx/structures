package structs

import (
	"fmt"
)

// ErrNotFound node not found
type ErrNotFound[T int | int64 | int32 | uint | uint8 | uint32 | uint64] struct {
	Payload T
}

// Error stringify error message
func (enf *ErrNotFound[T]) Error() string {
	return fmt.Sprintf("failed to find node with payload %d", enf.Payload)
}

// ErrNotInitialized structure is empty
type ErrNotInitialized struct {
	// empty struct for error wrapping
}

// Error stringify error message
func (eni *ErrNotInitialized) Error() string {
	return "root not is not initalized add values to structure before calling queries/mutations"
}
