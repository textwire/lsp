package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/textwire/lsp/analisis"
	"github.com/textwire/lsp/internal/logger"
	"github.com/textwire/lsp/lsp"
	"github.com/textwire/lsp/rpc"
)

func main() {
	logger.Info.Println("Textwire LSP server is running...")

	defer func() {
		if r := recover(); r != nil {
			logger.Info.Println("Recovered from panic: ", r)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analisis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Info.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(writer, state, method, content)
	}
}

func handleMessage(writer io.Writer, state analisis.State, method string, content []byte) {
	switch method {
	case "initialize":
		var req lsp.InitializeRequest
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Info.Printf("initialize error: %s", err)
			return
		}

		msg := lsp.NewInitializeResponse(req.ID)
		writeResponse(writer, msg)
	case "textDocument/didOpen":
		var req lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Info.Printf("textDocument/didOpen error: %s", err)
			return
		}

		logger.Info.Printf("Opened: %s", req.Params.TextDocument.URI)

		state.OpenDocument(req.Params.TextDocument.URI, req.Params.TextDocument.Text)
	case "textDocument/didChange":
		var req lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Info.Printf("textDocument/didChange error: %s", err)
			return
		}

		logger.Info.Printf("Changed: %s", req.Params.TextDocument.URI)

		for _, change := range req.Params.ContentChanges {
			state.UpdateDocument(req.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var req lsp.HoverRequest
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Info.Printf("textDocument/hover error: %s", err)
			return
		}

		resp, err := state.Hover(req.ID, req.Params.TextDocument.URI, req.Params.Position)
		if err != nil {
			logger.Info.Printf("textDocument/hover error: %s", err)
			return
		}

		logger.Info.Printf("Hover, giving response")
		writeResponse(writer, resp)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}
