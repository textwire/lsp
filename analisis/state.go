package analisis

import (
	"github.com/textwire/lsp/lsp"
	"github.com/textwire/textwire/v2/lexer"
	"github.com/textwire/textwire/v2/token"
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

func (s *State) Hover(id int, uri string, pos lsp.Position) (lsp.HoverResponse, error) {
	// todo: check if document exists
	doc := s.Documents[uri]

	l := lexer.New(doc)

	var matchingTok *token.Token

	for {
		tok := l.NextToken()

		if tok.Type == token.EOF {
			break
		}

		if tok.Pos.Contains(pos.Line, pos.Character) {
			matchingTok = &tok
		}
	}

	// todo: we don't need to parse the Document
	// we just need to find a matching token and
	// get the hover meta infromation from the token.
	// Hover meta informatin should be in a separate package
	// and should not be connected to the parser or lexer.
	// Because we don't want to add this meta information into
	// binary when users of the language compile their code.

	if matchingTok == nil {
		return s.response(id, ""), nil
	}

	return s.response(id, matchingTok.Literal), nil
}

func (s *State) response(id int, contents string) lsp.HoverResponse {
	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{Contents: contents},
	}
}
