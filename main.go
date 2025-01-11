package main

import (
	"fmt"

	"github.com/thisisnitish/nimbusdb/server"
	"github.com/thisisnitish/nimbusdb/store"
)

func main() {
	fmt.Println("Hello world, I'm nimbus db")

	listener := server.CreateConnection()

	store := store.NewStore()

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
			// fmt.Println("Error accepting connection:", err)
			// continue
		}

		go server.HandleConnection(conn, store)
	}

}
