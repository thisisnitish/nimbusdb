package cmd

import (
	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func Get(key string, store *store.Store, writeToFile ...bool) string {
	// store := store.Store{}

	response := store.Get(key)

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "GET " + key
		utils.AppendToAOF("file.aof", command)
	}
	// command := "GET " + key
	// utils.AppendToAOF("file.aof", command)
	return response
}
