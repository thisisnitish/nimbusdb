package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func LLen(key string, store *store.Store, writeToFile ...bool) string {
	if store.List[key] == nil {
		return "0"
	}

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "LLEN " + key
		utils.AppendToAOF("file.aof", command)
	}

	return strconv.Itoa(len(store.List[key]))
}
