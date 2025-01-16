package cmd

import "github.com/thisisnitish/nimbusdb/store"

func SIsMember(key string, value string, store *store.Store) string {
	if store.Sets[key] == nil {
		return "0"
	}

	_, ok := store.Sets[key][value]

	if ok {
		return "1"
	}

	return "0"
}
