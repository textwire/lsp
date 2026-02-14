package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/textwire/lsp/analysis"
	"github.com/textwire/lsp/internal/logger"
	"github.com/textwire/lsp/lsp"
	"github.com/textwire/lsp/rpc"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error.Println("Recovered from panic: ", r)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 4096), rpc.MaxContentLength)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Error.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(writer, state, method, content)
	}
}

func handleMessage(writer io.Writer, state analysis.State, method string, content []byte) {
	switch method {
	case "initialize":
		var req lsp.InitializeRequest
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Error.Printf("initialize error: %s", err)
			return
		}

		msg := lsp.NewInitializeResponse(req.ID)
		writeResponse(writer, msg)
	case "textDocument/didOpen":
		var req lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Error.Printf("textDocument/didOpen error: %s", err)
			return
		}

		state.OpenDocument(req.Params.TextDocument.URI, req.Params.TextDocument.Text)
	case "textDocument/didChange":
		var req lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Error.Printf("textDocument/didChange error: %s", err)
			return
		}

		for _, change := range req.Params.ContentChanges {
			state.UpdateDocument(req.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var req lsp.HoverRequest
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Error.Printf("textDocument/hover error: %s", err)
			return
		}

		resp, err := state.Hover(req.ID, req.Params.TextDocument.URI, req.Params.Position)
		if err != nil {
			logger.Error.Printf("textDocument/hover error: %s", err)
		}

		writeResponse(writer, resp)
	case "textDocument/completion":
		var req lsp.CompletionRequest
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Error.Printf("textDocument/completion error: %s", err)
			return
		}

		resp, err := state.Completion(req.ID, req.Params.TextDocument.URI, req.Params.Position)
		if err != nil {
			logger.Error.Printf("textDocument/completion error: %s", err)
		}

		writeResponse(writer, resp)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}
