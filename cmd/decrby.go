package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

// TODO: Do error handling
func DecrBy(key string, decr string, store *store.Store, writeToFile ...bool) string {
	num, _ := strconv.Atoi(decr)
	value, _ := strconv.Atoi(store.Get(key))

	value -= num
	store.Set(key, strconv.Itoa(value))

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "DECRBY " + key + " " + decr
		utils.AppendToAOF("file.aof", command)
	}

	return strconv.Itoa(value)
}
