package utils

import "os"

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
