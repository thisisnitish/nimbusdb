package cmd

import (
	"fmt"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func Publish(channel string, message string, store *store.Store, writeToFile ...bool) string {
	if len(writeToFile) > 0 && writeToFile[0] {
		command := "PUBLISH " + channel + " " + message
		utils.AppendToAOF("file.aof", command)
	}

	subscribers, ok := store.Subscribers[channel]

	if ok {
		for _, subscriber := range subscribers {
			fmt.Fprintf(subscriber.Conn, "+%s\n", message)
			fmt.Fprint(subscriber.Conn, "nimbusdb > ")
		}
		return "PUBLISHED"
	}

	return "NOT PUBLISHED"
}
