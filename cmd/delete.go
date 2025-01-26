package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func Delete(key string, store *store.Store, writeToFile ...bool) string {
	// store := store.Store{}

	response := store.Delete(key)
	if len(writeToFile) > 0 && writeToFile[0] {
		command := "DEL " + key
		utils.AppendToAOF("file.aof", command)
	}
	return response
}
