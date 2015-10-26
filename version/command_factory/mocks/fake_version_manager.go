// This file was generated by counterfeiter
package mocks

import (
	"sync"

	config_package "github.com/cloudfoundry-incubator/ltc/config"
	"github.com/cloudfoundry-incubator/ltc/version"
	"github.com/cloudfoundry-incubator/ltc/version/command_factory"
)

type FakeVersionManager struct {
	SyncLTCStub        func(ltcPath string, arch string, config *config_package.Config) error
	syncLTCMutex       sync.RWMutex
	syncLTCArgsForCall []struct {
		ltcPath string
		arch    string
		config  *config_package.Config
	}
	syncLTCReturns struct {
		result1 error
	}
	ServerVersionsStub        func() (version.ServerVersions, error)
	serverVersionsMutex       sync.RWMutex
	serverVersionsArgsForCall []struct{}
	serverVersionsReturns     struct {
		result1 version.ServerVersions
		result2 error
	}
}

func (fake *FakeVersionManager) SyncLTC(ltcPath string, arch string, config *config_package.Config) error {
	fake.syncLTCMutex.Lock()
	fake.syncLTCArgsForCall = append(fake.syncLTCArgsForCall, struct {
		ltcPath string
		arch    string
		config  *config_package.Config
	}{ltcPath, arch, config})
	fake.syncLTCMutex.Unlock()
	if fake.SyncLTCStub != nil {
		return fake.SyncLTCStub(ltcPath, arch, config)
	} else {
		return fake.syncLTCReturns.result1
	}
}

func (fake *FakeVersionManager) SyncLTCCallCount() int {
	fake.syncLTCMutex.RLock()
	defer fake.syncLTCMutex.RUnlock()
	return len(fake.syncLTCArgsForCall)
}

func (fake *FakeVersionManager) SyncLTCArgsForCall(i int) (string, string, *config_package.Config) {
	fake.syncLTCMutex.RLock()
	defer fake.syncLTCMutex.RUnlock()
	return fake.syncLTCArgsForCall[i].ltcPath, fake.syncLTCArgsForCall[i].arch, fake.syncLTCArgsForCall[i].config
}

func (fake *FakeVersionManager) SyncLTCReturns(result1 error) {
	fake.SyncLTCStub = nil
	fake.syncLTCReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVersionManager) ServerVersions() (version.ServerVersions, error) {
	fake.serverVersionsMutex.Lock()
	fake.serverVersionsArgsForCall = append(fake.serverVersionsArgsForCall, struct{}{})
	fake.serverVersionsMutex.Unlock()
	if fake.ServerVersionsStub != nil {
		return fake.ServerVersionsStub()
	} else {
		return fake.serverVersionsReturns.result1, fake.serverVersionsReturns.result2
	}
}

func (fake *FakeVersionManager) ServerVersionsCallCount() int {
	fake.serverVersionsMutex.RLock()
	defer fake.serverVersionsMutex.RUnlock()
	return len(fake.serverVersionsArgsForCall)
}

func (fake *FakeVersionManager) ServerVersionsReturns(result1 version.ServerVersions, result2 error) {
	fake.ServerVersionsStub = nil
	fake.serverVersionsReturns = struct {
		result1 version.ServerVersions
		result2 error
	}{result1, result2}
}

var _ command_factory.VersionManager = new(FakeVersionManager)
