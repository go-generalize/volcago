package model

//go:generate volcago LockMeta
//go:generate volcago -o different_dir LockMeta

// LockMeta ID自動生成あり
type LockMeta struct {
	ID   string             `firestore:"-" firestore_key:"auto"`
	Text string             `firestore:"text"`
	Flag map[string]float64 `firestore:"flag"`
	Meta
}
