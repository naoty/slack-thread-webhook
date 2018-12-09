package datastore

// Client writes and reads data.
type Client interface {
	Get(key string) (string, error)
	Set(key, value string) error
}
