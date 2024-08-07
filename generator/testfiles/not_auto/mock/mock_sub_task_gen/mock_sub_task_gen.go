// Code generated by MockGen. DO NOT EDIT.
// Source: sub_task_gen.go
//
// Generated by this command:
//
//	mockgen -source sub_task_gen.go -destination mock/mock_sub_task_gen/mock_sub_task_gen.go
//

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	reflect "reflect"

	firestore "cloud.google.com/go/firestore"
	model "github.com/go-generalize/volcago/generator/testfiles/not_auto"
	gomock "go.uber.org/mock/gomock"
)

// MockSubTaskRepository is a mock of SubTaskRepository interface.
type MockSubTaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSubTaskRepositoryMockRecorder
}

// MockSubTaskRepositoryMockRecorder is the mock recorder for MockSubTaskRepository.
type MockSubTaskRepositoryMockRecorder struct {
	mock *MockSubTaskRepository
}

// NewMockSubTaskRepository creates a new mock instance.
func NewMockSubTaskRepository(ctrl *gomock.Controller) *MockSubTaskRepository {
	mock := &MockSubTaskRepository{ctrl: ctrl}
	mock.recorder = &MockSubTaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubTaskRepository) EXPECT() *MockSubTaskRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSubTaskRepository) Delete(ctx context.Context, subject *model.SubTask, opts ...model.DeleteOption) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) Delete(ctx, subject any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSubTaskRepository)(nil).Delete), varargs...)
}

// DeleteByID mocks base method.
func (m *MockSubTaskRepository) DeleteByID(ctx context.Context, id string, opts ...model.DeleteOption) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) DeleteByID(ctx, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockSubTaskRepository)(nil).DeleteByID), varargs...)
}

// DeleteByIDWithTx mocks base method.
func (m *MockSubTaskRepository) DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...model.DeleteOption) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) DeleteByIDWithTx(ctx, tx, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, tx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByIDWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).DeleteByIDWithTx), varargs...)
}

// DeleteMulti mocks base method.
func (m *MockSubTaskRepository) DeleteMulti(ctx context.Context, subjects []*model.SubTask, opts ...model.DeleteOption) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) DeleteMulti(ctx, subjects any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, subjects}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMulti", reflect.TypeOf((*MockSubTaskRepository)(nil).DeleteMulti), varargs...)
}

// DeleteMultiByIDs mocks base method.
func (m *MockSubTaskRepository) DeleteMultiByIDs(ctx context.Context, ids []string, opts ...model.DeleteOption) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) DeleteMultiByIDs(ctx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiByIDs", reflect.TypeOf((*MockSubTaskRepository)(nil).DeleteMultiByIDs), varargs...)
}

// DeleteMultiByIDsWithTx mocks base method.
func (m *MockSubTaskRepository) DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...model.DeleteOption) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) DeleteMultiByIDsWithTx(ctx, tx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, tx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiByIDsWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).DeleteMultiByIDsWithTx), varargs...)
}

// DeleteMultiWithTx mocks base method.
func (m *MockSubTaskRepository) DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.SubTask, opts ...model.DeleteOption) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) DeleteMultiWithTx(ctx, tx, subjects any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, tx, subjects}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).DeleteMultiWithTx), varargs...)
}

// DeleteWithTx mocks base method.
func (m *MockSubTaskRepository) DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.SubTask, opts ...model.DeleteOption) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) DeleteWithTx(ctx, tx, subject any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, tx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).DeleteWithTx), varargs...)
}

// Free mocks base method.
func (m *MockSubTaskRepository) Free() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Free")
}

// Free indicates an expected call of Free.
func (mr *MockSubTaskRepositoryMockRecorder) Free() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Free", reflect.TypeOf((*MockSubTaskRepository)(nil).Free))
}

// Get mocks base method.
func (m *MockSubTaskRepository) Get(ctx context.Context, id string, opts ...model.GetOption) (*model.SubTask, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*model.SubTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSubTaskRepositoryMockRecorder) Get(ctx, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubTaskRepository)(nil).Get), varargs...)
}

// GetCollection mocks base method.
func (m *MockSubTaskRepository) GetCollection() *firestore.CollectionRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection")
	ret0, _ := ret[0].(*firestore.CollectionRef)
	return ret0
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockSubTaskRepositoryMockRecorder) GetCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockSubTaskRepository)(nil).GetCollection))
}

// GetCollectionName mocks base method.
func (m *MockSubTaskRepository) GetCollectionName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCollectionName indicates an expected call of GetCollectionName.
func (mr *MockSubTaskRepositoryMockRecorder) GetCollectionName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionName", reflect.TypeOf((*MockSubTaskRepository)(nil).GetCollectionName))
}

// GetDocRef mocks base method.
func (m *MockSubTaskRepository) GetDocRef(id string) *firestore.DocumentRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocRef", id)
	ret0, _ := ret[0].(*firestore.DocumentRef)
	return ret0
}

// GetDocRef indicates an expected call of GetDocRef.
func (mr *MockSubTaskRepositoryMockRecorder) GetDocRef(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocRef", reflect.TypeOf((*MockSubTaskRepository)(nil).GetDocRef), id)
}

// GetMulti mocks base method.
func (m *MockSubTaskRepository) GetMulti(ctx context.Context, ids []string, opts ...model.GetOption) ([]*model.SubTask, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMulti", varargs...)
	ret0, _ := ret[0].([]*model.SubTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMulti indicates an expected call of GetMulti.
func (mr *MockSubTaskRepositoryMockRecorder) GetMulti(ctx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockSubTaskRepository)(nil).GetMulti), varargs...)
}

// GetMultiWithTx mocks base method.
func (m *MockSubTaskRepository) GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...model.GetOption) ([]*model.SubTask, error) {
	m.ctrl.T.Helper()
	varargs := []any{tx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMultiWithTx", varargs...)
	ret0, _ := ret[0].([]*model.SubTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiWithTx indicates an expected call of GetMultiWithTx.
func (mr *MockSubTaskRepositoryMockRecorder) GetMultiWithTx(tx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{tx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).GetMultiWithTx), varargs...)
}

// GetWithDoc mocks base method.
func (m *MockSubTaskRepository) GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...model.GetOption) (*model.SubTask, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, doc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithDoc", varargs...)
	ret0, _ := ret[0].(*model.SubTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDoc indicates an expected call of GetWithDoc.
func (mr *MockSubTaskRepositoryMockRecorder) GetWithDoc(ctx, doc any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, doc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDoc", reflect.TypeOf((*MockSubTaskRepository)(nil).GetWithDoc), varargs...)
}

// GetWithDocWithTx mocks base method.
func (m *MockSubTaskRepository) GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...model.GetOption) (*model.SubTask, error) {
	m.ctrl.T.Helper()
	varargs := []any{tx, doc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithDocWithTx", varargs...)
	ret0, _ := ret[0].(*model.SubTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDocWithTx indicates an expected call of GetWithDocWithTx.
func (mr *MockSubTaskRepositoryMockRecorder) GetWithDocWithTx(tx, doc any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{tx, doc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDocWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).GetWithDocWithTx), varargs...)
}

// GetWithTx mocks base method.
func (m *MockSubTaskRepository) GetWithTx(tx *firestore.Transaction, id string, opts ...model.GetOption) (*model.SubTask, error) {
	m.ctrl.T.Helper()
	varargs := []any{tx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithTx", varargs...)
	ret0, _ := ret[0].(*model.SubTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithTx indicates an expected call of GetWithTx.
func (mr *MockSubTaskRepositoryMockRecorder) GetWithTx(tx, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{tx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).GetWithTx), varargs...)
}

// Insert mocks base method.
func (m *MockSubTaskRepository) Insert(ctx context.Context, subject *model.SubTask) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, subject)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockSubTaskRepositoryMockRecorder) Insert(ctx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSubTaskRepository)(nil).Insert), ctx, subject)
}

// InsertMulti mocks base method.
func (m *MockSubTaskRepository) InsertMulti(ctx context.Context, subjects []*model.SubTask) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMulti", ctx, subjects)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMulti indicates an expected call of InsertMulti.
func (mr *MockSubTaskRepositoryMockRecorder) InsertMulti(ctx, subjects any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMulti", reflect.TypeOf((*MockSubTaskRepository)(nil).InsertMulti), ctx, subjects)
}

// InsertMultiWithTx mocks base method.
func (m *MockSubTaskRepository) InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.SubTask) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMultiWithTx", ctx, tx, subjects)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMultiWithTx indicates an expected call of InsertMultiWithTx.
func (mr *MockSubTaskRepositoryMockRecorder) InsertMultiWithTx(ctx, tx, subjects any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMultiWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).InsertMultiWithTx), ctx, tx, subjects)
}

// InsertWithTx mocks base method.
func (m *MockSubTaskRepository) InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.SubTask) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertWithTx", ctx, tx, subject)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertWithTx indicates an expected call of InsertWithTx.
func (mr *MockSubTaskRepositoryMockRecorder) InsertWithTx(ctx, tx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).InsertWithTx), ctx, tx, subject)
}

// NewRepositoryByParent mocks base method.
func (m *MockSubTaskRepository) NewRepositoryByParent(doc *firestore.DocumentRef) model.SubTaskRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRepositoryByParent", doc)
	ret0, _ := ret[0].(model.SubTaskRepository)
	return ret0
}

// NewRepositoryByParent indicates an expected call of NewRepositoryByParent.
func (mr *MockSubTaskRepositoryMockRecorder) NewRepositoryByParent(doc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRepositoryByParent", reflect.TypeOf((*MockSubTaskRepository)(nil).NewRepositoryByParent), doc)
}

// RunInTransaction mocks base method.
func (m *MockSubTaskRepository) RunInTransaction() func(context.Context, func(context.Context, *firestore.Transaction) error, ...firestore.TransactionOption) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInTransaction")
	ret0, _ := ret[0].(func(context.Context, func(context.Context, *firestore.Transaction) error, ...firestore.TransactionOption) error)
	return ret0
}

// RunInTransaction indicates an expected call of RunInTransaction.
func (mr *MockSubTaskRepositoryMockRecorder) RunInTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInTransaction", reflect.TypeOf((*MockSubTaskRepository)(nil).RunInTransaction))
}

// Search mocks base method.
func (m *MockSubTaskRepository) Search(ctx context.Context, param *model.SubTaskSearchParam, q *firestore.Query) ([]*model.SubTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, param, q)
	ret0, _ := ret[0].([]*model.SubTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockSubTaskRepositoryMockRecorder) Search(ctx, param, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockSubTaskRepository)(nil).Search), ctx, param, q)
}

// SearchByParam mocks base method.
func (m *MockSubTaskRepository) SearchByParam(ctx context.Context, param *model.SubTaskSearchParam) ([]*model.SubTask, *model.PagingResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByParam", ctx, param)
	ret0, _ := ret[0].([]*model.SubTask)
	ret1, _ := ret[1].(*model.PagingResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchByParam indicates an expected call of SearchByParam.
func (mr *MockSubTaskRepositoryMockRecorder) SearchByParam(ctx, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByParam", reflect.TypeOf((*MockSubTaskRepository)(nil).SearchByParam), ctx, param)
}

// SearchByParamWithTx mocks base method.
func (m *MockSubTaskRepository) SearchByParamWithTx(tx *firestore.Transaction, param *model.SubTaskSearchParam) ([]*model.SubTask, *model.PagingResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByParamWithTx", tx, param)
	ret0, _ := ret[0].([]*model.SubTask)
	ret1, _ := ret[1].(*model.PagingResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchByParamWithTx indicates an expected call of SearchByParamWithTx.
func (mr *MockSubTaskRepositoryMockRecorder) SearchByParamWithTx(tx, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByParamWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).SearchByParamWithTx), tx, param)
}

// SearchWithTx mocks base method.
func (m *MockSubTaskRepository) SearchWithTx(tx *firestore.Transaction, param *model.SubTaskSearchParam, q *firestore.Query) ([]*model.SubTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithTx", tx, param, q)
	ret0, _ := ret[0].([]*model.SubTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithTx indicates an expected call of SearchWithTx.
func (mr *MockSubTaskRepositoryMockRecorder) SearchWithTx(tx, param, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).SearchWithTx), tx, param, q)
}

// SetParentDoc mocks base method.
func (m *MockSubTaskRepository) SetParentDoc(doc *firestore.DocumentRef) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetParentDoc", doc)
}

// SetParentDoc indicates an expected call of SetParentDoc.
func (mr *MockSubTaskRepositoryMockRecorder) SetParentDoc(doc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParentDoc", reflect.TypeOf((*MockSubTaskRepository)(nil).SetParentDoc), doc)
}

// StrictUpdate mocks base method.
func (m *MockSubTaskRepository) StrictUpdate(ctx context.Context, id string, param *model.SubTaskUpdateParam, opts ...firestore.Precondition) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) StrictUpdate(ctx, id, param any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, id, param}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrictUpdate", reflect.TypeOf((*MockSubTaskRepository)(nil).StrictUpdate), varargs...)
}

// StrictUpdateWithTx mocks base method.
func (m *MockSubTaskRepository) StrictUpdateWithTx(tx *firestore.Transaction, id string, param *model.SubTaskUpdateParam, opts ...firestore.Precondition) error {
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
func (mr *MockSubTaskRepositoryMockRecorder) StrictUpdateWithTx(tx, id, param any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{tx, id, param}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrictUpdateWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).StrictUpdateWithTx), varargs...)
}

// Update mocks base method.
func (m *MockSubTaskRepository) Update(ctx context.Context, subject *model.SubTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSubTaskRepositoryMockRecorder) Update(ctx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSubTaskRepository)(nil).Update), ctx, subject)
}

// UpdateMulti mocks base method.
func (m *MockSubTaskRepository) UpdateMulti(ctx context.Context, subjects []*model.SubTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMulti", ctx, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMulti indicates an expected call of UpdateMulti.
func (mr *MockSubTaskRepositoryMockRecorder) UpdateMulti(ctx, subjects any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMulti", reflect.TypeOf((*MockSubTaskRepository)(nil).UpdateMulti), ctx, subjects)
}

// UpdateMultiWithTx mocks base method.
func (m *MockSubTaskRepository) UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.SubTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMultiWithTx", ctx, tx, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMultiWithTx indicates an expected call of UpdateMultiWithTx.
func (mr *MockSubTaskRepositoryMockRecorder) UpdateMultiWithTx(ctx, tx, subjects any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMultiWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).UpdateMultiWithTx), ctx, tx, subjects)
}

// UpdateWithTx mocks base method.
func (m *MockSubTaskRepository) UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.SubTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithTx", ctx, tx, subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWithTx indicates an expected call of UpdateWithTx.
func (mr *MockSubTaskRepositoryMockRecorder) UpdateWithTx(ctx, tx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithTx", reflect.TypeOf((*MockSubTaskRepository)(nil).UpdateWithTx), ctx, tx, subject)
}

// MockSubTaskRepositoryMiddleware is a mock of SubTaskRepositoryMiddleware interface.
type MockSubTaskRepositoryMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockSubTaskRepositoryMiddlewareMockRecorder
}

// MockSubTaskRepositoryMiddlewareMockRecorder is the mock recorder for MockSubTaskRepositoryMiddleware.
type MockSubTaskRepositoryMiddlewareMockRecorder struct {
	mock *MockSubTaskRepositoryMiddleware
}

// NewMockSubTaskRepositoryMiddleware creates a new mock instance.
func NewMockSubTaskRepositoryMiddleware(ctrl *gomock.Controller) *MockSubTaskRepositoryMiddleware {
	mock := &MockSubTaskRepositoryMiddleware{ctrl: ctrl}
	mock.recorder = &MockSubTaskRepositoryMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubTaskRepositoryMiddleware) EXPECT() *MockSubTaskRepositoryMiddlewareMockRecorder {
	return m.recorder
}

// BeforeDelete mocks base method.
func (m *MockSubTaskRepositoryMiddleware) BeforeDelete(ctx context.Context, subject *model.SubTask, opts ...model.DeleteOption) (bool, error) {
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
func (mr *MockSubTaskRepositoryMiddlewareMockRecorder) BeforeDelete(ctx, subject any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelete", reflect.TypeOf((*MockSubTaskRepositoryMiddleware)(nil).BeforeDelete), varargs...)
}

// BeforeDeleteByID mocks base method.
func (m *MockSubTaskRepositoryMiddleware) BeforeDeleteByID(ctx context.Context, ids []string, opts ...model.DeleteOption) (bool, error) {
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
func (mr *MockSubTaskRepositoryMiddlewareMockRecorder) BeforeDeleteByID(ctx, ids any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDeleteByID", reflect.TypeOf((*MockSubTaskRepositoryMiddleware)(nil).BeforeDeleteByID), varargs...)
}

// BeforeInsert mocks base method.
func (m *MockSubTaskRepositoryMiddleware) BeforeInsert(ctx context.Context, subject *model.SubTask) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeInsert", ctx, subject)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeInsert indicates an expected call of BeforeInsert.
func (mr *MockSubTaskRepositoryMiddlewareMockRecorder) BeforeInsert(ctx, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeInsert", reflect.TypeOf((*MockSubTaskRepositoryMiddleware)(nil).BeforeInsert), ctx, subject)
}

// BeforeUpdate mocks base method.
func (m *MockSubTaskRepositoryMiddleware) BeforeUpdate(ctx context.Context, old, subject *model.SubTask) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeUpdate", ctx, old, subject)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeUpdate indicates an expected call of BeforeUpdate.
func (mr *MockSubTaskRepositoryMiddlewareMockRecorder) BeforeUpdate(ctx, old, subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeUpdate", reflect.TypeOf((*MockSubTaskRepositoryMiddleware)(nil).BeforeUpdate), ctx, old, subject)
}
