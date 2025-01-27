package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func LIndex(key string, index string, store *store.Store, writeToFile ...bool) string {
	if len(writeToFile) > 0 && writeToFile[0] {
		command := "LINDEX " + key + " " + index
		utils.AppendToAOF("file.aof", command)
	}

	idx, _ := strconv.Atoi(index)

	if store.List[key] == nil || idx >= len(store.List[key]) || idx < 0 {
		return "Invalid Index"
	}

	return store.List[key][idx]
}
