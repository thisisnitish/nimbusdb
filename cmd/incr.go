package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

// TODO: Do error handling
func Incr(key string, store *store.Store, writeToFile ...bool) string {
	value, _ := strconv.Atoi(store.Get(key))

	value += 1
	store.Set(key, strconv.Itoa(value))

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "INCR " + key
		utils.AppendToAOF("file.aof", command)
	}

	return strconv.Itoa(value)
}
