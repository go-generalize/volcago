package model

import (
	"time"

	"cloud.google.com/go/firestore"
)

//go:generate volcago Task

// Task ID自動生成あり
type Task struct {
	ID         string                 `firestore:"-" firestore_key:"auto"`
	Desc       string                 `firestore:"description" indexer:"e,p,s,l"`
	Created    time.Time              `firestore:"created"`
	Done       bool                   `firestore:"done"`
	Done2      bool                   `firestore:"done2"`
	Count      int                    `firestore:"count"`
	Count64    int64                  `firestore:"count64"`
	NameList   []string               `firestore:"nameList"`
	Proportion float64                `firestore:"proportion" indexer:"e"`
	Flag       map[string]float64     `firestore:"flag"`
	Indexes    map[string]bool        `firestore:"indexes"`
	Sub        *firestore.DocumentRef `firestore:"sub"`
}
