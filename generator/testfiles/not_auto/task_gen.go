// Code generated by volcago. DO NOT EDIT.
// generated version: devel
package model

import (
	"context"

	"cloud.google.com/go/firestore"
	"golang.org/x/xerrors"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate mockgen -source $GOFILE -destination mock/mock_task_gen/mock_task_gen.go

// TaskRepository - Repository of Task
type TaskRepository interface {
	// Single
	Get(ctx context.Context, identity string, opts ...GetOption) (*Task, error)
	GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...GetOption) (*Task, error)
	Insert(ctx context.Context, subject *Task) (_ string, err error)
	Update(ctx context.Context, subject *Task) (err error)
	StrictUpdate(ctx context.Context, id string, param *TaskUpdateParam, opts ...firestore.Precondition) error
	Delete(ctx context.Context, subject *Task, opts ...DeleteOption) (err error)
	DeleteByIdentity(ctx context.Context, identity string, opts ...DeleteOption) (err error)
	// Multiple
	GetMulti(ctx context.Context, identities []string, opts ...GetOption) ([]*Task, error)
	InsertMulti(ctx context.Context, subjects []*Task) (_ []string, er error)
	UpdateMulti(ctx context.Context, subjects []*Task) (er error)
	DeleteMulti(ctx context.Context, subjects []*Task, opts ...DeleteOption) (er error)
	DeleteMultiByIdentities(ctx context.Context, identities []string, opts ...DeleteOption) (er error)
	// Single(Transaction)
	GetWithTx(tx *firestore.Transaction, identity string, opts ...GetOption) (*Task, error)
	GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...GetOption) (*Task, error)
	InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *Task) (_ string, err error)
	UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *Task) (err error)
	StrictUpdateWithTx(tx *firestore.Transaction, id string, param *TaskUpdateParam, opts ...firestore.Precondition) error
	DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *Task, opts ...DeleteOption) (err error)
	DeleteByIdentityWithTx(ctx context.Context, tx *firestore.Transaction, identity string, opts ...DeleteOption) (err error)
	// Multiple(Transaction)
	GetMultiWithTx(tx *firestore.Transaction, identities []string, opts ...GetOption) ([]*Task, error)
	InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*Task) (_ []string, er error)
	UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*Task) (er error)
	DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*Task, opts ...DeleteOption) (er error)
	DeleteMultiByIdentitiesWithTx(ctx context.Context, tx *firestore.Transaction, identities []string, opts ...DeleteOption) (er error)
	// Search
	Search(ctx context.Context, param *TaskSearchParam, q *firestore.Query) ([]*Task, error)
	SearchWithTx(tx *firestore.Transaction, param *TaskSearchParam, q *firestore.Query) ([]*Task, error)
	SearchByParam(ctx context.Context, param *TaskSearchParam) ([]*Task, *PagingResult, error)
	SearchByParamWithTx(tx *firestore.Transaction, param *TaskSearchParam) ([]*Task, *PagingResult, error)
	// misc
	GetCollection() *firestore.CollectionRef
	GetCollectionName() string
	GetDocRef(identity string) *firestore.DocumentRef
	RunInTransaction() func(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error)
	// get by unique field
	GetByDesc(ctx context.Context, description string) (*Task, error)
	GetByDescWithTx(tx *firestore.Transaction, description string) (*Task, error)
}

// TaskRepositoryMiddleware - middleware of TaskRepository
type TaskRepositoryMiddleware interface {
	BeforeInsert(ctx context.Context, subject *Task) (bool, error)
	BeforeUpdate(ctx context.Context, old, subject *Task) (bool, error)
	BeforeDelete(ctx context.Context, subject *Task, opts ...DeleteOption) (bool, error)
	BeforeDeleteByIdentity(ctx context.Context, identities []string, opts ...DeleteOption) (bool, error)
}

type taskRepository struct {
	collectionName   string
	firestoreClient  *firestore.Client
	middleware       []TaskRepositoryMiddleware
	uniqueRepository *uniqueRepository
}

// NewTaskRepository - constructor
func NewTaskRepository(firestoreClient *firestore.Client, middleware ...TaskRepositoryMiddleware) TaskRepository {
	return &taskRepository{
		collectionName:   "Task",
		firestoreClient:  firestoreClient,
		middleware:       middleware,
		uniqueRepository: newUniqueRepository(firestoreClient, "Task"),
	}
}

func (repo *taskRepository) beforeInsert(ctx context.Context, subject *Task) (RollbackFunc, error) {
	repo.uniqueRepository.setMiddleware(ctx)
	rb, err := repo.uniqueRepository.CheckUnique(ctx, nil, subject)
	if err != nil {
		return nil, xerrors.Errorf("unique.middleware error: %w", err)
	}

	for _, m := range repo.middleware {
		c, err := m.BeforeInsert(ctx, subject)
		if err != nil {
			return nil, xerrors.Errorf("beforeInsert.middleware error(uniqueRB=%t): %w", rb(ctx) == nil, err)
		}
		if !c {
			continue
		}
	}

	return rb, nil
}

func (repo *taskRepository) beforeUpdate(ctx context.Context, old, subject *Task) (RollbackFunc, error) {
	if ctx.Value(transactionInProgressKey{}) != nil && old == nil {
		var err error
		doc := repo.GetDocRef(subject.Identity)
		old, err = repo.get(context.Background(), doc)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return nil, ErrNotFound
			}
			return nil, xerrors.Errorf("error in Get method: %w", err)
		}
	}
	repo.uniqueRepository.setMiddleware(ctx)
	rb, err := repo.uniqueRepository.CheckUnique(ctx, old, subject)
	if err != nil {
		return nil, xerrors.Errorf("unique.middleware error: %w", err)
	}

	for _, m := range repo.middleware {
		c, err := m.BeforeUpdate(ctx, old, subject)
		if err != nil {
			return nil, xerrors.Errorf("beforeUpdate.middleware error: %w", err)
		}
		if !c {
			continue
		}
	}

	return rb, nil
}

func (repo *taskRepository) beforeDelete(ctx context.Context, subject *Task, opts ...DeleteOption) (RollbackFunc, error) {
	repo.uniqueRepository.setMiddleware(ctx)
	rb, err := repo.uniqueRepository.DeleteUnique(ctx, subject)
	if err != nil {
		return nil, xerrors.Errorf("unique.middleware error: %w", err)
	}

	for _, m := range repo.middleware {
		c, err := m.BeforeDelete(ctx, subject, opts...)
		if err != nil {
			return nil, xerrors.Errorf("beforeDelete.middleware error: %w", err)
		}
		if !c {
			continue
		}
	}

	return rb, nil
}

// GetCollection - *firestore.CollectionRef getter
func (repo *taskRepository) GetCollection() *firestore.CollectionRef {
	return repo.firestoreClient.Collection(repo.collectionName)
}

// GetCollectionName - CollectionName getter
func (repo *taskRepository) GetCollectionName() string {
	return repo.collectionName
}

// GetDocRef - *firestore.DocumentRef getter
func (repo *taskRepository) GetDocRef(identity string) *firestore.DocumentRef {
	return repo.GetCollection().Doc(identity)
}

// RunInTransaction - (*firestore.Client).RunTransaction getter
func (repo *taskRepository) RunInTransaction() func(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error) {
	return repo.firestoreClient.RunTransaction
}

// TaskSearchParam - params for search
type TaskSearchParam struct {
	Desc         *QueryChainer
	Desc2        *QueryChainer
	Created      *QueryChainer
	ReservedDate *QueryChainer
	Done         *QueryChainer
	Done2        *QueryChainer
	Count        *QueryChainer
	Count64      *QueryChainer
	NameList     *QueryChainer
	Proportion   *QueryChainer
	Geo          *QueryChainer
	Sub          *QueryChainer
	Flag         *QueryChainer

	CursorKey   string
	CursorLimit int
}

// TaskUpdateParam - params for strict updates
type TaskUpdateParam struct {
	Desc2        interface{}
	Created      interface{}
	ReservedDate interface{}
	Done         interface{}
	Done2        interface{}
	Count        interface{}
	Count64      interface{}
	NameList     interface{}
	Proportion   interface{}
	Geo          interface{}
	Sub          interface{}
	Flag         interface{}
}

// Search - search documents
// The third argument is firestore.Query, basically you can pass nil
func (repo *taskRepository) Search(ctx context.Context, param *TaskSearchParam, q *firestore.Query) ([]*Task, error) {
	return repo.search(ctx, param, q)
}

// SearchByParam - search documents by search param
func (repo *taskRepository) SearchByParam(ctx context.Context, param *TaskSearchParam) ([]*Task, *PagingResult, error) {
	return repo.searchByParam(ctx, param)
}

// Get - get `Task` by `Task.Identity`
func (repo *taskRepository) Get(ctx context.Context, identity string, opts ...GetOption) (*Task, error) {
	doc := repo.GetDocRef(identity)
	return repo.get(ctx, doc, opts...)
}

// GetWithDoc - get `Task` by *firestore.DocumentRef
func (repo *taskRepository) GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...GetOption) (*Task, error) {
	return repo.get(ctx, doc, opts...)
}

// Insert - insert of `Task`
func (repo *taskRepository) Insert(ctx context.Context, subject *Task) (_ string, err error) {
	rb, err := repo.beforeInsert(ctx, subject)
	if err != nil {
		return "", xerrors.Errorf("before insert error: %w", err)
	}
	defer func() {
		if err != nil {
			if er := rb(ctx); er != nil {
				err = xerrors.Errorf("unique check error %+v, original error: %w", er, err)
			}
		}
	}()

	return repo.insert(ctx, subject)
}

// Update - update of `Task`
func (repo *taskRepository) Update(ctx context.Context, subject *Task) (err error) {
	doc := repo.GetDocRef(subject.Identity)

	old, err := repo.get(ctx, doc)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return ErrNotFound
		}
		return xerrors.Errorf("error in Get method: %w", err)
	}

	rb, err := repo.beforeUpdate(ctx, old, subject)
	if err != nil {
		return xerrors.Errorf("before update error: %w", err)
	}
	defer func() {
		if err != nil {
			if er := rb(ctx); er != nil {
				err = xerrors.Errorf("unique check error %+v, original error: %w", er, err)
			}
		}
	}()

	return repo.update(ctx, subject)
}

// StrictUpdate - strict update of `Task`
func (repo *taskRepository) StrictUpdate(ctx context.Context, id string, param *TaskUpdateParam, opts ...firestore.Precondition) error {
	return repo.strictUpdate(ctx, id, param, opts...)
}

// Delete - delete of `Task`
func (repo *taskRepository) Delete(ctx context.Context, subject *Task, opts ...DeleteOption) (err error) {
	rb, err := repo.beforeDelete(ctx, subject, opts...)
	if err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}
	defer func() {
		if err != nil {
			if er := rb(ctx); er != nil {
				err = xerrors.Errorf("unique delete error %+v, original error: %w", er, err)
			}
		}
	}()

	return repo.deleteByIdentity(ctx, subject.Identity)
}

// DeleteByIdentity - delete `Task` by `Task.Identity`
func (repo *taskRepository) DeleteByIdentity(ctx context.Context, identity string, opts ...DeleteOption) (err error) {
	subject, err := repo.Get(ctx, identity)
	if err != nil {
		return xerrors.Errorf("error in Get method: %w", err)
	}

	return repo.Delete(ctx, subject, opts...)
}

// GetMulti - get `Task` in bulk by array of `Task.Identity`
func (repo *taskRepository) GetMulti(ctx context.Context, identities []string, opts ...GetOption) ([]*Task, error) {
	return repo.getMulti(ctx, identities, opts...)
}

// InsertMulti - bulk insert of `Task`
func (repo *taskRepository) InsertMulti(ctx context.Context, subjects []*Task) (_ []string, er error) {
	var rbs []RollbackFunc
	defer func() {
		if er == nil {
			return
		}
		if len(rbs) == 0 {
			return
		}
		errs := make([]error, 0)
		for _, rb := range rbs {
			if err := rb(ctx); err != nil {
				errs = append(errs, err)
			}
		}
		er = xerrors.Errorf("unique check error %+v, original error: %w", errs, er)
	}()

	identities := make([]string, 0, len(subjects))
	batches := make([]*firestore.WriteBatch, 0)
	batch := repo.firestoreClient.Batch()
	collect := repo.GetCollection()

	for i, subject := range subjects {
		var ref *firestore.DocumentRef
		if subject.Identity == "" {
			ref = collect.NewDoc()
			subject.Identity = ref.ID
		} else {
			ref = collect.Doc(subject.Identity)
			if s, err := ref.Get(ctx); err == nil {
				return nil, xerrors.Errorf("already exists [%v]: %#v", subject.Identity, s)
			}
		}

		rb, err := repo.beforeInsert(ctx, subject)
		if err != nil {
			return nil, xerrors.Errorf("before insert error(%d) [%v]: %w", i, subject.Identity, err)
		}
		rbs = append(rbs, rb)

		batch.Set(ref, subject)
		identities = append(identities, ref.ID)
		i++
		if (i%500) == 0 && len(subjects) != i {
			batches = append(batches, batch)
			batch = repo.firestoreClient.Batch()
		}
	}
	batches = append(batches, batch)

	for _, b := range batches {
		if _, err := b.Commit(ctx); err != nil {
			return nil, xerrors.Errorf("error in Commit method: %w", err)
		}
	}

	return identities, nil
}

// UpdateMulti - bulk update of `Task`
func (repo *taskRepository) UpdateMulti(ctx context.Context, subjects []*Task) (er error) {
	var rbs []RollbackFunc
	defer func() {
		if er == nil {
			return
		}
		if len(rbs) == 0 {
			return
		}
		errs := make([]error, 0)
		for _, rb := range rbs {
			if err := rb(ctx); err != nil {
				errs = append(errs, err)
			}
		}
		er = xerrors.Errorf("unique check error %+v, original error: %w", errs, er)
	}()

	batches := make([]*firestore.WriteBatch, 0)
	batch := repo.firestoreClient.Batch()
	collect := repo.GetCollection()

	for i, subject := range subjects {
		ref := collect.Doc(subject.Identity)
		snapShot, err := ref.Get(ctx)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found [%v]: %w", subject.Identity, err)
			}
			return xerrors.Errorf("error in Get method [%v]: %w", subject.Identity, err)
		}

		old := new(Task)
		if err = snapShot.DataTo(&old); err != nil {
			return xerrors.Errorf("error in DataTo method: %w", err)
		}

		rb, err := repo.beforeUpdate(ctx, old, subject)
		if err != nil {
			return xerrors.Errorf("before update error(%d) [%v]: %w", i, subject.Identity, err)
		}
		rbs = append(rbs, rb)

		batch.Set(ref, subject)
		i++
		if (i%500) == 0 && len(subjects) != i {
			batches = append(batches, batch)
			batch = repo.firestoreClient.Batch()
		}
	}
	batches = append(batches, batch)

	for _, b := range batches {
		if _, err := b.Commit(ctx); err != nil {
			return xerrors.Errorf("error in Commit method: %w", err)
		}
	}

	return nil
}

// DeleteMulti - bulk delete of `Task`
func (repo *taskRepository) DeleteMulti(ctx context.Context, subjects []*Task, opts ...DeleteOption) (er error) {
	var rbs []RollbackFunc
	defer func() {
		if er == nil {
			return
		}
		if len(rbs) == 0 {
			return
		}
		errs := make([]error, 0)
		for _, rb := range rbs {
			if err := rb(ctx); err != nil {
				errs = append(errs, err)
			}
		}
		er = xerrors.Errorf("unique delete error %+v, original error: %w", errs, er)
	}()

	batches := make([]*firestore.WriteBatch, 0)
	batch := repo.firestoreClient.Batch()
	collect := repo.GetCollection()

	for i, subject := range subjects {
		ref := collect.Doc(subject.Identity)
		if _, err := ref.Get(ctx); err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found [%v]: %w", subject.Identity, err)
			}
			return xerrors.Errorf("error in Get method [%v]: %w", subject.Identity, err)
		}

		rb, err := repo.beforeDelete(ctx, subject, opts...)
		if err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subject.Identity, err)
		}
		rbs = append(rbs, rb)

		batch.Delete(ref)

		i++
		if (i%500) == 0 && len(subjects) != i {
			batches = append(batches, batch)
			batch = repo.firestoreClient.Batch()
		}
	}
	batches = append(batches, batch)

	for _, b := range batches {
		if _, err := b.Commit(ctx); err != nil {
			return xerrors.Errorf("error in Commit method: %w", err)
		}
	}

	return nil
}

// DeleteMultiByIdentities - delete `Task` in bulk by array of `Task.Identity`
func (repo *taskRepository) DeleteMultiByIdentities(ctx context.Context, identities []string, opts ...DeleteOption) (er error) {
	subjects := make([]*Task, len(identities))

	opt := GetOption{}
	if len(opts) > 0 {
		opt.IncludeSoftDeleted = opts[0].Mode == DeleteModeHard
	}
	for i, identity := range identities {
		subject, err := repo.Get(ctx, identity, opt)
		if err != nil {
			return xerrors.Errorf("error in Get method: %w", err)
		}
		subjects[i] = subject
	}

	return repo.DeleteMulti(ctx, subjects, opts...)
}

// SearchWithTx - search documents in transaction
func (repo *taskRepository) SearchWithTx(tx *firestore.Transaction, param *TaskSearchParam, q *firestore.Query) ([]*Task, error) {
	return repo.search(tx, param, q)
}

// SearchByParamWithTx - search documents by search param in transaction
func (repo *taskRepository) SearchByParamWithTx(tx *firestore.Transaction, param *TaskSearchParam) ([]*Task, *PagingResult, error) {
	return repo.searchByParam(tx, param)
}

// GetWithTx - get `Task` by `Task.Identity` in transaction
func (repo *taskRepository) GetWithTx(tx *firestore.Transaction, identity string, opts ...GetOption) (*Task, error) {
	doc := repo.GetDocRef(identity)
	return repo.get(tx, doc, opts...)
}

// GetWithDocWithTx - get `Task` by *firestore.DocumentRef in transaction
func (repo *taskRepository) GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...GetOption) (*Task, error) {
	return repo.get(tx, doc, opts...)
}

// InsertWithTx - insert of `Task` in transaction
func (repo *taskRepository) InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *Task) (_ string, err error) {
	rb, err := repo.beforeInsert(context.WithValue(ctx, transactionInProgressKey{}, 1), subject)
	if err != nil {
		return "", xerrors.Errorf("before insert error: %w", err)
	}
	defer func() {
		if err != nil {
			if er := rb(ctx); er != nil {
				err = xerrors.Errorf("unique check error %+v, original error: %w", er, err)
			}
		}
	}()

	return repo.insert(tx, subject)
}

// UpdateWithTx - update of `Task` in transaction
func (repo *taskRepository) UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *Task) (err error) {
	rb, err := repo.beforeUpdate(context.WithValue(ctx, transactionInProgressKey{}, 1), nil, subject)
	if err != nil {
		return xerrors.Errorf("before update error: %w", err)
	}
	defer func() {
		if err != nil {
			if er := rb(ctx); er != nil {
				err = xerrors.Errorf("unique check error %+v, original error: %w", er, err)
			}
		}
	}()

	return repo.update(tx, subject)
}

// StrictUpdateWithTx - strict update of `Task` in transaction
func (repo *taskRepository) StrictUpdateWithTx(tx *firestore.Transaction, id string, param *TaskUpdateParam, opts ...firestore.Precondition) error {
	return repo.strictUpdate(tx, id, param, opts...)
}

// DeleteWithTx - delete of `Task` in transaction
func (repo *taskRepository) DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *Task, opts ...DeleteOption) (err error) {
	rb, err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, 1), subject, opts...)
	if err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}
	defer func() {
		if err != nil {
			if er := rb(ctx); er != nil {
				err = xerrors.Errorf("unique check error %+v, original error: %w", er, err)
			}
		}
	}()

	return repo.deleteByIdentity(tx, subject.Identity)
}

// DeleteByIdentityWithTx - delete `Task` by `Task.Identity` in transaction
func (repo *taskRepository) DeleteByIdentityWithTx(ctx context.Context, tx *firestore.Transaction, identity string, opts ...DeleteOption) (err error) {
	subject, err := repo.Get(context.Background(), identity)
	if err != nil {
		return xerrors.Errorf("error in Get method: %w", err)
	}

	rb, err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, 1), subject, opts...)
	if err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}
	defer func() {
		if err != nil {
			if er := rb(ctx); er != nil {
				err = xerrors.Errorf("unique delete error %+v, original error: %w", er, err)
			}
		}
	}()

	return repo.deleteByIdentity(tx, identity)
}

// GetMultiWithTx - get `Task` in bulk by array of `Task.Identity` in transaction
func (repo *taskRepository) GetMultiWithTx(tx *firestore.Transaction, identities []string, opts ...GetOption) ([]*Task, error) {
	return repo.getMulti(tx, identities, opts...)
}

// InsertMultiWithTx - bulk insert of `Task` in transaction
func (repo *taskRepository) InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*Task) (_ []string, er error) {
	ctx = context.WithValue(ctx, transactionInProgressKey{}, 1)
	var rbs []RollbackFunc
	defer func() {
		if er == nil {
			return
		}
		if len(rbs) == 0 {
			return
		}
		errs := make([]error, 0)
		for _, rb := range rbs {
			if err := rb(ctx); err != nil {
				errs = append(errs, err)
			}
		}
		er = xerrors.Errorf("unique check error %+v, original error: %w", errs, er)
	}()

	for i := range subjects {
		if _, err := tx.Get(new(firestore.DocumentRef)); err == nil {
			return nil, xerrors.Errorf("already exists(%d) [%v]: %w", i, subjects[i].Identity, err)
		}
	}

	identities := make([]string, len(subjects))

	for i := range subjects {
		rb, err := repo.beforeInsert(ctx, subjects[i])
		if err != nil {
			return nil, xerrors.Errorf("before insert error(%d) [%v]: %w", i, subjects[i].Identity, err)
		}
		rbs = append(rbs, rb)

		identity, err := repo.insert(tx, subjects[i])
		if err != nil {
			return nil, xerrors.Errorf("error in insert method(%d) [%v]: %w", i, subjects[i].Identity, err)
		}
		identities[i] = identity
	}

	return identities, nil
}

// UpdateMultiWithTx - bulk update of `Task` in transaction
func (repo *taskRepository) UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*Task) (er error) {
	ctx = context.WithValue(ctx, transactionInProgressKey{}, 1)
	var rbs []RollbackFunc
	defer func() {
		if er == nil {
			return
		}
		if len(rbs) == 0 {
			return
		}
		errs := make([]error, 0)
		for _, rb := range rbs {
			if err := rb(ctx); err != nil {
				errs = append(errs, err)
			}
		}
		er = xerrors.Errorf("unique check error %+v, original error: %w", errs, er)
	}()

	ctx = context.WithValue(ctx, transactionInProgressKey{}, 1)
	for i := range subjects {
		rb, err := repo.beforeUpdate(ctx, nil, subjects[i])
		if err != nil {
			return xerrors.Errorf("before update error(%d) [%v]: %w", i, subjects[i].Identity, err)
		}
		rbs = append(rbs, rb)
	}

	for i := range subjects {
		if err := repo.update(tx, subjects[i]); err != nil {
			return xerrors.Errorf("error in update method(%d) [%v]: %w", i, subjects[i].Identity, err)
		}
	}

	return nil
}

// DeleteMultiWithTx - bulk delete of `Task` in transaction
func (repo *taskRepository) DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*Task, opts ...DeleteOption) (er error) {
	var rbs []RollbackFunc
	defer func() {
		if er == nil {
			return
		}
		if len(rbs) == 0 {
			return
		}
		errs := make([]error, 0)
		for _, rb := range rbs {
			if err := rb(ctx); err != nil {
				errs = append(errs, err)
			}
		}
		er = xerrors.Errorf("unique delete error %+v, original error: %w", errs, er)
	}()

	var isHardDeleteMode bool
	if len(opts) > 0 {
		isHardDeleteMode = opts[0].Mode == DeleteModeHard
	}
	opt := GetOption{
		IncludeSoftDeleted: isHardDeleteMode,
	}
	for i := range subjects {
		dr := repo.GetDocRef(subjects[i].Identity)
		if _, err := repo.get(context.Background(), dr, opt); err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found(%d) [%v]", i, subjects[i].Identity)
			}
			return xerrors.Errorf("error in get method(%d) [%v]: %w", i, subjects[i].Identity, err)
		}

		rb, err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, 1), subjects[i], opts...)
		if err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subjects[i].Identity, err)
		}
		rbs = append(rbs, rb)
	}

	for i := range subjects {
		if err := repo.deleteByIdentity(tx, subjects[i].Identity); err != nil {
			return xerrors.Errorf("error in delete method(%d) [%v]: %w", i, subjects[i].Identity, err)
		}
	}

	return nil
}

// DeleteMultiByIdentityWithTx - delete `Task` in bulk by array of `Task.Identity` in transaction
func (repo *taskRepository) DeleteMultiByIdentitiesWithTx(ctx context.Context, tx *firestore.Transaction, identities []string, opts ...DeleteOption) (er error) {
	var rbs []RollbackFunc
	defer func() {
		if er == nil {
			return
		}
		if len(rbs) == 0 {
			return
		}
		errs := make([]error, 0)
		for _, rb := range rbs {
			if err := rb(ctx); err != nil {
				errs = append(errs, err)
			}
		}
		er = xerrors.Errorf("unique delete error %+v, original error: %w", errs, er)
	}()

	for i := range identities {
		dr := repo.GetDocRef(identities[i])
		subject, err := repo.get(context.Background(), dr)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found(%d) [%v]", i, identities[i])
			}
			return xerrors.Errorf("error in get method(%d) [%v]: %w", i, identities[i], err)
		}

		rb, err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, 1), subject, opts...)
		if err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subject.Identity, err)
		}
		rbs = append(rbs, rb)
	}

	for i := range identities {
		if err := repo.deleteByIdentity(tx, identities[i]); err != nil {
			return xerrors.Errorf("error in delete method(%d) [%v]: %w", i, identities[i], err)
		}
	}

	return nil
}

func (repo *taskRepository) get(v interface{}, doc *firestore.DocumentRef, _ ...GetOption) (*Task, error) {
	var (
		snapShot *firestore.DocumentSnapshot
		err      error
	)

	switch x := v.(type) {
	case *firestore.Transaction:
		snapShot, err = x.Get(doc)
	case context.Context:
		snapShot, err = doc.Get(x)
	default:
		return nil, xerrors.Errorf("invalid type: %v", x)
	}

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, ErrNotFound
		}
		return nil, xerrors.Errorf("error in Get method: %w", err)
	}

	subject := new(Task)
	if err := snapShot.DataTo(&subject); err != nil {
		return nil, xerrors.Errorf("error in DataTo method: %w", err)
	}

	subject.Identity = snapShot.Ref.ID

	return subject, nil
}

func (repo *taskRepository) getMulti(v interface{}, identities []string, _ ...GetOption) ([]*Task, error) {
	var (
		snapShots []*firestore.DocumentSnapshot
		err       error
		collect   = repo.GetCollection()
		drs       = make([]*firestore.DocumentRef, len(identities))
	)

	for i, identity := range identities {
		ref := collect.Doc(identity)
		drs[i] = ref
	}

	switch x := v.(type) {
	case *firestore.Transaction:
		snapShots, err = x.GetAll(drs)
	case context.Context:
		snapShots, err = repo.firestoreClient.GetAll(x, drs)
	default:
		return nil, xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		return nil, xerrors.Errorf("error in GetAll method: %w", err)
	}

	subjects := make([]*Task, 0, len(identities))
	mErr := NewMultiErrors()
	for i, snapShot := range snapShots {
		if !snapShot.Exists() {
			mErr = append(mErr, NewMultiError(i, ErrNotFound))
			continue
		}

		subject := new(Task)
		if err = snapShot.DataTo(&subject); err != nil {
			return nil, xerrors.Errorf("error in DataTo method: %w", err)
		}

		subject.Identity = snapShot.Ref.ID
		subjects = append(subjects, subject)
	}

	if len(mErr) == 0 {
		return subjects, nil
	}

	return subjects, mErr
}

func (repo *taskRepository) insert(v interface{}, subject *Task) (string, error) {
	var (
		dr  = repo.GetDocRef(subject.Identity)
		err error
	)

	switch x := v.(type) {
	case *firestore.Transaction:
		err = x.Create(dr, subject)
	case context.Context:
		_, err = dr.Create(x, subject)
	default:
		return "", xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return "", xerrors.Errorf("error in Create method: err=%+v: %w", err, ErrAlreadyExists)
		}
		return "", xerrors.Errorf("error in Create method: %w", err)
	}

	subject.Identity = dr.ID

	return dr.ID, nil
}

func (repo *taskRepository) update(v interface{}, subject *Task) error {
	var (
		dr  = repo.GetDocRef(subject.Identity)
		err error
	)

	switch x := v.(type) {
	case *firestore.Transaction:
		err = x.Set(dr, subject)
	case context.Context:
		_, err = dr.Set(x, subject)
	default:
		return xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		return xerrors.Errorf("error in Set method: %w", err)
	}

	return nil
}

func (repo *taskRepository) strictUpdate(v interface{}, id string, param *TaskUpdateParam, opts ...firestore.Precondition) error {
	var (
		dr  = repo.GetDocRef(id)
		err error
	)

	updates := updater(Task{}, param)

	switch x := v.(type) {
	case *firestore.Transaction:
		err = x.Update(dr, updates, opts...)
	case context.Context:
		_, err = dr.Update(x, updates, opts...)
	default:
		return xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		return xerrors.Errorf("error in Update method: %w", err)
	}

	return nil
}

func (repo *taskRepository) deleteByIdentity(v interface{}, identity string) error {
	dr := repo.GetDocRef(identity)
	var err error

	switch x := v.(type) {
	case *firestore.Transaction:
		err = x.Delete(dr, firestore.Exists)
	case context.Context:
		_, err = dr.Delete(x, firestore.Exists)
	default:
		return xerrors.Errorf("invalid type: %v", v)
	}

	if err != nil {
		return xerrors.Errorf("error in Delete method: %w", err)
	}

	return nil
}

func (repo *taskRepository) runQuery(v interface{}, query firestore.Query) ([]*Task, error) {
	var iter *firestore.DocumentIterator

	switch x := v.(type) {
	case *firestore.Transaction:
		iter = x.Documents(query)
	case context.Context:
		iter = query.Documents(x)
	default:
		return nil, xerrors.Errorf("invalid type: %v", v)
	}

	defer iter.Stop()

	subjects := make([]*Task, 0)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, xerrors.Errorf("error in Next method: %w", err)
		}

		subject := new(Task)

		if err := doc.DataTo(&subject); err != nil {
			return nil, xerrors.Errorf("error in DataTo method: %w", err)
		}

		subject.Identity = doc.Ref.ID
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

// BUG(54m): there may be potential bugs
func (repo *taskRepository) searchByParam(v interface{}, param *TaskSearchParam) ([]*Task, *PagingResult, error) {
	query := repo.GetCollection().Query
	if param.Desc != nil {
		for _, chain := range param.Desc.QueryGroup {
			query = query.Where("description", chain.Operator, chain.Value)
		}
		if direction := param.Desc.OrderByDirection; direction > 0 {
			query = query.OrderBy("description", direction)
			query = param.Desc.BuildCursorQuery(query)
		}
	}
	if param.Desc2 != nil {
		for _, chain := range param.Desc2.QueryGroup {
			query = query.Where("desc2", chain.Operator, chain.Value)
		}
		if direction := param.Desc2.OrderByDirection; direction > 0 {
			query = query.OrderBy("desc2", direction)
			query = param.Desc2.BuildCursorQuery(query)
		}
	}
	if param.Created != nil {
		for _, chain := range param.Created.QueryGroup {
			query = query.Where("created", chain.Operator, chain.Value)
		}
		if direction := param.Created.OrderByDirection; direction > 0 {
			query = query.OrderBy("created", direction)
			query = param.Created.BuildCursorQuery(query)
		}
	}
	if param.ReservedDate != nil {
		for _, chain := range param.ReservedDate.QueryGroup {
			query = query.Where("reservedDate", chain.Operator, chain.Value)
		}
		if direction := param.ReservedDate.OrderByDirection; direction > 0 {
			query = query.OrderBy("reservedDate", direction)
			query = param.ReservedDate.BuildCursorQuery(query)
		}
	}
	if param.Done != nil {
		for _, chain := range param.Done.QueryGroup {
			query = query.Where("done", chain.Operator, chain.Value)
		}
		if direction := param.Done.OrderByDirection; direction > 0 {
			query = query.OrderBy("done", direction)
			query = param.Done.BuildCursorQuery(query)
		}
	}
	if param.Done2 != nil {
		for _, chain := range param.Done2.QueryGroup {
			query = query.Where("done2", chain.Operator, chain.Value)
		}
		if direction := param.Done2.OrderByDirection; direction > 0 {
			query = query.OrderBy("done2", direction)
			query = param.Done2.BuildCursorQuery(query)
		}
	}
	if param.Count != nil {
		for _, chain := range param.Count.QueryGroup {
			query = query.Where("count", chain.Operator, chain.Value)
		}
		if direction := param.Count.OrderByDirection; direction > 0 {
			query = query.OrderBy("count", direction)
			query = param.Count.BuildCursorQuery(query)
		}
	}
	if param.Count64 != nil {
		for _, chain := range param.Count64.QueryGroup {
			query = query.Where("count64", chain.Operator, chain.Value)
		}
		if direction := param.Count64.OrderByDirection; direction > 0 {
			query = query.OrderBy("count64", direction)
			query = param.Count64.BuildCursorQuery(query)
		}
	}
	if param.NameList != nil {
		for _, chain := range param.NameList.QueryGroup {
			query = query.Where("nameList", chain.Operator, chain.Value)
		}
	}
	if param.Proportion != nil {
		for _, chain := range param.Proportion.QueryGroup {
			query = query.Where("proportion", chain.Operator, chain.Value)
		}
		if direction := param.Proportion.OrderByDirection; direction > 0 {
			query = query.OrderBy("proportion", direction)
			query = param.Proportion.BuildCursorQuery(query)
		}
	}
	if param.Geo != nil {
		for _, chain := range param.Geo.QueryGroup {
			query = query.Where("geo", chain.Operator, chain.Value)
		}
		if direction := param.Geo.OrderByDirection; direction > 0 {
			query = query.OrderBy("geo", direction)
			query = param.Geo.BuildCursorQuery(query)
		}
	}
	if param.Sub != nil {
		for _, chain := range param.Sub.QueryGroup {
			query = query.Where("sub", chain.Operator, chain.Value)
		}
		if direction := param.Sub.OrderByDirection; direction > 0 {
			query = query.OrderBy("sub", direction)
			query = param.Sub.BuildCursorQuery(query)
		}
	}
	if param.Flag != nil {
		for _, chain := range param.Flag.QueryGroup {
			query = query.Where("flag", chain.Operator, chain.Value)
		}
		if direction := param.Flag.OrderByDirection; direction > 0 {
			query = query.OrderBy("flag", direction)
			query = param.Flag.BuildCursorQuery(query)
		}
	}

	if l := param.CursorLimit; l > 0 {
		query = query.Limit(l)
	}

	cursorKey := param.CursorKey
	if cursorKey == "" {
		if l := param.CursorLimit; l > 0 {
			query = query.Limit(l)
		}

		subjects, err := repo.runQuery(v, query)
		if err != nil {
			return nil, nil, xerrors.Errorf("error in runQuery method: %w", err)
		}

		return subjects, nil, nil
	}

	limit := param.CursorLimit + 1

	dr := repo.GetDocRef(cursorKey)
	query = query.StartAt(dr).Limit(limit)

	subjects, err := repo.runQuery(v, query)
	if err != nil {
		return nil, nil, xerrors.Errorf("error in runQuery method: %w", err)
	}

	pagingResult := &PagingResult{
		Length: len(subjects),
	}
	if limit == pagingResult.Length {
		pagingResult.NextCursorKey = subjects[pagingResult.Length-1].Identity
	}

	return subjects, pagingResult, nil
}

func (repo *taskRepository) search(v interface{}, param *TaskSearchParam, q *firestore.Query) ([]*Task, error) {
	if (param == nil && q == nil) || (param != nil && q != nil) {
		return nil, xerrors.New("either one should be nil")
	}

	query := func() firestore.Query {
		if q != nil {
			return *q
		}
		return repo.GetCollection().Query
	}()

	if q == nil {
		subjects, _, err := repo.searchByParam(v, param)
		if err != nil {
			return nil, xerrors.Errorf("error in searchByParam method: %w", err)
		}

		return subjects, nil
	}

	return repo.runQuery(v, query)
}

// GetByDesc - get by Desc
func (repo *taskRepository) GetByDesc(ctx context.Context, description string) (*Task, error) {
	return repo.getByXXX(ctx, "description", description)
}

// GetByDescWithTx - get by Desc in transaction
func (repo *taskRepository) GetByDescWithTx(tx *firestore.Transaction, description string) (*Task, error) {
	return repo.getByXXX(tx, "description", description)
}

func (repo *taskRepository) getByXXX(v interface{}, field, value string) (*Task, error) {
	query := repo.GetCollection().Query.Where(field, OpTypeEqual, value).Limit(1)
	results, err := repo.runQuery(v, query)
	if err != nil {
		return nil, xerrors.Errorf("failed to run query: %w", err)
	} else if len(results) == 0 {
		return nil, ErrNotFound
	}
	return results[0], nil
}
