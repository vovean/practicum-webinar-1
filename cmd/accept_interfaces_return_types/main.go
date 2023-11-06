package main

type MyStore struct {
	// Absolute path to the base directory for storing data files.
	directoryPath string
}

func (s *MyStore) Read(key string) (string, error) {
	// Implementation

	return "", nil
}

func (s *MyStore) Write(key, val string) error {
	return nil
}

type Storer interface {
	Read(key string) (value string, err error)
}

type Writer interface {
	Write(key, val string) error
}

func NewMyStore() *MyStore {
	return &MyStore{directoryPath: "/tmp"}
}

func main() {
	store := NewMyStore()
	WriteToStorage(store, "", "")
	ReadFromStorage(store, "")
}

func WriteToStorage(w Writer, k, v string) error {
	return w.Write(k, v)
}

func ReadFromStorage(r Storer, key string) (string, error) {
	return r.Read(key)
}
