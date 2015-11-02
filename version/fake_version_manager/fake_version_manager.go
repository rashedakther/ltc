// This file was generated by counterfeiter
package fake_version_manager

import (
	"sync"

	config_package "github.com/cloudfoundry-incubator/ltc/config"
	"github.com/cloudfoundry-incubator/ltc/version"
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
	LtcVersionStub        func() string
	ltcVersionMutex       sync.RWMutex
	ltcVersionArgsForCall []struct{}
	ltcVersionReturns     struct {
		result1 string
	}
	LtcMatchesServerStub        func() (bool, error)
	ltcMatchesServerMutex       sync.RWMutex
	ltcMatchesServerArgsForCall []struct{}
	ltcMatchesServerReturns     struct {
		result1 bool
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

func (fake *FakeVersionManager) LtcVersion() string {
	fake.ltcVersionMutex.Lock()
	fake.ltcVersionArgsForCall = append(fake.ltcVersionArgsForCall, struct{}{})
	fake.ltcVersionMutex.Unlock()
	if fake.LtcVersionStub != nil {
		return fake.LtcVersionStub()
	} else {
		return fake.ltcVersionReturns.result1
	}
}

func (fake *FakeVersionManager) LtcVersionCallCount() int {
	fake.ltcVersionMutex.RLock()
	defer fake.ltcVersionMutex.RUnlock()
	return len(fake.ltcVersionArgsForCall)
}

func (fake *FakeVersionManager) LtcVersionReturns(result1 string) {
	fake.LtcVersionStub = nil
	fake.ltcVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVersionManager) LtcMatchesServer() (bool, error) {
	fake.ltcMatchesServerMutex.Lock()
	fake.ltcMatchesServerArgsForCall = append(fake.ltcMatchesServerArgsForCall, struct{}{})
	fake.ltcMatchesServerMutex.Unlock()
	if fake.LtcMatchesServerStub != nil {
		return fake.LtcMatchesServerStub()
	} else {
		return fake.ltcMatchesServerReturns.result1, fake.ltcMatchesServerReturns.result2
	}
}

func (fake *FakeVersionManager) LtcMatchesServerCallCount() int {
	fake.ltcMatchesServerMutex.RLock()
	defer fake.ltcMatchesServerMutex.RUnlock()
	return len(fake.ltcMatchesServerArgsForCall)
}

func (fake *FakeVersionManager) LtcMatchesServerReturns(result1 bool, result2 error) {
	fake.LtcMatchesServerStub = nil
	fake.ltcMatchesServerReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

var _ version.VersionManager = new(FakeVersionManager)