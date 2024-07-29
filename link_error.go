package osi

import "os"

type LinkError struct {
	*os.LinkError
}

func (e *LinkError) Error() string {
	return e.LinkError.Error()
}

func (e *LinkError) Unwrap() error {
	return e.LinkError.Unwrap()
}
