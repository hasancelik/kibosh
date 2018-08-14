// Code generated by counterfeiter. DO NOT EDIT.
package statefakes

import (
	"sync"

	"github.com/cf-platform-eng/kibosh/pkg/state"
)

type FakeKeyValueStore struct {
	OpenStub        func(dataDir string) error
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		dataDir string
	}
	openReturns struct {
		result1 error
	}
	openReturnsOnCall map[int]struct {
		result1 error
	}
	PutJsonStub        func(key string, val interface{}) error
	putJsonMutex       sync.RWMutex
	putJsonArgsForCall []struct {
		key string
		val interface{}
	}
	putJsonReturns struct {
		result1 error
	}
	putJsonReturnsOnCall map[int]struct {
		result1 error
	}
	GetJsonStub        func(key string, val interface{}) error
	getJsonMutex       sync.RWMutex
	getJsonArgsForCall []struct {
		key string
		val interface{}
	}
	getJsonReturns struct {
		result1 error
	}
	getJsonReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(key string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		key string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKeyValueStore) Open(dataDir string) error {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		dataDir string
	}{dataDir})
	fake.recordInvocation("Open", []interface{}{dataDir})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(dataDir)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.openReturns.result1
}

func (fake *FakeKeyValueStore) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakeKeyValueStore) OpenArgsForCall(i int) string {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return fake.openArgsForCall[i].dataDir
}

func (fake *FakeKeyValueStore) OpenReturns(result1 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKeyValueStore) OpenReturnsOnCall(i int, result1 error) {
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKeyValueStore) PutJson(key string, val interface{}) error {
	fake.putJsonMutex.Lock()
	ret, specificReturn := fake.putJsonReturnsOnCall[len(fake.putJsonArgsForCall)]
	fake.putJsonArgsForCall = append(fake.putJsonArgsForCall, struct {
		key string
		val interface{}
	}{key, val})
	fake.recordInvocation("PutJson", []interface{}{key, val})
	fake.putJsonMutex.Unlock()
	if fake.PutJsonStub != nil {
		return fake.PutJsonStub(key, val)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.putJsonReturns.result1
}

func (fake *FakeKeyValueStore) PutJsonCallCount() int {
	fake.putJsonMutex.RLock()
	defer fake.putJsonMutex.RUnlock()
	return len(fake.putJsonArgsForCall)
}

func (fake *FakeKeyValueStore) PutJsonArgsForCall(i int) (string, interface{}) {
	fake.putJsonMutex.RLock()
	defer fake.putJsonMutex.RUnlock()
	return fake.putJsonArgsForCall[i].key, fake.putJsonArgsForCall[i].val
}

func (fake *FakeKeyValueStore) PutJsonReturns(result1 error) {
	fake.PutJsonStub = nil
	fake.putJsonReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKeyValueStore) PutJsonReturnsOnCall(i int, result1 error) {
	fake.PutJsonStub = nil
	if fake.putJsonReturnsOnCall == nil {
		fake.putJsonReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.putJsonReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKeyValueStore) GetJson(key string, val interface{}) error {
	fake.getJsonMutex.Lock()
	ret, specificReturn := fake.getJsonReturnsOnCall[len(fake.getJsonArgsForCall)]
	fake.getJsonArgsForCall = append(fake.getJsonArgsForCall, struct {
		key string
		val interface{}
	}{key, val})
	fake.recordInvocation("GetJson", []interface{}{key, val})
	fake.getJsonMutex.Unlock()
	if fake.GetJsonStub != nil {
		return fake.GetJsonStub(key, val)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getJsonReturns.result1
}

func (fake *FakeKeyValueStore) GetJsonCallCount() int {
	fake.getJsonMutex.RLock()
	defer fake.getJsonMutex.RUnlock()
	return len(fake.getJsonArgsForCall)
}

func (fake *FakeKeyValueStore) GetJsonArgsForCall(i int) (string, interface{}) {
	fake.getJsonMutex.RLock()
	defer fake.getJsonMutex.RUnlock()
	return fake.getJsonArgsForCall[i].key, fake.getJsonArgsForCall[i].val
}

func (fake *FakeKeyValueStore) GetJsonReturns(result1 error) {
	fake.GetJsonStub = nil
	fake.getJsonReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKeyValueStore) GetJsonReturnsOnCall(i int, result1 error) {
	fake.GetJsonStub = nil
	if fake.getJsonReturnsOnCall == nil {
		fake.getJsonReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getJsonReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKeyValueStore) Delete(key string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		key string
	}{key})
	fake.recordInvocation("Delete", []interface{}{key})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(key)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeKeyValueStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeKeyValueStore) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].key
}

func (fake *FakeKeyValueStore) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKeyValueStore) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeKeyValueStore) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeKeyValueStore) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeKeyValueStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.putJsonMutex.RLock()
	defer fake.putJsonMutex.RUnlock()
	fake.getJsonMutex.RLock()
	defer fake.getJsonMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKeyValueStore) recordInvocation(key string, args []interface{}) {
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

var _ state.KeyValueStore = new(FakeKeyValueStore)
