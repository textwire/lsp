package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/textwire/lsp/analisis"
	"github.com/textwire/lsp/lsp"
	"github.com/textwire/lsp/rpc"
)

func main() {
	logger := getLogger("/Users/serhiichornenkyi/www/open/textwire/lsp/log.txt")
	logger.Println("I've started!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analisis.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(logger, state, method, content)
	}
}

func handleMessage(logger *log.Logger, state analisis.State, method string, content []byte) {
	logger.Printf("Received a method: `%s`", method)

	switch method {
	case "initialize":
		var req lsp.InitializeRequest
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("initialize error: %s", err)
			return
		}

		replyOnInitialize(req)
	case "textDocument/didOpen":
		var req lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("textDocument/didOpen error: %s", err)
			return
		}

		logger.Printf("Opened: %s", req.Params.TextDocument.URI)

		state.OpenDocument(req.Params.TextDocument.URI, req.Params.TextDocument.Text)
	case "textDocument/didChange":
		var req lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("textDocument/didChange error: %s", err)
			return
		}

		logger.Printf("Changed: %s", req.Params.TextDocument.URI)

		for _, change := range req.Params.ContentChanges {
			state.UpdateDocument(req.Params.TextDocument.URI, change.Text)
		}
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
