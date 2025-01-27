package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func LPop(key string, store *store.Store, writeToFile ...bool) string {
	if store.List[key] == nil {
		return "List doesn't exists"
	}

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "LPOP " + key
		utils.AppendToAOF("file.aof", command)
	}

	value := store.List[key][0]
	store.List[key] = store.List[key][1:]

	return value
}
