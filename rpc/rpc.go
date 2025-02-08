package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Did not find a separator")
	}

	headerLabelLen := len("Content-Length: ")
	contentLenBytes := header[headerLabelLen:]
	contentLen, err := strconv.Atoi(string(contentLenBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMsg BaseMessage

	if err := json.Unmarshal(content[:contentLen], &baseMsg); err != nil {
		return "", nil, err
	}

	return baseMsg.Method, content[:contentLen], nil
}
