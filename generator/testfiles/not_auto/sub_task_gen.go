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

//go:generate mockgen -source $GOFILE -destination mock/mock_sub_task_gen/mock_sub_task_gen.go

// SubTaskRepository - Repository of SubTask
type SubTaskRepository interface {
	// Single
	Get(ctx context.Context, id string, opts ...GetOption) (*SubTask, error)
	GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...GetOption) (*SubTask, error)
	Insert(ctx context.Context, subject *SubTask) (_ string, err error)
	Update(ctx context.Context, subject *SubTask) (err error)
	StrictUpdate(ctx context.Context, id string, param *SubTaskUpdateParam, opts ...firestore.Precondition) error
	Delete(ctx context.Context, subject *SubTask, opts ...DeleteOption) (err error)
	DeleteByID(ctx context.Context, id string, opts ...DeleteOption) (err error)
	// Multiple
	GetMulti(ctx context.Context, ids []string, opts ...GetOption) ([]*SubTask, error)
	InsertMulti(ctx context.Context, subjects []*SubTask) (_ []string, er error)
	UpdateMulti(ctx context.Context, subjects []*SubTask) (er error)
	DeleteMulti(ctx context.Context, subjects []*SubTask, opts ...DeleteOption) (er error)
	DeleteMultiByIDs(ctx context.Context, ids []string, opts ...DeleteOption) (er error)
	// Single(Transaction)
	GetWithTx(tx *firestore.Transaction, id string, opts ...GetOption) (*SubTask, error)
	GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...GetOption) (*SubTask, error)
	InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *SubTask) (_ string, err error)
	UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *SubTask) (err error)
	StrictUpdateWithTx(tx *firestore.Transaction, id string, param *SubTaskUpdateParam, opts ...firestore.Precondition) error
	DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *SubTask, opts ...DeleteOption) (err error)
	DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...DeleteOption) (err error)
	// Multiple(Transaction)
	GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...GetOption) ([]*SubTask, error)
	InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*SubTask) (_ []string, er error)
	UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*SubTask) (er error)
	DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*SubTask, opts ...DeleteOption) (er error)
	DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...DeleteOption) (er error)
	// Search
	Search(ctx context.Context, param *SubTaskSearchParam, q *firestore.Query) ([]*SubTask, error)
	SearchWithTx(tx *firestore.Transaction, param *SubTaskSearchParam, q *firestore.Query) ([]*SubTask, error)
	SearchByParam(ctx context.Context, param *SubTaskSearchParam) ([]*SubTask, *PagingResult, error)
	SearchByParamWithTx(tx *firestore.Transaction, param *SubTaskSearchParam) ([]*SubTask, *PagingResult, error)
	// misc
	GetCollection() *firestore.CollectionRef
	GetCollectionName() string
	GetDocRef(id string) *firestore.DocumentRef
	RunInTransaction() func(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error)
	SetParentDoc(doc *firestore.DocumentRef)
	NewRepositoryByParent(doc *firestore.DocumentRef) SubTaskRepository
	Free()
}

// SubTaskRepositoryMiddleware - middleware of SubTaskRepository
type SubTaskRepositoryMiddleware interface {
	BeforeInsert(ctx context.Context, subject *SubTask) (bool, error)
	BeforeUpdate(ctx context.Context, old, subject *SubTask) (bool, error)
	BeforeDelete(ctx context.Context, subject *SubTask, opts ...DeleteOption) (bool, error)
	BeforeDeleteByID(ctx context.Context, ids []string, opts ...DeleteOption) (bool, error)
}

type subTaskRepository struct {
	collectionName   string
	firestoreClient  *firestore.Client
	parentDocument   *firestore.DocumentRef
	collectionGroup  *firestore.CollectionGroupRef
	middleware       []SubTaskRepositoryMiddleware
	uniqueRepository *uniqueRepository
}

// NewSubTaskRepository - constructor
func NewSubTaskRepository(firestoreClient *firestore.Client, parentDocument *firestore.DocumentRef, middleware ...SubTaskRepositoryMiddleware) SubTaskRepository {
	return &subTaskRepository{
		collectionName:   "SubTask",
		firestoreClient:  firestoreClient,
		parentDocument:   parentDocument,
		middleware:       middleware,
		uniqueRepository: newUniqueRepository(firestoreClient, "SubTask"),
	}
}

// NewSubTaskCollectionGroupRepository - constructor
func NewSubTaskCollectionGroupRepository(firestoreClient *firestore.Client) SubTaskRepository {
	return &subTaskRepository{
		collectionName:  "SubTask",
		collectionGroup: firestoreClient.CollectionGroup("SubTask"),
	}
}

func (repo *subTaskRepository) beforeInsert(ctx context.Context, subject *SubTask) error {
	repo.uniqueRepository.setMiddleware(ctx)
	err := repo.uniqueRepository.CheckUnique(ctx, nil, subject)
	if err != nil {
		return xerrors.Errorf("unique.middleware error: %w", err)
	}

	for _, m := range repo.middleware {
		c, err := m.BeforeInsert(ctx, subject)
		if err != nil {
			return xerrors.Errorf("beforeInsert.middleware error: %w", err)
		}
		if !c {
			continue
		}
	}

	return nil
}

func (repo *subTaskRepository) beforeUpdate(ctx context.Context, old, subject *SubTask) error {
	if ctx.Value(transactionInProgressKey{}) != nil && old == nil {
		var err error
		doc := repo.GetDocRef(subject.ID)
		old, err = repo.get(context.Background(), doc)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return ErrNotFound
			}
			return xerrors.Errorf("error in Get method: %w", err)
		}
	}
	repo.uniqueRepository.setMiddleware(ctx)
	err := repo.uniqueRepository.CheckUnique(ctx, old, subject)
	if err != nil {
		return xerrors.Errorf("unique.middleware error: %w", err)
	}

	for _, m := range repo.middleware {
		c, err := m.BeforeUpdate(ctx, old, subject)
		if err != nil {
			return xerrors.Errorf("beforeUpdate.middleware error: %w", err)
		}
		if !c {
			continue
		}
	}

	return nil
}

func (repo *subTaskRepository) beforeDelete(ctx context.Context, subject *SubTask, opts ...DeleteOption) error {
	repo.uniqueRepository.setMiddleware(ctx)
	err := repo.uniqueRepository.DeleteUnique(ctx, subject)
	if err != nil {
		return xerrors.Errorf("unique.middleware error: %w", err)
	}

	for _, m := range repo.middleware {
		c, err := m.BeforeDelete(ctx, subject, opts...)
		if err != nil {
			return xerrors.Errorf("beforeDelete.middleware error: %w", err)
		}
		if !c {
			continue
		}
	}

	return nil
}

// GetCollection - *firestore.CollectionRef getter
func (repo *subTaskRepository) GetCollection() *firestore.CollectionRef {
	if repo.collectionGroup != nil {
		return nil
	}
	return repo.parentDocument.Collection(repo.collectionName)
}

// GetCollectionName - CollectionName getter
func (repo *subTaskRepository) GetCollectionName() string {
	return repo.collectionName
}

// GetDocRef - *firestore.DocumentRef getter
func (repo *subTaskRepository) GetDocRef(id string) *firestore.DocumentRef {
	if repo.collectionGroup != nil {
		return nil
	}
	return repo.GetCollection().Doc(id)
}

// RunInTransaction - (*firestore.Client).RunTransaction getter
func (repo *subTaskRepository) RunInTransaction() func(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error) {
	return repo.firestoreClient.RunTransaction
}

// SetParentDoc - parent document setter
func (repo *subTaskRepository) SetParentDoc(doc *firestore.DocumentRef) {
	if doc == nil {
		return
	}
	repo.parentDocument = doc
}

// NewRepositoryByParent - Returns new instance with setting parent document
func (repo subTaskRepository) NewRepositoryByParent(doc *firestore.DocumentRef) SubTaskRepository {
	if doc == nil {
		return &repo
	}
	repo.parentDocument = doc
	return &repo
}

// Free - parent document releaser
func (repo *subTaskRepository) Free() {
	repo.parentDocument = nil
}

// SubTaskSearchParam - params for search
type SubTaskSearchParam struct {
	ID              *QueryChainer
	IsSubCollection *QueryChainer

	CursorKey   string
	CursorLimit int
}

// SubTaskUpdateParam - params for strict updates
type SubTaskUpdateParam struct {
	IsSubCollection interface{}
}

// Search - search documents
// The third argument is firestore.Query, basically you can pass nil
func (repo *subTaskRepository) Search(ctx context.Context, param *SubTaskSearchParam, q *firestore.Query) ([]*SubTask, error) {
	return repo.search(ctx, param, q)
}

// SearchByParam - search documents by search param
func (repo *subTaskRepository) SearchByParam(ctx context.Context, param *SubTaskSearchParam) ([]*SubTask, *PagingResult, error) {
	return repo.searchByParam(ctx, param)
}

// Get - get `SubTask` by `SubTask.ID`
func (repo *subTaskRepository) Get(ctx context.Context, id string, opts ...GetOption) (*SubTask, error) {
	if repo.collectionGroup != nil {
		return nil, ErrNotAvailableCG
	}
	doc := repo.GetDocRef(id)
	return repo.get(ctx, doc, opts...)
}

// GetWithDoc - get `SubTask` by *firestore.DocumentRef
func (repo *subTaskRepository) GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...GetOption) (*SubTask, error) {
	if repo.collectionGroup != nil {
		return nil, ErrNotAvailableCG
	}
	return repo.get(ctx, doc, opts...)
}

// Insert - insert of `SubTask`
func (repo *subTaskRepository) Insert(ctx context.Context, subject *SubTask) (_ string, err error) {
	if repo.collectionGroup != nil {
		return "", ErrNotAvailableCG
	}
	if err := repo.beforeInsert(ctx, subject); err != nil {
		return "", xerrors.Errorf("before insert error: %w", err)
	}

	return repo.insert(ctx, subject)
}

// Update - update of `SubTask`
func (repo *subTaskRepository) Update(ctx context.Context, subject *SubTask) (err error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	doc := repo.GetDocRef(subject.ID)

	old, err := repo.get(ctx, doc)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return ErrNotFound
		}
		return xerrors.Errorf("error in Get method: %w", err)
	}

	if err := repo.beforeUpdate(ctx, old, subject); err != nil {
		return xerrors.Errorf("before update error: %w", err)
	}

	return repo.update(ctx, subject)
}

// StrictUpdate - strict update of `SubTask`
func (repo *subTaskRepository) StrictUpdate(ctx context.Context, id string, param *SubTaskUpdateParam, opts ...firestore.Precondition) error {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	return repo.strictUpdate(ctx, id, param, opts...)
}

// Delete - delete of `SubTask`
func (repo *subTaskRepository) Delete(ctx context.Context, subject *SubTask, opts ...DeleteOption) (err error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	if err := repo.beforeDelete(ctx, subject, opts...); err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}

	return repo.deleteByID(ctx, subject.ID)
}

// DeleteByID - delete `SubTask` by `SubTask.ID`
func (repo *subTaskRepository) DeleteByID(ctx context.Context, id string, opts ...DeleteOption) (err error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	subject, err := repo.Get(ctx, id)
	if err != nil {
		return xerrors.Errorf("error in Get method: %w", err)
	}

	if err := repo.beforeDelete(ctx, subject, opts...); err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}

	return repo.Delete(ctx, subject, opts...)
}

// GetMulti - get `SubTask` in bulk by array of `SubTask.ID`
func (repo *subTaskRepository) GetMulti(ctx context.Context, ids []string, opts ...GetOption) ([]*SubTask, error) {
	if repo.collectionGroup != nil {
		return nil, ErrNotAvailableCG
	}
	return repo.getMulti(ctx, ids, opts...)
}

// InsertMulti - bulk insert of `SubTask`
func (repo *subTaskRepository) InsertMulti(ctx context.Context, subjects []*SubTask) (_ []string, er error) {
	if repo.collectionGroup != nil {
		return nil, ErrNotAvailableCG
	}

	ids := make([]string, 0, len(subjects))
	batches := make([]*firestore.WriteBatch, 0)
	batch := repo.firestoreClient.Batch()
	collect := repo.GetCollection()

	for i, subject := range subjects {
		var ref *firestore.DocumentRef
		if subject.ID == "" {
			ref = collect.NewDoc()
			subject.ID = ref.ID
		} else {
			ref = collect.Doc(subject.ID)
			if s, err := ref.Get(ctx); err == nil {
				return nil, xerrors.Errorf("already exists [%v]: %#v", subject.ID, s)
			}
		}

		if err := repo.beforeInsert(ctx, subject); err != nil {
			return nil, xerrors.Errorf("before insert error(%d) [%v]: %w", i, subject.ID, err)
		}

		batch.Set(ref, subject)
		ids = append(ids, ref.ID)
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

	return ids, nil
}

// UpdateMulti - bulk update of `SubTask`
func (repo *subTaskRepository) UpdateMulti(ctx context.Context, subjects []*SubTask) (er error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}

	batches := make([]*firestore.WriteBatch, 0)
	batch := repo.firestoreClient.Batch()
	collect := repo.GetCollection()

	for i, subject := range subjects {
		ref := collect.Doc(subject.ID)
		snapShot, err := ref.Get(ctx)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found [%v]: %w", subject.ID, err)
			}
			return xerrors.Errorf("error in Get method [%v]: %w", subject.ID, err)
		}

		old := new(SubTask)
		if err = snapShot.DataTo(&old); err != nil {
			return xerrors.Errorf("error in DataTo method: %w", err)
		}

		if err := repo.beforeUpdate(ctx, old, subject); err != nil {
			return xerrors.Errorf("before update error(%d) [%v]: %w", i, subject.ID, err)
		}

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

// DeleteMulti - bulk delete of `SubTask`
func (repo *subTaskRepository) DeleteMulti(ctx context.Context, subjects []*SubTask, opts ...DeleteOption) (er error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}

	batches := make([]*firestore.WriteBatch, 0)
	batch := repo.firestoreClient.Batch()
	collect := repo.GetCollection()

	for i, subject := range subjects {
		ref := collect.Doc(subject.ID)
		if _, err := ref.Get(ctx); err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found [%v]: %w", subject.ID, err)
			}
			return xerrors.Errorf("error in Get method [%v]: %w", subject.ID, err)
		}

		if err := repo.beforeDelete(ctx, subject, opts...); err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subject.ID, err)
		}

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

// DeleteMultiByIDs - delete `SubTask` in bulk by array of `SubTask.ID`
func (repo *subTaskRepository) DeleteMultiByIDs(ctx context.Context, ids []string, opts ...DeleteOption) (er error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	subjects := make([]*SubTask, len(ids))

	opt := GetOption{}
	if len(opts) > 0 {
		opt.IncludeSoftDeleted = opts[0].Mode == DeleteModeHard
	}
	for i, id := range ids {
		subject, err := repo.Get(ctx, id, opt)
		if err != nil {
			return xerrors.Errorf("error in Get method: %w", err)
		}
		subjects[i] = subject
	}

	return repo.DeleteMulti(ctx, subjects, opts...)
}

// SearchWithTx - search documents in transaction
func (repo *subTaskRepository) SearchWithTx(tx *firestore.Transaction, param *SubTaskSearchParam, q *firestore.Query) ([]*SubTask, error) {
	return repo.search(tx, param, q)
}

// SearchByParamWithTx - search documents by search param in transaction
func (repo *subTaskRepository) SearchByParamWithTx(tx *firestore.Transaction, param *SubTaskSearchParam) ([]*SubTask, *PagingResult, error) {
	return repo.searchByParam(tx, param)
}

// GetWithTx - get `SubTask` by `SubTask.ID` in transaction
func (repo *subTaskRepository) GetWithTx(tx *firestore.Transaction, id string, opts ...GetOption) (*SubTask, error) {
	if repo.collectionGroup != nil {
		return nil, ErrNotAvailableCG
	}
	doc := repo.GetDocRef(id)
	return repo.get(tx, doc, opts...)
}

// GetWithDocWithTx - get `SubTask` by *firestore.DocumentRef in transaction
func (repo *subTaskRepository) GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...GetOption) (*SubTask, error) {
	if repo.collectionGroup != nil {
		return nil, ErrNotAvailableCG
	}
	return repo.get(tx, doc, opts...)
}

// InsertWithTx - insert of `SubTask` in transaction
func (repo *subTaskRepository) InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *SubTask) (_ string, err error) {
	if repo.collectionGroup != nil {
		return "", ErrNotAvailableCG
	}
	if err := repo.beforeInsert(context.WithValue(ctx, transactionInProgressKey{}, tx), subject); err != nil {
		return "", xerrors.Errorf("before insert error: %w", err)
	}

	return repo.insert(tx, subject)
}

// UpdateWithTx - update of `SubTask` in transaction
func (repo *subTaskRepository) UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *SubTask) (err error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	if err := repo.beforeUpdate(context.WithValue(ctx, transactionInProgressKey{}, tx), nil, subject); err != nil {
		return xerrors.Errorf("before update error: %w", err)
	}

	return repo.update(tx, subject)
}

// StrictUpdateWithTx - strict update of `SubTask` in transaction
func (repo *subTaskRepository) StrictUpdateWithTx(tx *firestore.Transaction, id string, param *SubTaskUpdateParam, opts ...firestore.Precondition) error {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	return repo.strictUpdate(tx, id, param, opts...)
}

// DeleteWithTx - delete of `SubTask` in transaction
func (repo *subTaskRepository) DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *SubTask, opts ...DeleteOption) (err error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	if err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, tx), subject, opts...); err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}

	return repo.deleteByID(tx, subject.ID)
}

// DeleteByIDWithTx - delete `SubTask` by `SubTask.ID` in transaction
func (repo *subTaskRepository) DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...DeleteOption) (err error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	subject, err := repo.Get(context.Background(), id)
	if err != nil {
		return xerrors.Errorf("error in Get method: %w", err)
	}

	if err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, tx), subject, opts...); err != nil {
		return xerrors.Errorf("before delete error: %w", err)
	}

	return repo.deleteByID(tx, id)
}

// GetMultiWithTx - get `SubTask` in bulk by array of `SubTask.ID` in transaction
func (repo *subTaskRepository) GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...GetOption) ([]*SubTask, error) {
	if repo.collectionGroup != nil {
		return nil, ErrNotAvailableCG
	}
	return repo.getMulti(tx, ids, opts...)
}

// InsertMultiWithTx - bulk insert of `SubTask` in transaction
func (repo *subTaskRepository) InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*SubTask) (_ []string, er error) {
	if repo.collectionGroup != nil {
		return nil, ErrNotAvailableCG
	}

	ids := make([]string, len(subjects))

	for i := range subjects {
		if err := repo.beforeInsert(ctx, subjects[i]); err != nil {
			return nil, xerrors.Errorf("before insert error(%d) [%v]: %w", i, subjects[i].ID, err)
		}

		id, err := repo.insert(tx, subjects[i])
		if err != nil {
			return nil, xerrors.Errorf("error in insert method(%d) [%v]: %w", i, subjects[i].ID, err)
		}
		ids[i] = id
	}

	return ids, nil
}

// UpdateMultiWithTx - bulk update of `SubTask` in transaction
func (repo *subTaskRepository) UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*SubTask) (er error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}
	ctx = context.WithValue(ctx, transactionInProgressKey{}, tx)

	for i := range subjects {
		if err := repo.beforeUpdate(ctx, nil, subjects[i]); err != nil {
			return xerrors.Errorf("before update error(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	for i := range subjects {
		if err := repo.update(tx, subjects[i]); err != nil {
			return xerrors.Errorf("error in update method(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	return nil
}

// DeleteMultiWithTx - bulk delete of `SubTask` in transaction
func (repo *subTaskRepository) DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*SubTask, opts ...DeleteOption) (er error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}

	var isHardDeleteMode bool
	if len(opts) > 0 {
		isHardDeleteMode = opts[0].Mode == DeleteModeHard
	}
	opt := GetOption{
		IncludeSoftDeleted: isHardDeleteMode,
	}
	for i := range subjects {
		dr := repo.GetDocRef(subjects[i].ID)
		if _, err := repo.get(context.Background(), dr, opt); err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found(%d) [%v]", i, subjects[i].ID)
			}
			return xerrors.Errorf("error in get method(%d) [%v]: %w", i, subjects[i].ID, err)
		}

		if err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, tx), subjects[i], opts...); err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	for i := range subjects {
		if err := repo.deleteByID(tx, subjects[i].ID); err != nil {
			return xerrors.Errorf("error in delete method(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	return nil
}

// DeleteMultiByIDWithTx - delete `SubTask` in bulk by array of `SubTask.ID` in transaction
func (repo *subTaskRepository) DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...DeleteOption) (er error) {
	if repo.collectionGroup != nil {
		return ErrNotAvailableCG
	}

	for i := range ids {
		dr := repo.GetDocRef(ids[i])
		subject, err := repo.get(context.Background(), dr)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found(%d) [%v]", i, ids[i])
			}
			return xerrors.Errorf("error in get method(%d) [%v]: %w", i, ids[i], err)
		}

		if err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, tx), subject, opts...); err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subject.ID, err)
		}
	}

	for i := range ids {
		if err := repo.deleteByID(tx, ids[i]); err != nil {
			return xerrors.Errorf("error in delete method(%d) [%v]: %w", i, ids[i], err)
		}
	}

	return nil
}

func (repo *subTaskRepository) get(v interface{}, doc *firestore.DocumentRef, _ ...GetOption) (*SubTask, error) {
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

	subject := new(SubTask)
	if err := snapShot.DataTo(&subject); err != nil {
		return nil, xerrors.Errorf("error in DataTo method: %w", err)
	}

	subject.ID = snapShot.Ref.ID

	return subject, nil
}

func (repo *subTaskRepository) getMulti(v interface{}, ids []string, _ ...GetOption) ([]*SubTask, error) {
	var (
		snapShots []*firestore.DocumentSnapshot
		err       error
		collect   = repo.GetCollection()
		drs       = make([]*firestore.DocumentRef, len(ids))
	)

	for i, id := range ids {
		ref := collect.Doc(id)
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

	subjects := make([]*SubTask, 0, len(ids))
	mErr := NewMultiErrors()
	for i, snapShot := range snapShots {
		if !snapShot.Exists() {
			mErr = append(mErr, NewMultiError(i, ErrNotFound))
			continue
		}

		subject := new(SubTask)
		if err = snapShot.DataTo(&subject); err != nil {
			return nil, xerrors.Errorf("error in DataTo method: %w", err)
		}

		subject.ID = snapShot.Ref.ID
		subjects = append(subjects, subject)
	}

	if len(mErr) == 0 {
		return subjects, nil
	}

	return subjects, mErr
}

func (repo *subTaskRepository) insert(v interface{}, subject *SubTask) (string, error) {
	var (
		dr  = repo.GetCollection().NewDoc()
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

	subject.ID = dr.ID

	return dr.ID, nil
}

func (repo *subTaskRepository) update(v interface{}, subject *SubTask) error {
	var (
		dr  = repo.GetDocRef(subject.ID)
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

func (repo *subTaskRepository) strictUpdate(v interface{}, id string, param *SubTaskUpdateParam, opts ...firestore.Precondition) error {
	var (
		dr  = repo.GetDocRef(id)
		err error
	)

	updates := updater(SubTask{}, param)

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

func (repo *subTaskRepository) deleteByID(v interface{}, id string) error {
	dr := repo.GetDocRef(id)
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

func (repo *subTaskRepository) runQuery(v interface{}, query firestore.Query) ([]*SubTask, error) {
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

	subjects := make([]*SubTask, 0)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, xerrors.Errorf("error in Next method: %w", err)
		}

		subject := new(SubTask)

		if err = doc.DataTo(&subject); err != nil {
			return nil, xerrors.Errorf("error in DataTo method: %w", err)
		}

		subject.ID = doc.Ref.ID
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

// BUG(54m): there may be potential bugs
func (repo *subTaskRepository) searchByParam(v interface{}, param *SubTaskSearchParam) ([]*SubTask, *PagingResult, error) {
	query := func() firestore.Query {
		if repo.collectionGroup != nil {
			return repo.collectionGroup.Query
		}
		return repo.GetCollection().Query
	}()
	if param.ID != nil {
		for _, chain := range param.ID.QueryGroup {
			var value interface{}
			switch val := chain.Value.(type) {
			case string:
				value = repo.GetDocRef(val)
			case []string:
				docRefs := make([]*firestore.DocumentRef, len(val))
				for i := range val {
					docRefs[i] = repo.GetDocRef(val[i])
				}
				value = docRefs
			default:
				return nil, nil, xerrors.Errorf("document id can only be of type `string` and `[]string`. value: %#v", chain.Value)
			}
			query = query.Where(firestore.DocumentID, chain.Operator, value)
		}
		if direction := param.ID.OrderByDirection; direction > 0 {
			query = query.OrderBy(firestore.DocumentID, direction)
			query = param.ID.BuildCursorQuery(query)
		}
	}
	if param.IsSubCollection != nil {
		for _, chain := range param.IsSubCollection.QueryGroup {
			query = query.Where("IsSubCollection", chain.Operator, chain.Value)
		}
		if direction := param.IsSubCollection.OrderByDirection; direction > 0 {
			query = query.OrderBy("IsSubCollection", direction)
			query = param.IsSubCollection.BuildCursorQuery(query)
		}
	}

	limit := param.CursorLimit + 1

	if param.CursorKey != "" {
		var (
			ds  *firestore.DocumentSnapshot
			err error
		)
		switch x := v.(type) {
		case *firestore.Transaction:
			ds, err = x.Get(repo.GetDocRef(param.CursorKey))
		case context.Context:
			ds, err = repo.GetDocRef(param.CursorKey).Get(x)
		default:
			return nil, nil, xerrors.Errorf("invalid x type: %v", v)
		}
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return nil, nil, ErrNotFound
			}
			return nil, nil, xerrors.Errorf("error in Get method: %w", err)
		}
		query = query.StartAt(ds)
	}

	if limit > 1 {
		query = query.Limit(limit)
	}

	subjects, err := repo.runQuery(v, query)
	if err != nil {
		return nil, nil, xerrors.Errorf("error in runQuery method: %w", err)
	}

	pagingResult := &PagingResult{
		Length: len(subjects),
	}
	if limit > 1 && limit == pagingResult.Length {
		next := pagingResult.Length - 1
		pagingResult.NextCursorKey = subjects[next].ID
		subjects = subjects[:next]
		pagingResult.Length--
	}

	return subjects, pagingResult, nil
}

func (repo *subTaskRepository) search(v interface{}, param *SubTaskSearchParam, q *firestore.Query) ([]*SubTask, error) {
	if (param == nil && q == nil) || (param != nil && q != nil) {
		return nil, xerrors.New("either one should be nil")
	}

	query := func() firestore.Query {
		if q != nil {
			return *q
		}
		if repo.collectionGroup != nil {
			return repo.collectionGroup.Query
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
