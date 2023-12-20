// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	list "github.com/superlinkx/HomeList/app/list"

	listitem "github.com/superlinkx/HomeList/app/listitem"

	mock "github.com/stretchr/testify/mock"
)

// MockApplication is an autogenerated mock type for the application type
type MockApplication struct {
	mock.Mock
}

type MockApplication_Expecter struct {
	mock *mock.Mock
}

func (_m *MockApplication) EXPECT() *MockApplication_Expecter {
	return &MockApplication_Expecter{mock: &_m.Mock}
}

// AddItemToList provides a mock function with given fields: ctx, listID, content, sort
func (_m *MockApplication) AddItemToList(ctx context.Context, listID int64, content string, sort int64) (listitem.ListItem, error) {
	ret := _m.Called(ctx, listID, content, sort)

	if len(ret) == 0 {
		panic("no return value specified for AddItemToList")
	}

	var r0 listitem.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, int64) (listitem.ListItem, error)); ok {
		return rf(ctx, listID, content, sort)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, int64) listitem.ListItem); ok {
		r0 = rf(ctx, listID, content, sort)
	} else {
		r0 = ret.Get(0).(listitem.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, int64) error); ok {
		r1 = rf(ctx, listID, content, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_AddItemToList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddItemToList'
type MockApplication_AddItemToList_Call struct {
	*mock.Call
}

// AddItemToList is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - content string
//   - sort int64
func (_e *MockApplication_Expecter) AddItemToList(ctx interface{}, listID interface{}, content interface{}, sort interface{}) *MockApplication_AddItemToList_Call {
	return &MockApplication_AddItemToList_Call{Call: _e.mock.On("AddItemToList", ctx, listID, content, sort)}
}

func (_c *MockApplication_AddItemToList_Call) Run(run func(ctx context.Context, listID int64, content string, sort int64)) *MockApplication_AddItemToList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string), args[3].(int64))
	})
	return _c
}

func (_c *MockApplication_AddItemToList_Call) Return(_a0 listitem.ListItem, _a1 error) *MockApplication_AddItemToList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_AddItemToList_Call) RunAndReturn(run func(context.Context, int64, string, int64) (listitem.ListItem, error)) *MockApplication_AddItemToList_Call {
	_c.Call.Return(run)
	return _c
}

// AllLists provides a mock function with given fields: ctx, limit
func (_m *MockApplication) AllLists(ctx context.Context, limit int64) ([]list.List, error) {
	ret := _m.Called(ctx, limit)

	if len(ret) == 0 {
		panic("no return value specified for AllLists")
	}

	var r0 []list.List
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]list.List, error)); ok {
		return rf(ctx, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []list.List); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]list.List)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_AllLists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllLists'
type MockApplication_AllLists_Call struct {
	*mock.Call
}

// AllLists is a helper method to define mock.On call
//   - ctx context.Context
//   - limit int64
func (_e *MockApplication_Expecter) AllLists(ctx interface{}, limit interface{}) *MockApplication_AllLists_Call {
	return &MockApplication_AllLists_Call{Call: _e.mock.On("AllLists", ctx, limit)}
}

func (_c *MockApplication_AllLists_Call) Run(run func(ctx context.Context, limit int64)) *MockApplication_AllLists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockApplication_AllLists_Call) Return(_a0 []list.List, _a1 error) *MockApplication_AllLists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_AllLists_Call) RunAndReturn(run func(context.Context, int64) ([]list.List, error)) *MockApplication_AllLists_Call {
	_c.Call.Return(run)
	return _c
}

// CreateList provides a mock function with given fields: ctx, name
func (_m *MockApplication) CreateList(ctx context.Context, name string) (list.List, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for CreateList")
	}

	var r0 list.List
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (list.List, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) list.List); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(list.List)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_CreateList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateList'
type MockApplication_CreateList_Call struct {
	*mock.Call
}

// CreateList is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockApplication_Expecter) CreateList(ctx interface{}, name interface{}) *MockApplication_CreateList_Call {
	return &MockApplication_CreateList_Call{Call: _e.mock.On("CreateList", ctx, name)}
}

func (_c *MockApplication_CreateList_Call) Run(run func(ctx context.Context, name string)) *MockApplication_CreateList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockApplication_CreateList_Call) Return(_a0 list.List, _a1 error) *MockApplication_CreateList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_CreateList_Call) RunAndReturn(run func(context.Context, string) (list.List, error)) *MockApplication_CreateList_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteList provides a mock function with given fields: ctx, id
func (_m *MockApplication) DeleteList(ctx context.Context, id int64) error {
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

// MockApplication_DeleteList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteList'
type MockApplication_DeleteList_Call struct {
	*mock.Call
}

// DeleteList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockApplication_Expecter) DeleteList(ctx interface{}, id interface{}) *MockApplication_DeleteList_Call {
	return &MockApplication_DeleteList_Call{Call: _e.mock.On("DeleteList", ctx, id)}
}

func (_c *MockApplication_DeleteList_Call) Run(run func(ctx context.Context, id int64)) *MockApplication_DeleteList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockApplication_DeleteList_Call) Return(_a0 error) *MockApplication_DeleteList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockApplication_DeleteList_Call) RunAndReturn(run func(context.Context, int64) error) *MockApplication_DeleteList_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteListItem provides a mock function with given fields: ctx, id
func (_m *MockApplication) DeleteListItem(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteListItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockApplication_DeleteListItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteListItem'
type MockApplication_DeleteListItem_Call struct {
	*mock.Call
}

// DeleteListItem is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockApplication_Expecter) DeleteListItem(ctx interface{}, id interface{}) *MockApplication_DeleteListItem_Call {
	return &MockApplication_DeleteListItem_Call{Call: _e.mock.On("DeleteListItem", ctx, id)}
}

func (_c *MockApplication_DeleteListItem_Call) Run(run func(ctx context.Context, id int64)) *MockApplication_DeleteListItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockApplication_DeleteListItem_Call) Return(_a0 error) *MockApplication_DeleteListItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockApplication_DeleteListItem_Call) RunAndReturn(run func(context.Context, int64) error) *MockApplication_DeleteListItem_Call {
	_c.Call.Return(run)
	return _c
}

// FetchAllItemsFromList provides a mock function with given fields: ctx, listID, limit
func (_m *MockApplication) FetchAllItemsFromList(ctx context.Context, listID int64, limit int64) ([]listitem.ListItem, error) {
	ret := _m.Called(ctx, listID, limit)

	if len(ret) == 0 {
		panic("no return value specified for FetchAllItemsFromList")
	}

	var r0 []listitem.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) ([]listitem.ListItem, error)); ok {
		return rf(ctx, listID, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []listitem.ListItem); ok {
		r0 = rf(ctx, listID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]listitem.ListItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, listID, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_FetchAllItemsFromList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchAllItemsFromList'
type MockApplication_FetchAllItemsFromList_Call struct {
	*mock.Call
}

// FetchAllItemsFromList is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - limit int64
func (_e *MockApplication_Expecter) FetchAllItemsFromList(ctx interface{}, listID interface{}, limit interface{}) *MockApplication_FetchAllItemsFromList_Call {
	return &MockApplication_FetchAllItemsFromList_Call{Call: _e.mock.On("FetchAllItemsFromList", ctx, listID, limit)}
}

func (_c *MockApplication_FetchAllItemsFromList_Call) Run(run func(ctx context.Context, listID int64, limit int64)) *MockApplication_FetchAllItemsFromList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *MockApplication_FetchAllItemsFromList_Call) Return(_a0 []listitem.ListItem, _a1 error) *MockApplication_FetchAllItemsFromList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_FetchAllItemsFromList_Call) RunAndReturn(run func(context.Context, int64, int64) ([]listitem.ListItem, error)) *MockApplication_FetchAllItemsFromList_Call {
	_c.Call.Return(run)
	return _c
}

// FetchListItem provides a mock function with given fields: ctx, id
func (_m *MockApplication) FetchListItem(ctx context.Context, id int64) (listitem.ListItem, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FetchListItem")
	}

	var r0 listitem.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (listitem.ListItem, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) listitem.ListItem); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(listitem.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_FetchListItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchListItem'
type MockApplication_FetchListItem_Call struct {
	*mock.Call
}

// FetchListItem is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockApplication_Expecter) FetchListItem(ctx interface{}, id interface{}) *MockApplication_FetchListItem_Call {
	return &MockApplication_FetchListItem_Call{Call: _e.mock.On("FetchListItem", ctx, id)}
}

func (_c *MockApplication_FetchListItem_Call) Run(run func(ctx context.Context, id int64)) *MockApplication_FetchListItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockApplication_FetchListItem_Call) Return(_a0 listitem.ListItem, _a1 error) *MockApplication_FetchListItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_FetchListItem_Call) RunAndReturn(run func(context.Context, int64) (listitem.ListItem, error)) *MockApplication_FetchListItem_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: ctx, id
func (_m *MockApplication) GetList(ctx context.Context, id int64) (list.List, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetList")
	}

	var r0 list.List
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (list.List, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) list.List); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(list.List)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockApplication_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockApplication_Expecter) GetList(ctx interface{}, id interface{}) *MockApplication_GetList_Call {
	return &MockApplication_GetList_Call{Call: _e.mock.On("GetList", ctx, id)}
}

func (_c *MockApplication_GetList_Call) Run(run func(ctx context.Context, id int64)) *MockApplication_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockApplication_GetList_Call) Return(_a0 list.List, _a1 error) *MockApplication_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_GetList_Call) RunAndReturn(run func(context.Context, int64) (list.List, error)) *MockApplication_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateList provides a mock function with given fields: ctx, id, name
func (_m *MockApplication) UpdateList(ctx context.Context, id int64, name string) (list.List, error) {
	ret := _m.Called(ctx, id, name)

	if len(ret) == 0 {
		panic("no return value specified for UpdateList")
	}

	var r0 list.List
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (list.List, error)); ok {
		return rf(ctx, id, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) list.List); ok {
		r0 = rf(ctx, id, name)
	} else {
		r0 = ret.Get(0).(list.List)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, id, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_UpdateList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateList'
type MockApplication_UpdateList_Call struct {
	*mock.Call
}

// UpdateList is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
//   - name string
func (_e *MockApplication_Expecter) UpdateList(ctx interface{}, id interface{}, name interface{}) *MockApplication_UpdateList_Call {
	return &MockApplication_UpdateList_Call{Call: _e.mock.On("UpdateList", ctx, id, name)}
}

func (_c *MockApplication_UpdateList_Call) Run(run func(ctx context.Context, id int64, name string)) *MockApplication_UpdateList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MockApplication_UpdateList_Call) Return(_a0 list.List, _a1 error) *MockApplication_UpdateList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_UpdateList_Call) RunAndReturn(run func(context.Context, int64, string) (list.List, error)) *MockApplication_UpdateList_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateListItemChecked provides a mock function with given fields: ctx, id, checked
func (_m *MockApplication) UpdateListItemChecked(ctx context.Context, id int64, checked bool) (listitem.ListItem, error) {
	ret := _m.Called(ctx, id, checked)

	if len(ret) == 0 {
		panic("no return value specified for UpdateListItemChecked")
	}

	var r0 listitem.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, bool) (listitem.ListItem, error)); ok {
		return rf(ctx, id, checked)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, bool) listitem.ListItem); ok {
		r0 = rf(ctx, id, checked)
	} else {
		r0 = ret.Get(0).(listitem.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, bool) error); ok {
		r1 = rf(ctx, id, checked)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_UpdateListItemChecked_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateListItemChecked'
type MockApplication_UpdateListItemChecked_Call struct {
	*mock.Call
}

// UpdateListItemChecked is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
//   - checked bool
func (_e *MockApplication_Expecter) UpdateListItemChecked(ctx interface{}, id interface{}, checked interface{}) *MockApplication_UpdateListItemChecked_Call {
	return &MockApplication_UpdateListItemChecked_Call{Call: _e.mock.On("UpdateListItemChecked", ctx, id, checked)}
}

func (_c *MockApplication_UpdateListItemChecked_Call) Run(run func(ctx context.Context, id int64, checked bool)) *MockApplication_UpdateListItemChecked_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(bool))
	})
	return _c
}

func (_c *MockApplication_UpdateListItemChecked_Call) Return(_a0 listitem.ListItem, _a1 error) *MockApplication_UpdateListItemChecked_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_UpdateListItemChecked_Call) RunAndReturn(run func(context.Context, int64, bool) (listitem.ListItem, error)) *MockApplication_UpdateListItemChecked_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateListItemContent provides a mock function with given fields: ctx, id, content
func (_m *MockApplication) UpdateListItemContent(ctx context.Context, id int64, content string) (listitem.ListItem, error) {
	ret := _m.Called(ctx, id, content)

	if len(ret) == 0 {
		panic("no return value specified for UpdateListItemContent")
	}

	var r0 listitem.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (listitem.ListItem, error)); ok {
		return rf(ctx, id, content)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) listitem.ListItem); ok {
		r0 = rf(ctx, id, content)
	} else {
		r0 = ret.Get(0).(listitem.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, id, content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_UpdateListItemContent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateListItemContent'
type MockApplication_UpdateListItemContent_Call struct {
	*mock.Call
}

// UpdateListItemContent is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
//   - content string
func (_e *MockApplication_Expecter) UpdateListItemContent(ctx interface{}, id interface{}, content interface{}) *MockApplication_UpdateListItemContent_Call {
	return &MockApplication_UpdateListItemContent_Call{Call: _e.mock.On("UpdateListItemContent", ctx, id, content)}
}

func (_c *MockApplication_UpdateListItemContent_Call) Run(run func(ctx context.Context, id int64, content string)) *MockApplication_UpdateListItemContent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MockApplication_UpdateListItemContent_Call) Return(_a0 listitem.ListItem, _a1 error) *MockApplication_UpdateListItemContent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_UpdateListItemContent_Call) RunAndReturn(run func(context.Context, int64, string) (listitem.ListItem, error)) *MockApplication_UpdateListItemContent_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateListItemSort provides a mock function with given fields: ctx, id, sort
func (_m *MockApplication) UpdateListItemSort(ctx context.Context, id int64, sort int64) (listitem.ListItem, error) {
	ret := _m.Called(ctx, id, sort)

	if len(ret) == 0 {
		panic("no return value specified for UpdateListItemSort")
	}

	var r0 listitem.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) (listitem.ListItem, error)); ok {
		return rf(ctx, id, sort)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) listitem.ListItem); ok {
		r0 = rf(ctx, id, sort)
	} else {
		r0 = ret.Get(0).(listitem.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, id, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockApplication_UpdateListItemSort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateListItemSort'
type MockApplication_UpdateListItemSort_Call struct {
	*mock.Call
}

// UpdateListItemSort is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
//   - sort int64
func (_e *MockApplication_Expecter) UpdateListItemSort(ctx interface{}, id interface{}, sort interface{}) *MockApplication_UpdateListItemSort_Call {
	return &MockApplication_UpdateListItemSort_Call{Call: _e.mock.On("UpdateListItemSort", ctx, id, sort)}
}

func (_c *MockApplication_UpdateListItemSort_Call) Run(run func(ctx context.Context, id int64, sort int64)) *MockApplication_UpdateListItemSort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *MockApplication_UpdateListItemSort_Call) Return(_a0 listitem.ListItem, _a1 error) *MockApplication_UpdateListItemSort_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockApplication_UpdateListItemSort_Call) RunAndReturn(run func(context.Context, int64, int64) (listitem.ListItem, error)) *MockApplication_UpdateListItemSort_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockApplication creates a new instance of MockApplication. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockApplication(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockApplication {
	mock := &MockApplication{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
