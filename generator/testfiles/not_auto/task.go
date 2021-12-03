package model

import (
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/genproto/googleapis/type/latlng"
)

//go:generate volcago Task

// Task ID自動生成なし
type Task struct {
	Identity     string                 `firestore:"-" firestore_key:""`
	Desc         string                 `firestore:"description" unique:""`
	Desc2        string                 `firestore:"desc2"`
	Created      time.Time              `firestore:"created"`
	ReservedDate *time.Time             `firestore:"reservedDate"`
	Done         bool                   `firestore:"done"`
	Done2        bool                   `firestore:"done2"`
	Count        int                    `firestore:"count"`
	Count64      int64                  `firestore:"count64"`
	NameList     []string               `firestore:"nameList"`
	Proportion   float64                `firestore:"proportion"`
	Geo          *latlng.LatLng         `firestore:"geo"`
	Sub          *firestore.DocumentRef `firestore:"sub"`
	Flag         Flag                   `firestore:"flag"`
}

type Flag bool
