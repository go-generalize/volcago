package model

//go:generate volcago -sub-collection SubTask

type SubTask struct {
	ID              string `firestore:"-" firestore_key:"auto"`
	IsSubCollection bool   ``
}
