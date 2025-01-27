package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func Set(key string, value string, store *store.Store, writeToFile ...bool) string {
	// store := store.Store{}

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "SET " + key + " " + value
		utils.AppendToAOF("file.aof", command)
	}
	response := store.Set(key, value)
	// command := "SET " + key + " " + value
	// utils.AppendToAOF("file.aof", command)
	return response
}
