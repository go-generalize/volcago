package examples

import "time"

// Meta - meta information for optimistic exclusive lock
type Meta struct {
	CreatedAt time.Time  `firestore:"createdAt"`
	CreatedBy string     `firestore:"createdBy"`
	UpdatedAt time.Time  `firestore:"updatedAt"`
	UpdatedBy string     `firestore:"updatedBy"`
	DeletedAt *time.Time `firestore:"deletedAt"`
	DeletedBy string     `firestore:"deletedBy"`
	Version   int        `firestore:"version"`
}
