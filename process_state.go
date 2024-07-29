package osi

import (
	"os"
	"time"
)

type ProcessState struct {
	*os.ProcessState
}

func (ps *ProcessState) ExitCode() int {
	return ps.ProcessState.ExitCode()
}

func (ps *ProcessState) Exited() bool {
	return ps.ProcessState.Exited()
}

func (ps *ProcessState) Pid() int {
	return ps.ProcessState.Pid()
}

func (ps *ProcessState) String() string {
	return ps.ProcessState.String()
}

func (ps *ProcessState) Success() bool {
	return ps.ProcessState.Success()
}

func (ps *ProcessState) Sys() interface{} {
	return ps.ProcessState.Sys()
}

func (ps *ProcessState) SysUsage() interface{} {
	return ps.ProcessState.SysUsage()
}

func (ps *ProcessState) SystemTime() time.Duration {
	return ps.ProcessState.SystemTime()
}

func (ps *ProcessState) UserTime() time.Duration {
	return ps.ProcessState.UserTime()
}
