package tinfer

import "fmt"

// InferenceError is an type inference error.
type InferenceError struct {
	message,
	location string
}

func newInferenceError(m, l string) error {
	return InferenceError{m, l}
}

func (e InferenceError) Error() string {
	return fmt.Sprintf("%v: %v", e.location, e.message)
}
