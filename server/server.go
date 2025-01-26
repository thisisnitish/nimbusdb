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

	for {
		// Display the prompt to the user
		fmt.Fprint(conn, "nimbusdb > ")
		// Read the user's input
		if !scanner.Scan() {
			// If scanning fails (e.g., client disconnects or sends EOF), break out of the loop
			break
		}

		input := scanner.Text()
		parts := strings.Split(input, " ")

		if len(parts) < 2 {
			// fmt.Println("Invalid Command")
			fmt.Fprintln(conn, "Invalid Command")
			continue
		}

		command := parts[0]
		args := parts[1:]

		response := parser.ParseCommands(command, args, store, conn, true)

		fmt.Fprintln(conn, response)
	}

	// Handle any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(conn, "Error reading input:", err)
	}
}
