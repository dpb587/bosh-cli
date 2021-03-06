// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/release"
)

type FakeReader struct {
	ReadStub        func(string) (release.Release, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 string
	}
	readReturns struct {
		result1 release.Release
		result2 error
	}
}

func (fake *FakeReader) Read(arg1 string) (release.Release, error) {
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeReader) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeReader) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].arg1
}

func (fake *FakeReader) ReadReturns(result1 release.Release, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 release.Release
		result2 error
	}{result1, result2}
}

var _ release.Reader = new(FakeReader)
