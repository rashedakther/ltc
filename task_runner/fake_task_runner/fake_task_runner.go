// This file was generated by counterfeiter
package fake_task_runner

import (
	"sync"

	"github.com/cloudfoundry-incubator/ltc/task_runner"
)

type FakeTaskRunner struct {
	CreateTaskStub        func(createTaskParams task_runner.CreateTaskParams) error
	createTaskMutex       sync.RWMutex
	createTaskArgsForCall []struct {
		createTaskParams task_runner.CreateTaskParams
	}
	createTaskReturns struct {
		result1 error
	}
	SubmitTaskStub        func(submitTaskJson []byte) (string, error)
	submitTaskMutex       sync.RWMutex
	submitTaskArgsForCall []struct {
		submitTaskJson []byte
	}
	submitTaskReturns struct {
		result1 string
		result2 error
	}
	DeleteTaskStub        func(taskGuid string) error
	deleteTaskMutex       sync.RWMutex
	deleteTaskArgsForCall []struct {
		taskGuid string
	}
	deleteTaskReturns struct {
		result1 error
	}
	CancelTaskStub        func(taskGuid string) error
	cancelTaskMutex       sync.RWMutex
	cancelTaskArgsForCall []struct {
		taskGuid string
	}
	cancelTaskReturns struct {
		result1 error
	}
}

func (fake *FakeTaskRunner) CreateTask(createTaskParams task_runner.CreateTaskParams) error {
	fake.createTaskMutex.Lock()
	fake.createTaskArgsForCall = append(fake.createTaskArgsForCall, struct {
		createTaskParams task_runner.CreateTaskParams
	}{createTaskParams})
	fake.createTaskMutex.Unlock()
	if fake.CreateTaskStub != nil {
		return fake.CreateTaskStub(createTaskParams)
	} else {
		return fake.createTaskReturns.result1
	}
}

func (fake *FakeTaskRunner) CreateTaskCallCount() int {
	fake.createTaskMutex.RLock()
	defer fake.createTaskMutex.RUnlock()
	return len(fake.createTaskArgsForCall)
}

func (fake *FakeTaskRunner) CreateTaskArgsForCall(i int) task_runner.CreateTaskParams {
	fake.createTaskMutex.RLock()
	defer fake.createTaskMutex.RUnlock()
	return fake.createTaskArgsForCall[i].createTaskParams
}

func (fake *FakeTaskRunner) CreateTaskReturns(result1 error) {
	fake.CreateTaskStub = nil
	fake.createTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskRunner) SubmitTask(submitTaskJson []byte) (string, error) {
	fake.submitTaskMutex.Lock()
	fake.submitTaskArgsForCall = append(fake.submitTaskArgsForCall, struct {
		submitTaskJson []byte
	}{submitTaskJson})
	fake.submitTaskMutex.Unlock()
	if fake.SubmitTaskStub != nil {
		return fake.SubmitTaskStub(submitTaskJson)
	} else {
		return fake.submitTaskReturns.result1, fake.submitTaskReturns.result2
	}
}

func (fake *FakeTaskRunner) SubmitTaskCallCount() int {
	fake.submitTaskMutex.RLock()
	defer fake.submitTaskMutex.RUnlock()
	return len(fake.submitTaskArgsForCall)
}

func (fake *FakeTaskRunner) SubmitTaskArgsForCall(i int) []byte {
	fake.submitTaskMutex.RLock()
	defer fake.submitTaskMutex.RUnlock()
	return fake.submitTaskArgsForCall[i].submitTaskJson
}

func (fake *FakeTaskRunner) SubmitTaskReturns(result1 string, result2 error) {
	fake.SubmitTaskStub = nil
	fake.submitTaskReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskRunner) DeleteTask(taskGuid string) error {
	fake.deleteTaskMutex.Lock()
	fake.deleteTaskArgsForCall = append(fake.deleteTaskArgsForCall, struct {
		taskGuid string
	}{taskGuid})
	fake.deleteTaskMutex.Unlock()
	if fake.DeleteTaskStub != nil {
		return fake.DeleteTaskStub(taskGuid)
	} else {
		return fake.deleteTaskReturns.result1
	}
}

func (fake *FakeTaskRunner) DeleteTaskCallCount() int {
	fake.deleteTaskMutex.RLock()
	defer fake.deleteTaskMutex.RUnlock()
	return len(fake.deleteTaskArgsForCall)
}

func (fake *FakeTaskRunner) DeleteTaskArgsForCall(i int) string {
	fake.deleteTaskMutex.RLock()
	defer fake.deleteTaskMutex.RUnlock()
	return fake.deleteTaskArgsForCall[i].taskGuid
}

func (fake *FakeTaskRunner) DeleteTaskReturns(result1 error) {
	fake.DeleteTaskStub = nil
	fake.deleteTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskRunner) CancelTask(taskGuid string) error {
	fake.cancelTaskMutex.Lock()
	fake.cancelTaskArgsForCall = append(fake.cancelTaskArgsForCall, struct {
		taskGuid string
	}{taskGuid})
	fake.cancelTaskMutex.Unlock()
	if fake.CancelTaskStub != nil {
		return fake.CancelTaskStub(taskGuid)
	} else {
		return fake.cancelTaskReturns.result1
	}
}

func (fake *FakeTaskRunner) CancelTaskCallCount() int {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return len(fake.cancelTaskArgsForCall)
}

func (fake *FakeTaskRunner) CancelTaskArgsForCall(i int) string {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return fake.cancelTaskArgsForCall[i].taskGuid
}

func (fake *FakeTaskRunner) CancelTaskReturns(result1 error) {
	fake.CancelTaskStub = nil
	fake.cancelTaskReturns = struct {
		result1 error
	}{result1}
}

var _ task_runner.TaskRunner = new(FakeTaskRunner)