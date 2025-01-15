package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
)

func RPop(key string, store *store.Store) string {
	if store.List[key] == nil {
		return "List doesn't exists"
	}

	n := len(store.List[key])
	value := store.List[key][n-1]
	store.List[key] = store.List[key][:n-1]

	return value
}
