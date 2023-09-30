// Code generated by MockGen. DO NOT EDIT.
// Source: ../interface.go
//
// Generated by this command:
//
//	mockgen -source=../interface.go -destination=./bigset_mock.go -package=mock
//
// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	generic "github.com/OpenStars/BackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// BsGetItem mocks base method.
func (m *MockClient) BsGetItem(bskey, itemkey string) (*generic.TItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsGetItem", bskey, itemkey)
	ret0, _ := ret[0].(*generic.TItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsGetItem indicates an expected call of BsGetItem.
func (mr *MockClientMockRecorder) BsGetItem(bskey, itemkey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsGetItem", reflect.TypeOf((*MockClient)(nil).BsGetItem), bskey, itemkey)
}

// BsGetSlice mocks base method.
func (m *MockClient) BsGetSlice(bskey string, fromPos, count int32) ([]*generic.TItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsGetSlice", bskey, fromPos, count)
	ret0, _ := ret[0].([]*generic.TItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsGetSlice indicates an expected call of BsGetSlice.
func (mr *MockClientMockRecorder) BsGetSlice(bskey, fromPos, count any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsGetSlice", reflect.TypeOf((*MockClient)(nil).BsGetSlice), bskey, fromPos, count)
}

// BsGetSliceFromItem mocks base method.
func (m *MockClient) BsGetSliceFromItem(bskey, fromKey string, count int32) ([]*generic.TItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsGetSliceFromItem", bskey, fromKey, count)
	ret0, _ := ret[0].([]*generic.TItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsGetSliceFromItem indicates an expected call of BsGetSliceFromItem.
func (mr *MockClientMockRecorder) BsGetSliceFromItem(bskey, fromKey, count any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsGetSliceFromItem", reflect.TypeOf((*MockClient)(nil).BsGetSliceFromItem), bskey, fromKey, count)
}

// BsGetSliceFromItemR mocks base method.
func (m *MockClient) BsGetSliceFromItemR(bskey, fromKey string, count int32) ([]*generic.TItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsGetSliceFromItemR", bskey, fromKey, count)
	ret0, _ := ret[0].([]*generic.TItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsGetSliceFromItemR indicates an expected call of BsGetSliceFromItemR.
func (mr *MockClientMockRecorder) BsGetSliceFromItemR(bskey, fromKey, count any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsGetSliceFromItemR", reflect.TypeOf((*MockClient)(nil).BsGetSliceFromItemR), bskey, fromKey, count)
}

// BsGetSliceR mocks base method.
func (m *MockClient) BsGetSliceR(bskey string, fromPos, count int32) ([]*generic.TItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsGetSliceR", bskey, fromPos, count)
	ret0, _ := ret[0].([]*generic.TItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsGetSliceR indicates an expected call of BsGetSliceR.
func (mr *MockClientMockRecorder) BsGetSliceR(bskey, fromPos, count any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsGetSliceR", reflect.TypeOf((*MockClient)(nil).BsGetSliceR), bskey, fromPos, count)
}

// BsMultiPut mocks base method.
func (m *MockClient) BsMultiPut(bskey string, lsItems []*generic.TItem) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsMultiPut", bskey, lsItems)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsMultiPut indicates an expected call of BsMultiPut.
func (mr *MockClientMockRecorder) BsMultiPut(bskey, lsItems any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsMultiPut", reflect.TypeOf((*MockClient)(nil).BsMultiPut), bskey, lsItems)
}

// BsMultiPutBsItem mocks base method.
func (m *MockClient) BsMultiPutBsItem(lsItem []*generic.TBigsetItem) ([]*generic.TBigsetItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsMultiPutBsItem", lsItem)
	ret0, _ := ret[0].([]*generic.TBigsetItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsMultiPutBsItem indicates an expected call of BsMultiPutBsItem.
func (mr *MockClientMockRecorder) BsMultiPutBsItem(lsItem any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsMultiPutBsItem", reflect.TypeOf((*MockClient)(nil).BsMultiPutBsItem), lsItem)
}

// BsMultiRemoveBsItem mocks base method.
func (m *MockClient) BsMultiRemoveBsItem(listItems []*generic.TBigsetItem) ([]*generic.TBigsetItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsMultiRemoveBsItem", listItems)
	ret0, _ := ret[0].([]*generic.TBigsetItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsMultiRemoveBsItem indicates an expected call of BsMultiRemoveBsItem.
func (mr *MockClientMockRecorder) BsMultiRemoveBsItem(listItems any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsMultiRemoveBsItem", reflect.TypeOf((*MockClient)(nil).BsMultiRemoveBsItem), listItems)
}

// BsPutItem mocks base method.
func (m *MockClient) BsPutItem(bskey, itemKey, itemVal string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsPutItem", bskey, itemKey, itemVal)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsPutItem indicates an expected call of BsPutItem.
func (mr *MockClientMockRecorder) BsPutItem(bskey, itemKey, itemVal any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsPutItem", reflect.TypeOf((*MockClient)(nil).BsPutItem), bskey, itemKey, itemVal)
}

// BsRangeQuery mocks base method.
func (m *MockClient) BsRangeQuery(bskey, startKey, endKey string) ([]*generic.TItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsRangeQuery", bskey, startKey, endKey)
	ret0, _ := ret[0].([]*generic.TItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsRangeQuery indicates an expected call of BsRangeQuery.
func (mr *MockClientMockRecorder) BsRangeQuery(bskey, startKey, endKey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsRangeQuery", reflect.TypeOf((*MockClient)(nil).BsRangeQuery), bskey, startKey, endKey)
}

// BsRangeQueryAll mocks base method.
func (m *MockClient) BsRangeQueryAll(bskey string) ([]*generic.TItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsRangeQueryAll", bskey)
	ret0, _ := ret[0].([]*generic.TItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsRangeQueryAll indicates an expected call of BsRangeQueryAll.
func (mr *MockClientMockRecorder) BsRangeQueryAll(bskey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsRangeQueryAll", reflect.TypeOf((*MockClient)(nil).BsRangeQueryAll), bskey)
}

// BsRangeQueryByPage mocks base method.
func (m *MockClient) BsRangeQueryByPage(bskey, startKey, endKey string, begin, end int64) ([]*generic.TItem, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsRangeQueryByPage", bskey, startKey, endKey, begin, end)
	ret0, _ := ret[0].([]*generic.TItem)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BsRangeQueryByPage indicates an expected call of BsRangeQueryByPage.
func (mr *MockClientMockRecorder) BsRangeQueryByPage(bskey, startKey, endKey, begin, end any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsRangeQueryByPage", reflect.TypeOf((*MockClient)(nil).BsRangeQueryByPage), bskey, startKey, endKey, begin, end)
}

// BsRemoveItem mocks base method.
func (m *MockClient) BsRemoveItem(bskey, itemkey string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BsRemoveItem", bskey, itemkey)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BsRemoveItem indicates an expected call of BsRemoveItem.
func (mr *MockClientMockRecorder) BsRemoveItem(bskey, itemkey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BsRemoveItem", reflect.TypeOf((*MockClient)(nil).BsRemoveItem), bskey, itemkey)
}

// CreateStringBigSet mocks base method.
func (m *MockClient) CreateStringBigSet(bskey string) (*generic.TStringBigSetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStringBigSet", bskey)
	ret0, _ := ret[0].(*generic.TStringBigSetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStringBigSet indicates an expected call of CreateStringBigSet.
func (mr *MockClientMockRecorder) CreateStringBigSet(bskey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStringBigSet", reflect.TypeOf((*MockClient)(nil).CreateStringBigSet), bskey)
}

// GetBigSetInfoByName mocks base method.
func (m *MockClient) GetBigSetInfoByName(bskey string) (*generic.TStringBigSetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBigSetInfoByName", bskey)
	ret0, _ := ret[0].(*generic.TStringBigSetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBigSetInfoByName indicates an expected call of GetBigSetInfoByName.
func (mr *MockClientMockRecorder) GetBigSetInfoByName(bskey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBigSetInfoByName", reflect.TypeOf((*MockClient)(nil).GetBigSetInfoByName), bskey)
}

// GetListKey mocks base method.
func (m *MockClient) GetListKey(fromIndex int64, count int32) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListKey", fromIndex, count)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListKey indicates an expected call of GetListKey.
func (mr *MockClientMockRecorder) GetListKey(fromIndex, count any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListKey", reflect.TypeOf((*MockClient)(nil).GetListKey), fromIndex, count)
}

// GetTotalCount mocks base method.
func (m *MockClient) GetTotalCount(bskey string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalCount", bskey)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalCount indicates an expected call of GetTotalCount.
func (mr *MockClientMockRecorder) GetTotalCount(bskey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalCount", reflect.TypeOf((*MockClient)(nil).GetTotalCount), bskey)
}

// RemoveAll mocks base method.
func (m *MockClient) RemoveAll(bskey string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", bskey)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockClientMockRecorder) RemoveAll(bskey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockClient)(nil).RemoveAll), bskey)
}

// TotalStringKeyCount mocks base method.
func (m *MockClient) TotalStringKeyCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalStringKeyCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalStringKeyCount indicates an expected call of TotalStringKeyCount.
func (mr *MockClientMockRecorder) TotalStringKeyCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalStringKeyCount", reflect.TypeOf((*MockClient)(nil).TotalStringKeyCount))
}
