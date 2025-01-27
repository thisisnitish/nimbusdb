package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func SIsMember(key string, value string, store *store.Store, writeToFile ...bool) string {

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "SISMEMBER " + key + " " + value
		utils.AppendToAOF("file.aof", command)
	}

	if store.Sets[key] == nil {
		return "0"
	}

	_, ok := store.Sets[key][value]

	if ok {
		return "1"
	}

	return "0"
}
