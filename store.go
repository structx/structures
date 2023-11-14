package structs

// Store structures repository
type Store[T string | int8 | int64] interface {
	// Get get item from store
	Get(key T) (interface{}, error)
	// Set set item to store
	Set(key T, value interface{}) error
	// Delete delete item from store
	Delete(key T) error
	// Exists check if item exists in store
	Exists(key T) (bool, error)
}
