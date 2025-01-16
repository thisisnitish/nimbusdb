package store

// TODO: Change it to sync map -> concurrent map
type Store struct {
	data map[string]string
	// list map[string][]string
	List map[string][]string
	Sets map[string]map[string]bool
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
		List: make(map[string][]string),
		Sets: make(map[string]map[string]bool),
	}
}

func (s *Store) Set(key string, value string) string {
	s.data[key] = value

	return "OK"
}

func (s *Store) Get(key string) string {
	val, ok := s.data[key]
	if ok {
		return val
	}
	return "ERR, Invalid Key"
}

func (s *Store) Delete(key string) string {
	val, ok := s.data[key]
	if ok {
		delete(s.data, key)
		return val + " Deleted"
	}

	return "ERR, Invaid key"
}
