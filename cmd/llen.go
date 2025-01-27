package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func LLen(key string, store *store.Store, writeToFile ...bool) string {
	if len(writeToFile) > 0 && writeToFile[0] {
		command := "LLEN " + key
		utils.AppendToAOF("file.aof", command)
	}

	if store.List[key] == nil {
		return "0"
	}

	return strconv.Itoa(len(store.List[key]))
}
