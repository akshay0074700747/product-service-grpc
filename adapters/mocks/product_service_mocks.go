// Code generated by MockGen. DO NOT EDIT.
// Source: adapter_interfaces.go

// Package mock_adapters is a generated GoMock package.
package mock_adapters

import (
	reflect "reflect"

	entities "github.com/akshay0074700747/products-service/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockAdapterInterface is a mock of AdapterInterface interface.
type MockAdapterInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterInterfaceMockRecorder
}

// MockAdapterInterfaceMockRecorder is the mock recorder for MockAdapterInterface.
type MockAdapterInterfaceMockRecorder struct {
	mock *MockAdapterInterface
}

// NewMockAdapterInterface creates a new mock instance.
func NewMockAdapterInterface(ctrl *gomock.Controller) *MockAdapterInterface {
	mock := &MockAdapterInterface{ctrl: ctrl}
	mock.recorder = &MockAdapterInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdapterInterface) EXPECT() *MockAdapterInterfaceMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockAdapterInterface) AddProduct(req entities.Products) (entities.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", req)
	ret0, _ := ret[0].(entities.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockAdapterInterfaceMockRecorder) AddProduct(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockAdapterInterface)(nil).AddProduct), req)
}

// DecrementStock mocks base method.
func (m *MockAdapterInterface) DecrementStock(id uint, stock int) (entities.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrementStock", id, stock)
	ret0, _ := ret[0].(entities.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecrementStock indicates an expected call of DecrementStock.
func (mr *MockAdapterInterfaceMockRecorder) DecrementStock(id, stock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrementStock", reflect.TypeOf((*MockAdapterInterface)(nil).DecrementStock), id, stock)
}

// GetAllProducts mocks base method.
func (m *MockAdapterInterface) GetAllProducts() ([]entities.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts")
	ret0, _ := ret[0].([]entities.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockAdapterInterfaceMockRecorder) GetAllProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockAdapterInterface)(nil).GetAllProducts))
}

// GetProduct mocks base method.
func (m *MockAdapterInterface) GetProduct(id uint) (entities.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", id)
	ret0, _ := ret[0].(entities.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockAdapterInterfaceMockRecorder) GetProduct(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockAdapterInterface)(nil).GetProduct), id)
}

// IncrementStock mocks base method.
func (m *MockAdapterInterface) IncrementStock(id uint, stock int) (entities.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementStock", id, stock)
	ret0, _ := ret[0].(entities.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrementStock indicates an expected call of IncrementStock.
func (mr *MockAdapterInterfaceMockRecorder) IncrementStock(id, stock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementStock", reflect.TypeOf((*MockAdapterInterface)(nil).IncrementStock), id, stock)
}
