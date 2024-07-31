package model

import (
	"time"
)

//go:generate volcago Task

// Task ID自動生成あり
type Task struct {
	ID               string             `firestore:"-" firestore_key:"auto"`
	Desc             string             `firestore:"description" indexer:"e,p,s,l"`
	Created          time.Time          `firestore:"created"`
	Done             bool               `firestore:"done"`
	Done2            bool               `firestore:"done2"`
	Count            int                `firestore:"count"`
	Count64          int64              `firestore:"count64"`
	NameList         []string           `firestore:"nameList"`
	Proportion       float64            `firestore:"proportion" indexer:"e"`
	Flag             map[string]float64 `firestore:"flag"`
	SliceSubTask     []*SubTask         `firestore:"slice_sub_task"`
	NestedSubTask    SubTask            `firestore:"nested_sub_task"`
	NestedRefSubTask *SubTask           `firestore:"nested_ref_sub_task"`
	Indexes          map[string]bool    `firestore:"indexes"`
}

// SubTask - nested struct
type SubTask struct {
	Name                string      `firestore:"name" indexer:"e"`
	NestedSubSubTask    SubSubTask  `firestore:"nested_sub_task"`
	NestedRefSubSubTask *SubSubTask `firestore:"nested_ref_sub_task"`
}

// SubSubTask - nested struct
type SubSubTask struct {
	Name string `firestore:"name" indexer:"e"`
}
