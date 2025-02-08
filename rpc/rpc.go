package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (int, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, errors.New("Did not find a separator")
	}

	headerLabelLen := len("Content-Length: ")
	contentLenBytes := header[headerLabelLen:]
	contentLen, err := strconv.Atoi(string(contentLenBytes))
	if err != nil {
		return 0, err
	}

	_ = content

	return contentLen, nil
}
