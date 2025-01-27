package cmd

import (
	"net"

	"github.com/thisisnitish/nimbusdb/store"
	"github.com/thisisnitish/nimbusdb/utils"
)

func Subscribe(channel string, conn net.Conn, st *store.Store, writeToFile ...bool) string {
	if len(writeToFile) > 0 && writeToFile[0] {
		command := "SUBSCRIBE " + channel
		utils.AppendToAOF("file.aof", command)
	}
	client := store.Client{Conn: conn}
	st.Subscribers[channel] = append(st.Subscribers[channel], client)

	return "SUBSCRIBED"
}
