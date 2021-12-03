package examples

//go:generate ../bin/volcago -sub-collection History

// History - Task sub-collection
type History struct {
	ID              string `firestore:"-" firestore_key:"auto"`
	IsSubCollection bool   ``
	IsBool          bool   `firestore:"-"`
}
