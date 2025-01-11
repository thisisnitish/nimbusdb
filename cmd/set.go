package cmd

import "github.com/thisisnitish/nimbusdb/store"

func Set(key string, value string, store *store.Store) string {
	// store := store.Store{}

	response := store.Set(key, value)
	return response
}
