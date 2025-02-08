package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

const headerLabel = "Content-Length: "
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
		// not ready to execute a Split yet
		return 0, nil, nil
	}

	headerLabelLen := len(headerLabel)
	contentLenBytes := header[headerLabelLen:]
	contentLen, err := strconv.Atoi(string(contentLenBytes))
	if err != nil {
		// returns err because doesn't know what do with the content
		return 0, nil, err
	}

	if !isContentFull(content, contentLen) {
		return 0, nil, nil
	}

	totalLen := len(header) + len(contentSeparator) + contentLen

	return totalLen, data[:totalLen], nil
}

func isContentFull(content []byte, contentLen int) bool {
	return len(content) < contentLen
}
