package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
)

// TODO: Do error handling
func Decr(key string, store *store.Store) string {
	value, _ := strconv.Atoi(store.Get(key))

	value -= 1
	store.Set(key, strconv.Itoa(value))

	return strconv.Itoa(value)
}
