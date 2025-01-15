package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
)

func LLen(key string, store *store.Store) string {
	if store.List[key] == nil {
		return "0"
	}

	return strconv.Itoa(len(store.List[key]))
}
