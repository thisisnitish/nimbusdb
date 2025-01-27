package cmd

import (
	"strings"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func SMembers(key string, store *store.Store, writeToFile ...bool) string {

	if len(writeToFile) > 0 && writeToFile[0] {
		command := "SMEMBERS " + key
		utils.AppendToAOF("file.aof", command)
	}

	if store.Sets[key] == nil {
		return "[]"
	}

	members := make([]string, 0, len(store.Sets[key]))

	for member := range store.Sets[key] {
		members = append(members, member)
	}

	result := ""
	for _, member := range members {
		result += " " + member + " "
	}

	return "[" + strings.TrimSpace(result) + "]"

}
