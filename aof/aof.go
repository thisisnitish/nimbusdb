package aof

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thisisnitish/nimbusdb/parser"
	"github.com/thisisnitish/nimbusdb/store"
)

func AppendToAOF(filename string, command string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(command + "\n")
	if err != nil {
		return err
	}

	return nil
}

func ReplayAOF(filename string, store *store.Store) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// If the AOF file doesn't exist, no commands to replay
			return nil
		}
		return err
	}
	defer file.Close()

	fmt.Println("file: ", file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Split(input, " ")
		if len(parts) < 2 {
			// fmt.Println("Invalid Command")
			fmt.Println("Corrupted AOF file")
			break
		}
		command := parts[0]
		args := parts[1:]
		fmt.Println("command: ", command, args)
		// TODO: Handle conn parameter here
		parser.ParseCommands(command, args, store, nil)
	}

	return scanner.Err()
}
