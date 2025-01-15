package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
)

func LPush(key string, value string, store *store.Store) string {
	if store.List[key] == nil {
		store.List[key] = make([]string, 0)
	}

	store.List[key] = append([]string{value}, store.List[key]...)

	return strconv.Itoa(len(store.List[key]))
}
