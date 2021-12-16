package examples

//go:generate ../bin/volcago -disable-meta -sub-collection -c histories History

// History - Task sub-collection
type History struct {
	ID              string `firestore:"-" firestore_key:"auto"`
	IsSubCollection bool   ``
	IsBool          bool   `firestore:"-"`
}
