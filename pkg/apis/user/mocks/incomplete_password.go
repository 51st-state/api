// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/51st-state/api/pkg/apis/user"
)

type FakeIncompletePassword struct {
	PasswordStub        func() string
	passwordMutex       sync.RWMutex
	passwordArgsForCall []struct{}
	passwordReturns     struct {
		result1 string
	}
	passwordReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIncompletePassword) Password() string {
	fake.passwordMutex.Lock()
	ret, specificReturn := fake.passwordReturnsOnCall[len(fake.passwordArgsForCall)]
	fake.passwordArgsForCall = append(fake.passwordArgsForCall, struct{}{})
	fake.recordInvocation("Password", []interface{}{})
	fake.passwordMutex.Unlock()
	if fake.PasswordStub != nil {
		return fake.PasswordStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.passwordReturns.result1
}

func (fake *FakeIncompletePassword) PasswordCallCount() int {
	fake.passwordMutex.RLock()
	defer fake.passwordMutex.RUnlock()
	return len(fake.passwordArgsForCall)
}

func (fake *FakeIncompletePassword) PasswordReturns(result1 string) {
	fake.PasswordStub = nil
	fake.passwordReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIncompletePassword) PasswordReturnsOnCall(i int, result1 string) {
	fake.PasswordStub = nil
	if fake.passwordReturnsOnCall == nil {
		fake.passwordReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.passwordReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeIncompletePassword) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.passwordMutex.RLock()
	defer fake.passwordMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIncompletePassword) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ user.IncompletePassword = new(FakeIncompletePassword)
