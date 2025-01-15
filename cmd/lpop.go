package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
)

func LPop(key string, store *store.Store) string {
	if store.List[key] == nil {
		return "List doesn't exists"
	}

	value := store.List[key][0]
	store.List[key] = store.List[key][1:]

	return value
}
