// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	mapper "github.com/superlinkx/HomeList/data/mapper"

	model "github.com/superlinkx/HomeList/app/model"
)

// MockListAdapter is an autogenerated mock type for the ListAdapter type
type MockListAdapter struct {
	mock.Mock
}

type MockListAdapter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockListAdapter) EXPECT() *MockListAdapter_Expecter {
	return &MockListAdapter_Expecter{mock: &_m.Mock}
}

// AllLists provides a mock function with given fields: ctx, limit, offset
func (_m *MockListAdapter) AllLists(ctx context.Context, limit int32, offset int32) ([]model.List, error) {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for AllLists")
	}

	var r0 []model.List
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) ([]model.List, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) []model.List); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.List)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, int32) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockListAdapter_AllLists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllLists'
type MockListAdapter_AllLists_Call struct {
	*mock.Call
}

// AllLists is a helper method to define mock.On call
//   - ctx context.Context
//   - limit int32
//   - offset int32
func (_e *MockListAdapter_Expecter) AllLists(ctx interface{}, limit interface{}, offset interface{}) *MockListAdapter_AllLists_Call {
	return &MockListAdapter_AllLists_Call{Call: _e.mock.On("AllLists", ctx, limit, offset)}
}

func (_c *MockListAdapter_AllLists_Call) Run(run func(ctx context.Context, limit int32, offset int32)) *MockListAdapter_AllLists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(int32))
	})
	return _c
}

func (_c *MockListAdapter_AllLists_Call) Return(_a0 []model.List, _a1 error) *MockListAdapter_AllLists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListAdapter_AllLists_Call) RunAndReturn(run func(context.Context, int32, int32) ([]model.List, error)) *MockListAdapter_AllLists_Call {
	_c.Call.Return(run)
	return _c
}

// CreateList provides a mock function with given fields: ctx, name
func (_m *MockListAdapter) CreateList(ctx context.Context, name string) (model.List, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for CreateList")
	}

	var r0 model.List
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.List, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.List); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(model.List)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockListAdapter_CreateList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateList'
type MockListAdapter_CreateList_Call struct {
	*mock.Call
}

// CreateList is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockListAdapter_Expecter) CreateList(ctx interface{}, name interface{}) *MockListAdapter_CreateList_Call {
	return &MockListAdapter_CreateList_Call{Call: _e.mock.On("CreateList", ctx, name)}
}

func (_c *MockListAdapter_CreateList_Call) Run(run func(ctx context.Context, name string)) *MockListAdapter_CreateList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockListAdapter_CreateList_Call) Return(_a0 model.List, _a1 error) *MockListAdapter_CreateList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListAdapter_CreateList_Call) RunAndReturn(run func(context.Context, string) (model.List, error)) *MockListAdapter_CreateList_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteList provides a mock function with given fields: ctx, id
func (_m *MockListAdapter) DeleteList(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockListAdapter_DeleteList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteList'
type MockListAdapter_DeleteList_Call struct {
	*mock.Call
}

// DeleteList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockListAdapter_Expecter) DeleteList(ctx interface{}, id interface{}) *MockListAdapter_DeleteList_Call {
	return &MockListAdapter_DeleteList_Call{Call: _e.mock.On("DeleteList", ctx, id)}
}

func (_c *MockListAdapter_DeleteList_Call) Run(run func(ctx context.Context, id int64)) *MockListAdapter_DeleteList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockListAdapter_DeleteList_Call) Return(_a0 error) *MockListAdapter_DeleteList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockListAdapter_DeleteList_Call) RunAndReturn(run func(context.Context, int64) error) *MockListAdapter_DeleteList_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: ctx, id
func (_m *MockListAdapter) GetList(ctx context.Context, id int64) (model.List, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetList")
	}

	var r0 model.List
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.List, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.List); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.List)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockListAdapter_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockListAdapter_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockListAdapter_Expecter) GetList(ctx interface{}, id interface{}) *MockListAdapter_GetList_Call {
	return &MockListAdapter_GetList_Call{Call: _e.mock.On("GetList", ctx, id)}
}

func (_c *MockListAdapter_GetList_Call) Run(run func(ctx context.Context, id int64)) *MockListAdapter_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockListAdapter_GetList_Call) Return(_a0 model.List, _a1 error) *MockListAdapter_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListAdapter_GetList_Call) RunAndReturn(run func(context.Context, int64) (model.List, error)) *MockListAdapter_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// ReflowList provides a mock function with given fields: ctx, id, reflowMapper
func (_m *MockListAdapter) ReflowList(ctx context.Context, id int64, reflowMapper mapper.Reflow) error {
	ret := _m.Called(ctx, id, reflowMapper)

	if len(ret) == 0 {
		panic("no return value specified for ReflowList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, mapper.Reflow) error); ok {
		r0 = rf(ctx, id, reflowMapper)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockListAdapter_ReflowList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReflowList'
type MockListAdapter_ReflowList_Call struct {
	*mock.Call
}

// ReflowList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
//   - reflowMapper mapper.Reflow
func (_e *MockListAdapter_Expecter) ReflowList(ctx interface{}, id interface{}, reflowMapper interface{}) *MockListAdapter_ReflowList_Call {
	return &MockListAdapter_ReflowList_Call{Call: _e.mock.On("ReflowList", ctx, id, reflowMapper)}
}

func (_c *MockListAdapter_ReflowList_Call) Run(run func(ctx context.Context, id int64, reflowMapper mapper.Reflow)) *MockListAdapter_ReflowList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(mapper.Reflow))
	})
	return _c
}

func (_c *MockListAdapter_ReflowList_Call) Return(_a0 error) *MockListAdapter_ReflowList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockListAdapter_ReflowList_Call) RunAndReturn(run func(context.Context, int64, mapper.Reflow) error) *MockListAdapter_ReflowList_Call {
	_c.Call.Return(run)
	return _c
}

// RenameList provides a mock function with given fields: ctx, id, name
func (_m *MockListAdapter) RenameList(ctx context.Context, id int64, name string) (model.List, error) {
	ret := _m.Called(ctx, id, name)

	if len(ret) == 0 {
		panic("no return value specified for RenameList")
	}

	var r0 model.List
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (model.List, error)); ok {
		return rf(ctx, id, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) model.List); ok {
		r0 = rf(ctx, id, name)
	} else {
		r0 = ret.Get(0).(model.List)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, id, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockListAdapter_RenameList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RenameList'
type MockListAdapter_RenameList_Call struct {
	*mock.Call
}

// RenameList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
//   - name string
func (_e *MockListAdapter_Expecter) RenameList(ctx interface{}, id interface{}, name interface{}) *MockListAdapter_RenameList_Call {
	return &MockListAdapter_RenameList_Call{Call: _e.mock.On("RenameList", ctx, id, name)}
}

func (_c *MockListAdapter_RenameList_Call) Run(run func(ctx context.Context, id int64, name string)) *MockListAdapter_RenameList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MockListAdapter_RenameList_Call) Return(_a0 model.List, _a1 error) *MockListAdapter_RenameList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListAdapter_RenameList_Call) RunAndReturn(run func(context.Context, int64, string) (model.List, error)) *MockListAdapter_RenameList_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockListAdapter creates a new instance of MockListAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockListAdapter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockListAdapter {
	mock := &MockListAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
