// Code generated by MockGen. DO NOT EDIT.
// Source: ./customer-service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/nuno/nunes-jumia/src/dto"
)

// MockCustomerService is a mock of CustomerService interface.
type MockCustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerServiceMockRecorder
}

// MockCustomerServiceMockRecorder is the mock recorder for MockCustomerService.
type MockCustomerServiceMockRecorder struct {
	mock *MockCustomerService
}

// NewMockCustomerService creates a new mock instance.
func NewMockCustomerService(ctrl *gomock.Controller) *MockCustomerService {
	mock := &MockCustomerService{ctrl: ctrl}
	mock.recorder = &MockCustomerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerService) EXPECT() *MockCustomerServiceMockRecorder {
	return m.recorder
}

// GetCustomers mocks base method.
func (m *MockCustomerService) GetCustomers(limit, offset int, params map[string]string) (dto.CustomerOutputDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomers", limit, offset, params)
	ret0, _ := ret[0].(dto.CustomerOutputDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomers indicates an expected call of GetCustomers.
func (mr *MockCustomerServiceMockRecorder) GetCustomers(limit, offset, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomers", reflect.TypeOf((*MockCustomerService)(nil).GetCustomers), limit, offset, params)
}
