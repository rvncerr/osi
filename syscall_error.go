package osi

import "os"

type SyscallError struct {
	*os.SyscallError
}

func (e *SyscallError) Error() string {
	return e.SyscallError.Error()
}

func (e *SyscallError) Timeout() bool {
	return e.SyscallError.Timeout()
}

func (e *SyscallError) Unwrap() error {
	return e.SyscallError.Unwrap()
}
