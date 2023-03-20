// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/reviewpad/api/go/services (interfaces: HostsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	services "github.com/reviewpad/api/go/services"
	grpc "google.golang.org/grpc"
)

// MockHostsClient is a mock of HostsClient interface.
type MockHostsClient struct {
	ctrl     *gomock.Controller
	recorder *MockHostsClientMockRecorder
}

// MockHostsClientMockRecorder is the mock recorder for MockHostsClient.
type MockHostsClientMockRecorder struct {
	mock *MockHostsClient
}

// NewMockHostsClient creates a new mock instance.
func NewMockHostsClient(ctrl *gomock.Controller) *MockHostsClient {
	mock := &MockHostsClient{ctrl: ctrl}
	mock.recorder = &MockHostsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostsClient) EXPECT() *MockHostsClientMockRecorder {
	return m.recorder
}

// GetCodeReview mocks base method.
func (m *MockHostsClient) GetCodeReview(arg0 context.Context, arg1 *services.GetCodeReviewRequest, arg2 ...grpc.CallOption) (*services.GetCodeReviewReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCodeReview", varargs...)
	ret0, _ := ret[0].(*services.GetCodeReviewReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCodeReview indicates an expected call of GetCodeReview.
func (mr *MockHostsClientMockRecorder) GetCodeReview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeReview", reflect.TypeOf((*MockHostsClient)(nil).GetCodeReview), varargs...)
}

// PostGeneralComment mocks base method.
func (m *MockHostsClient) PostGeneralComment(arg0 context.Context, arg1 *services.PostGeneralCommentRequest, arg2 ...grpc.CallOption) (*services.PostGeneralCommentReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostGeneralComment", varargs...)
	ret0, _ := ret[0].(*services.PostGeneralCommentReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostGeneralComment indicates an expected call of PostGeneralComment.
func (mr *MockHostsClientMockRecorder) PostGeneralComment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostGeneralComment", reflect.TypeOf((*MockHostsClient)(nil).PostGeneralComment), varargs...)
}
