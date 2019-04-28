// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/51st-state/api/pkg/pubsub"
)

type FakeConsumer struct {
	ConsumeStub        func(context.Context, pubsub.HandlerFunc) error
	consumeMutex       sync.RWMutex
	consumeArgsForCall []struct {
		arg1 context.Context
		arg2 pubsub.HandlerFunc
	}
	consumeReturns struct {
		result1 error
	}
	consumeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConsumer) Consume(arg1 context.Context, arg2 pubsub.HandlerFunc) error {
	fake.consumeMutex.Lock()
	ret, specificReturn := fake.consumeReturnsOnCall[len(fake.consumeArgsForCall)]
	fake.consumeArgsForCall = append(fake.consumeArgsForCall, struct {
		arg1 context.Context
		arg2 pubsub.HandlerFunc
	}{arg1, arg2})
	fake.recordInvocation("Consume", []interface{}{arg1, arg2})
	fake.consumeMutex.Unlock()
	if fake.ConsumeStub != nil {
		return fake.ConsumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.consumeReturns.result1
}

func (fake *FakeConsumer) ConsumeCallCount() int {
	fake.consumeMutex.RLock()
	defer fake.consumeMutex.RUnlock()
	return len(fake.consumeArgsForCall)
}

func (fake *FakeConsumer) ConsumeArgsForCall(i int) (context.Context, pubsub.HandlerFunc) {
	fake.consumeMutex.RLock()
	defer fake.consumeMutex.RUnlock()
	return fake.consumeArgsForCall[i].arg1, fake.consumeArgsForCall[i].arg2
}

func (fake *FakeConsumer) ConsumeReturns(result1 error) {
	fake.ConsumeStub = nil
	fake.consumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumer) ConsumeReturnsOnCall(i int, result1 error) {
	fake.ConsumeStub = nil
	if fake.consumeReturnsOnCall == nil {
		fake.consumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.consumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.consumeMutex.RLock()
	defer fake.consumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConsumer) recordInvocation(key string, args []interface{}) {
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

var _ pubsub.Consumer = new(FakeConsumer)
