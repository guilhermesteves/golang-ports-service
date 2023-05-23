package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"guilhermesteves/golang-ports-service/client/pkg/models"
)

func StreamPorts(fileName string, stream chan models.Port) error {
	file, err := os.Open(fileName)
	if err != nil {
		printError("failed to open the file", err.Error())
		return err
	}

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	_, err = decoder.Token()
	if err != nil {
		printError("failed to read initial the token", err.Error())
		return err
	}

	var token json.Token
	for decoder.More() {
		token, err = decoder.Token()
		if err != nil {
			if err == io.EOF {
				return err
			}

			printError("failed to read code the token", err.Error())
			return err
		}

		var port = &models.Port{}
		err = decoder.Decode(&port)
		if err != nil {
			printError("failed to decode json", err.Error())
			return err
		}

		port.ID = fmt.Sprint(token)

		stream <- *port
	}
	return nil
}

func printError(message string, e string) {
	log.Printf("ERROR: %s", e)
	log.Printf("MSG: %s", message)
}
