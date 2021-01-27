// Code generated by counterfeiter. DO NOT EDIT.
package basefakes

import (
	context "context"
	sync "sync"

	models "github.com/cloudfoundry-incubator/cloud-service-broker/db_service/models"
	base "github.com/cloudfoundry-incubator/cloud-service-broker/pkg/providers/builtin/base"
	varcontext "github.com/cloudfoundry-incubator/cloud-service-broker/pkg/varcontext"
)

type FakeServiceAccountManager struct {
	CreateCredentialsStub        func(context.Context, *varcontext.VarContext) (map[string]interface{}, error)
	createCredentialsMutex       sync.RWMutex
	createCredentialsArgsForCall []struct {
		arg1 context.Context
		arg2 *varcontext.VarContext
	}
	createCredentialsReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	createCredentialsReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	DeleteCredentialsStub        func(context.Context, models.ServiceBindingCredentials) error
	deleteCredentialsMutex       sync.RWMutex
	deleteCredentialsArgsForCall []struct {
		arg1 context.Context
		arg2 models.ServiceBindingCredentials
	}
	deleteCredentialsReturns struct {
		result1 error
	}
	deleteCredentialsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceAccountManager) CreateCredentials(arg1 context.Context, arg2 *varcontext.VarContext) (map[string]interface{}, error) {
	fake.createCredentialsMutex.Lock()
	ret, specificReturn := fake.createCredentialsReturnsOnCall[len(fake.createCredentialsArgsForCall)]
	fake.createCredentialsArgsForCall = append(fake.createCredentialsArgsForCall, struct {
		arg1 context.Context
		arg2 *varcontext.VarContext
	}{arg1, arg2})
	fake.recordInvocation("CreateCredentials", []interface{}{arg1, arg2})
	fake.createCredentialsMutex.Unlock()
	if fake.CreateCredentialsStub != nil {
		return fake.CreateCredentialsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createCredentialsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceAccountManager) CreateCredentialsCallCount() int {
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	return len(fake.createCredentialsArgsForCall)
}

func (fake *FakeServiceAccountManager) CreateCredentialsCalls(stub func(context.Context, *varcontext.VarContext) (map[string]interface{}, error)) {
	fake.createCredentialsMutex.Lock()
	defer fake.createCredentialsMutex.Unlock()
	fake.CreateCredentialsStub = stub
}

func (fake *FakeServiceAccountManager) CreateCredentialsArgsForCall(i int) (context.Context, *varcontext.VarContext) {
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	argsForCall := fake.createCredentialsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceAccountManager) CreateCredentialsReturns(result1 map[string]interface{}, result2 error) {
	fake.createCredentialsMutex.Lock()
	defer fake.createCredentialsMutex.Unlock()
	fake.CreateCredentialsStub = nil
	fake.createCredentialsReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAccountManager) CreateCredentialsReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.createCredentialsMutex.Lock()
	defer fake.createCredentialsMutex.Unlock()
	fake.CreateCredentialsStub = nil
	if fake.createCredentialsReturnsOnCall == nil {
		fake.createCredentialsReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.createCredentialsReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAccountManager) DeleteCredentials(arg1 context.Context, arg2 models.ServiceBindingCredentials) error {
	fake.deleteCredentialsMutex.Lock()
	ret, specificReturn := fake.deleteCredentialsReturnsOnCall[len(fake.deleteCredentialsArgsForCall)]
	fake.deleteCredentialsArgsForCall = append(fake.deleteCredentialsArgsForCall, struct {
		arg1 context.Context
		arg2 models.ServiceBindingCredentials
	}{arg1, arg2})
	fake.recordInvocation("DeleteCredentials", []interface{}{arg1, arg2})
	fake.deleteCredentialsMutex.Unlock()
	if fake.DeleteCredentialsStub != nil {
		return fake.DeleteCredentialsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteCredentialsReturns
	return fakeReturns.result1
}

func (fake *FakeServiceAccountManager) DeleteCredentialsCallCount() int {
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	return len(fake.deleteCredentialsArgsForCall)
}

func (fake *FakeServiceAccountManager) DeleteCredentialsCalls(stub func(context.Context, models.ServiceBindingCredentials) error) {
	fake.deleteCredentialsMutex.Lock()
	defer fake.deleteCredentialsMutex.Unlock()
	fake.DeleteCredentialsStub = stub
}

func (fake *FakeServiceAccountManager) DeleteCredentialsArgsForCall(i int) (context.Context, models.ServiceBindingCredentials) {
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	argsForCall := fake.deleteCredentialsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceAccountManager) DeleteCredentialsReturns(result1 error) {
	fake.deleteCredentialsMutex.Lock()
	defer fake.deleteCredentialsMutex.Unlock()
	fake.DeleteCredentialsStub = nil
	fake.deleteCredentialsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAccountManager) DeleteCredentialsReturnsOnCall(i int, result1 error) {
	fake.deleteCredentialsMutex.Lock()
	defer fake.deleteCredentialsMutex.Unlock()
	fake.DeleteCredentialsStub = nil
	if fake.deleteCredentialsReturnsOnCall == nil {
		fake.deleteCredentialsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCredentialsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAccountManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceAccountManager) recordInvocation(key string, args []interface{}) {
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

var _ base.ServiceAccountManager = new(FakeServiceAccountManager)
