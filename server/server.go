package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/thisisnitish/nimbusdb/parser"
	"github.com/thisisnitish/nimbusdb/store"
)

func CreateConnection() net.Listener {
	const PORT string = ":6379"
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error in listerning...")
		panic(err)
	}

	fmt.Println("Server is listening on ", PORT)

	return listener
}

func HandleConnection(conn net.Conn, store *store.Store) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Split(input, " ")

		if len(parts) < 2 {
			fmt.Println("Invalid Command")
			continue
		}

		command := parts[0]
		args := parts[1:]

		response := parser.ParseCommands(command, args, store)

		fmt.Fprintln(conn, response)
	}

}
