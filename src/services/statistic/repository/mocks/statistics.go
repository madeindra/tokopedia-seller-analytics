// Code generated by MockGen. DO NOT EDIT.
// Source: statistics.go

// Package mock_repository is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/tokopedia-workshop-2022/seller-analytics-solution/src/services/statistic/domain"
)

// MockStatisticsRepository is a mock of StatisticsRepository interface.
type MockStatisticsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStatisticsRepositoryMockRecorder
}

// MockStatisticsRepositoryMockRecorder is the mock recorder for MockStatisticsRepository.
type MockStatisticsRepositoryMockRecorder struct {
	mock *MockStatisticsRepository
}

// NewMockStatisticsRepository creates a new mock instance.
func NewMockStatisticsRepository(ctrl *gomock.Controller) *MockStatisticsRepository {
	mock := &MockStatisticsRepository{ctrl: ctrl}
	mock.recorder = &MockStatisticsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatisticsRepository) EXPECT() *MockStatisticsRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStatisticsRepository) Create(ctx context.Context, stat domain.Statistics) (*domain.Statistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, stat)
	ret0, _ := ret[0].(*domain.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStatisticsRepositoryMockRecorder) Create(ctx, stat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStatisticsRepository)(nil).Create), ctx, stat)
}

// GetByDate mocks base method.
func (m *MockStatisticsRepository) GetByDate(ctx context.Context, date time.Time) (*domain.Statistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDate", ctx, date)
	ret0, _ := ret[0].(*domain.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDate indicates an expected call of GetByDate.
func (mr *MockStatisticsRepositoryMockRecorder) GetByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDate", reflect.TypeOf((*MockStatisticsRepository)(nil).GetByDate), ctx, date)
}

// PublishEvent mocks base method.
func (m *MockStatisticsRepository) PublishEvent(ctx context.Context, event domain.PayloadEventStatistic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishEvent", ctx, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishEvent indicates an expected call of PublishEvent.
func (mr *MockStatisticsRepositoryMockRecorder) PublishEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishEvent", reflect.TypeOf((*MockStatisticsRepository)(nil).PublishEvent), ctx, event)
}

// Update mocks base method.
func (m *MockStatisticsRepository) Update(ctx context.Context, stat domain.Statistics) (*domain.Statistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, stat)
	ret0, _ := ret[0].(*domain.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockStatisticsRepositoryMockRecorder) Update(ctx, stat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStatisticsRepository)(nil).Update), ctx, stat)
}