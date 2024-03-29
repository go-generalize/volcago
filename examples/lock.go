package examples

//go:generate ../bin/volcago -p repository -o different_dir -mockgen ../../bin/mockgen -mock-output mock/lock_gen.go Lock

// Lock - with automatic id generation
type Lock struct {
	ID           string                 `firestore:"-" firestore_key:"auto"`
	Text         string                 `firestore:"text" unique:""`
	Email        string                 `unique:""`
	Flag         map[string]float64     `firestore:"flag"`
	Interface    interface{}            `firestore:"interface"`
	MapInterface map[string]interface{} `firestore:"map_interface"`
	Nested       Nested                 `firestore:"nested"`
	NestedPtr    *Nested                `firestore:"nested_ptr"`
	SliceString  []string               `firestore:"slice_string"`
	SliceNested  []*Nested              `firestore:"slice_nested"`
	Meta
}

// Nested - nested struct
type Nested struct {
	Name string `firestore:"name"`
}
