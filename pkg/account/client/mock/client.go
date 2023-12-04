// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"

	account "github.com/bucketeer-io/bucketeer/proto/account"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ChangeAPIKeyName mocks base method.
func (m *MockClient) ChangeAPIKeyName(ctx context.Context, in *account.ChangeAPIKeyNameRequest, opts ...grpc.CallOption) (*account.ChangeAPIKeyNameResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeAPIKeyName", varargs...)
	ret0, _ := ret[0].(*account.ChangeAPIKeyNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeAPIKeyName indicates an expected call of ChangeAPIKeyName.
func (mr *MockClientMockRecorder) ChangeAPIKeyName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAPIKeyName", reflect.TypeOf((*MockClient)(nil).ChangeAPIKeyName), varargs...)
}

// ChangeAccountRole mocks base method.
func (m *MockClient) ChangeAccountRole(ctx context.Context, in *account.ChangeAccountRoleRequest, opts ...grpc.CallOption) (*account.ChangeAccountRoleResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeAccountRole", varargs...)
	ret0, _ := ret[0].(*account.ChangeAccountRoleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeAccountRole indicates an expected call of ChangeAccountRole.
func (mr *MockClientMockRecorder) ChangeAccountRole(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAccountRole", reflect.TypeOf((*MockClient)(nil).ChangeAccountRole), varargs...)
}

// Close mocks base method.
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// ConvertAccount mocks base method.
func (m *MockClient) ConvertAccount(ctx context.Context, in *account.ConvertAccountRequest, opts ...grpc.CallOption) (*account.ConvertAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConvertAccount", varargs...)
	ret0, _ := ret[0].(*account.ConvertAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertAccount indicates an expected call of ConvertAccount.
func (mr *MockClientMockRecorder) ConvertAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertAccount", reflect.TypeOf((*MockClient)(nil).ConvertAccount), varargs...)
}

// CreateAPIKey mocks base method.
func (m *MockClient) CreateAPIKey(ctx context.Context, in *account.CreateAPIKeyRequest, opts ...grpc.CallOption) (*account.CreateAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAPIKey", varargs...)
	ret0, _ := ret[0].(*account.CreateAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAPIKey indicates an expected call of CreateAPIKey.
func (mr *MockClientMockRecorder) CreateAPIKey(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAPIKey", reflect.TypeOf((*MockClient)(nil).CreateAPIKey), varargs...)
}

// CreateAccount mocks base method.
func (m *MockClient) CreateAccount(ctx context.Context, in *account.CreateAccountRequest, opts ...grpc.CallOption) (*account.CreateAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAccount", varargs...)
	ret0, _ := ret[0].(*account.CreateAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockClientMockRecorder) CreateAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockClient)(nil).CreateAccount), varargs...)
}

// CreateAccountV2 mocks base method.
func (m *MockClient) CreateAccountV2(ctx context.Context, in *account.CreateAccountV2Request, opts ...grpc.CallOption) (*account.CreateAccountV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAccountV2", varargs...)
	ret0, _ := ret[0].(*account.CreateAccountV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccountV2 indicates an expected call of CreateAccountV2.
func (mr *MockClientMockRecorder) CreateAccountV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccountV2", reflect.TypeOf((*MockClient)(nil).CreateAccountV2), varargs...)
}

// CreateAdminAccount mocks base method.
func (m *MockClient) CreateAdminAccount(ctx context.Context, in *account.CreateAdminAccountRequest, opts ...grpc.CallOption) (*account.CreateAdminAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAdminAccount", varargs...)
	ret0, _ := ret[0].(*account.CreateAdminAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdminAccount indicates an expected call of CreateAdminAccount.
func (mr *MockClientMockRecorder) CreateAdminAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdminAccount", reflect.TypeOf((*MockClient)(nil).CreateAdminAccount), varargs...)
}

// DeleteAccountV2 mocks base method.
func (m *MockClient) DeleteAccountV2(ctx context.Context, in *account.DeleteAccountV2Request, opts ...grpc.CallOption) (*account.DeleteAccountV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAccountV2", varargs...)
	ret0, _ := ret[0].(*account.DeleteAccountV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccountV2 indicates an expected call of DeleteAccountV2.
func (mr *MockClientMockRecorder) DeleteAccountV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccountV2", reflect.TypeOf((*MockClient)(nil).DeleteAccountV2), varargs...)
}

// DisableAPIKey mocks base method.
func (m *MockClient) DisableAPIKey(ctx context.Context, in *account.DisableAPIKeyRequest, opts ...grpc.CallOption) (*account.DisableAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableAPIKey", varargs...)
	ret0, _ := ret[0].(*account.DisableAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableAPIKey indicates an expected call of DisableAPIKey.
func (mr *MockClientMockRecorder) DisableAPIKey(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableAPIKey", reflect.TypeOf((*MockClient)(nil).DisableAPIKey), varargs...)
}

// DisableAccount mocks base method.
func (m *MockClient) DisableAccount(ctx context.Context, in *account.DisableAccountRequest, opts ...grpc.CallOption) (*account.DisableAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableAccount", varargs...)
	ret0, _ := ret[0].(*account.DisableAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableAccount indicates an expected call of DisableAccount.
func (mr *MockClientMockRecorder) DisableAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableAccount", reflect.TypeOf((*MockClient)(nil).DisableAccount), varargs...)
}

// DisableAccountV2 mocks base method.
func (m *MockClient) DisableAccountV2(ctx context.Context, in *account.DisableAccountV2Request, opts ...grpc.CallOption) (*account.DisableAccountV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableAccountV2", varargs...)
	ret0, _ := ret[0].(*account.DisableAccountV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableAccountV2 indicates an expected call of DisableAccountV2.
func (mr *MockClientMockRecorder) DisableAccountV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableAccountV2", reflect.TypeOf((*MockClient)(nil).DisableAccountV2), varargs...)
}

// DisableAdminAccount mocks base method.
func (m *MockClient) DisableAdminAccount(ctx context.Context, in *account.DisableAdminAccountRequest, opts ...grpc.CallOption) (*account.DisableAdminAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableAdminAccount", varargs...)
	ret0, _ := ret[0].(*account.DisableAdminAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableAdminAccount indicates an expected call of DisableAdminAccount.
func (mr *MockClientMockRecorder) DisableAdminAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableAdminAccount", reflect.TypeOf((*MockClient)(nil).DisableAdminAccount), varargs...)
}

// EnableAPIKey mocks base method.
func (m *MockClient) EnableAPIKey(ctx context.Context, in *account.EnableAPIKeyRequest, opts ...grpc.CallOption) (*account.EnableAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableAPIKey", varargs...)
	ret0, _ := ret[0].(*account.EnableAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableAPIKey indicates an expected call of EnableAPIKey.
func (mr *MockClientMockRecorder) EnableAPIKey(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAPIKey", reflect.TypeOf((*MockClient)(nil).EnableAPIKey), varargs...)
}

// EnableAccount mocks base method.
func (m *MockClient) EnableAccount(ctx context.Context, in *account.EnableAccountRequest, opts ...grpc.CallOption) (*account.EnableAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableAccount", varargs...)
	ret0, _ := ret[0].(*account.EnableAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableAccount indicates an expected call of EnableAccount.
func (mr *MockClientMockRecorder) EnableAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAccount", reflect.TypeOf((*MockClient)(nil).EnableAccount), varargs...)
}

// EnableAccountV2 mocks base method.
func (m *MockClient) EnableAccountV2(ctx context.Context, in *account.EnableAccountV2Request, opts ...grpc.CallOption) (*account.EnableAccountV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableAccountV2", varargs...)
	ret0, _ := ret[0].(*account.EnableAccountV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableAccountV2 indicates an expected call of EnableAccountV2.
func (mr *MockClientMockRecorder) EnableAccountV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAccountV2", reflect.TypeOf((*MockClient)(nil).EnableAccountV2), varargs...)
}

// EnableAdminAccount mocks base method.
func (m *MockClient) EnableAdminAccount(ctx context.Context, in *account.EnableAdminAccountRequest, opts ...grpc.CallOption) (*account.EnableAdminAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableAdminAccount", varargs...)
	ret0, _ := ret[0].(*account.EnableAdminAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableAdminAccount indicates an expected call of EnableAdminAccount.
func (mr *MockClientMockRecorder) EnableAdminAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAdminAccount", reflect.TypeOf((*MockClient)(nil).EnableAdminAccount), varargs...)
}

// GetAPIKey mocks base method.
func (m *MockClient) GetAPIKey(ctx context.Context, in *account.GetAPIKeyRequest, opts ...grpc.CallOption) (*account.GetAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAPIKey", varargs...)
	ret0, _ := ret[0].(*account.GetAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIKey indicates an expected call of GetAPIKey.
func (mr *MockClientMockRecorder) GetAPIKey(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIKey", reflect.TypeOf((*MockClient)(nil).GetAPIKey), varargs...)
}

// GetAPIKeyBySearchingAllEnvironments mocks base method.
func (m *MockClient) GetAPIKeyBySearchingAllEnvironments(ctx context.Context, in *account.GetAPIKeyBySearchingAllEnvironmentsRequest, opts ...grpc.CallOption) (*account.GetAPIKeyBySearchingAllEnvironmentsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAPIKeyBySearchingAllEnvironments", varargs...)
	ret0, _ := ret[0].(*account.GetAPIKeyBySearchingAllEnvironmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIKeyBySearchingAllEnvironments indicates an expected call of GetAPIKeyBySearchingAllEnvironments.
func (mr *MockClientMockRecorder) GetAPIKeyBySearchingAllEnvironments(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIKeyBySearchingAllEnvironments", reflect.TypeOf((*MockClient)(nil).GetAPIKeyBySearchingAllEnvironments), varargs...)
}

// GetAccount mocks base method.
func (m *MockClient) GetAccount(ctx context.Context, in *account.GetAccountRequest, opts ...grpc.CallOption) (*account.GetAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccount", varargs...)
	ret0, _ := ret[0].(*account.GetAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockClientMockRecorder) GetAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockClient)(nil).GetAccount), varargs...)
}

// GetAccountV2 mocks base method.
func (m *MockClient) GetAccountV2(ctx context.Context, in *account.GetAccountV2Request, opts ...grpc.CallOption) (*account.GetAccountV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountV2", varargs...)
	ret0, _ := ret[0].(*account.GetAccountV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountV2 indicates an expected call of GetAccountV2.
func (mr *MockClientMockRecorder) GetAccountV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountV2", reflect.TypeOf((*MockClient)(nil).GetAccountV2), varargs...)
}

// GetAdminAccount mocks base method.
func (m *MockClient) GetAdminAccount(ctx context.Context, in *account.GetAdminAccountRequest, opts ...grpc.CallOption) (*account.GetAdminAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAdminAccount", varargs...)
	ret0, _ := ret[0].(*account.GetAdminAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminAccount indicates an expected call of GetAdminAccount.
func (mr *MockClientMockRecorder) GetAdminAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminAccount", reflect.TypeOf((*MockClient)(nil).GetAdminAccount), varargs...)
}

// GetMeByEmailV2 mocks base method.
func (m *MockClient) GetMeByEmailV2(ctx context.Context, in *account.GetMeByEmailV2Request, opts ...grpc.CallOption) (*account.GetMeV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMeByEmailV2", varargs...)
	ret0, _ := ret[0].(*account.GetMeV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeByEmailV2 indicates an expected call of GetMeByEmailV2.
func (mr *MockClientMockRecorder) GetMeByEmailV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeByEmailV2", reflect.TypeOf((*MockClient)(nil).GetMeByEmailV2), varargs...)
}

// GetMeV2 mocks base method.
func (m *MockClient) GetMeV2(ctx context.Context, in *account.GetMeV2Request, opts ...grpc.CallOption) (*account.GetMeV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMeV2", varargs...)
	ret0, _ := ret[0].(*account.GetMeV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeV2 indicates an expected call of GetMeV2.
func (mr *MockClientMockRecorder) GetMeV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeV2", reflect.TypeOf((*MockClient)(nil).GetMeV2), varargs...)
}

// ListAPIKeys mocks base method.
func (m *MockClient) ListAPIKeys(ctx context.Context, in *account.ListAPIKeysRequest, opts ...grpc.CallOption) (*account.ListAPIKeysResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAPIKeys", varargs...)
	ret0, _ := ret[0].(*account.ListAPIKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAPIKeys indicates an expected call of ListAPIKeys.
func (mr *MockClientMockRecorder) ListAPIKeys(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAPIKeys", reflect.TypeOf((*MockClient)(nil).ListAPIKeys), varargs...)
}

// ListAccounts mocks base method.
func (m *MockClient) ListAccounts(ctx context.Context, in *account.ListAccountsRequest, opts ...grpc.CallOption) (*account.ListAccountsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccounts", varargs...)
	ret0, _ := ret[0].(*account.ListAccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockClientMockRecorder) ListAccounts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockClient)(nil).ListAccounts), varargs...)
}

// ListAccountsV2 mocks base method.
func (m *MockClient) ListAccountsV2(ctx context.Context, in *account.ListAccountsV2Request, opts ...grpc.CallOption) (*account.ListAccountsV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountsV2", varargs...)
	ret0, _ := ret[0].(*account.ListAccountsV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountsV2 indicates an expected call of ListAccountsV2.
func (mr *MockClientMockRecorder) ListAccountsV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsV2", reflect.TypeOf((*MockClient)(nil).ListAccountsV2), varargs...)
}

// ListAdminAccounts mocks base method.
func (m *MockClient) ListAdminAccounts(ctx context.Context, in *account.ListAdminAccountsRequest, opts ...grpc.CallOption) (*account.ListAdminAccountsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAdminAccounts", varargs...)
	ret0, _ := ret[0].(*account.ListAdminAccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAdminAccounts indicates an expected call of ListAdminAccounts.
func (mr *MockClientMockRecorder) ListAdminAccounts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdminAccounts", reflect.TypeOf((*MockClient)(nil).ListAdminAccounts), varargs...)
}

// UpdateAccountV2 mocks base method.
func (m *MockClient) UpdateAccountV2(ctx context.Context, in *account.UpdateAccountV2Request, opts ...grpc.CallOption) (*account.UpdateAccountV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAccountV2", varargs...)
	ret0, _ := ret[0].(*account.UpdateAccountV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccountV2 indicates an expected call of UpdateAccountV2.
func (mr *MockClientMockRecorder) UpdateAccountV2(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountV2", reflect.TypeOf((*MockClient)(nil).UpdateAccountV2), varargs...)
}
