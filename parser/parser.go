package parser

import (
	"fmt"
	"strings"

	"github.com/thisisnitish/nimbusdb/cmd"
	"github.com/thisisnitish/nimbusdb/store"
)

func ParseCommands(command string, args []string, store *store.Store) string {
	fmt.Println("cmd: ", command, "agrs: ", args)

	command = strings.ToUpper(command)
	switch command {
	case "SET":
		return cmd.Set(args[0], strings.Join(args[1:], " "), store)
	case "GET":
		return cmd.Get(args[0], store)
	case "DELETE":
		return cmd.Delete(args[0], store)
	default:
		return "ERR: Unknown command"
	}
}
