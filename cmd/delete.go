package cmd

import "github.com/thisisnitish/nimbusdb/store"

func Delete(key string, store *store.Store) string {
	// store := store.Store{}

	response := store.Delete(key)
	return response
}
