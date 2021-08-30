package abc

import (
	"errors"
	"fmt"
)

func createError(context string) error {
	return errors.New(context)
}

type sizeError struct {
	err error
}

func (s *sizeError) Error() string {
	return fmt.Sprintf("%v", s.err)
}

func sizeErr(err string) *sizeError {
	return &sizeError{createError(err)}
}

type negativeValueError struct {
	err error
}

func (n *negativeValueError) Error() string {
	return fmt.Sprintf("%v", n.err)
}

func negativeValueErr(err string) *negativeValueError {
	return &negativeValueError{createError(err)}
}