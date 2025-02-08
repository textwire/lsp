package rpc

import "testing"

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	incomingMsg := "Content-Length: 16\r\n\r\n{\"Testing\":true}"

	contentLen, err := DecodeMessage([]byte(incomingMsg) )
	if err != nil {
		t.Fatal(err)
	}

	if contentLen != 16 {
		t.Fatalf("Expected: 16, Actual: %d", contentLen)
	}
}
