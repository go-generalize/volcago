package model

//go:generate volcago -sub-collection SubTask

type SubTask struct {
	ID   string `firestore:"-" firestore_key:"auto"`
	Flag bool   `firestore:"flag"`
}
