package parser

import (
	"fmt"
	"net"
	"strings"

	"github.com/thisisnitish/nimbusdb/cmd"
	"github.com/thisisnitish/nimbusdb/store"
)

func ParseCommands(command string, args []string, store *store.Store, conn net.Conn, writeToFile ...bool) string {
	fmt.Println("cmd: ", command, "agrs: ", args)

	command = strings.ToUpper(command)
	switch command {
	case "SET":
		return cmd.Set(args[0], strings.Join(args[1:], " "), store, writeToFile...)
	case "GET":
		return cmd.Get(args[0], store, writeToFile...)
	case "DEL":
		return cmd.Delete(args[0], store, writeToFile...)
	case "INCR":
		return cmd.Incr(args[0], store, writeToFile...)
	case "DECR":
		return cmd.Decr(args[0], store, writeToFile...)
	case "INCRBY":
		return cmd.IncrBy(args[0], args[1], store, writeToFile...)
	case "DECRBY":
		return cmd.DecrBy(args[0], args[1], store, writeToFile...)
	case "LPUSH":
		return cmd.LPush(args[0], args[1], store, writeToFile...)
	case "RPUSH":
		return cmd.RPush(args[0], args[1], store, writeToFile...)
	case "LPOP":
		return cmd.LPop(args[0], store, writeToFile...)
	case "RPOP":
		return cmd.RPop(args[0], store, writeToFile...)
	case "LINDEX":
		return cmd.LIndex(args[0], args[1], store, writeToFile...)
	case "LLEN":
		return cmd.LLen(args[0], store, writeToFile...)
	case "SADD":
		return cmd.SAdd(args[0], args[1], store, writeToFile...)
	case "SREM":
		return cmd.SRem(args[0], args[1], store, writeToFile...)
	case "SMEMBERS":
		return cmd.SMembers(args[0], store, writeToFile...)
	case "SISMEMBER":
		return cmd.SIsMember(args[0], args[1], store, writeToFile...)
	case "SUBSCRIBE":
		if conn != nil {
			return cmd.Subscribe(args[0], conn, store)
		}
		return "Pub/Sub has no support for logs."
	case "PUBLISH":
		return cmd.Publish(args[0], strings.Join(args[1:], " "), store)
	default:
		return "ERR: Unknown command"
	}
}
