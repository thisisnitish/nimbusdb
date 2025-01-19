package cmd

import (
	"net"

	"github.com/thisisnitish/nimbusdb/store"
)

func Subscribe(channel string, conn net.Conn, st *store.Store) string {
	client := store.Client{Conn: conn}
	st.Subscribers[channel] = append(st.Subscribers[channel], client)

	return "SUBSCRIBED"
}
