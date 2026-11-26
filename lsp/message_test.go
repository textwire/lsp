package lsp

import (
	"encoding/json"
	"testing"
)

func TestNewErrorResponse(t *testing.T) {
	resp := NewErrorResponse(3, ErrCodeRequestCancelled, "request timed out")

	if resp.RPC != JSONRPCVersion {
		t.Fatalf("expected jsonrpc %s, got %q", JSONRPCVersion, resp.RPC)
	}

	if resp.ID == nil || *resp.ID != 3 {
		t.Fatalf("expected id 3, got %v", resp.ID)
	}

	if resp.Error.Code != ErrCodeRequestCancelled {
		t.Fatalf("expected code %d, got %d", ErrCodeRequestCancelled, resp.Error.Code)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}

	var got map[string]any
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}

	if _, ok := got["error"]; !ok {
		t.Fatal("expected error field")
	}

	if _, ok := got["result"]; ok {
		t.Fatal("did not expect result field")
	}
}
