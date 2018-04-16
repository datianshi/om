// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type CredentialReferencesService struct {
	ListDeployedProductCredentialsStub        func(deployedProductGUID string) (api.CredentialReferencesOutput, error)
	listDeployedProductCredentialsMutex       sync.RWMutex
	listDeployedProductCredentialsArgsForCall []struct {
		deployedProductGUID string
	}
	listDeployedProductCredentialsReturns struct {
		result1 api.CredentialReferencesOutput
		result2 error
	}
	listDeployedProductCredentialsReturnsOnCall map[int]struct {
		result1 api.CredentialReferencesOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CredentialReferencesService) ListDeployedProductCredentials(deployedProductGUID string) (api.CredentialReferencesOutput, error) {
	fake.listDeployedProductCredentialsMutex.Lock()
	ret, specificReturn := fake.listDeployedProductCredentialsReturnsOnCall[len(fake.listDeployedProductCredentialsArgsForCall)]
	fake.listDeployedProductCredentialsArgsForCall = append(fake.listDeployedProductCredentialsArgsForCall, struct {
		deployedProductGUID string
	}{deployedProductGUID})
	fake.recordInvocation("ListDeployedProductCredentials", []interface{}{deployedProductGUID})
	fake.listDeployedProductCredentialsMutex.Unlock()
	if fake.ListDeployedProductCredentialsStub != nil {
		return fake.ListDeployedProductCredentialsStub(deployedProductGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDeployedProductCredentialsReturns.result1, fake.listDeployedProductCredentialsReturns.result2
}

func (fake *CredentialReferencesService) ListDeployedProductCredentialsCallCount() int {
	fake.listDeployedProductCredentialsMutex.RLock()
	defer fake.listDeployedProductCredentialsMutex.RUnlock()
	return len(fake.listDeployedProductCredentialsArgsForCall)
}

func (fake *CredentialReferencesService) ListDeployedProductCredentialsArgsForCall(i int) string {
	fake.listDeployedProductCredentialsMutex.RLock()
	defer fake.listDeployedProductCredentialsMutex.RUnlock()
	return fake.listDeployedProductCredentialsArgsForCall[i].deployedProductGUID
}

func (fake *CredentialReferencesService) ListDeployedProductCredentialsReturns(result1 api.CredentialReferencesOutput, result2 error) {
	fake.ListDeployedProductCredentialsStub = nil
	fake.listDeployedProductCredentialsReturns = struct {
		result1 api.CredentialReferencesOutput
		result2 error
	}{result1, result2}
}

func (fake *CredentialReferencesService) ListDeployedProductCredentialsReturnsOnCall(i int, result1 api.CredentialReferencesOutput, result2 error) {
	fake.ListDeployedProductCredentialsStub = nil
	if fake.listDeployedProductCredentialsReturnsOnCall == nil {
		fake.listDeployedProductCredentialsReturnsOnCall = make(map[int]struct {
			result1 api.CredentialReferencesOutput
			result2 error
		})
	}
	fake.listDeployedProductCredentialsReturnsOnCall[i] = struct {
		result1 api.CredentialReferencesOutput
		result2 error
	}{result1, result2}
}

func (fake *CredentialReferencesService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listDeployedProductCredentialsMutex.RLock()
	defer fake.listDeployedProductCredentialsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CredentialReferencesService) recordInvocation(key string, args []interface{}) {
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
