// Code generated by MockGen. DO NOT EDIT.
// Source: analytic.go

// Package mock_usecase is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/tokopedia-workshop-2022/seller-analytics-solution/src/services/analytic/domain"
)

// MockAnalyticUsecase is a mock of AnalyticUsecase interface.
type MockAnalyticUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyticUsecaseMockRecorder
}

// MockAnalyticUsecaseMockRecorder is the mock recorder for MockAnalyticUsecase.
type MockAnalyticUsecaseMockRecorder struct {
	mock *MockAnalyticUsecase
}

// NewMockAnalyticUsecase creates a new mock instance.
func NewMockAnalyticUsecase(ctrl *gomock.Controller) *MockAnalyticUsecase {
	mock := &MockAnalyticUsecase{ctrl: ctrl}
	mock.recorder = &MockAnalyticUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyticUsecase) EXPECT() *MockAnalyticUsecaseMockRecorder {
	return m.recorder
}

// GetAnalyticByDate mocks base method.
func (m *MockAnalyticUsecase) GetAnalyticByDate(ctx context.Context, date time.Time) (*domain.Analytic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnalyticByDate", ctx, date)
	ret0, _ := ret[0].(*domain.Analytic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnalyticByDate indicates an expected call of GetAnalyticByDate.
func (mr *MockAnalyticUsecaseMockRecorder) GetAnalyticByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnalyticByDate", reflect.TypeOf((*MockAnalyticUsecase)(nil).GetAnalyticByDate), ctx, date)
}

// HandleStatisticEvent mocks base method.
func (m *MockAnalyticUsecase) HandleStatisticEvent(statisticEvent domain.StatisticEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleStatisticEvent", statisticEvent)
}

// HandleStatisticEvent indicates an expected call of HandleStatisticEvent.
func (mr *MockAnalyticUsecaseMockRecorder) HandleStatisticEvent(statisticEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStatisticEvent", reflect.TypeOf((*MockAnalyticUsecase)(nil).HandleStatisticEvent), statisticEvent)
}