package cmd

import "github.com/thisisnitish/nimbusdb/store"

func Get(key string) string {
	store := store.Store{}

	response := store.Get(key)
	return response
}
