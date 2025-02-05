package rpc

import "testing"

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected %s, Actual: %s", expected, actual)
	}
}
