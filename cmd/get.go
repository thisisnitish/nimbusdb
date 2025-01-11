package cmd

import "github.com/thisisnitish/nimbusdb/store"

func Get(key string, store *store.Store) string {
	// store := store.Store{}

	response := store.Get(key)
	return response
}
