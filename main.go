package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"os"
	"time"

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
	reqTimeout := 5 * time.Second

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Error.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(writer, state, reqTimeout, method, content)
	}

	if err := scanner.Err(); err != nil {
		logger.Error.Printf("scanner error: %s", err)
	}
}

// handleWithTimeout runs fn with a request timeout. If fn exceeds timeout,
// it returns a cancelled error response so the server can keep processing messages.
// The handler goroutine may continue running until fn returns because Go cannot
// cancel it, but its result is discarded through the buffered channel.
func handleWithTimeout(reqID int, reqTimeout time.Duration, fn func(ctx context.Context) any) any {
	ctx, cancel := context.WithTimeout(context.Background(), reqTimeout)
	defer cancel()

	ch := make(chan any, 1)
	go func() {
		ch <- fn(ctx)
	}()

	select {
	case resp := <-ch:
		return resp
	case <-ctx.Done():
		logger.Error.Printf("request %d timed out after %s", reqID, reqTimeout)
		return lsp.NewErrorResponse(reqID, lsp.ErrCodeRequestCancelled, "request timed out")
	}
}

func handleMessage(
	writer io.Writer,
	state *analysis.State,
	reqTimeout time.Duration,
	method string,
	content []byte,
) {
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

		state.SetDocument(req.Params.TextDocument.URI, req.Params.TextDocument.Text)
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

		resp := handleWithTimeout(req.ID, reqTimeout, func(_ context.Context) any {
			resp, err := state.Hover(req.ID, req.Params.TextDocument.URI, req.Params.Position)
			if err != nil {
				logger.Error.Printf("textDocument/hover error: %s", err)
			}

			return resp
		})

		writeResponse(writer, resp)
	case "textDocument/completion":
		var req lsp.CompletionRequest
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Error.Printf("textDocument/completion error: %s", err)
			return
		}

		resp := handleWithTimeout(req.ID, reqTimeout, func(_ context.Context) any {
			resp, err := state.Completion(req.ID, req.Params.TextDocument.URI, req.Params.Position)
			if err != nil {
				logger.Error.Printf("textDocument/completion error: %s", err)
			}

			return resp
		})

		writeResponse(writer, resp)
	case "textDocument/didClose":
		var notification lsp.DidCloseTextDocumentNotification
		if err := json.Unmarshal(content, &notification); err != nil {
			logger.Error.Printf("textDocument/didClose error: %s", err)
			return
		}

		state.CloseDocument(notification.Params.TextDocument.URI)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	_, err := writer.Write([]byte(reply))

	if err != nil {
		logger.Error.Printf("writeResponse error: %s", err)
	}
}
