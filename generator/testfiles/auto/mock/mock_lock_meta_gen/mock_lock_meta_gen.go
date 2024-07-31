// Code generated by MockGen. DO NOT EDIT.
// Source: lock_meta_gen.go
//
// Generated by this command:
//
//	mockgen -source lock_meta_gen.go -destination mock/mock_lock_meta_gen/mock_lock_meta_gen.go
//

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	reflect "reflect"

	firestore "cloud.google.com/go/firestore"
	model "github.com/go-generalize/volcago/generator/testfiles/auto"
	gomock "go.uber.org/mock/gomock"
)

// MockLockMetaRepository is a mock of LockMetaRepository interface.
type MockLockMetaRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLockMetaRepositoryMockRecorder
}

// MockLockMetaRepositoryMockRecorder is the mock recorder for MockLockMetaRepository.
type MockLockMetaRepositoryMockRecorder struct {
	mock *MockLockMetaRepository
}

// NewMockLockMetaRepository creates a new mock instance.
func NewMockLockMetaRepository(ctrl *gomock.Controller) *MockLockMetaRepository {
	mock := &MockLockMetaRepository{ctrl: ctrl}
	mock.recorder = &MockLockMetaRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLockMetaRepository) EXPECT() *MockLockMetaRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockLockMetaRepository) Delete(ctx context.Context, subject *model.LockMeta, opts ...model.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, subject}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLockMetaRepositoryMockRecorder) Delete(ctx, subject any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLockMetaRepository)(nil).Delete), varargs...)
}

// DeleteByID mocks base method.
func (m *MockLockMetaRepository) DeleteByID(ctx context.Context, id string, opts ...model.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteByID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockLockMetaRepositoryMockRecorder) DeleteByID(ctx, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockLockMetaRepository)(nil).DeleteByID), varargs...)
}

// DeleteByIDWithTx mocks base method.
func (m *MockLockMetaRepository) DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...model.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, tx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteByIDWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByIDWithTx indicates an expected call of DeleteByIDWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) DeleteByIDWithTx(ctx, tx, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, tx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByIDWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).DeleteByIDWithTx), varargs...)
}

// DeleteMulti mocks base method.
func (m *MockLockMetaRepository) DeleteMulti(ctx context.Context, subjects []*model.LockMeta, opts ...model.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, subjects}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMulti", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMulti indicates an expected call of DeleteMulti.
func (mr *MockLockMetaRepositoryMockRecorder) DeleteMulti(ctx, subjects any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, subjects}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMulti", reflect.TypeOf((*MockLockMetaRepository)(nil).DeleteMulti), varargs...)
}

// DeleteMultiByIDs mocks base method.
func (m *MockLockMetaRepository) DeleteMultiByIDs(ctx context.Context, ids []string, opts ...model.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiByIDs", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiByIDs indicates an expected call of DeleteMultiByIDs.
func (mr *MockLockMetaRepositoryMockRecorder) DeleteMultiByIDs(ctx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiByIDs", reflect.TypeOf((*MockLockMetaRepository)(nil).DeleteMultiByIDs), varargs...)
}

// DeleteMultiByIDsWithTx mocks base method.
func (m *MockLockMetaRepository) DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...model.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, tx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiByIDsWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiByIDsWithTx indicates an expected call of DeleteMultiByIDsWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) DeleteMultiByIDsWithTx(ctx, tx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, tx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiByIDsWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).DeleteMultiByIDsWithTx), varargs...)
}

// DeleteMultiWithTx mocks base method.
func (m *MockLockMetaRepository) DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.LockMeta, opts ...model.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, tx, subjects}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiWithTx indicates an expected call of DeleteMultiWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) DeleteMultiWithTx(ctx, tx, subjects any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, tx, subjects}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).DeleteMultiWithTx), varargs...)
}

// DeleteWithTx mocks base method.
func (m *MockLockMetaRepository) DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.LockMeta, opts ...model.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, tx, subject}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWithTx indicates an expected call of DeleteWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) DeleteWithTx(ctx, tx, subject any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, tx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).DeleteWithTx), varargs...)
}

// Get mocks base method.
func (m *MockLockMetaRepository) Get(ctx context.Context, id string, opts ...model.GetOption) (*model.LockMeta, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*model.LockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLockMetaRepositoryMockRecorder) Get(ctx, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLockMetaRepository)(nil).Get), varargs...)
}

// GetCollection mocks base method.
func (m *MockLockMetaRepository) GetCollection() *firestore.CollectionRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection")
	ret0, _ := ret[0].(*firestore.CollectionRef)
	return ret0
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockLockMetaRepositoryMockRecorder) GetCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockLockMetaRepository)(nil).GetCollection))
}

// GetCollectionName mocks base method.
func (m *MockLockMetaRepository) GetCollectionName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCollectionName indicates an expected call of GetCollectionName.
func (mr *MockLockMetaRepositoryMockRecorder) GetCollectionName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionName", reflect.TypeOf((*MockLockMetaRepository)(nil).GetCollectionName))
}

// GetDocRef mocks base method.
func (m *MockLockMetaRepository) GetDocRef(id string) *firestore.DocumentRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocRef", id)
	ret0, _ := ret[0].(*firestore.DocumentRef)
	return ret0
}

// GetDocRef indicates an expected call of GetDocRef.
func (mr *MockLockMetaRepositoryMockRecorder) GetDocRef(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocRef", reflect.TypeOf((*MockLockMetaRepository)(nil).GetDocRef), id)
}

// GetMulti mocks base method.
func (m *MockLockMetaRepository) GetMulti(ctx context.Context, ids []string, opts ...model.GetOption) ([]*model.LockMeta, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMulti", varargs...)
	ret0, _ := ret[0].([]*model.LockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMulti indicates an expected call of GetMulti.
func (mr *MockLockMetaRepositoryMockRecorder) GetMulti(ctx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockLockMetaRepository)(nil).GetMulti), varargs...)
}

// GetMultiWithTx mocks base method.
func (m *MockLockMetaRepository) GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...model.GetOption) ([]*model.LockMeta, error) {
	m.ctrl.T.Helper()
	varargs := []any{tx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMultiWithTx", varargs...)
	ret0, _ := ret[0].([]*model.LockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiWithTx indicates an expected call of GetMultiWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) GetMultiWithTx(tx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{tx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).GetMultiWithTx), varargs...)
}

// GetWithDoc mocks base method.
func (m *MockLockMetaRepository) GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...model.GetOption) (*model.LockMeta, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, doc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithDoc", varargs...)
	ret0, _ := ret[0].(*model.LockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDoc indicates an expected call of GetWithDoc.
func (mr *MockLockMetaRepositoryMockRecorder) GetWithDoc(ctx, doc any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, doc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDoc", reflect.TypeOf((*MockLockMetaRepository)(nil).GetWithDoc), varargs...)
}

// GetWithDocWithTx mocks base method.
func (m *MockLockMetaRepository) GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...model.GetOption) (*model.LockMeta, error) {
	m.ctrl.T.Helper()
	varargs := []any{tx, doc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithDocWithTx", varargs...)
	ret0, _ := ret[0].(*model.LockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDocWithTx indicates an expected call of GetWithDocWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) GetWithDocWithTx(tx, doc any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{tx, doc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDocWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).GetWithDocWithTx), varargs...)
}

// GetWithTx mocks base method.
func (m *MockLockMetaRepository) GetWithTx(tx *firestore.Transaction, id string, opts ...model.GetOption) (*model.LockMeta, error) {
	m.ctrl.T.Helper()
	varargs := []any{tx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithTx", varargs...)
	ret0, _ := ret[0].(*model.LockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithTx indicates an expected call of GetWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) GetWithTx(tx, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{tx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).GetWithTx), varargs...)
}

// Insert mocks base method.
func (m *MockLockMetaRepository) Insert(ctx context.Context, subject *model.LockMeta) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, subject)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockLockMetaRepositoryMockRecorder) Insert(ctx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockLockMetaRepository)(nil).Insert), ctx, subject)
}

// InsertMulti mocks base method.
func (m *MockLockMetaRepository) InsertMulti(ctx context.Context, subjects []*model.LockMeta) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMulti", ctx, subjects)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMulti indicates an expected call of InsertMulti.
func (mr *MockLockMetaRepositoryMockRecorder) InsertMulti(ctx, subjects any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMulti", reflect.TypeOf((*MockLockMetaRepository)(nil).InsertMulti), ctx, subjects)
}

// InsertMultiWithTx mocks base method.
func (m *MockLockMetaRepository) InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.LockMeta) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMultiWithTx", ctx, tx, subjects)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMultiWithTx indicates an expected call of InsertMultiWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) InsertMultiWithTx(ctx, tx, subjects any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMultiWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).InsertMultiWithTx), ctx, tx, subjects)
}

// InsertWithTx mocks base method.
func (m *MockLockMetaRepository) InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.LockMeta) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertWithTx", ctx, tx, subject)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertWithTx indicates an expected call of InsertWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) InsertWithTx(ctx, tx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).InsertWithTx), ctx, tx, subject)
}

// RunInTransaction mocks base method.
func (m *MockLockMetaRepository) RunInTransaction() func(context.Context, func(context.Context, *firestore.Transaction) error, ...firestore.TransactionOption) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInTransaction")
	ret0, _ := ret[0].(func(context.Context, func(context.Context, *firestore.Transaction) error, ...firestore.TransactionOption) error)
	return ret0
}

// RunInTransaction indicates an expected call of RunInTransaction.
func (mr *MockLockMetaRepositoryMockRecorder) RunInTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInTransaction", reflect.TypeOf((*MockLockMetaRepository)(nil).RunInTransaction))
}

// Search mocks base method.
func (m *MockLockMetaRepository) Search(ctx context.Context, param *model.LockMetaSearchParam, q *firestore.Query) ([]*model.LockMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, param, q)
	ret0, _ := ret[0].([]*model.LockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockLockMetaRepositoryMockRecorder) Search(ctx, param, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockLockMetaRepository)(nil).Search), ctx, param, q)
}

// SearchByParam mocks base method.
func (m *MockLockMetaRepository) SearchByParam(ctx context.Context, param *model.LockMetaSearchParam) ([]*model.LockMeta, *model.PagingResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByParam", ctx, param)
	ret0, _ := ret[0].([]*model.LockMeta)
	ret1, _ := ret[1].(*model.PagingResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchByParam indicates an expected call of SearchByParam.
func (mr *MockLockMetaRepositoryMockRecorder) SearchByParam(ctx, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByParam", reflect.TypeOf((*MockLockMetaRepository)(nil).SearchByParam), ctx, param)
}

// SearchByParamWithTx mocks base method.
func (m *MockLockMetaRepository) SearchByParamWithTx(tx *firestore.Transaction, param *model.LockMetaSearchParam) ([]*model.LockMeta, *model.PagingResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByParamWithTx", tx, param)
	ret0, _ := ret[0].([]*model.LockMeta)
	ret1, _ := ret[1].(*model.PagingResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchByParamWithTx indicates an expected call of SearchByParamWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) SearchByParamWithTx(tx, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByParamWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).SearchByParamWithTx), tx, param)
}

// SearchWithTx mocks base method.
func (m *MockLockMetaRepository) SearchWithTx(tx *firestore.Transaction, param *model.LockMetaSearchParam, q *firestore.Query) ([]*model.LockMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithTx", tx, param, q)
	ret0, _ := ret[0].([]*model.LockMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithTx indicates an expected call of SearchWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) SearchWithTx(tx, param, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).SearchWithTx), tx, param, q)
}

// StrictUpdate mocks base method.
func (m *MockLockMetaRepository) StrictUpdate(ctx context.Context, id string, param *model.LockMetaUpdateParam, opts ...firestore.Precondition) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, id, param}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StrictUpdate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StrictUpdate indicates an expected call of StrictUpdate.
func (mr *MockLockMetaRepositoryMockRecorder) StrictUpdate(ctx, id, param any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, id, param}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrictUpdate", reflect.TypeOf((*MockLockMetaRepository)(nil).StrictUpdate), varargs...)
}

// StrictUpdateWithTx mocks base method.
func (m *MockLockMetaRepository) StrictUpdateWithTx(tx *firestore.Transaction, id string, param *model.LockMetaUpdateParam, opts ...firestore.Precondition) error {
	m.ctrl.T.Helper()
	varargs := []any{tx, id, param}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StrictUpdateWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StrictUpdateWithTx indicates an expected call of StrictUpdateWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) StrictUpdateWithTx(tx, id, param any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{tx, id, param}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrictUpdateWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).StrictUpdateWithTx), varargs...)
}

// Update mocks base method.
func (m *MockLockMetaRepository) Update(ctx context.Context, subject *model.LockMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockLockMetaRepositoryMockRecorder) Update(ctx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLockMetaRepository)(nil).Update), ctx, subject)
}

// UpdateMulti mocks base method.
func (m *MockLockMetaRepository) UpdateMulti(ctx context.Context, subjects []*model.LockMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMulti", ctx, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMulti indicates an expected call of UpdateMulti.
func (mr *MockLockMetaRepositoryMockRecorder) UpdateMulti(ctx, subjects any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMulti", reflect.TypeOf((*MockLockMetaRepository)(nil).UpdateMulti), ctx, subjects)
}

// UpdateMultiWithTx mocks base method.
func (m *MockLockMetaRepository) UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.LockMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMultiWithTx", ctx, tx, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMultiWithTx indicates an expected call of UpdateMultiWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) UpdateMultiWithTx(ctx, tx, subjects any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMultiWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).UpdateMultiWithTx), ctx, tx, subjects)
}

// UpdateWithTx mocks base method.
func (m *MockLockMetaRepository) UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.LockMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithTx", ctx, tx, subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWithTx indicates an expected call of UpdateWithTx.
func (mr *MockLockMetaRepositoryMockRecorder) UpdateWithTx(ctx, tx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithTx", reflect.TypeOf((*MockLockMetaRepository)(nil).UpdateWithTx), ctx, tx, subject)
}

// MockLockMetaRepositoryMiddleware is a mock of LockMetaRepositoryMiddleware interface.
type MockLockMetaRepositoryMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockLockMetaRepositoryMiddlewareMockRecorder
}

// MockLockMetaRepositoryMiddlewareMockRecorder is the mock recorder for MockLockMetaRepositoryMiddleware.
type MockLockMetaRepositoryMiddlewareMockRecorder struct {
	mock *MockLockMetaRepositoryMiddleware
}

// NewMockLockMetaRepositoryMiddleware creates a new mock instance.
func NewMockLockMetaRepositoryMiddleware(ctrl *gomock.Controller) *MockLockMetaRepositoryMiddleware {
	mock := &MockLockMetaRepositoryMiddleware{ctrl: ctrl}
	mock.recorder = &MockLockMetaRepositoryMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLockMetaRepositoryMiddleware) EXPECT() *MockLockMetaRepositoryMiddlewareMockRecorder {
	return m.recorder
}

// BeforeDelete mocks base method.
func (m *MockLockMetaRepositoryMiddleware) BeforeDelete(ctx context.Context, subject *model.LockMeta, opts ...model.DeleteOption) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, subject}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BeforeDelete", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeDelete indicates an expected call of BeforeDelete.
func (mr *MockLockMetaRepositoryMiddlewareMockRecorder) BeforeDelete(ctx, subject any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelete", reflect.TypeOf((*MockLockMetaRepositoryMiddleware)(nil).BeforeDelete), varargs...)
}

// BeforeDeleteByID mocks base method.
func (m *MockLockMetaRepositoryMiddleware) BeforeDeleteByID(ctx context.Context, ids []string, opts ...model.DeleteOption) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BeforeDeleteByID", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeDeleteByID indicates an expected call of BeforeDeleteByID.
func (mr *MockLockMetaRepositoryMiddlewareMockRecorder) BeforeDeleteByID(ctx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDeleteByID", reflect.TypeOf((*MockLockMetaRepositoryMiddleware)(nil).BeforeDeleteByID), varargs...)
}

// BeforeInsert mocks base method.
func (m *MockLockMetaRepositoryMiddleware) BeforeInsert(ctx context.Context, subject *model.LockMeta) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeInsert", ctx, subject)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeInsert indicates an expected call of BeforeInsert.
func (mr *MockLockMetaRepositoryMiddlewareMockRecorder) BeforeInsert(ctx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeInsert", reflect.TypeOf((*MockLockMetaRepositoryMiddleware)(nil).BeforeInsert), ctx, subject)
}

// BeforeUpdate mocks base method.
func (m *MockLockMetaRepositoryMiddleware) BeforeUpdate(ctx context.Context, old, subject *model.LockMeta) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeUpdate", ctx, old, subject)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeUpdate indicates an expected call of BeforeUpdate.
func (mr *MockLockMetaRepositoryMiddlewareMockRecorder) BeforeUpdate(ctx, old, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeUpdate", reflect.TypeOf((*MockLockMetaRepositoryMiddleware)(nil).BeforeUpdate), ctx, old, subject)
}
