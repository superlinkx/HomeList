// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sqlite "github.com/superlinkx/HomeList/db/sqlite"
)

// Queries is an autogenerated mock type for the queries type
type Queries struct {
	mock.Mock
}

type Queries_Expecter struct {
	mock *mock.Mock
}

func (_m *Queries) EXPECT() *Queries_Expecter {
	return &Queries_Expecter{mock: &_m.Mock}
}

// AllItemsFromList provides a mock function with given fields: ctx, params
func (_m *Queries) AllItemsFromList(ctx context.Context, params sqlite.AllItemsFromListParams) ([]sqlite.ListItem, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for AllItemsFromList")
	}

	var r0 []sqlite.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.AllItemsFromListParams) ([]sqlite.ListItem, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.AllItemsFromListParams) []sqlite.ListItem); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sqlite.ListItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sqlite.AllItemsFromListParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Queries_AllItemsFromList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllItemsFromList'
type Queries_AllItemsFromList_Call struct {
	*mock.Call
}

// AllItemsFromList is a helper method to define mock.On call
//   - ctx context.Context
//   - params sqlite.AllItemsFromListParams
func (_e *Queries_Expecter) AllItemsFromList(ctx interface{}, params interface{}) *Queries_AllItemsFromList_Call {
	return &Queries_AllItemsFromList_Call{Call: _e.mock.On("AllItemsFromList", ctx, params)}
}

func (_c *Queries_AllItemsFromList_Call) Run(run func(ctx context.Context, params sqlite.AllItemsFromListParams)) *Queries_AllItemsFromList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sqlite.AllItemsFromListParams))
	})
	return _c
}

func (_c *Queries_AllItemsFromList_Call) Return(_a0 []sqlite.ListItem, _a1 error) *Queries_AllItemsFromList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Queries_AllItemsFromList_Call) RunAndReturn(run func(context.Context, sqlite.AllItemsFromListParams) ([]sqlite.ListItem, error)) *Queries_AllItemsFromList_Call {
	_c.Call.Return(run)
	return _c
}

// CreateListItem provides a mock function with given fields: ctx, params
func (_m *Queries) CreateListItem(ctx context.Context, params sqlite.CreateListItemParams) (sqlite.ListItem, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for CreateListItem")
	}

	var r0 sqlite.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.CreateListItemParams) (sqlite.ListItem, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.CreateListItemParams) sqlite.ListItem); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(sqlite.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, sqlite.CreateListItemParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Queries_CreateListItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateListItem'
type Queries_CreateListItem_Call struct {
	*mock.Call
}

// CreateListItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params sqlite.CreateListItemParams
func (_e *Queries_Expecter) CreateListItem(ctx interface{}, params interface{}) *Queries_CreateListItem_Call {
	return &Queries_CreateListItem_Call{Call: _e.mock.On("CreateListItem", ctx, params)}
}

func (_c *Queries_CreateListItem_Call) Run(run func(ctx context.Context, params sqlite.CreateListItemParams)) *Queries_CreateListItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sqlite.CreateListItemParams))
	})
	return _c
}

func (_c *Queries_CreateListItem_Call) Return(_a0 sqlite.ListItem, _a1 error) *Queries_CreateListItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Queries_CreateListItem_Call) RunAndReturn(run func(context.Context, sqlite.CreateListItemParams) (sqlite.ListItem, error)) *Queries_CreateListItem_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteListItem provides a mock function with given fields: ctx, id
func (_m *Queries) DeleteListItem(ctx context.Context, id int64) error {
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

// Queries_DeleteListItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteListItem'
type Queries_DeleteListItem_Call struct {
	*mock.Call
}

// DeleteListItem is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *Queries_Expecter) DeleteListItem(ctx interface{}, id interface{}) *Queries_DeleteListItem_Call {
	return &Queries_DeleteListItem_Call{Call: _e.mock.On("DeleteListItem", ctx, id)}
}

func (_c *Queries_DeleteListItem_Call) Run(run func(ctx context.Context, id int64)) *Queries_DeleteListItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Queries_DeleteListItem_Call) Return(_a0 error) *Queries_DeleteListItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Queries_DeleteListItem_Call) RunAndReturn(run func(context.Context, int64) error) *Queries_DeleteListItem_Call {
	_c.Call.Return(run)
	return _c
}

// GetListItem provides a mock function with given fields: ctx, id
func (_m *Queries) GetListItem(ctx context.Context, id int64) (sqlite.ListItem, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetListItem")
	}

	var r0 sqlite.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (sqlite.ListItem, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) sqlite.ListItem); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(sqlite.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Queries_GetListItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListItem'
type Queries_GetListItem_Call struct {
	*mock.Call
}

// GetListItem is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *Queries_Expecter) GetListItem(ctx interface{}, id interface{}) *Queries_GetListItem_Call {
	return &Queries_GetListItem_Call{Call: _e.mock.On("GetListItem", ctx, id)}
}

func (_c *Queries_GetListItem_Call) Run(run func(ctx context.Context, id int64)) *Queries_GetListItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Queries_GetListItem_Call) Return(_a0 sqlite.ListItem, _a1 error) *Queries_GetListItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Queries_GetListItem_Call) RunAndReturn(run func(context.Context, int64) (sqlite.ListItem, error)) *Queries_GetListItem_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateListItemChecked provides a mock function with given fields: ctx, params
func (_m *Queries) UpdateListItemChecked(ctx context.Context, params sqlite.UpdateListItemCheckedParams) (sqlite.ListItem, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for UpdateListItemChecked")
	}

	var r0 sqlite.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.UpdateListItemCheckedParams) (sqlite.ListItem, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.UpdateListItemCheckedParams) sqlite.ListItem); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(sqlite.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, sqlite.UpdateListItemCheckedParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Queries_UpdateListItemChecked_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateListItemChecked'
type Queries_UpdateListItemChecked_Call struct {
	*mock.Call
}

// UpdateListItemChecked is a helper method to define mock.On call
//   - ctx context.Context
//   - params sqlite.UpdateListItemCheckedParams
func (_e *Queries_Expecter) UpdateListItemChecked(ctx interface{}, params interface{}) *Queries_UpdateListItemChecked_Call {
	return &Queries_UpdateListItemChecked_Call{Call: _e.mock.On("UpdateListItemChecked", ctx, params)}
}

func (_c *Queries_UpdateListItemChecked_Call) Run(run func(ctx context.Context, params sqlite.UpdateListItemCheckedParams)) *Queries_UpdateListItemChecked_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sqlite.UpdateListItemCheckedParams))
	})
	return _c
}

func (_c *Queries_UpdateListItemChecked_Call) Return(_a0 sqlite.ListItem, _a1 error) *Queries_UpdateListItemChecked_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Queries_UpdateListItemChecked_Call) RunAndReturn(run func(context.Context, sqlite.UpdateListItemCheckedParams) (sqlite.ListItem, error)) *Queries_UpdateListItemChecked_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateListItemSort provides a mock function with given fields: ctx, params
func (_m *Queries) UpdateListItemSort(ctx context.Context, params sqlite.UpdateListItemSortParams) (sqlite.ListItem, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for UpdateListItemSort")
	}

	var r0 sqlite.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.UpdateListItemSortParams) (sqlite.ListItem, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.UpdateListItemSortParams) sqlite.ListItem); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(sqlite.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, sqlite.UpdateListItemSortParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Queries_UpdateListItemSort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateListItemSort'
type Queries_UpdateListItemSort_Call struct {
	*mock.Call
}

// UpdateListItemSort is a helper method to define mock.On call
//   - ctx context.Context
//   - params sqlite.UpdateListItemSortParams
func (_e *Queries_Expecter) UpdateListItemSort(ctx interface{}, params interface{}) *Queries_UpdateListItemSort_Call {
	return &Queries_UpdateListItemSort_Call{Call: _e.mock.On("UpdateListItemSort", ctx, params)}
}

func (_c *Queries_UpdateListItemSort_Call) Run(run func(ctx context.Context, params sqlite.UpdateListItemSortParams)) *Queries_UpdateListItemSort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sqlite.UpdateListItemSortParams))
	})
	return _c
}

func (_c *Queries_UpdateListItemSort_Call) Return(_a0 sqlite.ListItem, _a1 error) *Queries_UpdateListItemSort_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Queries_UpdateListItemSort_Call) RunAndReturn(run func(context.Context, sqlite.UpdateListItemSortParams) (sqlite.ListItem, error)) *Queries_UpdateListItemSort_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateListItemText provides a mock function with given fields: ctx, params
func (_m *Queries) UpdateListItemText(ctx context.Context, params sqlite.UpdateListItemTextParams) (sqlite.ListItem, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for UpdateListItemText")
	}

	var r0 sqlite.ListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.UpdateListItemTextParams) (sqlite.ListItem, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sqlite.UpdateListItemTextParams) sqlite.ListItem); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(sqlite.ListItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, sqlite.UpdateListItemTextParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Queries_UpdateListItemText_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateListItemText'
type Queries_UpdateListItemText_Call struct {
	*mock.Call
}

// UpdateListItemText is a helper method to define mock.On call
//   - ctx context.Context
//   - params sqlite.UpdateListItemTextParams
func (_e *Queries_Expecter) UpdateListItemText(ctx interface{}, params interface{}) *Queries_UpdateListItemText_Call {
	return &Queries_UpdateListItemText_Call{Call: _e.mock.On("UpdateListItemText", ctx, params)}
}

func (_c *Queries_UpdateListItemText_Call) Run(run func(ctx context.Context, params sqlite.UpdateListItemTextParams)) *Queries_UpdateListItemText_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sqlite.UpdateListItemTextParams))
	})
	return _c
}

func (_c *Queries_UpdateListItemText_Call) Return(_a0 sqlite.ListItem, _a1 error) *Queries_UpdateListItemText_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Queries_UpdateListItemText_Call) RunAndReturn(run func(context.Context, sqlite.UpdateListItemTextParams) (sqlite.ListItem, error)) *Queries_UpdateListItemText_Call {
	_c.Call.Return(run)
	return _c
}

// NewQueries creates a new instance of Queries. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueries(t interface {
	mock.TestingT
	Cleanup(func())
}) *Queries {
	mock := &Queries{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}