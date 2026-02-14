package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

const (
	// MaxContentLength sets the maximum size we can process for a single file.
	// When file is larger, we stop the script to prevent infinite loop.
	MaxContentLength = 50 * 1024 * 1024 // 50MB
	headerLabel      = "Content-Length: "
)

var contentSeparator = []byte{'\r', '\n', '\r', '\n'}

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
	header, content, found := bytes.Cut(msg, contentSeparator)
	if !found {
		return "", nil, errors.New("Did not find a separator")
	}

	headerLabelLen := len(headerLabel)
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

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, contentSeparator)
	if !found {
		// Not ready to execute a Split yet
		return 0, nil, nil
	}

	headerLabelLen := len(headerLabel)
	contentLenBytes := header[headerLabelLen:]
	contentLen, err := strconv.Atoi(string(contentLenBytes))
	if err != nil {
		// Returns err because doesn't know what do with the content
		return 0, nil, err
	}

	// The file is too large, break the script
	if contentLen > MaxContentLength {
		return 0, nil, fmt.Errorf("content length %d exceeds maximum %d", contentLen, MaxContentLength)
	}

	if len(content) < contentLen {
		return 0, nil, nil
	}

	totalLen := len(header) + len(contentSeparator) + contentLen

	return totalLen, data[:totalLen], nil
}
