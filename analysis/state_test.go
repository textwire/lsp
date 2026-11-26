package analysis

import "testing"

func TestStateDocument(t *testing.T) {
	s := NewState()

	s.SetDocument("file:///a.tw", "hello")
	got, ok := s.Document("file:///a.tw")
	if !ok || got != "hello" {
		t.Fatalf("expected hello, got %q, ok=%v", got, ok)
	}

	s.UpdateDocument("file:///a.tw", "world")
	got, ok = s.Document("file:///a.tw")
	if !ok || got != "world" {
		t.Fatalf("expected world, got %q, ok=%v", got, ok)
	}

	_, ok = s.Document("file:///missing.tw")
	if ok {
		t.Fatal("expected missing document")
	}
}
