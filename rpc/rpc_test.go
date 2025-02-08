package rpc

import (
	"fmt"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected string: %s, Actual: %s", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	expectLen := 15
	expectMethod := "hi"
	incomingMsg := fmt.Sprintf("Content-Length: %d\r\n\r\n{\"Method\":\"%s\"}", expectLen, expectMethod)

	method, content, err := DecodeMessage([]byte(incomingMsg) )
	if err != nil {
		t.Fatal(err)
	}

	if len(content) != expectLen {
		t.Fatalf("Expected content length: %d, Actual: %d", expectLen, len(content))
	}

	if method != expectMethod {
		t.Fatalf("Expected method name: %s, Actual: %s", method, expectMethod)
	}
}
