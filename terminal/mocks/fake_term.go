// This file was generated by counterfeiter
package mocks

import (
	"sync"

	"github.com/cloudfoundry-incubator/ltc/terminal"
	"github.com/docker/docker/pkg/term"
)

type FakeTerm struct {
	SaveStateStub        func(fd uintptr) (*term.State, error)
	saveStateMutex       sync.RWMutex
	saveStateArgsForCall []struct {
		fd uintptr
	}
	saveStateReturns struct {
		result1 *term.State
		result2 error
	}
	RestoreTerminalStub        func(fd uintptr, state *term.State) error
	restoreTerminalMutex       sync.RWMutex
	restoreTerminalArgsForCall []struct {
		fd    uintptr
		state *term.State
	}
	restoreTerminalReturns struct {
		result1 error
	}
	DisableEchoStub        func(fd uintptr, state *term.State) error
	disableEchoMutex       sync.RWMutex
	disableEchoArgsForCall []struct {
		fd    uintptr
		state *term.State
	}
	disableEchoReturns struct {
		result1 error
	}
}

func (fake *FakeTerm) SaveState(fd uintptr) (*term.State, error) {
	fake.saveStateMutex.Lock()
	fake.saveStateArgsForCall = append(fake.saveStateArgsForCall, struct {
		fd uintptr
	}{fd})
	fake.saveStateMutex.Unlock()
	if fake.SaveStateStub != nil {
		return fake.SaveStateStub(fd)
	} else {
		return fake.saveStateReturns.result1, fake.saveStateReturns.result2
	}
}

func (fake *FakeTerm) SaveStateCallCount() int {
	fake.saveStateMutex.RLock()
	defer fake.saveStateMutex.RUnlock()
	return len(fake.saveStateArgsForCall)
}

func (fake *FakeTerm) SaveStateArgsForCall(i int) uintptr {
	fake.saveStateMutex.RLock()
	defer fake.saveStateMutex.RUnlock()
	return fake.saveStateArgsForCall[i].fd
}

func (fake *FakeTerm) SaveStateReturns(result1 *term.State, result2 error) {
	fake.SaveStateStub = nil
	fake.saveStateReturns = struct {
		result1 *term.State
		result2 error
	}{result1, result2}
}

func (fake *FakeTerm) RestoreTerminal(fd uintptr, state *term.State) error {
	fake.restoreTerminalMutex.Lock()
	fake.restoreTerminalArgsForCall = append(fake.restoreTerminalArgsForCall, struct {
		fd    uintptr
		state *term.State
	}{fd, state})
	fake.restoreTerminalMutex.Unlock()
	if fake.RestoreTerminalStub != nil {
		return fake.RestoreTerminalStub(fd, state)
	} else {
		return fake.restoreTerminalReturns.result1
	}
}

func (fake *FakeTerm) RestoreTerminalCallCount() int {
	fake.restoreTerminalMutex.RLock()
	defer fake.restoreTerminalMutex.RUnlock()
	return len(fake.restoreTerminalArgsForCall)
}

func (fake *FakeTerm) RestoreTerminalArgsForCall(i int) (uintptr, *term.State) {
	fake.restoreTerminalMutex.RLock()
	defer fake.restoreTerminalMutex.RUnlock()
	return fake.restoreTerminalArgsForCall[i].fd, fake.restoreTerminalArgsForCall[i].state
}

func (fake *FakeTerm) RestoreTerminalReturns(result1 error) {
	fake.RestoreTerminalStub = nil
	fake.restoreTerminalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTerm) DisableEcho(fd uintptr, state *term.State) error {
	fake.disableEchoMutex.Lock()
	fake.disableEchoArgsForCall = append(fake.disableEchoArgsForCall, struct {
		fd    uintptr
		state *term.State
	}{fd, state})
	fake.disableEchoMutex.Unlock()
	if fake.DisableEchoStub != nil {
		return fake.DisableEchoStub(fd, state)
	} else {
		return fake.disableEchoReturns.result1
	}
}

func (fake *FakeTerm) DisableEchoCallCount() int {
	fake.disableEchoMutex.RLock()
	defer fake.disableEchoMutex.RUnlock()
	return len(fake.disableEchoArgsForCall)
}

func (fake *FakeTerm) DisableEchoArgsForCall(i int) (uintptr, *term.State) {
	fake.disableEchoMutex.RLock()
	defer fake.disableEchoMutex.RUnlock()
	return fake.disableEchoArgsForCall[i].fd, fake.disableEchoArgsForCall[i].state
}

func (fake *FakeTerm) DisableEchoReturns(result1 error) {
	fake.DisableEchoStub = nil
	fake.disableEchoReturns = struct {
		result1 error
	}{result1}
}

var _ terminal.Term = new(FakeTerm)
