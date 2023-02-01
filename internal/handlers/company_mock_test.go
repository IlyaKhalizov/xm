// Code generated by MockGen. DO NOT EDIT.
// Source: company.go

// Package handlers is a generated GoMock package.
package handlers_test

import (
	context "context"
	reflect "reflect"
	dto "xm/internal/handlers/dto"
	entities "xm/internal/repositories/entities"

	gomock "github.com/golang/mock/gomock"
)

// MockcompanyRepository is a mock of companyRepository interface.
type MockcompanyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockcompanyRepositoryMockRecorder
}

// MockcompanyRepositoryMockRecorder is the mock recorder for MockcompanyRepository.
type MockcompanyRepositoryMockRecorder struct {
	mock *MockcompanyRepository
}

// NewMockcompanyRepository creates a new mock instance.
func NewMockcompanyRepository(ctrl *gomock.Controller) *MockcompanyRepository {
	mock := &MockcompanyRepository{ctrl: ctrl}
	mock.recorder = &MockcompanyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcompanyRepository) EXPECT() *MockcompanyRepositoryMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockcompanyRepository) CreateCompany(ctx context.Context, company dto.Company) (entities.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, company)
	ret0, _ := ret[0].(entities.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockcompanyRepositoryMockRecorder) CreateCompany(ctx, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockcompanyRepository)(nil).CreateCompany), ctx, company)
}

// DeleteCompany mocks base method.
func (m *MockcompanyRepository) DeleteCompany(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockcompanyRepositoryMockRecorder) DeleteCompany(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockcompanyRepository)(nil).DeleteCompany), ctx, id)
}

// GetCompany mocks base method.
func (m *MockcompanyRepository) GetCompany(ctx context.Context, id int) (entities.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompany", ctx, id)
	ret0, _ := ret[0].(entities.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompany indicates an expected call of GetCompany.
func (mr *MockcompanyRepositoryMockRecorder) GetCompany(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompany", reflect.TypeOf((*MockcompanyRepository)(nil).GetCompany), ctx, id)
}

// UpdateCompany mocks base method.
func (m *MockcompanyRepository) UpdateCompany(ctx context.Context, id int, company dto.Company) (entities.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompany", ctx, id, company)
	ret0, _ := ret[0].(entities.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCompany indicates an expected call of UpdateCompany.
func (mr *MockcompanyRepositoryMockRecorder) UpdateCompany(ctx, id, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompany", reflect.TypeOf((*MockcompanyRepository)(nil).UpdateCompany), ctx, id, company)
}