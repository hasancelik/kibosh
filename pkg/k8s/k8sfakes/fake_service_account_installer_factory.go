// Code generated by counterfeiter. DO NOT EDIT.
package k8sfakes

import (
	"sync"

	"github.com/cf-platform-eng/kibosh/pkg/k8s"
)

type FakeServiceAccountInstallerFactory struct {
	ServiceAccountInstallerStub        func(cluster k8s.Cluster) k8s.ServiceAccountInstaller
	serviceAccountInstallerMutex       sync.RWMutex
	serviceAccountInstallerArgsForCall []struct {
		cluster k8s.Cluster
	}
	serviceAccountInstallerReturns struct {
		result1 k8s.ServiceAccountInstaller
	}
	serviceAccountInstallerReturnsOnCall map[int]struct {
		result1 k8s.ServiceAccountInstaller
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceAccountInstallerFactory) ServiceAccountInstaller(cluster k8s.Cluster) k8s.ServiceAccountInstaller {
	fake.serviceAccountInstallerMutex.Lock()
	ret, specificReturn := fake.serviceAccountInstallerReturnsOnCall[len(fake.serviceAccountInstallerArgsForCall)]
	fake.serviceAccountInstallerArgsForCall = append(fake.serviceAccountInstallerArgsForCall, struct {
		cluster k8s.Cluster
	}{cluster})
	fake.recordInvocation("ServiceAccountInstaller", []interface{}{cluster})
	fake.serviceAccountInstallerMutex.Unlock()
	if fake.ServiceAccountInstallerStub != nil {
		return fake.ServiceAccountInstallerStub(cluster)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.serviceAccountInstallerReturns.result1
}

func (fake *FakeServiceAccountInstallerFactory) ServiceAccountInstallerCallCount() int {
	fake.serviceAccountInstallerMutex.RLock()
	defer fake.serviceAccountInstallerMutex.RUnlock()
	return len(fake.serviceAccountInstallerArgsForCall)
}

func (fake *FakeServiceAccountInstallerFactory) ServiceAccountInstallerArgsForCall(i int) k8s.Cluster {
	fake.serviceAccountInstallerMutex.RLock()
	defer fake.serviceAccountInstallerMutex.RUnlock()
	return fake.serviceAccountInstallerArgsForCall[i].cluster
}

func (fake *FakeServiceAccountInstallerFactory) ServiceAccountInstallerReturns(result1 k8s.ServiceAccountInstaller) {
	fake.ServiceAccountInstallerStub = nil
	fake.serviceAccountInstallerReturns = struct {
		result1 k8s.ServiceAccountInstaller
	}{result1}
}

func (fake *FakeServiceAccountInstallerFactory) ServiceAccountInstallerReturnsOnCall(i int, result1 k8s.ServiceAccountInstaller) {
	fake.ServiceAccountInstallerStub = nil
	if fake.serviceAccountInstallerReturnsOnCall == nil {
		fake.serviceAccountInstallerReturnsOnCall = make(map[int]struct {
			result1 k8s.ServiceAccountInstaller
		})
	}
	fake.serviceAccountInstallerReturnsOnCall[i] = struct {
		result1 k8s.ServiceAccountInstaller
	}{result1}
}

func (fake *FakeServiceAccountInstallerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.serviceAccountInstallerMutex.RLock()
	defer fake.serviceAccountInstallerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceAccountInstallerFactory) recordInvocation(key string, args []interface{}) {
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

var _ k8s.ServiceAccountInstallerFactory = new(FakeServiceAccountInstallerFactory)
