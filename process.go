package osi

import "os"

type Process struct {
	*os.Process
}

func FindProcess(pid int) (*Process, error) {
	p, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	return &Process{p}, nil
}

func StartProcess(name string, argv []string, attr *os.ProcAttr) (*Process, error) {
	p, err := os.StartProcess(name, argv, attr)
	if err != nil {
		return nil, err
	}
	return &Process{p}, nil
}

func (p *Process) Kill() error {
	return p.Process.Kill()
}

func (p *Process) Release() error {
	return p.Process.Release()
}

func (p *Process) Signal(sig os.Signal) error {
	return p.Process.Signal(sig)
}

func (p *Process) Wait() (*os.ProcessState, error) {
	return p.Process.Wait()
}
