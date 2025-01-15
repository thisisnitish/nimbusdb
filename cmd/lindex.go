package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
)

func LIndex(key string, index string, store *store.Store) string {
	idx, _ := strconv.Atoi(index)

	if store.List[key] == nil || idx >= len(store.List[key]) || idx < 0 {
		return "Invalid Index"
	}

	return store.List[key][idx]
}
