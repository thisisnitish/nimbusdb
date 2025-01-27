package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func SRem(key string, value string, store *store.Store, writeToFile ...bool) string {
	if len(writeToFile) > 0 && writeToFile[0] {
		command := "SREM " + key + " " + value
		utils.AppendToAOF("file.aof", command)
	}

	if store.Sets[key] == nil {
		return "0"
	}

	if !store.Sets[key][value] {
		return "0"
	}

	delete(store.Sets[key], value) // delete the value from the set

	return "1"
}
