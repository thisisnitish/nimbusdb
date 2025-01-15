package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
)

func RPush(key string, value string, store *store.Store) string {
	if store.List[key] == nil {
		store.List[key] = make([]string, 0)
	}

	store.List[key] = append(store.List[key], value)

	return strconv.Itoa(len(store.List[key]))
}
