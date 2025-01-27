package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func RPush(key string, value string, store *store.Store, writeToFile ...bool) string {
	if len(writeToFile) > 0 && writeToFile[0] {
		command := "RPUSH " + key + " " + value
		utils.AppendToAOF("file.aof", command)

	}

	if store.List[key] == nil {
		store.List[key] = make([]string, 0)
	}

	store.List[key] = append(store.List[key], value)

	return strconv.Itoa(len(store.List[key]))
}
