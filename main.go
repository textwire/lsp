package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/textwire/lsp/lsp"
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

func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Received a method: `%s`", method)

	switch method {
	case "initialize":
		var req lsp.InitializeRequest
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("Couldn't parse this: %s", err)
			break
		}

		replyOnInitialize(req)
		logger.Println("Sent the initialize reply")
	case "textDocument/didOpen":
		var req lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("Couldn't parse this: %s", err)
		}

		logger.Printf("Opened: %s %s", req.Params.TextDocument.URI, req.Params.TextDocument.Text)
		logger.Println("Sent the didOpen reply")
	}
}

func replyOnInitialize(req lsp.InitializeRequest) {
	msg := lsp.NewInitializeResponse(req.ID)
	reply := rpc.EncodeMessage(msg)

	writer := os.Stdout
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	const fileMode = 0666

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileMode)
	if err != nil {
		log.Panicf("The filepath %s is missing a file", filename)
	}

	return log.New(logfile, "[textwire lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
