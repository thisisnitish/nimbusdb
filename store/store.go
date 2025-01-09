package store

// TODO: Change it to sync map -> concurrent map
type Store struct {
	data map[string]string
}

func (s *Store) Set(key string, value string) string {
	s.data[key] = value

	return "OK"
}

func (s *Store) Get(key string) string {
	val, exists := s.data[key]
	if exists {
		return val
	}
	return "ERR, Invalid Key"
}

func (s *Store) Delete(key string) string {
	val, exists := s.data[key]
	if exists {
		delete(s.data, key)
		return val + " Deleted"
	}

	return "ERR, Invaid key"
}
