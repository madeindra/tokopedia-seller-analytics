// Code generated by MockGen. DO NOT EDIT.
// Source: analytic.go

// Package mock_repository is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/tokopedia-workshop-2022/seller-analytics-solution/src/services/analytic/domain"
)

// MockAnalyticRepository is a mock of AnalyticRepository interface.
type MockAnalyticRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyticRepositoryMockRecorder
}

// MockAnalyticRepositoryMockRecorder is the mock recorder for MockAnalyticRepository.
type MockAnalyticRepositoryMockRecorder struct {
	mock *MockAnalyticRepository
}

// NewMockAnalyticRepository creates a new mock instance.
func NewMockAnalyticRepository(ctrl *gomock.Controller) *MockAnalyticRepository {
	mock := &MockAnalyticRepository{ctrl: ctrl}
	mock.recorder = &MockAnalyticRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyticRepository) EXPECT() *MockAnalyticRepositoryMockRecorder {
	return m.recorder
}

// CreateAnalytic mocks base method.
func (m *MockAnalyticRepository) CreateAnalytic(ctx context.Context, analytic domain.Analytic) (*domain.Analytic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnalytic", ctx, analytic)
	ret0, _ := ret[0].(*domain.Analytic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAnalytic indicates an expected call of CreateAnalytic.
func (mr *MockAnalyticRepositoryMockRecorder) CreateAnalytic(ctx, analytic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnalytic", reflect.TypeOf((*MockAnalyticRepository)(nil).CreateAnalytic), ctx, analytic)
}

// GetAnalyticByDate mocks base method.
func (m *MockAnalyticRepository) GetAnalyticByDate(ctx context.Context, date time.Time) (*domain.Analytic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnalyticByDate", ctx, date)
	ret0, _ := ret[0].(*domain.Analytic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnalyticByDate indicates an expected call of GetAnalyticByDate.
func (mr *MockAnalyticRepositoryMockRecorder) GetAnalyticByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnalyticByDate", reflect.TypeOf((*MockAnalyticRepository)(nil).GetAnalyticByDate), ctx, date)
}

// UpdateAnalytic mocks base method.
func (m *MockAnalyticRepository) UpdateAnalytic(ctx context.Context, analytic domain.Analytic) (*domain.Analytic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnalytic", ctx, analytic)
	ret0, _ := ret[0].(*domain.Analytic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAnalytic indicates an expected call of UpdateAnalytic.
func (mr *MockAnalyticRepositoryMockRecorder) UpdateAnalytic(ctx, analytic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnalytic", reflect.TypeOf((*MockAnalyticRepository)(nil).UpdateAnalytic), ctx, analytic)
}
