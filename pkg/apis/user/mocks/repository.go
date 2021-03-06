// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/51st-state/api/pkg/apis/user"
)

type FakeRepository struct {
	GetStub        func(context.Context, user.Identifier) (user.Complete, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 user.Identifier
	}
	getReturns struct {
		result1 user.Complete
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 user.Complete
		result2 error
	}
	GetByGameSerialHashStub        func(context.Context, string) (user.Complete, error)
	getByGameSerialHashMutex       sync.RWMutex
	getByGameSerialHashArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getByGameSerialHashReturns struct {
		result1 user.Complete
		result2 error
	}
	getByGameSerialHashReturnsOnCall map[int]struct {
		result1 user.Complete
		result2 error
	}
	GetByWCFUserIDStub        func(context.Context, user.WCFUserID) (user.Complete, error)
	getByWCFUserIDMutex       sync.RWMutex
	getByWCFUserIDArgsForCall []struct {
		arg1 context.Context
		arg2 user.WCFUserID
	}
	getByWCFUserIDReturns struct {
		result1 user.Complete
		result2 error
	}
	getByWCFUserIDReturnsOnCall map[int]struct {
		result1 user.Complete
		result2 error
	}
	CreateStub        func(context.Context, user.Incomplete) (user.Complete, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 user.Incomplete
	}
	createReturns struct {
		result1 user.Complete
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 user.Complete
		result2 error
	}
	UpdateStub        func(context.Context, user.Complete) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 user.Complete
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(context.Context, user.Identifier) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 user.Identifier
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

func (fake *FakeRepository) Get(arg1 context.Context, arg2 user.Identifier) (user.Complete, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 user.Identifier
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

func (fake *FakeRepository) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeRepository) GetArgsForCall(i int) (context.Context, user.Identifier) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2
}

func (fake *FakeRepository) GetReturns(result1 user.Complete, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 user.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetReturnsOnCall(i int, result1 user.Complete, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 user.Complete
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 user.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetByGameSerialHash(arg1 context.Context, arg2 string) (user.Complete, error) {
	fake.getByGameSerialHashMutex.Lock()
	ret, specificReturn := fake.getByGameSerialHashReturnsOnCall[len(fake.getByGameSerialHashArgsForCall)]
	fake.getByGameSerialHashArgsForCall = append(fake.getByGameSerialHashArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetByGameSerialHash", []interface{}{arg1, arg2})
	fake.getByGameSerialHashMutex.Unlock()
	if fake.GetByGameSerialHashStub != nil {
		return fake.GetByGameSerialHashStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getByGameSerialHashReturns.result1, fake.getByGameSerialHashReturns.result2
}

func (fake *FakeRepository) GetByGameSerialHashCallCount() int {
	fake.getByGameSerialHashMutex.RLock()
	defer fake.getByGameSerialHashMutex.RUnlock()
	return len(fake.getByGameSerialHashArgsForCall)
}

func (fake *FakeRepository) GetByGameSerialHashArgsForCall(i int) (context.Context, string) {
	fake.getByGameSerialHashMutex.RLock()
	defer fake.getByGameSerialHashMutex.RUnlock()
	return fake.getByGameSerialHashArgsForCall[i].arg1, fake.getByGameSerialHashArgsForCall[i].arg2
}

func (fake *FakeRepository) GetByGameSerialHashReturns(result1 user.Complete, result2 error) {
	fake.GetByGameSerialHashStub = nil
	fake.getByGameSerialHashReturns = struct {
		result1 user.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetByGameSerialHashReturnsOnCall(i int, result1 user.Complete, result2 error) {
	fake.GetByGameSerialHashStub = nil
	if fake.getByGameSerialHashReturnsOnCall == nil {
		fake.getByGameSerialHashReturnsOnCall = make(map[int]struct {
			result1 user.Complete
			result2 error
		})
	}
	fake.getByGameSerialHashReturnsOnCall[i] = struct {
		result1 user.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetByWCFUserID(arg1 context.Context, arg2 user.WCFUserID) (user.Complete, error) {
	fake.getByWCFUserIDMutex.Lock()
	ret, specificReturn := fake.getByWCFUserIDReturnsOnCall[len(fake.getByWCFUserIDArgsForCall)]
	fake.getByWCFUserIDArgsForCall = append(fake.getByWCFUserIDArgsForCall, struct {
		arg1 context.Context
		arg2 user.WCFUserID
	}{arg1, arg2})
	fake.recordInvocation("GetByWCFUserID", []interface{}{arg1, arg2})
	fake.getByWCFUserIDMutex.Unlock()
	if fake.GetByWCFUserIDStub != nil {
		return fake.GetByWCFUserIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getByWCFUserIDReturns.result1, fake.getByWCFUserIDReturns.result2
}

func (fake *FakeRepository) GetByWCFUserIDCallCount() int {
	fake.getByWCFUserIDMutex.RLock()
	defer fake.getByWCFUserIDMutex.RUnlock()
	return len(fake.getByWCFUserIDArgsForCall)
}

func (fake *FakeRepository) GetByWCFUserIDArgsForCall(i int) (context.Context, user.WCFUserID) {
	fake.getByWCFUserIDMutex.RLock()
	defer fake.getByWCFUserIDMutex.RUnlock()
	return fake.getByWCFUserIDArgsForCall[i].arg1, fake.getByWCFUserIDArgsForCall[i].arg2
}

func (fake *FakeRepository) GetByWCFUserIDReturns(result1 user.Complete, result2 error) {
	fake.GetByWCFUserIDStub = nil
	fake.getByWCFUserIDReturns = struct {
		result1 user.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetByWCFUserIDReturnsOnCall(i int, result1 user.Complete, result2 error) {
	fake.GetByWCFUserIDStub = nil
	if fake.getByWCFUserIDReturnsOnCall == nil {
		fake.getByWCFUserIDReturnsOnCall = make(map[int]struct {
			result1 user.Complete
			result2 error
		})
	}
	fake.getByWCFUserIDReturnsOnCall[i] = struct {
		result1 user.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Create(arg1 context.Context, arg2 user.Incomplete) (user.Complete, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 user.Incomplete
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

func (fake *FakeRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeRepository) CreateArgsForCall(i int) (context.Context, user.Incomplete) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2
}

func (fake *FakeRepository) CreateReturns(result1 user.Complete, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 user.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) CreateReturnsOnCall(i int, result1 user.Complete, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 user.Complete
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 user.Complete
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Update(arg1 context.Context, arg2 user.Complete) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 user.Complete
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

func (fake *FakeRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeRepository) UpdateArgsForCall(i int) (context.Context, user.Complete) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1, fake.updateArgsForCall[i].arg2
}

func (fake *FakeRepository) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) UpdateReturnsOnCall(i int, result1 error) {
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

func (fake *FakeRepository) Delete(arg1 context.Context, arg2 user.Identifier) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 user.Identifier
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

func (fake *FakeRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeRepository) DeleteArgsForCall(i int) (context.Context, user.Identifier) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1, fake.deleteArgsForCall[i].arg2
}

func (fake *FakeRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getByGameSerialHashMutex.RLock()
	defer fake.getByGameSerialHashMutex.RUnlock()
	fake.getByWCFUserIDMutex.RLock()
	defer fake.getByWCFUserIDMutex.RUnlock()
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

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
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

var _ user.Repository = new(FakeRepository)
