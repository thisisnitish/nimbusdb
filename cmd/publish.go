package cmd

import (
	"fmt"

	"github.com/thisisnitish/nimbusdb/store"
)

func Publish(channel string, message string, store *store.Store) string {

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
