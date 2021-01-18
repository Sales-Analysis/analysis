package abc

import (
	"errors"
	"fmt"
)

type sizeError struct {
	err error
}

func (s *sizeError) Error() string {
	return fmt.Sprintf("%v", s.err)
}

func sizeErr(err string) *sizeError {
	return &sizeError{createError(err)}
}

func createError(context string) error {
	return errors.New(context)
}
