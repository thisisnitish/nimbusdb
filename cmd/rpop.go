package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func RPop(key string, store *store.Store, writeToFile ...bool) string {
	if store.List[key] == nil {
		return "List doesn't exists"
	}

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "RPOP " + key
		utils.AppendToAOF("file.aof", command)
	}

	n := len(store.List[key])
	value := store.List[key][n-1]
	store.List[key] = store.List[key][:n-1]

	return value
}
