package cmd

import "github.com/thisisnitish/nimbusdb/store"

func Set(key string, value string) string {
	store := store.Store{}

	response := store.Set(key, value)
	return response
}
