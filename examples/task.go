package examples

import "time"

//go:generate ../bin/volcago Task
//go:generate ../bin/volcago -o repository Task

type Inner struct {
	A     string `firestore:"a" indexer:"e"`
	Code  string `firestore:"code" indexer:"e" unique:""`
	Email string `indexer:"e" unique:""`
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
	TaskKind     TaskKind           `firestore:"taskKind"   indexer:"e"`
	Flag         map[string]float64 `firestore:"flag"`
	Indexes      map[string]bool    ``
	Inner        Inner              `firestore:"inner"`
	InnerRef     *Inner             `firestore:"innerRef"`
	InnerMap     map[string]Inner   `firestore:"innerMap"`
}

type TaskKind string

const (
	TaskKindTodo     TaskKind = "todo"
	TaskKindSchedule TaskKind = "schedule"
)
