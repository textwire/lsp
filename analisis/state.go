package analisis

import (
	"fmt"

	"github.com/textwire/lsp/lsp"
)

type State struct {
	// Documents is a map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(id int, uri string, pos lsp.Position) lsp.HoverResponse {
	// todo: look up the type through our parser

	// todo: check if document exists
	doc := s.Documents[uri]

	// todo: give the correct response
	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Chars: %d, Line: %d, Char: %d", uri, len(doc), pos.Line, pos.Character),
		},
	}
}
