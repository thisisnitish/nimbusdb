package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

// TODO: Do error handling
func IncrBy(key string, incr string, store *store.Store, writeToFile ...bool) string {
	if len(writeToFile) > 0 && writeToFile[0] {
		command := "INCRBY " + key + " " + incr
		utils.AppendToAOF("file.aof", command)
	}

	num, _ := strconv.Atoi(incr)
	value, _ := strconv.Atoi(store.Get(key))

	value += num
	store.Set(key, strconv.Itoa(value))

	return strconv.Itoa(value)
}
