package cmd

import "github.com/thisisnitish/nimbusdb/store"

func SRem(key string, value string, store *store.Store) string {
	if store.Sets[key] == nil {
		return "0"
	}

	if !store.Sets[key][value] {
		return "0"
	}

	delete(store.Sets[key], value) // delete the value from the set

	return "1"
}
