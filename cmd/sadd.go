package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
)

func SAdd(key string, value string, store *store.Store) string {
	if store.Sets[key] == nil {
		store.Sets[key] = make(map[string]bool)
	}

	if store.Sets[key][value] {
		return "0"
	}

	store.Sets[key][value] = true

	return "1"
}
