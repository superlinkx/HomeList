// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/superlinkx/HomeList/model"
)

// MockAdapter is an autogenerated mock type for the Adapter type
type MockAdapter struct {
	mock.Mock
}

type MockAdapter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAdapter) EXPECT() *MockAdapter_Expecter {
	return &MockAdapter_Expecter{mock: &_m.Mock}
}

// AllItemsFromList provides a mock function with given fields: ctx, listID, limit, offset
func (_m *MockAdapter) AllItemsFromList(ctx context.Context, listID int64, limit int32, offset int32) ([]model.Item, error) {
	ret := _m.Called(ctx, listID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for AllItemsFromList")
	}

	var r0 []model.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int32, int32) ([]model.Item, error)); ok {
		return rf(ctx, listID, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int32, int32) []model.Item); ok {
		r0 = rf(ctx, listID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int32, int32) error); ok {
		r1 = rf(ctx, listID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_AllItemsFromList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllItemsFromList'
type MockAdapter_AllItemsFromList_Call struct {
	*mock.Call
}

// AllItemsFromList is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - limit int32
//   - offset int32
func (_e *MockAdapter_Expecter) AllItemsFromList(ctx interface{}, listID interface{}, limit interface{}, offset interface{}) *MockAdapter_AllItemsFromList_Call {
	return &MockAdapter_AllItemsFromList_Call{Call: _e.mock.On("AllItemsFromList", ctx, listID, limit, offset)}
}

func (_c *MockAdapter_AllItemsFromList_Call) Run(run func(ctx context.Context, listID int64, limit int32, offset int32)) *MockAdapter_AllItemsFromList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int32), args[3].(int32))
	})
	return _c
}

func (_c *MockAdapter_AllItemsFromList_Call) Return(_a0 []model.Item, _a1 error) *MockAdapter_AllItemsFromList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdapter_AllItemsFromList_Call) RunAndReturn(run func(context.Context, int64, int32, int32) ([]model.Item, error)) *MockAdapter_AllItemsFromList_Call {
	_c.Call.Return(run)
	return _c
}

// AllLists provides a mock function with given fields: ctx, limit, offset
func (_m *MockAdapter) AllLists(ctx context.Context, limit int32, offset int32) ([]model.List, error) {
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

// MockAdapter_AllLists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllLists'
type MockAdapter_AllLists_Call struct {
	*mock.Call
}

// AllLists is a helper method to define mock.On call
//   - ctx context.Context
//   - limit int32
//   - offset int32
func (_e *MockAdapter_Expecter) AllLists(ctx interface{}, limit interface{}, offset interface{}) *MockAdapter_AllLists_Call {
	return &MockAdapter_AllLists_Call{Call: _e.mock.On("AllLists", ctx, limit, offset)}
}

func (_c *MockAdapter_AllLists_Call) Run(run func(ctx context.Context, limit int32, offset int32)) *MockAdapter_AllLists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(int32))
	})
	return _c
}

func (_c *MockAdapter_AllLists_Call) Return(_a0 []model.List, _a1 error) *MockAdapter_AllLists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdapter_AllLists_Call) RunAndReturn(run func(context.Context, int32, int32) ([]model.List, error)) *MockAdapter_AllLists_Call {
	_c.Call.Return(run)
	return _c
}

// CreateItemOnList provides a mock function with given fields: ctx, listID, content, sort
func (_m *MockAdapter) CreateItemOnList(ctx context.Context, listID int64, content string, sort int64) (model.Item, error) {
	ret := _m.Called(ctx, listID, content, sort)

	if len(ret) == 0 {
		panic("no return value specified for CreateItemOnList")
	}

	var r0 model.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, int64) (model.Item, error)); ok {
		return rf(ctx, listID, content, sort)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, int64) model.Item); ok {
		r0 = rf(ctx, listID, content, sort)
	} else {
		r0 = ret.Get(0).(model.Item)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, int64) error); ok {
		r1 = rf(ctx, listID, content, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_CreateItemOnList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateItemOnList'
type MockAdapter_CreateItemOnList_Call struct {
	*mock.Call
}

// CreateItemOnList is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - content string
//   - sort int64
func (_e *MockAdapter_Expecter) CreateItemOnList(ctx interface{}, listID interface{}, content interface{}, sort interface{}) *MockAdapter_CreateItemOnList_Call {
	return &MockAdapter_CreateItemOnList_Call{Call: _e.mock.On("CreateItemOnList", ctx, listID, content, sort)}
}

func (_c *MockAdapter_CreateItemOnList_Call) Run(run func(ctx context.Context, listID int64, content string, sort int64)) *MockAdapter_CreateItemOnList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string), args[3].(int64))
	})
	return _c
}

func (_c *MockAdapter_CreateItemOnList_Call) Return(_a0 model.Item, _a1 error) *MockAdapter_CreateItemOnList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdapter_CreateItemOnList_Call) RunAndReturn(run func(context.Context, int64, string, int64) (model.Item, error)) *MockAdapter_CreateItemOnList_Call {
	_c.Call.Return(run)
	return _c
}

// CreateList provides a mock function with given fields: ctx, name
func (_m *MockAdapter) CreateList(ctx context.Context, name string) (model.List, error) {
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

// MockAdapter_CreateList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateList'
type MockAdapter_CreateList_Call struct {
	*mock.Call
}

// CreateList is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockAdapter_Expecter) CreateList(ctx interface{}, name interface{}) *MockAdapter_CreateList_Call {
	return &MockAdapter_CreateList_Call{Call: _e.mock.On("CreateList", ctx, name)}
}

func (_c *MockAdapter_CreateList_Call) Run(run func(ctx context.Context, name string)) *MockAdapter_CreateList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAdapter_CreateList_Call) Return(_a0 model.List, _a1 error) *MockAdapter_CreateList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdapter_CreateList_Call) RunAndReturn(run func(context.Context, string) (model.List, error)) *MockAdapter_CreateList_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteItemFromList provides a mock function with given fields: ctx, listID, itemID
func (_m *MockAdapter) DeleteItemFromList(ctx context.Context, listID int64, itemID int64) error {
	ret := _m.Called(ctx, listID, itemID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteItemFromList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, listID, itemID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAdapter_DeleteItemFromList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteItemFromList'
type MockAdapter_DeleteItemFromList_Call struct {
	*mock.Call
}

// DeleteItemFromList is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - itemID int64
func (_e *MockAdapter_Expecter) DeleteItemFromList(ctx interface{}, listID interface{}, itemID interface{}) *MockAdapter_DeleteItemFromList_Call {
	return &MockAdapter_DeleteItemFromList_Call{Call: _e.mock.On("DeleteItemFromList", ctx, listID, itemID)}
}

func (_c *MockAdapter_DeleteItemFromList_Call) Run(run func(ctx context.Context, listID int64, itemID int64)) *MockAdapter_DeleteItemFromList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *MockAdapter_DeleteItemFromList_Call) Return(_a0 error) *MockAdapter_DeleteItemFromList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAdapter_DeleteItemFromList_Call) RunAndReturn(run func(context.Context, int64, int64) error) *MockAdapter_DeleteItemFromList_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteList provides a mock function with given fields: ctx, id
func (_m *MockAdapter) DeleteList(ctx context.Context, id int64) error {
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

// MockAdapter_DeleteList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteList'
type MockAdapter_DeleteList_Call struct {
	*mock.Call
}

// DeleteList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockAdapter_Expecter) DeleteList(ctx interface{}, id interface{}) *MockAdapter_DeleteList_Call {
	return &MockAdapter_DeleteList_Call{Call: _e.mock.On("DeleteList", ctx, id)}
}

func (_c *MockAdapter_DeleteList_Call) Run(run func(ctx context.Context, id int64)) *MockAdapter_DeleteList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockAdapter_DeleteList_Call) Return(_a0 error) *MockAdapter_DeleteList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAdapter_DeleteList_Call) RunAndReturn(run func(context.Context, int64) error) *MockAdapter_DeleteList_Call {
	_c.Call.Return(run)
	return _c
}

// GetItemFromList provides a mock function with given fields: ctx, listID, itemID
func (_m *MockAdapter) GetItemFromList(ctx context.Context, listID int64, itemID int64) (model.Item, error) {
	ret := _m.Called(ctx, listID, itemID)

	if len(ret) == 0 {
		panic("no return value specified for GetItemFromList")
	}

	var r0 model.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) (model.Item, error)); ok {
		return rf(ctx, listID, itemID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) model.Item); ok {
		r0 = rf(ctx, listID, itemID)
	} else {
		r0 = ret.Get(0).(model.Item)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, listID, itemID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetItemFromList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetItemFromList'
type MockAdapter_GetItemFromList_Call struct {
	*mock.Call
}

// GetItemFromList is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - itemID int64
func (_e *MockAdapter_Expecter) GetItemFromList(ctx interface{}, listID interface{}, itemID interface{}) *MockAdapter_GetItemFromList_Call {
	return &MockAdapter_GetItemFromList_Call{Call: _e.mock.On("GetItemFromList", ctx, listID, itemID)}
}

func (_c *MockAdapter_GetItemFromList_Call) Run(run func(ctx context.Context, listID int64, itemID int64)) *MockAdapter_GetItemFromList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *MockAdapter_GetItemFromList_Call) Return(_a0 model.Item, _a1 error) *MockAdapter_GetItemFromList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdapter_GetItemFromList_Call) RunAndReturn(run func(context.Context, int64, int64) (model.Item, error)) *MockAdapter_GetItemFromList_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: ctx, id
func (_m *MockAdapter) GetList(ctx context.Context, id int64) (model.List, error) {
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

// MockAdapter_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockAdapter_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockAdapter_Expecter) GetList(ctx interface{}, id interface{}) *MockAdapter_GetList_Call {
	return &MockAdapter_GetList_Call{Call: _e.mock.On("GetList", ctx, id)}
}

func (_c *MockAdapter_GetList_Call) Run(run func(ctx context.Context, id int64)) *MockAdapter_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockAdapter_GetList_Call) Return(_a0 model.List, _a1 error) *MockAdapter_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdapter_GetList_Call) RunAndReturn(run func(context.Context, int64) (model.List, error)) *MockAdapter_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// RenameList provides a mock function with given fields: ctx, id, name
func (_m *MockAdapter) RenameList(ctx context.Context, id int64, name string) (model.List, error) {
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

// MockAdapter_RenameList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RenameList'
type MockAdapter_RenameList_Call struct {
	*mock.Call
}

// RenameList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
//   - name string
func (_e *MockAdapter_Expecter) RenameList(ctx interface{}, id interface{}, name interface{}) *MockAdapter_RenameList_Call {
	return &MockAdapter_RenameList_Call{Call: _e.mock.On("RenameList", ctx, id, name)}
}

func (_c *MockAdapter_RenameList_Call) Run(run func(ctx context.Context, id int64, name string)) *MockAdapter_RenameList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MockAdapter_RenameList_Call) Return(_a0 model.List, _a1 error) *MockAdapter_RenameList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdapter_RenameList_Call) RunAndReturn(run func(context.Context, int64, string) (model.List, error)) *MockAdapter_RenameList_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateItemFromList provides a mock function with given fields: ctx, listID, itemID, content, checked, sort
func (_m *MockAdapter) UpdateItemFromList(ctx context.Context, listID int64, itemID int64, content string, checked bool, sort int64) (model.Item, error) {
	ret := _m.Called(ctx, listID, itemID, content, checked, sort)

	if len(ret) == 0 {
		panic("no return value specified for UpdateItemFromList")
	}

	var r0 model.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, bool, int64) (model.Item, error)); ok {
		return rf(ctx, listID, itemID, content, checked, sort)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, bool, int64) model.Item); ok {
		r0 = rf(ctx, listID, itemID, content, checked, sort)
	} else {
		r0 = ret.Get(0).(model.Item)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, string, bool, int64) error); ok {
		r1 = rf(ctx, listID, itemID, content, checked, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_UpdateItemFromList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateItemFromList'
type MockAdapter_UpdateItemFromList_Call struct {
	*mock.Call
}

// UpdateItemFromList is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - itemID int64
//   - content string
//   - checked bool
//   - sort int64
func (_e *MockAdapter_Expecter) UpdateItemFromList(ctx interface{}, listID interface{}, itemID interface{}, content interface{}, checked interface{}, sort interface{}) *MockAdapter_UpdateItemFromList_Call {
	return &MockAdapter_UpdateItemFromList_Call{Call: _e.mock.On("UpdateItemFromList", ctx, listID, itemID, content, checked, sort)}
}

func (_c *MockAdapter_UpdateItemFromList_Call) Run(run func(ctx context.Context, listID int64, itemID int64, content string, checked bool, sort int64)) *MockAdapter_UpdateItemFromList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64), args[3].(string), args[4].(bool), args[5].(int64))
	})
	return _c
}

func (_c *MockAdapter_UpdateItemFromList_Call) Return(_a0 model.Item, _a1 error) *MockAdapter_UpdateItemFromList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdapter_UpdateItemFromList_Call) RunAndReturn(run func(context.Context, int64, int64, string, bool, int64) (model.Item, error)) *MockAdapter_UpdateItemFromList_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAdapter creates a new instance of MockAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAdapter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAdapter {
	mock := &MockAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
