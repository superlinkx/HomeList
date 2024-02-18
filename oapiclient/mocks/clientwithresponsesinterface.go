// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	oapiclient "github.com/superlinkx/HomeList/oapiclient"
)

// MockClientWithResponsesInterface is an autogenerated mock type for the ClientWithResponsesInterface type
type MockClientWithResponsesInterface struct {
	mock.Mock
}

type MockClientWithResponsesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientWithResponsesInterface) EXPECT() *MockClientWithResponsesInterface_Expecter {
	return &MockClientWithResponsesInterface_Expecter{mock: &_m.Mock}
}

// CreateListWithBodyWithResponse provides a mock function with given fields: ctx, contentType, body, reqEditors
func (_m *MockClientWithResponsesInterface) CreateListWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...oapiclient.RequestEditorFn) (*oapiclient.CreateListResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateListWithBodyWithResponse")
	}

	var r0 *oapiclient.CreateListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...oapiclient.RequestEditorFn) (*oapiclient.CreateListResponse, error)); ok {
		return rf(ctx, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...oapiclient.RequestEditorFn) *oapiclient.CreateListResponse); ok {
		r0 = rf(ctx, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oapiclient.CreateListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, io.Reader, ...oapiclient.RequestEditorFn) error); ok {
		r1 = rf(ctx, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateListWithBodyWithResponse'
type MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call struct {
	*mock.Call
}

// CreateListWithBodyWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - contentType string
//   - body io.Reader
//   - reqEditors ...oapiclient.RequestEditorFn
func (_e *MockClientWithResponsesInterface_Expecter) CreateListWithBodyWithResponse(ctx interface{}, contentType interface{}, body interface{}, reqEditors ...interface{}) *MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call {
	return &MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call{Call: _e.mock.On("CreateListWithBodyWithResponse",
		append([]interface{}{ctx, contentType, body}, reqEditors...)...)}
}

func (_c *MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call) Run(run func(ctx context.Context, contentType string, body io.Reader, reqEditors ...oapiclient.RequestEditorFn)) *MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]oapiclient.RequestEditorFn, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(oapiclient.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(io.Reader), variadicArgs...)
	})
	return _c
}

func (_c *MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call) Return(_a0 *oapiclient.CreateListResponse, _a1 error) *MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call) RunAndReturn(run func(context.Context, string, io.Reader, ...oapiclient.RequestEditorFn) (*oapiclient.CreateListResponse, error)) *MockClientWithResponsesInterface_CreateListWithBodyWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// CreateListWithResponse provides a mock function with given fields: ctx, body, reqEditors
func (_m *MockClientWithResponsesInterface) CreateListWithResponse(ctx context.Context, body oapiclient.CreateListJSONRequestBody, reqEditors ...oapiclient.RequestEditorFn) (*oapiclient.CreateListResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateListWithResponse")
	}

	var r0 *oapiclient.CreateListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oapiclient.CreateListJSONRequestBody, ...oapiclient.RequestEditorFn) (*oapiclient.CreateListResponse, error)); ok {
		return rf(ctx, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oapiclient.CreateListJSONRequestBody, ...oapiclient.RequestEditorFn) *oapiclient.CreateListResponse); ok {
		r0 = rf(ctx, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oapiclient.CreateListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oapiclient.CreateListJSONRequestBody, ...oapiclient.RequestEditorFn) error); ok {
		r1 = rf(ctx, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientWithResponsesInterface_CreateListWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateListWithResponse'
type MockClientWithResponsesInterface_CreateListWithResponse_Call struct {
	*mock.Call
}

// CreateListWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - body oapiclient.CreateListJSONRequestBody
//   - reqEditors ...oapiclient.RequestEditorFn
func (_e *MockClientWithResponsesInterface_Expecter) CreateListWithResponse(ctx interface{}, body interface{}, reqEditors ...interface{}) *MockClientWithResponsesInterface_CreateListWithResponse_Call {
	return &MockClientWithResponsesInterface_CreateListWithResponse_Call{Call: _e.mock.On("CreateListWithResponse",
		append([]interface{}{ctx, body}, reqEditors...)...)}
}

func (_c *MockClientWithResponsesInterface_CreateListWithResponse_Call) Run(run func(ctx context.Context, body oapiclient.CreateListJSONRequestBody, reqEditors ...oapiclient.RequestEditorFn)) *MockClientWithResponsesInterface_CreateListWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]oapiclient.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(oapiclient.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(oapiclient.CreateListJSONRequestBody), variadicArgs...)
	})
	return _c
}

func (_c *MockClientWithResponsesInterface_CreateListWithResponse_Call) Return(_a0 *oapiclient.CreateListResponse, _a1 error) *MockClientWithResponsesInterface_CreateListWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientWithResponsesInterface_CreateListWithResponse_Call) RunAndReturn(run func(context.Context, oapiclient.CreateListJSONRequestBody, ...oapiclient.RequestEditorFn) (*oapiclient.CreateListResponse, error)) *MockClientWithResponsesInterface_CreateListWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteListWithResponse provides a mock function with given fields: ctx, listID, reqEditors
func (_m *MockClientWithResponsesInterface) DeleteListWithResponse(ctx context.Context, listID int64, reqEditors ...oapiclient.RequestEditorFn) (*oapiclient.DeleteListResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, listID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteListWithResponse")
	}

	var r0 *oapiclient.DeleteListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...oapiclient.RequestEditorFn) (*oapiclient.DeleteListResponse, error)); ok {
		return rf(ctx, listID, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...oapiclient.RequestEditorFn) *oapiclient.DeleteListResponse); ok {
		r0 = rf(ctx, listID, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oapiclient.DeleteListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...oapiclient.RequestEditorFn) error); ok {
		r1 = rf(ctx, listID, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientWithResponsesInterface_DeleteListWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteListWithResponse'
type MockClientWithResponsesInterface_DeleteListWithResponse_Call struct {
	*mock.Call
}

// DeleteListWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - reqEditors ...oapiclient.RequestEditorFn
func (_e *MockClientWithResponsesInterface_Expecter) DeleteListWithResponse(ctx interface{}, listID interface{}, reqEditors ...interface{}) *MockClientWithResponsesInterface_DeleteListWithResponse_Call {
	return &MockClientWithResponsesInterface_DeleteListWithResponse_Call{Call: _e.mock.On("DeleteListWithResponse",
		append([]interface{}{ctx, listID}, reqEditors...)...)}
}

func (_c *MockClientWithResponsesInterface_DeleteListWithResponse_Call) Run(run func(ctx context.Context, listID int64, reqEditors ...oapiclient.RequestEditorFn)) *MockClientWithResponsesInterface_DeleteListWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]oapiclient.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(oapiclient.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockClientWithResponsesInterface_DeleteListWithResponse_Call) Return(_a0 *oapiclient.DeleteListResponse, _a1 error) *MockClientWithResponsesInterface_DeleteListWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientWithResponsesInterface_DeleteListWithResponse_Call) RunAndReturn(run func(context.Context, int64, ...oapiclient.RequestEditorFn) (*oapiclient.DeleteListResponse, error)) *MockClientWithResponsesInterface_DeleteListWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// GetListWithResponse provides a mock function with given fields: ctx, listID, reqEditors
func (_m *MockClientWithResponsesInterface) GetListWithResponse(ctx context.Context, listID int64, reqEditors ...oapiclient.RequestEditorFn) (*oapiclient.GetListResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, listID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetListWithResponse")
	}

	var r0 *oapiclient.GetListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...oapiclient.RequestEditorFn) (*oapiclient.GetListResponse, error)); ok {
		return rf(ctx, listID, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...oapiclient.RequestEditorFn) *oapiclient.GetListResponse); ok {
		r0 = rf(ctx, listID, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oapiclient.GetListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...oapiclient.RequestEditorFn) error); ok {
		r1 = rf(ctx, listID, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientWithResponsesInterface_GetListWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListWithResponse'
type MockClientWithResponsesInterface_GetListWithResponse_Call struct {
	*mock.Call
}

// GetListWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - reqEditors ...oapiclient.RequestEditorFn
func (_e *MockClientWithResponsesInterface_Expecter) GetListWithResponse(ctx interface{}, listID interface{}, reqEditors ...interface{}) *MockClientWithResponsesInterface_GetListWithResponse_Call {
	return &MockClientWithResponsesInterface_GetListWithResponse_Call{Call: _e.mock.On("GetListWithResponse",
		append([]interface{}{ctx, listID}, reqEditors...)...)}
}

func (_c *MockClientWithResponsesInterface_GetListWithResponse_Call) Run(run func(ctx context.Context, listID int64, reqEditors ...oapiclient.RequestEditorFn)) *MockClientWithResponsesInterface_GetListWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]oapiclient.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(oapiclient.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockClientWithResponsesInterface_GetListWithResponse_Call) Return(_a0 *oapiclient.GetListResponse, _a1 error) *MockClientWithResponsesInterface_GetListWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientWithResponsesInterface_GetListWithResponse_Call) RunAndReturn(run func(context.Context, int64, ...oapiclient.RequestEditorFn) (*oapiclient.GetListResponse, error)) *MockClientWithResponsesInterface_GetListWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// GetListsWithResponse provides a mock function with given fields: ctx, params, reqEditors
func (_m *MockClientWithResponsesInterface) GetListsWithResponse(ctx context.Context, params *oapiclient.GetListsParams, reqEditors ...oapiclient.RequestEditorFn) (*oapiclient.GetListsResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetListsWithResponse")
	}

	var r0 *oapiclient.GetListsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *oapiclient.GetListsParams, ...oapiclient.RequestEditorFn) (*oapiclient.GetListsResponse, error)); ok {
		return rf(ctx, params, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *oapiclient.GetListsParams, ...oapiclient.RequestEditorFn) *oapiclient.GetListsResponse); ok {
		r0 = rf(ctx, params, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oapiclient.GetListsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *oapiclient.GetListsParams, ...oapiclient.RequestEditorFn) error); ok {
		r1 = rf(ctx, params, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientWithResponsesInterface_GetListsWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetListsWithResponse'
type MockClientWithResponsesInterface_GetListsWithResponse_Call struct {
	*mock.Call
}

// GetListsWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - params *oapiclient.GetListsParams
//   - reqEditors ...oapiclient.RequestEditorFn
func (_e *MockClientWithResponsesInterface_Expecter) GetListsWithResponse(ctx interface{}, params interface{}, reqEditors ...interface{}) *MockClientWithResponsesInterface_GetListsWithResponse_Call {
	return &MockClientWithResponsesInterface_GetListsWithResponse_Call{Call: _e.mock.On("GetListsWithResponse",
		append([]interface{}{ctx, params}, reqEditors...)...)}
}

func (_c *MockClientWithResponsesInterface_GetListsWithResponse_Call) Run(run func(ctx context.Context, params *oapiclient.GetListsParams, reqEditors ...oapiclient.RequestEditorFn)) *MockClientWithResponsesInterface_GetListsWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]oapiclient.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(oapiclient.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(*oapiclient.GetListsParams), variadicArgs...)
	})
	return _c
}

func (_c *MockClientWithResponsesInterface_GetListsWithResponse_Call) Return(_a0 *oapiclient.GetListsResponse, _a1 error) *MockClientWithResponsesInterface_GetListsWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientWithResponsesInterface_GetListsWithResponse_Call) RunAndReturn(run func(context.Context, *oapiclient.GetListsParams, ...oapiclient.RequestEditorFn) (*oapiclient.GetListsResponse, error)) *MockClientWithResponsesInterface_GetListsWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateListWithBodyWithResponse provides a mock function with given fields: ctx, listID, contentType, body, reqEditors
func (_m *MockClientWithResponsesInterface) UpdateListWithBodyWithResponse(ctx context.Context, listID int64, contentType string, body io.Reader, reqEditors ...oapiclient.RequestEditorFn) (*oapiclient.UpdateListResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, listID, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateListWithBodyWithResponse")
	}

	var r0 *oapiclient.UpdateListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, io.Reader, ...oapiclient.RequestEditorFn) (*oapiclient.UpdateListResponse, error)); ok {
		return rf(ctx, listID, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, io.Reader, ...oapiclient.RequestEditorFn) *oapiclient.UpdateListResponse); ok {
		r0 = rf(ctx, listID, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oapiclient.UpdateListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, io.Reader, ...oapiclient.RequestEditorFn) error); ok {
		r1 = rf(ctx, listID, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateListWithBodyWithResponse'
type MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call struct {
	*mock.Call
}

// UpdateListWithBodyWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - contentType string
//   - body io.Reader
//   - reqEditors ...oapiclient.RequestEditorFn
func (_e *MockClientWithResponsesInterface_Expecter) UpdateListWithBodyWithResponse(ctx interface{}, listID interface{}, contentType interface{}, body interface{}, reqEditors ...interface{}) *MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call {
	return &MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call{Call: _e.mock.On("UpdateListWithBodyWithResponse",
		append([]interface{}{ctx, listID, contentType, body}, reqEditors...)...)}
}

func (_c *MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call) Run(run func(ctx context.Context, listID int64, contentType string, body io.Reader, reqEditors ...oapiclient.RequestEditorFn)) *MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]oapiclient.RequestEditorFn, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(oapiclient.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(int64), args[2].(string), args[3].(io.Reader), variadicArgs...)
	})
	return _c
}

func (_c *MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call) Return(_a0 *oapiclient.UpdateListResponse, _a1 error) *MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call) RunAndReturn(run func(context.Context, int64, string, io.Reader, ...oapiclient.RequestEditorFn) (*oapiclient.UpdateListResponse, error)) *MockClientWithResponsesInterface_UpdateListWithBodyWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateListWithResponse provides a mock function with given fields: ctx, listID, body, reqEditors
func (_m *MockClientWithResponsesInterface) UpdateListWithResponse(ctx context.Context, listID int64, body oapiclient.UpdateListJSONRequestBody, reqEditors ...oapiclient.RequestEditorFn) (*oapiclient.UpdateListResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, listID, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateListWithResponse")
	}

	var r0 *oapiclient.UpdateListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, oapiclient.UpdateListJSONRequestBody, ...oapiclient.RequestEditorFn) (*oapiclient.UpdateListResponse, error)); ok {
		return rf(ctx, listID, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, oapiclient.UpdateListJSONRequestBody, ...oapiclient.RequestEditorFn) *oapiclient.UpdateListResponse); ok {
		r0 = rf(ctx, listID, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oapiclient.UpdateListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, oapiclient.UpdateListJSONRequestBody, ...oapiclient.RequestEditorFn) error); ok {
		r1 = rf(ctx, listID, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientWithResponsesInterface_UpdateListWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateListWithResponse'
type MockClientWithResponsesInterface_UpdateListWithResponse_Call struct {
	*mock.Call
}

// UpdateListWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - listID int64
//   - body oapiclient.UpdateListJSONRequestBody
//   - reqEditors ...oapiclient.RequestEditorFn
func (_e *MockClientWithResponsesInterface_Expecter) UpdateListWithResponse(ctx interface{}, listID interface{}, body interface{}, reqEditors ...interface{}) *MockClientWithResponsesInterface_UpdateListWithResponse_Call {
	return &MockClientWithResponsesInterface_UpdateListWithResponse_Call{Call: _e.mock.On("UpdateListWithResponse",
		append([]interface{}{ctx, listID, body}, reqEditors...)...)}
}

func (_c *MockClientWithResponsesInterface_UpdateListWithResponse_Call) Run(run func(ctx context.Context, listID int64, body oapiclient.UpdateListJSONRequestBody, reqEditors ...oapiclient.RequestEditorFn)) *MockClientWithResponsesInterface_UpdateListWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]oapiclient.RequestEditorFn, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(oapiclient.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(int64), args[2].(oapiclient.UpdateListJSONRequestBody), variadicArgs...)
	})
	return _c
}

func (_c *MockClientWithResponsesInterface_UpdateListWithResponse_Call) Return(_a0 *oapiclient.UpdateListResponse, _a1 error) *MockClientWithResponsesInterface_UpdateListWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientWithResponsesInterface_UpdateListWithResponse_Call) RunAndReturn(run func(context.Context, int64, oapiclient.UpdateListJSONRequestBody, ...oapiclient.RequestEditorFn) (*oapiclient.UpdateListResponse, error)) *MockClientWithResponsesInterface_UpdateListWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClientWithResponsesInterface creates a new instance of MockClientWithResponsesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientWithResponsesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientWithResponsesInterface {
	mock := &MockClientWithResponsesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
