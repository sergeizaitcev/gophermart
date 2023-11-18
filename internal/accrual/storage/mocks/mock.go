// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	storage "github.com/sergeizaitcev/gophermart/internal/accrual/storage"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// CreateInvalidOrder mocks base method.
func (m *MockStorage) CreateInvalidOrder(ctx context.Context, order string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvalidOrder", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInvalidOrder indicates an expected call of CreateInvalidOrder.
func (mr *MockStorageMockRecorder) CreateInvalidOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvalidOrder", reflect.TypeOf((*MockStorage)(nil).CreateInvalidOrder), ctx, order)
}

// CreateMatch mocks base method.
func (m *MockStorage) CreateMatch(ctx context.Context, match *storage.Match) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMatch", ctx, match)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMatch indicates an expected call of CreateMatch.
func (mr *MockStorageMockRecorder) CreateMatch(ctx, match interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMatch", reflect.TypeOf((*MockStorage)(nil).CreateMatch), ctx, match)
}

// CreateOrderWithGoods mocks base method.
func (m *MockStorage) CreateOrderWithGoods(ctx context.Context, order string, goods []*storage.Goods) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderWithGoods", ctx, order, goods)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderWithGoods indicates an expected call of CreateOrderWithGoods.
func (mr *MockStorageMockRecorder) CreateOrderWithGoods(ctx, order, goods interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderWithGoods", reflect.TypeOf((*MockStorage)(nil).CreateOrderWithGoods), ctx, order, goods)
}

// GetMatchByName mocks base method.
func (m *MockStorage) GetMatchByName(ctx context.Context, matchName string) (*storage.MatchOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchByName", ctx, matchName)
	ret0, _ := ret[0].(*storage.MatchOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchByName indicates an expected call of GetMatchByName.
func (mr *MockStorageMockRecorder) GetMatchByName(ctx, matchName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchByName", reflect.TypeOf((*MockStorage)(nil).GetMatchByName), ctx, matchName)
}

// GetOrderByNumber mocks base method.
func (m *MockStorage) GetOrderByNumber(ctx context.Context, orderNumber string) (*storage.OrderOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByNumber", ctx, orderNumber)
	ret0, _ := ret[0].(*storage.OrderOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByNumber indicates an expected call of GetOrderByNumber.
func (mr *MockStorageMockRecorder) GetOrderByNumber(ctx, orderNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByNumber", reflect.TypeOf((*MockStorage)(nil).GetOrderByNumber), ctx, orderNumber)
}

// UpdateGoodAccrual mocks base method.
func (m *MockStorage) UpdateGoodAccrual(ctx context.Context, orderID, matchID uuid.UUID, accrual float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGoodAccrual", ctx, orderID, matchID, accrual)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGoodAccrual indicates an expected call of UpdateGoodAccrual.
func (mr *MockStorageMockRecorder) UpdateGoodAccrual(ctx, orderID, matchID, accrual interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGoodAccrual", reflect.TypeOf((*MockStorage)(nil).UpdateGoodAccrual), ctx, orderID, matchID, accrual)
}

// UpdateOrder mocks base method.
func (m *MockStorage) UpdateOrder(ctx context.Context, order *storage.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockStorageMockRecorder) UpdateOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockStorage)(nil).UpdateOrder), ctx, order)
}

// MockAccrual is a mock of Accrual interface.
type MockAccrual struct {
	ctrl     *gomock.Controller
	recorder *MockAccrualMockRecorder
}

// MockAccrualMockRecorder is the mock recorder for MockAccrual.
type MockAccrualMockRecorder struct {
	mock *MockAccrual
}

// NewMockAccrual creates a new mock instance.
func NewMockAccrual(ctrl *gomock.Controller) *MockAccrual {
	mock := &MockAccrual{ctrl: ctrl}
	mock.recorder = &MockAccrualMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccrual) EXPECT() *MockAccrualMockRecorder {
	return m.recorder
}

// CreateInvalidOrder mocks base method.
func (m *MockAccrual) CreateInvalidOrder(ctx context.Context, order string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvalidOrder", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInvalidOrder indicates an expected call of CreateInvalidOrder.
func (mr *MockAccrualMockRecorder) CreateInvalidOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvalidOrder", reflect.TypeOf((*MockAccrual)(nil).CreateInvalidOrder), ctx, order)
}

// CreateMatch mocks base method.
func (m *MockAccrual) CreateMatch(ctx context.Context, match *storage.Match) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMatch", ctx, match)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMatch indicates an expected call of CreateMatch.
func (mr *MockAccrualMockRecorder) CreateMatch(ctx, match interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMatch", reflect.TypeOf((*MockAccrual)(nil).CreateMatch), ctx, match)
}

// CreateOrderWithGoods mocks base method.
func (m *MockAccrual) CreateOrderWithGoods(ctx context.Context, order string, goods []*storage.Goods) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderWithGoods", ctx, order, goods)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderWithGoods indicates an expected call of CreateOrderWithGoods.
func (mr *MockAccrualMockRecorder) CreateOrderWithGoods(ctx, order, goods interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderWithGoods", reflect.TypeOf((*MockAccrual)(nil).CreateOrderWithGoods), ctx, order, goods)
}

// GetMatchByName mocks base method.
func (m *MockAccrual) GetMatchByName(ctx context.Context, matchName string) (*storage.MatchOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchByName", ctx, matchName)
	ret0, _ := ret[0].(*storage.MatchOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchByName indicates an expected call of GetMatchByName.
func (mr *MockAccrualMockRecorder) GetMatchByName(ctx, matchName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchByName", reflect.TypeOf((*MockAccrual)(nil).GetMatchByName), ctx, matchName)
}

// GetOrderByNumber mocks base method.
func (m *MockAccrual) GetOrderByNumber(ctx context.Context, orderNumber string) (*storage.OrderOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByNumber", ctx, orderNumber)
	ret0, _ := ret[0].(*storage.OrderOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByNumber indicates an expected call of GetOrderByNumber.
func (mr *MockAccrualMockRecorder) GetOrderByNumber(ctx, orderNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByNumber", reflect.TypeOf((*MockAccrual)(nil).GetOrderByNumber), ctx, orderNumber)
}

// UpdateGoodAccrual mocks base method.
func (m *MockAccrual) UpdateGoodAccrual(ctx context.Context, orderID, matchID uuid.UUID, accrual float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGoodAccrual", ctx, orderID, matchID, accrual)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGoodAccrual indicates an expected call of UpdateGoodAccrual.
func (mr *MockAccrualMockRecorder) UpdateGoodAccrual(ctx, orderID, matchID, accrual interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGoodAccrual", reflect.TypeOf((*MockAccrual)(nil).UpdateGoodAccrual), ctx, orderID, matchID, accrual)
}

// UpdateOrder mocks base method.
func (m *MockAccrual) UpdateOrder(ctx context.Context, order *storage.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockAccrualMockRecorder) UpdateOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockAccrual)(nil).UpdateOrder), ctx, order)
}