// Code generated by volcago. DO NOT EDIT.
// generated version: (devel)
package model

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"golang.org/x/xerrors"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate mockgen -source $GOFILE -destination mock/mock_lock_meta2_gen/mock_lock_meta2_gen.go

// LockMeta2Repository - Repository of LockMeta2
type LockMeta2Repository interface {
	// Single
	Get(ctx context.Context, id string, opts ...GetOption) (*LockMeta2, error)
	GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...GetOption) (*LockMeta2, error)
	Insert(ctx context.Context, subject *LockMeta2) (_ string, err error)
	Update(ctx context.Context, subject *LockMeta2) (err error)
	StrictUpdate(ctx context.Context, id string, param *LockMeta2UpdateParam, opts ...firestore.Precondition) error
	Delete(ctx context.Context, subject *LockMeta2, opts ...DeleteOption) (err error)
	DeleteByID(ctx context.Context, id string, opts ...DeleteOption) (err error)
	// Multiple
	GetMulti(ctx context.Context, ids []string, opts ...GetOption) ([]*LockMeta2, error)
	InsertMulti(ctx context.Context, subjects []*LockMeta2) (_ []string, er error)
	UpdateMulti(ctx context.Context, subjects []*LockMeta2) (er error)
	DeleteMulti(ctx context.Context, subjects []*LockMeta2, opts ...DeleteOption) (er error)
	DeleteMultiByIDs(ctx context.Context, ids []string, opts ...DeleteOption) (er error)
	// Single(Transaction)
	GetWithTx(tx *firestore.Transaction, id string, opts ...GetOption) (*LockMeta2, error)
	GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...GetOption) (*LockMeta2, error)
	InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *LockMeta2) (_ string, err error)
	UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *LockMeta2) (err error)
	StrictUpdateWithTx(tx *firestore.Transaction, id string, param *LockMeta2UpdateParam, opts ...firestore.Precondition) error
	DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *LockMeta2, opts ...DeleteOption) (err error)
	DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...DeleteOption) (err error)
	// Multiple(Transaction)
	GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...GetOption) ([]*LockMeta2, error)
	InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*LockMeta2) (_ []string, er error)
	UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*LockMeta2) (er error)
	DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*LockMeta2, opts ...DeleteOption) (er error)
	DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...DeleteOption) (er error)
	// Search
	Search(ctx context.Context, param *LockMeta2SearchParam, q *firestore.Query) ([]*LockMeta2, error)
	SearchWithTx(tx *firestore.Transaction, param *LockMeta2SearchParam, q *firestore.Query) ([]*LockMeta2, error)
	// misc
	GetCollection() *firestore.CollectionRef
	GetCollectionName() string
	GetDocRef(id string) *firestore.DocumentRef
	RunInTransaction() func(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error)
}

// LockMeta2RepositoryMiddleware - middleware of LockMeta2Repository
type LockMeta2RepositoryMiddleware interface {
	BeforeInsert(ctx context.Context, subject *LockMeta2) (bool, error)
	BeforeUpdate(ctx context.Context, old, subject *LockMeta2) (bool, error)
	BeforeDelete(ctx context.Context, subject *LockMeta2, opts ...DeleteOption) (bool, error)
	BeforeDeleteByID(ctx context.Context, ids []string, opts ...DeleteOption) (bool, error)
}

type lockMeta2Repository struct {
	collectionName   string
	firestoreClient  *firestore.Client
	middleware       []LockMeta2RepositoryMiddleware
	uniqueRepository *uniqueRepository
}

// NewLockMeta2Repository - constructor
func NewLockMeta2Repository(firestoreClient *firestore.Client, middleware ...LockMeta2RepositoryMiddleware) LockMeta2Repository {
	return &lockMeta2Repository{
		collectionName:   "LockMeta2",
		firestoreClient:  firestoreClient,
		middleware:       middleware,
		uniqueRepository: newUniqueRepository(firestoreClient, "LockMeta2"),
	}
}

func (repo *lockMeta2Repository) setMeta(subject *LockMeta2, isInsert bool) {
	now := time.Now()

	if isInsert {
		subject.CreatedAt = now
	}
	subject.UpdatedAt = now
	subject.Version++
}

func (repo *lockMeta2Repository) setMetaWithStrictUpdate(param *LockMeta2UpdateParam) {
	param.UpdatedAt = firestore.ServerTimestamp
	param.Version = firestore.Increment(1)
}

func (repo *lockMeta2Repository) beforeInsert(ctx context.Context, subject *LockMeta2) (RollbackFunc, error) {
	if subject.Version != 0 {
		return nil, xerrors.Errorf("insert data must be Version == 0 %+v: %w", subject, ErrVersionConflict)
	}
	if subject.DeletedAt != nil {
		return nil, xerrors.Errorf("insert data must be DeletedAt == nil: %+v", subject)
	}
	repo.setMeta(subject, true)
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

func (repo *lockMeta2Repository) beforeUpdate(ctx context.Context, old, subject *LockMeta2) (RollbackFunc, error) {
	if ctx.Value(transactionInProgressKey{}) != nil && old == nil {
		var err error
		doc := repo.GetDocRef(subject.ID)
		old, err = repo.get(context.Background(), doc)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return nil, ErrNotFound
			}
			return nil, xerrors.Errorf("error in Get method: %w", err)
		}
	}
	if old.Version > subject.Version {
		return nil, xerrors.Errorf(
			"The data in the database is newer: (db version: %d, target version: %d) %+v: %w",
			old.Version, subject.Version, subject, ErrVersionConflict,
		)
	}
	if subject.DeletedAt != nil {
		return nil, xerrors.Errorf("update data must be DeletedAt == nil: %+v", subject)
	}
	repo.setMeta(subject, false)
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

func (repo *lockMeta2Repository) beforeDelete(ctx context.Context, subject *LockMeta2, opts ...DeleteOption) (RollbackFunc, error) {
	repo.setMeta(subject, false)
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
func (repo *lockMeta2Repository) GetCollection() *firestore.CollectionRef {
	return repo.firestoreClient.Collection(repo.collectionName)
}

// GetCollectionName - CollectionName getter
func (repo *lockMeta2Repository) GetCollectionName() string {
	return repo.collectionName
}

// GetDocRef - *firestore.DocumentRef getter
func (repo *lockMeta2Repository) GetDocRef(id string) *firestore.DocumentRef {
	return repo.GetCollection().Doc(id)
}

// RunInTransaction - (*firestore.Client).RunTransaction getter
func (repo *lockMeta2Repository) RunInTransaction() func(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error) {
	return repo.firestoreClient.RunTransaction
}

// LockMeta2SearchParam - params for search
type LockMeta2SearchParam struct {
	Text      *QueryChainer
	Flag      *QueryChainer
	CreatedAt *QueryChainer
	CreatedBy *QueryChainer
	UpdatedAt *QueryChainer
	UpdatedBy *QueryChainer
	DeletedAt *QueryChainer
	DeletedBy *QueryChainer
	Version   *QueryChainer

	IncludeSoftDeleted bool
	CursorLimit        int
}

// LockMeta2UpdateParam - params for strict updates
type LockMeta2UpdateParam struct {
	Text      interface{}
	Flag      interface{}
	CreatedAt interface{}
	CreatedBy interface{}
	UpdatedAt interface{}
	UpdatedBy interface{}
	DeletedAt interface{}
	DeletedBy interface{}
	Version   interface{}
}

// Search - search documents
// The third argument is firestore.Query, basically you can pass nil
func (repo *lockMeta2Repository) Search(ctx context.Context, param *LockMeta2SearchParam, q *firestore.Query) ([]*LockMeta2, error) {
	return repo.search(ctx, param, q)
}

// Get - get `LockMeta2` by `LockMeta2.ID`
func (repo *lockMeta2Repository) Get(ctx context.Context, id string, opts ...GetOption) (*LockMeta2, error) {
	doc := repo.GetDocRef(id)
	return repo.get(ctx, doc, opts...)
}

// GetWithDoc - get `LockMeta2` by *firestore.DocumentRef
func (repo *lockMeta2Repository) GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...GetOption) (*LockMeta2, error) {
	return repo.get(ctx, doc, opts...)
}

// Insert - insert of `LockMeta2`
func (repo *lockMeta2Repository) Insert(ctx context.Context, subject *LockMeta2) (_ string, err error) {
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

// Update - update of `LockMeta2`
func (repo *lockMeta2Repository) Update(ctx context.Context, subject *LockMeta2) (err error) {
	doc := repo.GetDocRef(subject.ID)

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

// StrictUpdate - strict update of `LockMeta2`
func (repo *lockMeta2Repository) StrictUpdate(ctx context.Context, id string, param *LockMeta2UpdateParam, opts ...firestore.Precondition) error {
	return repo.strictUpdate(ctx, id, param, opts...)
}

// Delete - delete of `LockMeta2`
func (repo *lockMeta2Repository) Delete(ctx context.Context, subject *LockMeta2, opts ...DeleteOption) (err error) {
	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		t := time.Now()
		subject.DeletedAt = &t
		if err := repo.update(ctx, subject); err != nil {
			return xerrors.Errorf("error in update method: %w", err)
		}
		return nil
	}

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

	return repo.deleteByID(ctx, subject.ID)
}

// DeleteByID - delete `LockMeta2` by `LockMeta2.ID`
func (repo *lockMeta2Repository) DeleteByID(ctx context.Context, id string, opts ...DeleteOption) (err error) {
	subject, err := repo.Get(ctx, id)
	if err != nil {
		return xerrors.Errorf("error in Get method: %w", err)
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		t := time.Now()
		subject.DeletedAt = &t
		if err := repo.update(ctx, subject); err != nil {
			return xerrors.Errorf("error in update method: %w", err)
		}
		return nil
	}

	return repo.Delete(ctx, subject, opts...)
}

// GetMulti - get `LockMeta2` in bulk by array of `LockMeta2.ID`
func (repo *lockMeta2Repository) GetMulti(ctx context.Context, ids []string, opts ...GetOption) ([]*LockMeta2, error) {
	return repo.getMulti(ctx, ids, opts...)
}

// InsertMulti - bulk insert of `LockMeta2`
func (repo *lockMeta2Repository) InsertMulti(ctx context.Context, subjects []*LockMeta2) (_ []string, er error) {
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

		rb, err := repo.beforeInsert(ctx, subject)
		if err != nil {
			return nil, xerrors.Errorf("before insert error(%d) [%v]: %w", i, subject.ID, err)
		}
		rbs = append(rbs, rb)

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

// UpdateMulti - bulk update of `LockMeta2`
func (repo *lockMeta2Repository) UpdateMulti(ctx context.Context, subjects []*LockMeta2) (er error) {
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
		ref := collect.Doc(subject.ID)
		snapShot, err := ref.Get(ctx)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found [%v]: %w", subject.ID, err)
			}
			return xerrors.Errorf("error in Get method [%v]: %w", subject.ID, err)
		}

		old := new(LockMeta2)
		if err = snapShot.DataTo(&old); err != nil {
			return xerrors.Errorf("error in DataTo method: %w", err)
		}

		rb, err := repo.beforeUpdate(ctx, old, subject)
		if err != nil {
			return xerrors.Errorf("before update error(%d) [%v]: %w", i, subject.ID, err)
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

// DeleteMulti - bulk delete of `LockMeta2`
func (repo *lockMeta2Repository) DeleteMulti(ctx context.Context, subjects []*LockMeta2, opts ...DeleteOption) (er error) {
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
		ref := collect.Doc(subject.ID)
		if _, err := ref.Get(ctx); err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found [%v]: %w", subject.ID, err)
			}
			return xerrors.Errorf("error in Get method [%v]: %w", subject.ID, err)
		}

		rb, err := repo.beforeDelete(ctx, subject, opts...)
		if err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subject.ID, err)
		}
		rbs = append(rbs, rb)

		if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
			t := time.Now()
			subject.DeletedAt = &t
			batch.Set(ref, subject)
		} else {
			batch.Delete(ref)
		}

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

// DeleteMultiByIDs - delete `LockMeta2` in bulk by array of `LockMeta2.ID`
func (repo *lockMeta2Repository) DeleteMultiByIDs(ctx context.Context, ids []string, opts ...DeleteOption) (er error) {
	subjects := make([]*LockMeta2, len(ids))

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

func (repo *lockMeta2Repository) SearchWithTx(tx *firestore.Transaction, param *LockMeta2SearchParam, q *firestore.Query) ([]*LockMeta2, error) {
	return repo.search(tx, param, q)
}

// GetWithTx - get `LockMeta2` by `LockMeta2.ID` in transaction
func (repo *lockMeta2Repository) GetWithTx(tx *firestore.Transaction, id string, opts ...GetOption) (*LockMeta2, error) {
	doc := repo.GetDocRef(id)
	return repo.get(tx, doc, opts...)
}

// GetWithDocWithTx - get `LockMeta2` by *firestore.DocumentRef in transaction
func (repo *lockMeta2Repository) GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...GetOption) (*LockMeta2, error) {
	return repo.get(tx, doc, opts...)
}

// InsertWithTx - insert of `LockMeta2` in transaction
func (repo *lockMeta2Repository) InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *LockMeta2) (_ string, err error) {
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

// UpdateWithTx - update of `LockMeta2` in transaction
func (repo *lockMeta2Repository) UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *LockMeta2) (err error) {
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

// StrictUpdateWithTx - strict update of `LockMeta2` in transaction
func (repo *lockMeta2Repository) StrictUpdateWithTx(tx *firestore.Transaction, id string, param *LockMeta2UpdateParam, opts ...firestore.Precondition) error {
	return repo.strictUpdate(tx, id, param, opts...)
}

// DeleteWithTx - delete of `LockMeta2` in transaction
func (repo *lockMeta2Repository) DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *LockMeta2, opts ...DeleteOption) (err error) {
	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		t := time.Now()
		subject.DeletedAt = &t
		if err := repo.update(tx, subject); err != nil {
			return xerrors.Errorf("error in update method: %w", err)
		}
		return nil
	}

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

	return repo.deleteByID(tx, subject.ID)
}

// DeleteByIDWithTx - delete `LockMeta2` by `LockMeta2.ID` in transaction
func (repo *lockMeta2Repository) DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...DeleteOption) (err error) {
	subject, err := repo.Get(context.Background(), id)
	if err != nil {
		return xerrors.Errorf("error in Get method: %w", err)
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		t := time.Now()
		subject.DeletedAt = &t
		if err := repo.update(tx, subject); err != nil {
			return xerrors.Errorf("error in update method: %w", err)
		}
		return nil
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

	return repo.deleteByID(tx, id)
}

// GetMultiWithTx - get `LockMeta2` in bulk by array of `LockMeta2.ID` in transaction
func (repo *lockMeta2Repository) GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...GetOption) ([]*LockMeta2, error) {
	return repo.getMulti(tx, ids, opts...)
}

// InsertMultiWithTx - bulk insert of `LockMeta2` in transaction
func (repo *lockMeta2Repository) InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*LockMeta2) (_ []string, er error) {
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

	ids := make([]string, len(subjects))

	for i := range subjects {
		rb, err := repo.beforeInsert(ctx, subjects[i])
		if err != nil {
			return nil, xerrors.Errorf("before insert error(%d) [%v]: %w", i, subjects[i].ID, err)
		}
		rbs = append(rbs, rb)

		id, err := repo.insert(tx, subjects[i])
		if err != nil {
			return nil, xerrors.Errorf("error in insert method(%d) [%v]: %w", i, subjects[i].ID, err)
		}
		ids[i] = id
	}

	return ids, nil
}

// UpdateMultiWithTx - bulk update of `LockMeta2` in transaction
func (repo *lockMeta2Repository) UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*LockMeta2) (er error) {
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
			return xerrors.Errorf("before update error(%d) [%v]: %w", i, subjects[i].ID, err)
		}
		rbs = append(rbs, rb)
	}

	for i := range subjects {
		if err := repo.update(tx, subjects[i]); err != nil {
			return xerrors.Errorf("error in update method(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	return nil
}

// DeleteMultiWithTx - bulk delete of `LockMeta2` in transaction
func (repo *lockMeta2Repository) DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*LockMeta2, opts ...DeleteOption) (er error) {
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

	t := time.Now()
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

		rb, err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, 1), subjects[i], opts...)
		if err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subjects[i].ID, err)
		}
		rbs = append(rbs, rb)

		if !isHardDeleteMode {
			subjects[i].DeletedAt = &t
			if err := repo.update(tx, subjects[i]); err != nil {
				return xerrors.Errorf("error in update method(%d) [%v]: %w", i, subjects[i].ID, err)
			}
		}
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		return nil
	}

	for i := range subjects {
		if err := repo.deleteByID(tx, subjects[i].ID); err != nil {
			return xerrors.Errorf("error in delete method(%d) [%v]: %w", i, subjects[i].ID, err)
		}
	}

	return nil
}

// DeleteMultiByIDWithTx - delete `LockMeta2` in bulk by array of `LockMeta2.ID` in transaction
func (repo *lockMeta2Repository) DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...DeleteOption) (er error) {
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

	t := time.Now()
	for i := range ids {
		dr := repo.GetDocRef(ids[i])
		subject, err := repo.get(context.Background(), dr)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return xerrors.Errorf("not found(%d) [%v]", i, ids[i])
			}
			return xerrors.Errorf("error in get method(%d) [%v]: %w", i, ids[i], err)
		}

		rb, err := repo.beforeDelete(context.WithValue(ctx, transactionInProgressKey{}, 1), subject, opts...)
		if err != nil {
			return xerrors.Errorf("before delete error(%d) [%v]: %w", i, subject.ID, err)
		}
		rbs = append(rbs, rb)

		if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
			subject.DeletedAt = &t
			if err := repo.update(tx, subject); err != nil {
				return xerrors.Errorf("error in update method(%d) [%v]: %w", i, ids[i], err)
			}
		}
	}

	if len(opts) > 0 && opts[0].Mode == DeleteModeSoft {
		return nil
	}

	for i := range ids {
		if err := repo.deleteByID(tx, ids[i]); err != nil {
			return xerrors.Errorf("error in delete method(%d) [%v]: %w", i, ids[i], err)
		}
	}

	return nil
}

func (repo *lockMeta2Repository) get(v interface{}, doc *firestore.DocumentRef, opts ...GetOption) (*LockMeta2, error) {
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

	subject := new(LockMeta2)
	if err := snapShot.DataTo(&subject); err != nil {
		return nil, xerrors.Errorf("error in DataTo method: %w", err)
	}

	if len(opts) == 0 || !opts[0].IncludeSoftDeleted {
		if subject.DeletedAt != nil {
			return nil, ErrAlreadyDeleted
		}
	}
	subject.ID = snapShot.Ref.ID

	return subject, nil
}

func (repo *lockMeta2Repository) getMulti(v interface{}, ids []string, opts ...GetOption) ([]*LockMeta2, error) {
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

	subjects := make([]*LockMeta2, 0, len(ids))
	mErr := NewMultiErrors()
	for i, snapShot := range snapShots {
		if !snapShot.Exists() {
			mErr = append(mErr, NewMultiError(i, ErrNotFound))
			continue
		}

		subject := new(LockMeta2)
		if err = snapShot.DataTo(&subject); err != nil {
			return nil, xerrors.Errorf("error in DataTo method: %w", err)
		}

		if len(opts) == 0 || !opts[0].IncludeSoftDeleted {
			if subject.DeletedAt != nil {
				mErr = append(mErr, NewMultiError(i, ErrLogicallyDeletedData))
				continue
			}
		}
		subject.ID = snapShot.Ref.ID
		subjects = append(subjects, subject)
	}

	if len(mErr) == 0 {
		return subjects, nil
	}

	return subjects, mErr
}

func (repo *lockMeta2Repository) insert(v interface{}, subject *LockMeta2) (string, error) {
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

func (repo *lockMeta2Repository) update(v interface{}, subject *LockMeta2) error {
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

func (repo *lockMeta2Repository) strictUpdate(v interface{}, id string, param *LockMeta2UpdateParam, opts ...firestore.Precondition) error {
	var (
		dr  = repo.GetDocRef(id)
		err error
	)

	repo.setMetaWithStrictUpdate(param)

	updates := updater(LockMeta2{}, param)

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

func (repo *lockMeta2Repository) deleteByID(v interface{}, id string) error {
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

func (repo *lockMeta2Repository) runQuery(v interface{}, query firestore.Query) ([]*LockMeta2, error) {
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

	subjects := make([]*LockMeta2, 0)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, xerrors.Errorf("error in Next method: %w", err)
		}

		subject := new(LockMeta2)

		if err := doc.DataTo(&subject); err != nil {
			return nil, xerrors.Errorf("error in DataTo method: %w", err)
		}

		subject.ID = doc.Ref.ID
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

// BUG(54m): there may be potential bugs
func (repo *lockMeta2Repository) search(v interface{}, param *LockMeta2SearchParam, q *firestore.Query) ([]*LockMeta2, error) {
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
		if param.Text != nil {
			for _, chain := range param.Text.QueryGroup {
				query = query.Where("text", chain.Operator, chain.Value)
			}
			if direction := param.Text.OrderByDirection; direction > 0 {
				query = query.OrderBy("text", direction)
				query = param.Text.BuildCursorQuery(query)
			}
		}
		if param.Flag != nil {
			for _, chain := range param.Flag.QueryGroup {
				items, ok := chain.Value.(map[string]float64)
				if !ok {
					continue
				}
				for key, value := range items {
					query = query.WherePath(firestore.FieldPath{"flag", key}, chain.Operator, value)
				}
			}
		}
		if param.CreatedAt != nil {
			for _, chain := range param.CreatedAt.QueryGroup {
				query = query.Where("CreatedAt", chain.Operator, chain.Value)
			}
			if direction := param.CreatedAt.OrderByDirection; direction > 0 {
				query = query.OrderBy("CreatedAt", direction)
				query = param.CreatedAt.BuildCursorQuery(query)
			}
		}
		if param.CreatedBy != nil {
			for _, chain := range param.CreatedBy.QueryGroup {
				query = query.Where("CreatedBy", chain.Operator, chain.Value)
			}
			if direction := param.CreatedBy.OrderByDirection; direction > 0 {
				query = query.OrderBy("CreatedBy", direction)
				query = param.CreatedBy.BuildCursorQuery(query)
			}
		}
		if param.UpdatedAt != nil {
			for _, chain := range param.UpdatedAt.QueryGroup {
				query = query.Where("UpdatedAt", chain.Operator, chain.Value)
			}
			if direction := param.UpdatedAt.OrderByDirection; direction > 0 {
				query = query.OrderBy("UpdatedAt", direction)
				query = param.UpdatedAt.BuildCursorQuery(query)
			}
		}
		if param.UpdatedBy != nil {
			for _, chain := range param.UpdatedBy.QueryGroup {
				query = query.Where("UpdatedBy", chain.Operator, chain.Value)
			}
			if direction := param.UpdatedBy.OrderByDirection; direction > 0 {
				query = query.OrderBy("UpdatedBy", direction)
				query = param.UpdatedBy.BuildCursorQuery(query)
			}
		}
		if param.DeletedAt != nil {
			for _, chain := range param.DeletedAt.QueryGroup {
				query = query.Where("DeletedAt", chain.Operator, chain.Value)
			}
			if direction := param.DeletedAt.OrderByDirection; direction > 0 {
				query = query.OrderBy("DeletedAt", direction)
				query = param.DeletedAt.BuildCursorQuery(query)
			}
		}
		if param.DeletedBy != nil {
			for _, chain := range param.DeletedBy.QueryGroup {
				query = query.Where("DeletedBy", chain.Operator, chain.Value)
			}
			if direction := param.DeletedBy.OrderByDirection; direction > 0 {
				query = query.OrderBy("DeletedBy", direction)
				query = param.DeletedBy.BuildCursorQuery(query)
			}
		}
		if param.Version != nil {
			for _, chain := range param.Version.QueryGroup {
				query = query.Where("Version", chain.Operator, chain.Value)
			}
			if direction := param.Version.OrderByDirection; direction > 0 {
				query = query.OrderBy("Version", direction)
				query = param.Version.BuildCursorQuery(query)
			}
		}
		if !param.IncludeSoftDeleted {
			query = query.Where("DeletedAt", OpTypeEqual, nil)
		}

		if l := param.CursorLimit; l > 0 {
			query = query.Limit(l)
		}
	}

	return repo.runQuery(v, query)
}
