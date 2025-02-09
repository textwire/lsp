package main

import (
	"bufio"
	"log"
	"os"

	"github.com/textwire/lsp/rpc"
)

func main() {
	logger := getLogger("/Users/serhiichornenkyi/www/open/textwire/lsp/log.txt")
	logger.Println("I've started!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, _ []byte) {
	logger.Printf("Got a message with the method: %s", method)
}

func getLogger(filename string) *log.Logger {
	const fileMode = 0666

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileMode)
	if err != nil {
		log.Panicf("The filepath %s is missing a file", filename)
	}

	return log.New(logfile, "[textwire lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
