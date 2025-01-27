package cmd

import (
	"strconv"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func LPush(key string, value string, store *store.Store, writeToFile ...bool) string {
	if store.List[key] == nil {
		store.List[key] = make([]string, 0)
	}

	store.List[key] = append([]string{value}, store.List[key]...)

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "LPUSH " + key + " " + value
		utils.AppendToAOF("file.aof", command)
	}

	return strconv.Itoa(len(store.List[key]))
}
