/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"

	"k8s.io/release/pkg/release"
)

type FakeArchiverImpl struct {
	CleanStagedBuildsStub        func(string, string) error
	cleanStagedBuildsMutex       sync.RWMutex
	cleanStagedBuildsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	cleanStagedBuildsReturns struct {
		result1 error
	}
	cleanStagedBuildsReturnsOnCall map[int]struct {
		result1 error
	}
	CopyReleaseLogsStub        func([]string, string, string) error
	copyReleaseLogsMutex       sync.RWMutex
	copyReleaseLogsArgsForCall []struct {
		arg1 []string
		arg2 string
		arg3 string
	}
	copyReleaseLogsReturns struct {
		result1 error
	}
	copyReleaseLogsReturnsOnCall map[int]struct {
		result1 error
	}
	CopyReleaseToBucketStub        func(string, string) error
	copyReleaseToBucketMutex       sync.RWMutex
	copyReleaseToBucketArgsForCall []struct {
		arg1 string
		arg2 string
	}
	copyReleaseToBucketReturns struct {
		result1 error
	}
	copyReleaseToBucketReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStalePasswordFilesStub        func(string) error
	deleteStalePasswordFilesMutex       sync.RWMutex
	deleteStalePasswordFilesArgsForCall []struct {
		arg1 string
	}
	deleteStalePasswordFilesReturns struct {
		result1 error
	}
	deleteStalePasswordFilesReturnsOnCall map[int]struct {
		result1 error
	}
	MakeFilesPrivateStub        func(string) error
	makeFilesPrivateMutex       sync.RWMutex
	makeFilesPrivateArgsForCall []struct {
		arg1 string
	}
	makeFilesPrivateReturns struct {
		result1 error
	}
	makeFilesPrivateReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateOptionsStub        func(*release.ArchiverOptions) error
	validateOptionsMutex       sync.RWMutex
	validateOptionsArgsForCall []struct {
		arg1 *release.ArchiverOptions
	}
	validateOptionsReturns struct {
		result1 error
	}
	validateOptionsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArchiverImpl) CleanStagedBuilds(arg1 string, arg2 string) error {
	fake.cleanStagedBuildsMutex.Lock()
	ret, specificReturn := fake.cleanStagedBuildsReturnsOnCall[len(fake.cleanStagedBuildsArgsForCall)]
	fake.cleanStagedBuildsArgsForCall = append(fake.cleanStagedBuildsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CleanStagedBuildsStub
	fakeReturns := fake.cleanStagedBuildsReturns
	fake.recordInvocation("CleanStagedBuilds", []interface{}{arg1, arg2})
	fake.cleanStagedBuildsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArchiverImpl) CleanStagedBuildsCallCount() int {
	fake.cleanStagedBuildsMutex.RLock()
	defer fake.cleanStagedBuildsMutex.RUnlock()
	return len(fake.cleanStagedBuildsArgsForCall)
}

func (fake *FakeArchiverImpl) CleanStagedBuildsCalls(stub func(string, string) error) {
	fake.cleanStagedBuildsMutex.Lock()
	defer fake.cleanStagedBuildsMutex.Unlock()
	fake.CleanStagedBuildsStub = stub
}

func (fake *FakeArchiverImpl) CleanStagedBuildsArgsForCall(i int) (string, string) {
	fake.cleanStagedBuildsMutex.RLock()
	defer fake.cleanStagedBuildsMutex.RUnlock()
	argsForCall := fake.cleanStagedBuildsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeArchiverImpl) CleanStagedBuildsReturns(result1 error) {
	fake.cleanStagedBuildsMutex.Lock()
	defer fake.cleanStagedBuildsMutex.Unlock()
	fake.CleanStagedBuildsStub = nil
	fake.cleanStagedBuildsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) CleanStagedBuildsReturnsOnCall(i int, result1 error) {
	fake.cleanStagedBuildsMutex.Lock()
	defer fake.cleanStagedBuildsMutex.Unlock()
	fake.CleanStagedBuildsStub = nil
	if fake.cleanStagedBuildsReturnsOnCall == nil {
		fake.cleanStagedBuildsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanStagedBuildsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) CopyReleaseLogs(arg1 []string, arg2 string, arg3 string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.copyReleaseLogsMutex.Lock()
	ret, specificReturn := fake.copyReleaseLogsReturnsOnCall[len(fake.copyReleaseLogsArgsForCall)]
	fake.copyReleaseLogsArgsForCall = append(fake.copyReleaseLogsArgsForCall, struct {
		arg1 []string
		arg2 string
		arg3 string
	}{arg1Copy, arg2, arg3})
	stub := fake.CopyReleaseLogsStub
	fakeReturns := fake.copyReleaseLogsReturns
	fake.recordInvocation("CopyReleaseLogs", []interface{}{arg1Copy, arg2, arg3})
	fake.copyReleaseLogsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArchiverImpl) CopyReleaseLogsCallCount() int {
	fake.copyReleaseLogsMutex.RLock()
	defer fake.copyReleaseLogsMutex.RUnlock()
	return len(fake.copyReleaseLogsArgsForCall)
}

func (fake *FakeArchiverImpl) CopyReleaseLogsCalls(stub func([]string, string, string) error) {
	fake.copyReleaseLogsMutex.Lock()
	defer fake.copyReleaseLogsMutex.Unlock()
	fake.CopyReleaseLogsStub = stub
}

func (fake *FakeArchiverImpl) CopyReleaseLogsArgsForCall(i int) ([]string, string, string) {
	fake.copyReleaseLogsMutex.RLock()
	defer fake.copyReleaseLogsMutex.RUnlock()
	argsForCall := fake.copyReleaseLogsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeArchiverImpl) CopyReleaseLogsReturns(result1 error) {
	fake.copyReleaseLogsMutex.Lock()
	defer fake.copyReleaseLogsMutex.Unlock()
	fake.CopyReleaseLogsStub = nil
	fake.copyReleaseLogsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) CopyReleaseLogsReturnsOnCall(i int, result1 error) {
	fake.copyReleaseLogsMutex.Lock()
	defer fake.copyReleaseLogsMutex.Unlock()
	fake.CopyReleaseLogsStub = nil
	if fake.copyReleaseLogsReturnsOnCall == nil {
		fake.copyReleaseLogsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyReleaseLogsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) CopyReleaseToBucket(arg1 string, arg2 string) error {
	fake.copyReleaseToBucketMutex.Lock()
	ret, specificReturn := fake.copyReleaseToBucketReturnsOnCall[len(fake.copyReleaseToBucketArgsForCall)]
	fake.copyReleaseToBucketArgsForCall = append(fake.copyReleaseToBucketArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CopyReleaseToBucketStub
	fakeReturns := fake.copyReleaseToBucketReturns
	fake.recordInvocation("CopyReleaseToBucket", []interface{}{arg1, arg2})
	fake.copyReleaseToBucketMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArchiverImpl) CopyReleaseToBucketCallCount() int {
	fake.copyReleaseToBucketMutex.RLock()
	defer fake.copyReleaseToBucketMutex.RUnlock()
	return len(fake.copyReleaseToBucketArgsForCall)
}

func (fake *FakeArchiverImpl) CopyReleaseToBucketCalls(stub func(string, string) error) {
	fake.copyReleaseToBucketMutex.Lock()
	defer fake.copyReleaseToBucketMutex.Unlock()
	fake.CopyReleaseToBucketStub = stub
}

func (fake *FakeArchiverImpl) CopyReleaseToBucketArgsForCall(i int) (string, string) {
	fake.copyReleaseToBucketMutex.RLock()
	defer fake.copyReleaseToBucketMutex.RUnlock()
	argsForCall := fake.copyReleaseToBucketArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeArchiverImpl) CopyReleaseToBucketReturns(result1 error) {
	fake.copyReleaseToBucketMutex.Lock()
	defer fake.copyReleaseToBucketMutex.Unlock()
	fake.CopyReleaseToBucketStub = nil
	fake.copyReleaseToBucketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) CopyReleaseToBucketReturnsOnCall(i int, result1 error) {
	fake.copyReleaseToBucketMutex.Lock()
	defer fake.copyReleaseToBucketMutex.Unlock()
	fake.CopyReleaseToBucketStub = nil
	if fake.copyReleaseToBucketReturnsOnCall == nil {
		fake.copyReleaseToBucketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyReleaseToBucketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) DeleteStalePasswordFiles(arg1 string) error {
	fake.deleteStalePasswordFilesMutex.Lock()
	ret, specificReturn := fake.deleteStalePasswordFilesReturnsOnCall[len(fake.deleteStalePasswordFilesArgsForCall)]
	fake.deleteStalePasswordFilesArgsForCall = append(fake.deleteStalePasswordFilesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteStalePasswordFilesStub
	fakeReturns := fake.deleteStalePasswordFilesReturns
	fake.recordInvocation("DeleteStalePasswordFiles", []interface{}{arg1})
	fake.deleteStalePasswordFilesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArchiverImpl) DeleteStalePasswordFilesCallCount() int {
	fake.deleteStalePasswordFilesMutex.RLock()
	defer fake.deleteStalePasswordFilesMutex.RUnlock()
	return len(fake.deleteStalePasswordFilesArgsForCall)
}

func (fake *FakeArchiverImpl) DeleteStalePasswordFilesCalls(stub func(string) error) {
	fake.deleteStalePasswordFilesMutex.Lock()
	defer fake.deleteStalePasswordFilesMutex.Unlock()
	fake.DeleteStalePasswordFilesStub = stub
}

func (fake *FakeArchiverImpl) DeleteStalePasswordFilesArgsForCall(i int) string {
	fake.deleteStalePasswordFilesMutex.RLock()
	defer fake.deleteStalePasswordFilesMutex.RUnlock()
	argsForCall := fake.deleteStalePasswordFilesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeArchiverImpl) DeleteStalePasswordFilesReturns(result1 error) {
	fake.deleteStalePasswordFilesMutex.Lock()
	defer fake.deleteStalePasswordFilesMutex.Unlock()
	fake.DeleteStalePasswordFilesStub = nil
	fake.deleteStalePasswordFilesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) DeleteStalePasswordFilesReturnsOnCall(i int, result1 error) {
	fake.deleteStalePasswordFilesMutex.Lock()
	defer fake.deleteStalePasswordFilesMutex.Unlock()
	fake.DeleteStalePasswordFilesStub = nil
	if fake.deleteStalePasswordFilesReturnsOnCall == nil {
		fake.deleteStalePasswordFilesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteStalePasswordFilesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) MakeFilesPrivate(arg1 string) error {
	fake.makeFilesPrivateMutex.Lock()
	ret, specificReturn := fake.makeFilesPrivateReturnsOnCall[len(fake.makeFilesPrivateArgsForCall)]
	fake.makeFilesPrivateArgsForCall = append(fake.makeFilesPrivateArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.MakeFilesPrivateStub
	fakeReturns := fake.makeFilesPrivateReturns
	fake.recordInvocation("MakeFilesPrivate", []interface{}{arg1})
	fake.makeFilesPrivateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArchiverImpl) MakeFilesPrivateCallCount() int {
	fake.makeFilesPrivateMutex.RLock()
	defer fake.makeFilesPrivateMutex.RUnlock()
	return len(fake.makeFilesPrivateArgsForCall)
}

func (fake *FakeArchiverImpl) MakeFilesPrivateCalls(stub func(string) error) {
	fake.makeFilesPrivateMutex.Lock()
	defer fake.makeFilesPrivateMutex.Unlock()
	fake.MakeFilesPrivateStub = stub
}

func (fake *FakeArchiverImpl) MakeFilesPrivateArgsForCall(i int) string {
	fake.makeFilesPrivateMutex.RLock()
	defer fake.makeFilesPrivateMutex.RUnlock()
	argsForCall := fake.makeFilesPrivateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeArchiverImpl) MakeFilesPrivateReturns(result1 error) {
	fake.makeFilesPrivateMutex.Lock()
	defer fake.makeFilesPrivateMutex.Unlock()
	fake.MakeFilesPrivateStub = nil
	fake.makeFilesPrivateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) MakeFilesPrivateReturnsOnCall(i int, result1 error) {
	fake.makeFilesPrivateMutex.Lock()
	defer fake.makeFilesPrivateMutex.Unlock()
	fake.MakeFilesPrivateStub = nil
	if fake.makeFilesPrivateReturnsOnCall == nil {
		fake.makeFilesPrivateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.makeFilesPrivateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) ValidateOptions(arg1 *release.ArchiverOptions) error {
	fake.validateOptionsMutex.Lock()
	ret, specificReturn := fake.validateOptionsReturnsOnCall[len(fake.validateOptionsArgsForCall)]
	fake.validateOptionsArgsForCall = append(fake.validateOptionsArgsForCall, struct {
		arg1 *release.ArchiverOptions
	}{arg1})
	stub := fake.ValidateOptionsStub
	fakeReturns := fake.validateOptionsReturns
	fake.recordInvocation("ValidateOptions", []interface{}{arg1})
	fake.validateOptionsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArchiverImpl) ValidateOptionsCallCount() int {
	fake.validateOptionsMutex.RLock()
	defer fake.validateOptionsMutex.RUnlock()
	return len(fake.validateOptionsArgsForCall)
}

func (fake *FakeArchiverImpl) ValidateOptionsCalls(stub func(*release.ArchiverOptions) error) {
	fake.validateOptionsMutex.Lock()
	defer fake.validateOptionsMutex.Unlock()
	fake.ValidateOptionsStub = stub
}

func (fake *FakeArchiverImpl) ValidateOptionsArgsForCall(i int) *release.ArchiverOptions {
	fake.validateOptionsMutex.RLock()
	defer fake.validateOptionsMutex.RUnlock()
	argsForCall := fake.validateOptionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeArchiverImpl) ValidateOptionsReturns(result1 error) {
	fake.validateOptionsMutex.Lock()
	defer fake.validateOptionsMutex.Unlock()
	fake.ValidateOptionsStub = nil
	fake.validateOptionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) ValidateOptionsReturnsOnCall(i int, result1 error) {
	fake.validateOptionsMutex.Lock()
	defer fake.validateOptionsMutex.Unlock()
	fake.ValidateOptionsStub = nil
	if fake.validateOptionsReturnsOnCall == nil {
		fake.validateOptionsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateOptionsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArchiverImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cleanStagedBuildsMutex.RLock()
	defer fake.cleanStagedBuildsMutex.RUnlock()
	fake.copyReleaseLogsMutex.RLock()
	defer fake.copyReleaseLogsMutex.RUnlock()
	fake.copyReleaseToBucketMutex.RLock()
	defer fake.copyReleaseToBucketMutex.RUnlock()
	fake.deleteStalePasswordFilesMutex.RLock()
	defer fake.deleteStalePasswordFilesMutex.RUnlock()
	fake.makeFilesPrivateMutex.RLock()
	defer fake.makeFilesPrivateMutex.RUnlock()
	fake.validateOptionsMutex.RLock()
	defer fake.validateOptionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeArchiverImpl) recordInvocation(key string, args []interface{}) {
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
