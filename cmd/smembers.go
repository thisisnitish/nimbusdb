package cmd

import (
	"strings"

	"github.com/thisisnitish/nimbusdb/store"
)

func SMembers(key string, store *store.Store) string {
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
