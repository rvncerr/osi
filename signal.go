package osi

import "syscall"

type Signal interface {
	String() string
	Signal()
}

var (
	Interrupt Signal = syscall.SIGINT
	Kill      Signal = syscall.SIGKILL
)
