package main

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/textwire/lsp/analysis"
	"github.com/textwire/lsp/lsp"
	"github.com/textwire/lsp/rpc"
	textwireLsp "github.com/textwire/textwire/v6/pkg/lsp"
	"github.com/textwire/textwire/v6/pkg/token"
)

func handleMessageResponse(
	t *testing.T,
	content []byte,
	method string,
	state *analysis.State,
	wantID int,
) json.RawMessage {
	t.Helper()

	var out bytes.Buffer
	handleMessage(&out, state, time.Second, method, content)

	_, responseContent, err := rpc.DecodeMessage(out.Bytes())
	if err != nil {
		t.Fatalf("decode response message: %v", err)
	}

	var response struct {
		lsp.Response
		Result json.RawMessage `json:"result"`
	}

	if err := json.Unmarshal(responseContent, &response); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}

	if response.RPC != lsp.JSONRPCVersion {
		t.Errorf("expected jsonrpc %q, got %q", lsp.JSONRPCVersion, response.RPC)
	}

	if response.ID == nil || *response.ID != wantID {
		t.Errorf("expected response id %d, got %v", wantID, response.ID)
	}

	if len(response.Result) == 0 {
		t.Error("expected response result")
	}

	return response.Result
}

func handleMessageNotification(t *testing.T, content []byte, method string, state *analysis.State) {
	t.Helper()

	var out bytes.Buffer
	handleMessage(&out, state, time.Second, method, content)

	if out.Len() != 0 {
		t.Errorf("expected no response for %s notification, got %q", method, out.String())
	}
}

func TestHandleMessage_InitializeResponse(t *testing.T) {
	content := []byte(`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`)
	response := handleMessageResponse(t, content, "initialize", analysis.NewState(), 1)

	var result lsp.InitializeResult
	if err := json.Unmarshal(response, &result); err != nil {
		t.Fatalf("unmarshal initialize result: %v", err)
	}

	if result.Capabilities.TextDocumentSync != 1 {
		t.Errorf("expected textDocumentSync 1, got %d", result.Capabilities.TextDocumentSync)
	}

	if !result.Capabilities.HoverProvider {
		t.Error("expected hover provider to be enabled")
	}
}

func TestHandleMessage_HoverResponse(t *testing.T) {
	content := []byte(
		`{"jsonrpc":"2.0","id":2,"method":"textDocument/hover","params":{"textDocument":{"uri":"file:///missing.tw"},"position":{"line":0,"character":0}}}`,
	)
	response := handleMessageResponse(t, content, "textDocument/hover", analysis.NewState(), 2)

	var result lsp.HoverResult
	if err := json.Unmarshal(response, &result); err != nil {
		t.Fatalf("unmarshal hover result: %v", err)
	}

	if result.Contents != "" {
		t.Errorf("expected empty hover contents, got %q", result.Contents)
	}
}

func TestHandleMessage_HoverResponseWithContents(t *testing.T) {
	state := analysis.NewState()
	state.SetDocument("file:///document.tw", "@if true")
	content := []byte(
		`{"jsonrpc":"2.0","id":4,"method":"textDocument/hover","params":{"textDocument":{"uri":"file:///document.tw"},"position":{"line":0,"character":1}}}`,
	)

	response := handleMessageResponse(t, content, "textDocument/hover", state, 4)

	var result lsp.HoverResult
	if err := json.Unmarshal(response, &result); err != nil {
		t.Fatalf("unmarshal hover result: %v", err)
	}

	want, err := textwireLsp.GetTokenMeta(token.IF, "en")
	if err != nil {
		t.Fatalf("get @if metadata: %v", err)
	}

	if result.Contents != want {
		t.Errorf("unexpected hover contents:\n got: %q\nwant: %q", result.Contents, want)
	}
}

func TestHandleMessage_CompletionResponse(t *testing.T) {
	state := analysis.NewState()
	state.SetDocument("file:///document.tw", "plain text")
	content := []byte(
		`{"jsonrpc":"2.0","id":3,"method":"textDocument/completion","params":{"textDocument":{"uri":"file:///document.tw"},"position":{"line":0,"character":1}}}`,
	)

	response := handleMessageResponse(t, content, "textDocument/completion", state, 3)

	var result []lsp.CompletionItem
	if err := json.Unmarshal(response, &result); err != nil {
		t.Fatalf("unmarshal completion result: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("expected no completion items, got %d", len(result))
	}
}

func TestHandleMessage_CompletionResponseWithItems(t *testing.T) {
	state := analysis.NewState()
	state.SetDocument("file:///document.tw", "@")
	content := []byte(
		`{"jsonrpc":"2.0","id":5,"method":"textDocument/completion","params":{"textDocument":{"uri":"file:///document.tw"},"position":{"line":0,"character":1}}}`,
	)

	response := handleMessageResponse(t, content, "textDocument/completion", state, 5)

	var result []lsp.CompletionItem
	if err := json.Unmarshal(response, &result); err != nil {
		t.Fatalf("unmarshal completion result: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("expected completion items after @")
	}

	for _, item := range result {
		if item.Label == "@if" {
			return
		}
	}

	t.Error("expected completion items to include @if")
}

func TestHandleMessage_DidOpen(t *testing.T) {
	state := analysis.NewState()
	content := []byte(
		`{"jsonrpc":"2.0","method":"textDocument/didOpen","params":{"textDocument":{"uri":"file:///document.tw","languageId":"textwire","version":1,"text":"opened"}}}`,
	)

	handleMessageNotification(t, content, "textDocument/didOpen", state)

	got, ok := state.Document("file:///document.tw")
	if !ok {
		t.Fatal("expected opened document to be stored")
	}

	if got != "opened" {
		t.Errorf("expected document text %q, got %q", "opened", got)
	}
}

func TestHandleMessage_DidChange(t *testing.T) {
	state := analysis.NewState()
	state.SetDocument("file:///document.tw", "before")
	content := []byte(
		`{"jsonrpc":"2.0","method":"textDocument/didChange","params":{"textDocument":{"uri":"file:///document.tw","version":2},"contentChanges":[{"text":"after"}]}}`,
	)

	handleMessageNotification(t, content, "textDocument/didChange", state)

	got, ok := state.Document("file:///document.tw")
	if !ok {
		t.Fatal("expected changed document to remain stored")
	}

	if got != "after" {
		t.Errorf("expected document text %q, got %q", "after", got)
	}
}

func TestHandleMessage_DidClose(t *testing.T) {
	state := analysis.NewState()
	state.SetDocument("file:///document.tw", "opened")
	content := []byte(
		`{"jsonrpc":"2.0","method":"textDocument/didClose","params":{"textDocument":{"uri":"file:///document.tw"}}}`,
	)

	handleMessageNotification(t, content, "textDocument/didClose", state)

	if _, ok := state.Document("file:///document.tw"); ok {
		t.Error("expected closed document to be removed")
	}
}

func TestHandleWithTimeout_ReturnsResponse(t *testing.T) {
	want := "ok"
	resp := handleWithTimeout(1, 100*time.Millisecond, func(_ context.Context) any {
		return want
	})

	got, ok := resp.(string)
	if !ok {
		t.Fatalf("expected string, got %T", resp)
	}

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestHandleWithTimeout_ReturnsErrorOnTimeout(t *testing.T) {
	resp := handleWithTimeout(7, 20*time.Millisecond, func(_ context.Context) any {
		time.Sleep(200 * time.Millisecond)
		return "late"
	})

	errResp, ok := resp.(lsp.ErrorResponse)
	if !ok {
		t.Fatalf("expected ErrorResponse, got %T", resp)
	}

	if errResp.ID == nil || *errResp.ID != 7 {
		t.Fatalf("expected id 7, got %v", errResp.ID)
	}

	if errResp.Error.Code != lsp.ErrCodeRequestCancelled {
		t.Fatalf("expected code %d, got %d", lsp.ErrCodeRequestCancelled, errResp.Error.Code)
	}

	if errResp.Error.Message == "" {
		t.Fatal("expected error message")
	}
}
