package model

//go:generate volcago -sub-collection SubCollection

type SubCollection struct {
	ID   string `firestore:"-" firestore_key:"auto"`
	Flag bool   `firestore:"flag"`
}
