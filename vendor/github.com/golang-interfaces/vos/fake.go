// Code generated by counterfeiter. DO NOT EDIT.
package ios

import (
	"os"
	"sync"
)

type Fake struct {
	ChmodStub        func(name string, mode os.FileMode) error
	chmodMutex       sync.RWMutex
	chmodArgsForCall []struct {
		name string
		mode os.FileMode
	}
	chmodReturns struct {
		result1 error
	}
	chmodReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(name string) (*os.File, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		name string
	}
	createReturns struct {
		result1 *os.File
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	ExitStub        func(code int)
	exitMutex       sync.RWMutex
	exitArgsForCall []struct {
		code int
	}
	FindProcessStub        func(pid int) (*os.Process, error)
	findProcessMutex       sync.RWMutex
	findProcessArgsForCall []struct {
		pid int
	}
	findProcessReturns struct {
		result1 *os.Process
		result2 error
	}
	findProcessReturnsOnCall map[int]struct {
		result1 *os.Process
		result2 error
	}
	GetenvStub        func(key string) string
	getenvMutex       sync.RWMutex
	getenvArgsForCall []struct {
		key string
	}
	getenvReturns struct {
		result1 string
	}
	getenvReturnsOnCall map[int]struct {
		result1 string
	}
	GetpidStub        func() int
	getpidMutex       sync.RWMutex
	getpidArgsForCall []struct{}
	getpidReturns     struct {
		result1 int
	}
	getpidReturnsOnCall map[int]struct {
		result1 int
	}
	GetwdStub        func() (string, error)
	getwdMutex       sync.RWMutex
	getwdArgsForCall []struct{}
	getwdReturns     struct {
		result1 string
		result2 error
	}
	getwdReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	MkdirAllStub        func(path string, perm os.FileMode) error
	mkdirAllMutex       sync.RWMutex
	mkdirAllArgsForCall []struct {
		path string
		perm os.FileMode
	}
	mkdirAllReturns struct {
		result1 error
	}
	mkdirAllReturnsOnCall map[int]struct {
		result1 error
	}
	OpenStub        func(name string) (*os.File, error)
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		name string
	}
	openReturns struct {
		result1 *os.File
		result2 error
	}
	openReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	OpenFileStub        func(name string, flag int, perm os.FileMode) (*os.File, error)
	openFileMutex       sync.RWMutex
	openFileArgsForCall []struct {
		name string
		flag int
		perm os.FileMode
	}
	openFileReturns struct {
		result1 *os.File
		result2 error
	}
	openFileReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	RemoveAllStub        func(path string) error
	removeAllMutex       sync.RWMutex
	removeAllArgsForCall []struct {
		path string
	}
	removeAllReturns struct {
		result1 error
	}
	removeAllReturnsOnCall map[int]struct {
		result1 error
	}
	SetenvStub        func(key, value string) error
	setenvMutex       sync.RWMutex
	setenvArgsForCall []struct {
		key   string
		value string
	}
	setenvReturns struct {
		result1 error
	}
	setenvReturnsOnCall map[int]struct {
		result1 error
	}
	StatStub        func(name string) (os.FileInfo, error)
	statMutex       sync.RWMutex
	statArgsForCall []struct {
		name string
	}
	statReturns struct {
		result1 os.FileInfo
		result2 error
	}
	statReturnsOnCall map[int]struct {
		result1 os.FileInfo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Chmod(name string, mode os.FileMode) error {
	fake.chmodMutex.Lock()
	ret, specificReturn := fake.chmodReturnsOnCall[len(fake.chmodArgsForCall)]
	fake.chmodArgsForCall = append(fake.chmodArgsForCall, struct {
		name string
		mode os.FileMode
	}{name, mode})
	fake.recordInvocation("Chmod", []interface{}{name, mode})
	fake.chmodMutex.Unlock()
	if fake.ChmodStub != nil {
		return fake.ChmodStub(name, mode)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.chmodReturns.result1
}

func (fake *Fake) ChmodCallCount() int {
	fake.chmodMutex.RLock()
	defer fake.chmodMutex.RUnlock()
	return len(fake.chmodArgsForCall)
}

func (fake *Fake) ChmodArgsForCall(i int) (string, os.FileMode) {
	fake.chmodMutex.RLock()
	defer fake.chmodMutex.RUnlock()
	return fake.chmodArgsForCall[i].name, fake.chmodArgsForCall[i].mode
}

func (fake *Fake) ChmodReturns(result1 error) {
	fake.ChmodStub = nil
	fake.chmodReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) ChmodReturnsOnCall(i int, result1 error) {
	fake.ChmodStub = nil
	if fake.chmodReturnsOnCall == nil {
		fake.chmodReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.chmodReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) Create(name string) (*os.File, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Create", []interface{}{name})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *Fake) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Fake) CreateArgsForCall(i int) string {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].name
}

func (fake *Fake) CreateReturns(result1 *os.File, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *Fake) CreateReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *Fake) Exit(code int) {
	fake.exitMutex.Lock()
	fake.exitArgsForCall = append(fake.exitArgsForCall, struct {
		code int
	}{code})
	fake.recordInvocation("Exit", []interface{}{code})
	fake.exitMutex.Unlock()
	if fake.ExitStub != nil {
		fake.ExitStub(code)
	}
}

func (fake *Fake) ExitCallCount() int {
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	return len(fake.exitArgsForCall)
}

func (fake *Fake) ExitArgsForCall(i int) int {
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	return fake.exitArgsForCall[i].code
}

func (fake *Fake) FindProcess(pid int) (*os.Process, error) {
	fake.findProcessMutex.Lock()
	ret, specificReturn := fake.findProcessReturnsOnCall[len(fake.findProcessArgsForCall)]
	fake.findProcessArgsForCall = append(fake.findProcessArgsForCall, struct {
		pid int
	}{pid})
	fake.recordInvocation("FindProcess", []interface{}{pid})
	fake.findProcessMutex.Unlock()
	if fake.FindProcessStub != nil {
		return fake.FindProcessStub(pid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findProcessReturns.result1, fake.findProcessReturns.result2
}

func (fake *Fake) FindProcessCallCount() int {
	fake.findProcessMutex.RLock()
	defer fake.findProcessMutex.RUnlock()
	return len(fake.findProcessArgsForCall)
}

func (fake *Fake) FindProcessArgsForCall(i int) int {
	fake.findProcessMutex.RLock()
	defer fake.findProcessMutex.RUnlock()
	return fake.findProcessArgsForCall[i].pid
}

func (fake *Fake) FindProcessReturns(result1 *os.Process, result2 error) {
	fake.FindProcessStub = nil
	fake.findProcessReturns = struct {
		result1 *os.Process
		result2 error
	}{result1, result2}
}

func (fake *Fake) FindProcessReturnsOnCall(i int, result1 *os.Process, result2 error) {
	fake.FindProcessStub = nil
	if fake.findProcessReturnsOnCall == nil {
		fake.findProcessReturnsOnCall = make(map[int]struct {
			result1 *os.Process
			result2 error
		})
	}
	fake.findProcessReturnsOnCall[i] = struct {
		result1 *os.Process
		result2 error
	}{result1, result2}
}

func (fake *Fake) Getenv(key string) string {
	fake.getenvMutex.Lock()
	ret, specificReturn := fake.getenvReturnsOnCall[len(fake.getenvArgsForCall)]
	fake.getenvArgsForCall = append(fake.getenvArgsForCall, struct {
		key string
	}{key})
	fake.recordInvocation("Getenv", []interface{}{key})
	fake.getenvMutex.Unlock()
	if fake.GetenvStub != nil {
		return fake.GetenvStub(key)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getenvReturns.result1
}

func (fake *Fake) GetenvCallCount() int {
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	return len(fake.getenvArgsForCall)
}

func (fake *Fake) GetenvArgsForCall(i int) string {
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	return fake.getenvArgsForCall[i].key
}

func (fake *Fake) GetenvReturns(result1 string) {
	fake.GetenvStub = nil
	fake.getenvReturns = struct {
		result1 string
	}{result1}
}

func (fake *Fake) GetenvReturnsOnCall(i int, result1 string) {
	fake.GetenvStub = nil
	if fake.getenvReturnsOnCall == nil {
		fake.getenvReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getenvReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Fake) Getpid() int {
	fake.getpidMutex.Lock()
	ret, specificReturn := fake.getpidReturnsOnCall[len(fake.getpidArgsForCall)]
	fake.getpidArgsForCall = append(fake.getpidArgsForCall, struct{}{})
	fake.recordInvocation("Getpid", []interface{}{})
	fake.getpidMutex.Unlock()
	if fake.GetpidStub != nil {
		return fake.GetpidStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getpidReturns.result1
}

func (fake *Fake) GetpidCallCount() int {
	fake.getpidMutex.RLock()
	defer fake.getpidMutex.RUnlock()
	return len(fake.getpidArgsForCall)
}

func (fake *Fake) GetpidReturns(result1 int) {
	fake.GetpidStub = nil
	fake.getpidReturns = struct {
		result1 int
	}{result1}
}

func (fake *Fake) GetpidReturnsOnCall(i int, result1 int) {
	fake.GetpidStub = nil
	if fake.getpidReturnsOnCall == nil {
		fake.getpidReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.getpidReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *Fake) Getwd() (string, error) {
	fake.getwdMutex.Lock()
	ret, specificReturn := fake.getwdReturnsOnCall[len(fake.getwdArgsForCall)]
	fake.getwdArgsForCall = append(fake.getwdArgsForCall, struct{}{})
	fake.recordInvocation("Getwd", []interface{}{})
	fake.getwdMutex.Unlock()
	if fake.GetwdStub != nil {
		return fake.GetwdStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getwdReturns.result1, fake.getwdReturns.result2
}

func (fake *Fake) GetwdCallCount() int {
	fake.getwdMutex.RLock()
	defer fake.getwdMutex.RUnlock()
	return len(fake.getwdArgsForCall)
}

func (fake *Fake) GetwdReturns(result1 string, result2 error) {
	fake.GetwdStub = nil
	fake.getwdReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Fake) GetwdReturnsOnCall(i int, result1 string, result2 error) {
	fake.GetwdStub = nil
	if fake.getwdReturnsOnCall == nil {
		fake.getwdReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getwdReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Fake) MkdirAll(path string, perm os.FileMode) error {
	fake.mkdirAllMutex.Lock()
	ret, specificReturn := fake.mkdirAllReturnsOnCall[len(fake.mkdirAllArgsForCall)]
	fake.mkdirAllArgsForCall = append(fake.mkdirAllArgsForCall, struct {
		path string
		perm os.FileMode
	}{path, perm})
	fake.recordInvocation("MkdirAll", []interface{}{path, perm})
	fake.mkdirAllMutex.Unlock()
	if fake.MkdirAllStub != nil {
		return fake.MkdirAllStub(path, perm)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.mkdirAllReturns.result1
}

func (fake *Fake) MkdirAllCallCount() int {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return len(fake.mkdirAllArgsForCall)
}

func (fake *Fake) MkdirAllArgsForCall(i int) (string, os.FileMode) {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return fake.mkdirAllArgsForCall[i].path, fake.mkdirAllArgsForCall[i].perm
}

func (fake *Fake) MkdirAllReturns(result1 error) {
	fake.MkdirAllStub = nil
	fake.mkdirAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) MkdirAllReturnsOnCall(i int, result1 error) {
	fake.MkdirAllStub = nil
	if fake.mkdirAllReturnsOnCall == nil {
		fake.mkdirAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mkdirAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) Open(name string) (*os.File, error) {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Open", []interface{}{name})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.openReturns.result1, fake.openReturns.result2
}

func (fake *Fake) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *Fake) OpenArgsForCall(i int) string {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return fake.openArgsForCall[i].name
}

func (fake *Fake) OpenReturns(result1 *os.File, result2 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *Fake) OpenReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *Fake) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	fake.openFileMutex.Lock()
	ret, specificReturn := fake.openFileReturnsOnCall[len(fake.openFileArgsForCall)]
	fake.openFileArgsForCall = append(fake.openFileArgsForCall, struct {
		name string
		flag int
		perm os.FileMode
	}{name, flag, perm})
	fake.recordInvocation("OpenFile", []interface{}{name, flag, perm})
	fake.openFileMutex.Unlock()
	if fake.OpenFileStub != nil {
		return fake.OpenFileStub(name, flag, perm)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.openFileReturns.result1, fake.openFileReturns.result2
}

func (fake *Fake) OpenFileCallCount() int {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return len(fake.openFileArgsForCall)
}

func (fake *Fake) OpenFileArgsForCall(i int) (string, int, os.FileMode) {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return fake.openFileArgsForCall[i].name, fake.openFileArgsForCall[i].flag, fake.openFileArgsForCall[i].perm
}

func (fake *Fake) OpenFileReturns(result1 *os.File, result2 error) {
	fake.OpenFileStub = nil
	fake.openFileReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *Fake) OpenFileReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.OpenFileStub = nil
	if fake.openFileReturnsOnCall == nil {
		fake.openFileReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.openFileReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *Fake) RemoveAll(path string) error {
	fake.removeAllMutex.Lock()
	ret, specificReturn := fake.removeAllReturnsOnCall[len(fake.removeAllArgsForCall)]
	fake.removeAllArgsForCall = append(fake.removeAllArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("RemoveAll", []interface{}{path})
	fake.removeAllMutex.Unlock()
	if fake.RemoveAllStub != nil {
		return fake.RemoveAllStub(path)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeAllReturns.result1
}

func (fake *Fake) RemoveAllCallCount() int {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return len(fake.removeAllArgsForCall)
}

func (fake *Fake) RemoveAllArgsForCall(i int) string {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return fake.removeAllArgsForCall[i].path
}

func (fake *Fake) RemoveAllReturns(result1 error) {
	fake.RemoveAllStub = nil
	fake.removeAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) RemoveAllReturnsOnCall(i int, result1 error) {
	fake.RemoveAllStub = nil
	if fake.removeAllReturnsOnCall == nil {
		fake.removeAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) Setenv(key string, value string) error {
	fake.setenvMutex.Lock()
	ret, specificReturn := fake.setenvReturnsOnCall[len(fake.setenvArgsForCall)]
	fake.setenvArgsForCall = append(fake.setenvArgsForCall, struct {
		key   string
		value string
	}{key, value})
	fake.recordInvocation("Setenv", []interface{}{key, value})
	fake.setenvMutex.Unlock()
	if fake.SetenvStub != nil {
		return fake.SetenvStub(key, value)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setenvReturns.result1
}

func (fake *Fake) SetenvCallCount() int {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	return len(fake.setenvArgsForCall)
}

func (fake *Fake) SetenvArgsForCall(i int) (string, string) {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	return fake.setenvArgsForCall[i].key, fake.setenvArgsForCall[i].value
}

func (fake *Fake) SetenvReturns(result1 error) {
	fake.SetenvStub = nil
	fake.setenvReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) SetenvReturnsOnCall(i int, result1 error) {
	fake.SetenvStub = nil
	if fake.setenvReturnsOnCall == nil {
		fake.setenvReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setenvReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) Stat(name string) (os.FileInfo, error) {
	fake.statMutex.Lock()
	ret, specificReturn := fake.statReturnsOnCall[len(fake.statArgsForCall)]
	fake.statArgsForCall = append(fake.statArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Stat", []interface{}{name})
	fake.statMutex.Unlock()
	if fake.StatStub != nil {
		return fake.StatStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.statReturns.result1, fake.statReturns.result2
}

func (fake *Fake) StatCallCount() int {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return len(fake.statArgsForCall)
}

func (fake *Fake) StatArgsForCall(i int) string {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return fake.statArgsForCall[i].name
}

func (fake *Fake) StatReturns(result1 os.FileInfo, result2 error) {
	fake.StatStub = nil
	fake.statReturns = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *Fake) StatReturnsOnCall(i int, result1 os.FileInfo, result2 error) {
	fake.StatStub = nil
	if fake.statReturnsOnCall == nil {
		fake.statReturnsOnCall = make(map[int]struct {
			result1 os.FileInfo
			result2 error
		})
	}
	fake.statReturnsOnCall[i] = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chmodMutex.RLock()
	defer fake.chmodMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	fake.findProcessMutex.RLock()
	defer fake.findProcessMutex.RUnlock()
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	fake.getpidMutex.RLock()
	defer fake.getpidMutex.RUnlock()
	fake.getwdMutex.RLock()
	defer fake.getwdMutex.RUnlock()
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return fake.invocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ IOS = new(Fake)
