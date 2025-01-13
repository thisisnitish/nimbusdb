package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
)

// TODO: Do error handling
func IncrBy(key string, incr string, store *store.Store) string {
	num, _ := strconv.Atoi(incr)
	value, _ := strconv.Atoi(store.Get(key))

	value += num
	store.Set(key, strconv.Itoa(value))

	return strconv.Itoa(value)
}
