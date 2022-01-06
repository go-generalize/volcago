package examples

//go:generate ../bin/volcago -p repository -o different_dir -mockgen ../../bin/mockgen -mock-output mock/lock_gen.go Lock

// Lock - with automatic id generation
type Lock struct {
	ID     string             `firestore:"-" firestore_key:"auto"`
	Text   string             `firestore:"text" unique:""`
	Flag   map[string]float64 `firestore:"flag"`
	Nested Nested             `firestore:"nested"`
	Meta
}

// Nested - nested struct
type Nested struct {
	Name string `firestore:"name"`
}
