package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func SAdd(key string, value string, store *store.Store, writeToFile ...bool) string {
	if len(writeToFile) > 0 && writeToFile[0] {
		command := "SADD " + key + " " + value
		utils.AppendToAOF("file.aof", command)
	}

	if store.Sets[key] == nil {
		store.Sets[key] = make(map[string]bool)
	}

	if store.Sets[key][value] {
		return "0"
	}

	store.Sets[key][value] = true

	return "1"
}
