// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/reviewpad/api/go/services (interfaces: HostClient)

// Package services_mocks is a generated GoMock package.
package services_mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	services "github.com/reviewpad/api/go/services"
	grpc "google.golang.org/grpc"
)

// MockHostClient is a mock of HostClient interface.
type MockHostClient struct {
	ctrl     *gomock.Controller
	recorder *MockHostClientMockRecorder
}

// MockHostClientMockRecorder is the mock recorder for MockHostClient.
type MockHostClientMockRecorder struct {
	mock *MockHostClient
}

// NewMockHostClient creates a new mock instance.
func NewMockHostClient(ctrl *gomock.Controller) *MockHostClient {
	mock := &MockHostClient{ctrl: ctrl}
	mock.recorder = &MockHostClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostClient) EXPECT() *MockHostClientMockRecorder {
	return m.recorder
}

// GetPullRequest mocks base method.
func (m *MockHostClient) GetPullRequest(arg0 context.Context, arg1 *services.GetPullRequestRequest, arg2 ...grpc.CallOption) (*services.GetPullRequestReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPullRequest", varargs...)
	ret0, _ := ret[0].(*services.GetPullRequestReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPullRequest indicates an expected call of GetPullRequest.
func (mr *MockHostClientMockRecorder) GetPullRequest(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequest", reflect.TypeOf((*MockHostClient)(nil).GetPullRequest), varargs...)
}

// GetPullRequestFiles mocks base method.
func (m *MockHostClient) GetPullRequestFiles(arg0 context.Context, arg1 *services.GetPullRequestFilesRequest, arg2 ...grpc.CallOption) (*services.GetPullRequestFilesReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPullRequestFiles", varargs...)
	ret0, _ := ret[0].(*services.GetPullRequestFilesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPullRequestFiles indicates an expected call of GetPullRequestFiles.
func (mr *MockHostClientMockRecorder) GetPullRequestFiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequestFiles", reflect.TypeOf((*MockHostClient)(nil).GetPullRequestFiles), varargs...)
}

// PostGeneralComment mocks base method.
func (m *MockHostClient) PostGeneralComment(arg0 context.Context, arg1 *services.PostGeneralCommentRequest, arg2 ...grpc.CallOption) (*services.PostGeneralCommentReply, error) {
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
func (mr *MockHostClientMockRecorder) PostGeneralComment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostGeneralComment", reflect.TypeOf((*MockHostClient)(nil).PostGeneralComment), varargs...)
}
