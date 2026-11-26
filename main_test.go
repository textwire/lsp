package main

import (
	"context"
	"testing"
	"time"

	"github.com/textwire/lsp/lsp"
)

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
