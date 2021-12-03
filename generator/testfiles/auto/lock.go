package model

import "time"

//go:generate volcago Lock

type Meta struct {
	CreatedAt time.Time  `firestore:"createdAt"`
	CreatedBy string     `firestore:"createdBy"`
	UpdatedAt time.Time  `firestore:"updatedAt"`
	UpdatedBy string     `firestore:"updatedBy"`
	DeletedAt *time.Time `firestore:"deletedAt"`
	DeletedBy string     `firestore:"deletedBy"`
	Version   int        `firestore:"version"`
}

// Lock ID自動生成あり
type Lock struct {
	ID   string             `firestore:"-" firestore_key:"auto"`
	Text string             `firestore:"text"`
	Flag map[string]float64 `firestore:"flag"`
	Meta
}
