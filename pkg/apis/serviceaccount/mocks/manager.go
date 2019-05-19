// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/51st-state/api/pkg/apis/serviceaccount"
)

type FakeManager struct {
	GetStub        func(context.Context, serviceaccount.Identifier) (serviceaccount.Complete, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 serviceaccount.Identifier
	}
	getReturns struct {
		result1 serviceaccount.Complete
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 serviceaccount.Complete
		result2 error
	}
	CreateStub        func(context.Context, serviceaccount.Incomplete) (serviceaccount.Complete, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 serviceaccount.Incomplete
	}
	createReturns struct {
		result1 serviceaccount.Complete
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 serviceaccount.Complete
		result2 error
	}
	UpdateStub        func(context.Context, serviceaccount.Complete) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 serviceaccount.Complete
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(context.Context, serviceaccount.Identifier) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 serviceaccount.Identifier
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) Get(arg1 context.Context, arg2 serviceaccount.Identifier) (serviceaccount.Complete, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 serviceaccount.Identifier
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *FakeManager) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeManager) GetArgsForCall(i int) (context.Context, serviceaccount.Identifier) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2
}

func (fake *FakeManager) GetReturns(result1 serviceaccount.Complete, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 serviceaccount.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetReturnsOnCall(i int, result1 serviceaccount.Complete, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 serviceaccount.Complete
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 serviceaccount.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) Create(arg1 context.Context, arg2 serviceaccount.Incomplete) (serviceaccount.Complete, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 serviceaccount.Incomplete
	}{arg1, arg2})
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeManager) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeManager) CreateArgsForCall(i int) (context.Context, serviceaccount.Incomplete) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2
}

func (fake *FakeManager) CreateReturns(result1 serviceaccount.Complete, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 serviceaccount.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) CreateReturnsOnCall(i int, result1 serviceaccount.Complete, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 serviceaccount.Complete
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 serviceaccount.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) Update(arg1 context.Context, arg2 serviceaccount.Complete) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 serviceaccount.Complete
	}{arg1, arg2})
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateReturns.result1
}

func (fake *FakeManager) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeManager) UpdateArgsForCall(i int) (context.Context, serviceaccount.Complete) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1, fake.updateArgsForCall[i].arg2
}

func (fake *FakeManager) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UpdateReturnsOnCall(i int, result1 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Delete(arg1 context.Context, arg2 serviceaccount.Identifier) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 serviceaccount.Identifier
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeManager) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeManager) DeleteArgsForCall(i int) (context.Context, serviceaccount.Identifier) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1, fake.deleteArgsForCall[i].arg2
}

func (fake *FakeManager) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
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

var _ serviceaccount.Manager = new(FakeManager)
