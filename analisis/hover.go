package analisis

import (
	"github.com/textwire/lsp/lsp"
	"github.com/textwire/textwire/v2/lexer"
	twLsp "github.com/textwire/textwire/v2/lsp"
	"github.com/textwire/textwire/v2/token"
)

func (s *State) Hover(id int, uri string, pos lsp.Position) (lsp.HoverResponse, error) {
	// TODO: check if document exists
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
			break
		}
	}

	if matchingTok == nil {
		return s.hoverResponse(id, ""), nil
	}

	res, err := twLsp.GetTokenMeta(matchingTok.Type, "en")
	if err != nil {
		return s.hoverResponse(id, ""), err
	}

	return s.hoverResponse(id, res), nil
}

func (s *State) hoverResponse(id int, contents string) lsp.HoverResponse {
	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{Contents: contents},
	}
}
