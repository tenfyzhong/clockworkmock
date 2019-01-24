// Code generated by MockGen. DO NOT EDIT.
// Source: clockwork.go

// Package cloclworkmock is a generated GoMock package.
package cloclworkmock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockClock is a mock of Clock interface
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// After mocks base method
func (m *MockClock) After(d time.Duration) <-chan time.Time {
	ret := m.ctrl.Call(m, "After", d)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// After indicates an expected call of After
func (mr *MockClockMockRecorder) After(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "After", reflect.TypeOf((*MockClock)(nil).After), d)
}

// Sleep mocks base method
func (m *MockClock) Sleep(d time.Duration) {
	m.ctrl.Call(m, "Sleep", d)
}

// Sleep indicates an expected call of Sleep
func (mr *MockClockMockRecorder) Sleep(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MockClock)(nil).Sleep), d)
}

// Now mocks base method
func (m *MockClock) Now() time.Time {
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now
func (mr *MockClockMockRecorder) Now() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClock)(nil).Now))
}

// Since mocks base method
func (m *MockClock) Since(t time.Time) time.Duration {
	ret := m.ctrl.Call(m, "Since", t)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Since indicates an expected call of Since
func (mr *MockClockMockRecorder) Since(t interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Since", reflect.TypeOf((*MockClock)(nil).Since), t)
}

// NewTicker mocks base method
func (m *MockClock) NewTicker(d time.Duration) Ticker {
	ret := m.ctrl.Call(m, "NewTicker", d)
	ret0, _ := ret[0].(Ticker)
	return ret0
}

// NewTicker indicates an expected call of NewTicker
func (mr *MockClockMockRecorder) NewTicker(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTicker", reflect.TypeOf((*MockClock)(nil).NewTicker), d)
}

// MockFakeClock is a mock of FakeClock interface
type MockFakeClock struct {
	ctrl     *gomock.Controller
	recorder *MockFakeClockMockRecorder
}

// MockFakeClockMockRecorder is the mock recorder for MockFakeClock
type MockFakeClockMockRecorder struct {
	mock *MockFakeClock
}

// NewMockFakeClock creates a new mock instance
func NewMockFakeClock(ctrl *gomock.Controller) *MockFakeClock {
	mock := &MockFakeClock{ctrl: ctrl}
	mock.recorder = &MockFakeClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFakeClock) EXPECT() *MockFakeClockMockRecorder {
	return m.recorder
}

// After mocks base method
func (m *MockFakeClock) After(d time.Duration) <-chan time.Time {
	ret := m.ctrl.Call(m, "After", d)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// After indicates an expected call of After
func (mr *MockFakeClockMockRecorder) After(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "After", reflect.TypeOf((*MockFakeClock)(nil).After), d)
}

// Sleep mocks base method
func (m *MockFakeClock) Sleep(d time.Duration) {
	m.ctrl.Call(m, "Sleep", d)
}

// Sleep indicates an expected call of Sleep
func (mr *MockFakeClockMockRecorder) Sleep(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MockFakeClock)(nil).Sleep), d)
}

// Now mocks base method
func (m *MockFakeClock) Now() time.Time {
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now
func (mr *MockFakeClockMockRecorder) Now() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockFakeClock)(nil).Now))
}

// Since mocks base method
func (m *MockFakeClock) Since(t time.Time) time.Duration {
	ret := m.ctrl.Call(m, "Since", t)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Since indicates an expected call of Since
func (mr *MockFakeClockMockRecorder) Since(t interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Since", reflect.TypeOf((*MockFakeClock)(nil).Since), t)
}

// NewTicker mocks base method
func (m *MockFakeClock) NewTicker(d time.Duration) Ticker {
	ret := m.ctrl.Call(m, "NewTicker", d)
	ret0, _ := ret[0].(Ticker)
	return ret0
}

// NewTicker indicates an expected call of NewTicker
func (mr *MockFakeClockMockRecorder) NewTicker(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTicker", reflect.TypeOf((*MockFakeClock)(nil).NewTicker), d)
}

// Advance mocks base method
func (m *MockFakeClock) Advance(d time.Duration) {
	m.ctrl.Call(m, "Advance", d)
}

// Advance indicates an expected call of Advance
func (mr *MockFakeClockMockRecorder) Advance(d interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Advance", reflect.TypeOf((*MockFakeClock)(nil).Advance), d)
}

// BlockUntil mocks base method
func (m *MockFakeClock) BlockUntil(n int) {
	m.ctrl.Call(m, "BlockUntil", n)
}

// BlockUntil indicates an expected call of BlockUntil
func (mr *MockFakeClockMockRecorder) BlockUntil(n interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockUntil", reflect.TypeOf((*MockFakeClock)(nil).BlockUntil), n)
}