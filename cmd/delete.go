package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func Delete(key string, store *store.Store, writeToFile ...bool) string {
	// store := store.Store{}

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "DEL " + key
		utils.AppendToAOF("file.aof", command)
	}

	response := store.Delete(key)
	return response
}
