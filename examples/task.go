package examples

import "time"

//go:generate ../bin/volcago Task

type Inner struct {
	A string `firestore:"a"`
}

type Embedded struct {
	ID   string `firestore:"-" firestore_key:"auto"`
	Desc string `firestore:"description" indexer:"e,p,s,l"`
}

// Task - with automatic id generation
type Task struct {
	Embedded
	Created      time.Time          `firestore:"created"`
	ReservedDate *time.Time         `firestore:"reservedDate"`
	Done         bool               `firestore:"done"`
	Done2        bool               `firestore:"done2"`
	Count        int                `firestore:"count"`
	Count64      int64              `firestore:"count64"`
	NameList     []string           `firestore:"nameList"`
	Proportion   float64            `firestore:"proportion" indexer:"e"`
	Flag         map[string]float64 `firestore:"flag"`
	Indexes      map[string]bool    ``
	Inner        Inner              `firestore:"inner"`
}
