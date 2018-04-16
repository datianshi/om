// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ProductUploader struct {
	UploadAvailableProductStub        func(api.UploadAvailableProductInput) (api.UploadAvailableProductOutput, error)
	uploadAvailableProductMutex       sync.RWMutex
	uploadAvailableProductArgsForCall []struct {
		arg1 api.UploadAvailableProductInput
	}
	uploadAvailableProductReturns struct {
		result1 api.UploadAvailableProductOutput
		result2 error
	}
	uploadAvailableProductReturnsOnCall map[int]struct {
		result1 api.UploadAvailableProductOutput
		result2 error
	}
	CheckProductAvailabilityStub        func(string, string) (bool, error)
	checkProductAvailabilityMutex       sync.RWMutex
	checkProductAvailabilityArgsForCall []struct {
		arg1 string
		arg2 string
	}
	checkProductAvailabilityReturns struct {
		result1 bool
		result2 error
	}
	checkProductAvailabilityReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProductUploader) UploadAvailableProduct(arg1 api.UploadAvailableProductInput) (api.UploadAvailableProductOutput, error) {
	fake.uploadAvailableProductMutex.Lock()
	ret, specificReturn := fake.uploadAvailableProductReturnsOnCall[len(fake.uploadAvailableProductArgsForCall)]
	fake.uploadAvailableProductArgsForCall = append(fake.uploadAvailableProductArgsForCall, struct {
		arg1 api.UploadAvailableProductInput
	}{arg1})
	fake.recordInvocation("UploadAvailableProduct", []interface{}{arg1})
	fake.uploadAvailableProductMutex.Unlock()
	if fake.UploadAvailableProductStub != nil {
		return fake.UploadAvailableProductStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uploadAvailableProductReturns.result1, fake.uploadAvailableProductReturns.result2
}

func (fake *ProductUploader) UploadAvailableProductCallCount() int {
	fake.uploadAvailableProductMutex.RLock()
	defer fake.uploadAvailableProductMutex.RUnlock()
	return len(fake.uploadAvailableProductArgsForCall)
}

func (fake *ProductUploader) UploadAvailableProductArgsForCall(i int) api.UploadAvailableProductInput {
	fake.uploadAvailableProductMutex.RLock()
	defer fake.uploadAvailableProductMutex.RUnlock()
	return fake.uploadAvailableProductArgsForCall[i].arg1
}

func (fake *ProductUploader) UploadAvailableProductReturns(result1 api.UploadAvailableProductOutput, result2 error) {
	fake.UploadAvailableProductStub = nil
	fake.uploadAvailableProductReturns = struct {
		result1 api.UploadAvailableProductOutput
		result2 error
	}{result1, result2}
}

func (fake *ProductUploader) UploadAvailableProductReturnsOnCall(i int, result1 api.UploadAvailableProductOutput, result2 error) {
	fake.UploadAvailableProductStub = nil
	if fake.uploadAvailableProductReturnsOnCall == nil {
		fake.uploadAvailableProductReturnsOnCall = make(map[int]struct {
			result1 api.UploadAvailableProductOutput
			result2 error
		})
	}
	fake.uploadAvailableProductReturnsOnCall[i] = struct {
		result1 api.UploadAvailableProductOutput
		result2 error
	}{result1, result2}
}

func (fake *ProductUploader) CheckProductAvailability(arg1 string, arg2 string) (bool, error) {
	fake.checkProductAvailabilityMutex.Lock()
	ret, specificReturn := fake.checkProductAvailabilityReturnsOnCall[len(fake.checkProductAvailabilityArgsForCall)]
	fake.checkProductAvailabilityArgsForCall = append(fake.checkProductAvailabilityArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CheckProductAvailability", []interface{}{arg1, arg2})
	fake.checkProductAvailabilityMutex.Unlock()
	if fake.CheckProductAvailabilityStub != nil {
		return fake.CheckProductAvailabilityStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.checkProductAvailabilityReturns.result1, fake.checkProductAvailabilityReturns.result2
}

func (fake *ProductUploader) CheckProductAvailabilityCallCount() int {
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	return len(fake.checkProductAvailabilityArgsForCall)
}

func (fake *ProductUploader) CheckProductAvailabilityArgsForCall(i int) (string, string) {
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	return fake.checkProductAvailabilityArgsForCall[i].arg1, fake.checkProductAvailabilityArgsForCall[i].arg2
}

func (fake *ProductUploader) CheckProductAvailabilityReturns(result1 bool, result2 error) {
	fake.CheckProductAvailabilityStub = nil
	fake.checkProductAvailabilityReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *ProductUploader) CheckProductAvailabilityReturnsOnCall(i int, result1 bool, result2 error) {
	fake.CheckProductAvailabilityStub = nil
	if fake.checkProductAvailabilityReturnsOnCall == nil {
		fake.checkProductAvailabilityReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkProductAvailabilityReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *ProductUploader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadAvailableProductMutex.RLock()
	defer fake.uploadAvailableProductMutex.RUnlock()
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ProductUploader) recordInvocation(key string, args []interface{}) {
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
