package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func Get(key string, store *store.Store, writeToFile ...bool) string {
	// store := store.Store{}

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "GET " + key
		utils.AppendToAOF("file.aof", command)
	}

	response := store.Get(key)

	// command := "GET " + key
	// utils.AppendToAOF("file.aof", command)
	return response
}
