// Code generated by MockGen. DO NOT EDIT.
// Source: lock_meta2_gen.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	reflect "reflect"

	firestore "cloud.google.com/go/firestore"
	auto "github.com/go-generalize/volcago/generator/testfiles/auto"
	gomock "github.com/golang/mock/gomock"
)

// MockLockMeta2Repository is a mock of LockMeta2Repository interface.
type MockLockMeta2Repository struct {
	ctrl     *gomock.Controller
	recorder *MockLockMeta2RepositoryMockRecorder
}

// MockLockMeta2RepositoryMockRecorder is the mock recorder for MockLockMeta2Repository.
type MockLockMeta2RepositoryMockRecorder struct {
	mock *MockLockMeta2Repository
}

// NewMockLockMeta2Repository creates a new mock instance.
func NewMockLockMeta2Repository(ctrl *gomock.Controller) *MockLockMeta2Repository {
	mock := &MockLockMeta2Repository{ctrl: ctrl}
	mock.recorder = &MockLockMeta2RepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLockMeta2Repository) EXPECT() *MockLockMeta2RepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockLockMeta2Repository) Delete(ctx context.Context, subject *auto.LockMeta2, opts ...auto.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, subject}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLockMeta2RepositoryMockRecorder) Delete(ctx, subject interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLockMeta2Repository)(nil).Delete), varargs...)
}

// DeleteByID mocks base method.
func (m *MockLockMeta2Repository) DeleteByID(ctx context.Context, id string, opts ...auto.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteByID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockLockMeta2RepositoryMockRecorder) DeleteByID(ctx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockLockMeta2Repository)(nil).DeleteByID), varargs...)
}

// DeleteByIDWithTx mocks base method.
func (m *MockLockMeta2Repository) DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...auto.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteByIDWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByIDWithTx indicates an expected call of DeleteByIDWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) DeleteByIDWithTx(ctx, tx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByIDWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).DeleteByIDWithTx), varargs...)
}

// DeleteMulti mocks base method.
func (m *MockLockMeta2Repository) DeleteMulti(ctx context.Context, subjects []*auto.LockMeta2, opts ...auto.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, subjects}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMulti", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMulti indicates an expected call of DeleteMulti.
func (mr *MockLockMeta2RepositoryMockRecorder) DeleteMulti(ctx, subjects interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, subjects}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMulti", reflect.TypeOf((*MockLockMeta2Repository)(nil).DeleteMulti), varargs...)
}

// DeleteMultiByIDs mocks base method.
func (m *MockLockMeta2Repository) DeleteMultiByIDs(ctx context.Context, ids []string, opts ...auto.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiByIDs", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiByIDs indicates an expected call of DeleteMultiByIDs.
func (mr *MockLockMeta2RepositoryMockRecorder) DeleteMultiByIDs(ctx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiByIDs", reflect.TypeOf((*MockLockMeta2Repository)(nil).DeleteMultiByIDs), varargs...)
}

// DeleteMultiByIDsWithTx mocks base method.
func (m *MockLockMeta2Repository) DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...auto.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiByIDsWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiByIDsWithTx indicates an expected call of DeleteMultiByIDsWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) DeleteMultiByIDsWithTx(ctx, tx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiByIDsWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).DeleteMultiByIDsWithTx), varargs...)
}

// DeleteMultiWithTx mocks base method.
func (m *MockLockMeta2Repository) DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*auto.LockMeta2, opts ...auto.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, subjects}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiWithTx indicates an expected call of DeleteMultiWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) DeleteMultiWithTx(ctx, tx, subjects interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, subjects}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).DeleteMultiWithTx), varargs...)
}

// DeleteWithTx mocks base method.
func (m *MockLockMeta2Repository) DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *auto.LockMeta2, opts ...auto.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, subject}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWithTx indicates an expected call of DeleteWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) DeleteWithTx(ctx, tx, subject interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).DeleteWithTx), varargs...)
}

// Get mocks base method.
func (m *MockLockMeta2Repository) Get(ctx context.Context, id string, opts ...auto.GetOption) (*auto.LockMeta2, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*auto.LockMeta2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLockMeta2RepositoryMockRecorder) Get(ctx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLockMeta2Repository)(nil).Get), varargs...)
}

// GetCollection mocks base method.
func (m *MockLockMeta2Repository) GetCollection() *firestore.CollectionRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection")
	ret0, _ := ret[0].(*firestore.CollectionRef)
	return ret0
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockLockMeta2RepositoryMockRecorder) GetCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockLockMeta2Repository)(nil).GetCollection))
}

// GetCollectionName mocks base method.
func (m *MockLockMeta2Repository) GetCollectionName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCollectionName indicates an expected call of GetCollectionName.
func (mr *MockLockMeta2RepositoryMockRecorder) GetCollectionName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionName", reflect.TypeOf((*MockLockMeta2Repository)(nil).GetCollectionName))
}

// GetDocRef mocks base method.
func (m *MockLockMeta2Repository) GetDocRef(id string) *firestore.DocumentRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocRef", id)
	ret0, _ := ret[0].(*firestore.DocumentRef)
	return ret0
}

// GetDocRef indicates an expected call of GetDocRef.
func (mr *MockLockMeta2RepositoryMockRecorder) GetDocRef(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocRef", reflect.TypeOf((*MockLockMeta2Repository)(nil).GetDocRef), id)
}

// GetMulti mocks base method.
func (m *MockLockMeta2Repository) GetMulti(ctx context.Context, ids []string, opts ...auto.GetOption) ([]*auto.LockMeta2, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMulti", varargs...)
	ret0, _ := ret[0].([]*auto.LockMeta2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMulti indicates an expected call of GetMulti.
func (mr *MockLockMeta2RepositoryMockRecorder) GetMulti(ctx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockLockMeta2Repository)(nil).GetMulti), varargs...)
}

// GetMultiWithTx mocks base method.
func (m *MockLockMeta2Repository) GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...auto.GetOption) ([]*auto.LockMeta2, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMultiWithTx", varargs...)
	ret0, _ := ret[0].([]*auto.LockMeta2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiWithTx indicates an expected call of GetMultiWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) GetMultiWithTx(tx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).GetMultiWithTx), varargs...)
}

// GetWithDoc mocks base method.
func (m *MockLockMeta2Repository) GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...auto.GetOption) (*auto.LockMeta2, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, doc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithDoc", varargs...)
	ret0, _ := ret[0].(*auto.LockMeta2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDoc indicates an expected call of GetWithDoc.
func (mr *MockLockMeta2RepositoryMockRecorder) GetWithDoc(ctx, doc interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, doc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDoc", reflect.TypeOf((*MockLockMeta2Repository)(nil).GetWithDoc), varargs...)
}

// GetWithDocWithTx mocks base method.
func (m *MockLockMeta2Repository) GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...auto.GetOption) (*auto.LockMeta2, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, doc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithDocWithTx", varargs...)
	ret0, _ := ret[0].(*auto.LockMeta2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDocWithTx indicates an expected call of GetWithDocWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) GetWithDocWithTx(tx, doc interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, doc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDocWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).GetWithDocWithTx), varargs...)
}

// GetWithTx mocks base method.
func (m *MockLockMeta2Repository) GetWithTx(tx *firestore.Transaction, id string, opts ...auto.GetOption) (*auto.LockMeta2, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithTx", varargs...)
	ret0, _ := ret[0].(*auto.LockMeta2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithTx indicates an expected call of GetWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) GetWithTx(tx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).GetWithTx), varargs...)
}

// Insert mocks base method.
func (m *MockLockMeta2Repository) Insert(ctx context.Context, subject *auto.LockMeta2) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, subject)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockLockMeta2RepositoryMockRecorder) Insert(ctx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockLockMeta2Repository)(nil).Insert), ctx, subject)
}

// InsertMulti mocks base method.
func (m *MockLockMeta2Repository) InsertMulti(ctx context.Context, subjects []*auto.LockMeta2) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMulti", ctx, subjects)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMulti indicates an expected call of InsertMulti.
func (mr *MockLockMeta2RepositoryMockRecorder) InsertMulti(ctx, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMulti", reflect.TypeOf((*MockLockMeta2Repository)(nil).InsertMulti), ctx, subjects)
}

// InsertMultiWithTx mocks base method.
func (m *MockLockMeta2Repository) InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*auto.LockMeta2) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMultiWithTx", ctx, tx, subjects)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMultiWithTx indicates an expected call of InsertMultiWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) InsertMultiWithTx(ctx, tx, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMultiWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).InsertMultiWithTx), ctx, tx, subjects)
}

// InsertWithTx mocks base method.
func (m *MockLockMeta2Repository) InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *auto.LockMeta2) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertWithTx", ctx, tx, subject)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertWithTx indicates an expected call of InsertWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) InsertWithTx(ctx, tx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).InsertWithTx), ctx, tx, subject)
}

// RunInTransaction mocks base method.
func (m *MockLockMeta2Repository) RunInTransaction() func(context.Context, func(context.Context, *firestore.Transaction) error, ...firestore.TransactionOption) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInTransaction")
	ret0, _ := ret[0].(func(context.Context, func(context.Context, *firestore.Transaction) error, ...firestore.TransactionOption) error)
	return ret0
}

// RunInTransaction indicates an expected call of RunInTransaction.
func (mr *MockLockMeta2RepositoryMockRecorder) RunInTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInTransaction", reflect.TypeOf((*MockLockMeta2Repository)(nil).RunInTransaction))
}

// Search mocks base method.
func (m *MockLockMeta2Repository) Search(ctx context.Context, param *auto.LockMeta2SearchParam, q *firestore.Query) ([]*auto.LockMeta2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, param, q)
	ret0, _ := ret[0].([]*auto.LockMeta2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockLockMeta2RepositoryMockRecorder) Search(ctx, param, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockLockMeta2Repository)(nil).Search), ctx, param, q)
}

// SearchWithTx mocks base method.
func (m *MockLockMeta2Repository) SearchWithTx(tx *firestore.Transaction, param *auto.LockMeta2SearchParam, q *firestore.Query) ([]*auto.LockMeta2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithTx", tx, param, q)
	ret0, _ := ret[0].([]*auto.LockMeta2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithTx indicates an expected call of SearchWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) SearchWithTx(tx, param, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).SearchWithTx), tx, param, q)
}

// StrictUpdate mocks base method.
func (m *MockLockMeta2Repository) StrictUpdate(ctx context.Context, id string, param *auto.LockMeta2UpdateParam, opts ...firestore.Precondition) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id, param}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StrictUpdate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StrictUpdate indicates an expected call of StrictUpdate.
func (mr *MockLockMeta2RepositoryMockRecorder) StrictUpdate(ctx, id, param interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id, param}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrictUpdate", reflect.TypeOf((*MockLockMeta2Repository)(nil).StrictUpdate), varargs...)
}

// StrictUpdateWithTx mocks base method.
func (m *MockLockMeta2Repository) StrictUpdateWithTx(tx *firestore.Transaction, id string, param *auto.LockMeta2UpdateParam, opts ...firestore.Precondition) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, id, param}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StrictUpdateWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StrictUpdateWithTx indicates an expected call of StrictUpdateWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) StrictUpdateWithTx(tx, id, param interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, id, param}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrictUpdateWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).StrictUpdateWithTx), varargs...)
}

// Update mocks base method.
func (m *MockLockMeta2Repository) Update(ctx context.Context, subject *auto.LockMeta2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockLockMeta2RepositoryMockRecorder) Update(ctx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLockMeta2Repository)(nil).Update), ctx, subject)
}

// UpdateMulti mocks base method.
func (m *MockLockMeta2Repository) UpdateMulti(ctx context.Context, subjects []*auto.LockMeta2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMulti", ctx, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMulti indicates an expected call of UpdateMulti.
func (mr *MockLockMeta2RepositoryMockRecorder) UpdateMulti(ctx, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMulti", reflect.TypeOf((*MockLockMeta2Repository)(nil).UpdateMulti), ctx, subjects)
}

// UpdateMultiWithTx mocks base method.
func (m *MockLockMeta2Repository) UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*auto.LockMeta2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMultiWithTx", ctx, tx, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMultiWithTx indicates an expected call of UpdateMultiWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) UpdateMultiWithTx(ctx, tx, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMultiWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).UpdateMultiWithTx), ctx, tx, subjects)
}

// UpdateWithTx mocks base method.
func (m *MockLockMeta2Repository) UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *auto.LockMeta2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithTx", ctx, tx, subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWithTx indicates an expected call of UpdateWithTx.
func (mr *MockLockMeta2RepositoryMockRecorder) UpdateWithTx(ctx, tx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithTx", reflect.TypeOf((*MockLockMeta2Repository)(nil).UpdateWithTx), ctx, tx, subject)
}

// MockLockMeta2RepositoryMiddleware is a mock of LockMeta2RepositoryMiddleware interface.
type MockLockMeta2RepositoryMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockLockMeta2RepositoryMiddlewareMockRecorder
}

// MockLockMeta2RepositoryMiddlewareMockRecorder is the mock recorder for MockLockMeta2RepositoryMiddleware.
type MockLockMeta2RepositoryMiddlewareMockRecorder struct {
	mock *MockLockMeta2RepositoryMiddleware
}

// NewMockLockMeta2RepositoryMiddleware creates a new mock instance.
func NewMockLockMeta2RepositoryMiddleware(ctrl *gomock.Controller) *MockLockMeta2RepositoryMiddleware {
	mock := &MockLockMeta2RepositoryMiddleware{ctrl: ctrl}
	mock.recorder = &MockLockMeta2RepositoryMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLockMeta2RepositoryMiddleware) EXPECT() *MockLockMeta2RepositoryMiddlewareMockRecorder {
	return m.recorder
}

// BeforeDelete mocks base method.
func (m *MockLockMeta2RepositoryMiddleware) BeforeDelete(ctx context.Context, subject *auto.LockMeta2, opts ...auto.DeleteOption) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, subject}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BeforeDelete", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeDelete indicates an expected call of BeforeDelete.
func (mr *MockLockMeta2RepositoryMiddlewareMockRecorder) BeforeDelete(ctx, subject interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelete", reflect.TypeOf((*MockLockMeta2RepositoryMiddleware)(nil).BeforeDelete), varargs...)
}

// BeforeDeleteByID mocks base method.
func (m *MockLockMeta2RepositoryMiddleware) BeforeDeleteByID(ctx context.Context, ids []string, opts ...auto.DeleteOption) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BeforeDeleteByID", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeDeleteByID indicates an expected call of BeforeDeleteByID.
func (mr *MockLockMeta2RepositoryMiddlewareMockRecorder) BeforeDeleteByID(ctx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDeleteByID", reflect.TypeOf((*MockLockMeta2RepositoryMiddleware)(nil).BeforeDeleteByID), varargs...)
}

// BeforeInsert mocks base method.
func (m *MockLockMeta2RepositoryMiddleware) BeforeInsert(ctx context.Context, subject *auto.LockMeta2) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeInsert", ctx, subject)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeInsert indicates an expected call of BeforeInsert.
func (mr *MockLockMeta2RepositoryMiddlewareMockRecorder) BeforeInsert(ctx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeInsert", reflect.TypeOf((*MockLockMeta2RepositoryMiddleware)(nil).BeforeInsert), ctx, subject)
}

// BeforeUpdate mocks base method.
func (m *MockLockMeta2RepositoryMiddleware) BeforeUpdate(ctx context.Context, old, subject *auto.LockMeta2) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeUpdate", ctx, old, subject)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeUpdate indicates an expected call of BeforeUpdate.
func (mr *MockLockMeta2RepositoryMiddlewareMockRecorder) BeforeUpdate(ctx, old, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeUpdate", reflect.TypeOf((*MockLockMeta2RepositoryMiddleware)(nil).BeforeUpdate), ctx, old, subject)
}
