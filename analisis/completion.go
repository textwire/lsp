package analisis

import (
	"github.com/textwire/lsp/internal/logger"
	"github.com/textwire/lsp/lsp"
	"github.com/textwire/textwire/v2/lexer"
	"github.com/textwire/textwire/v2/token"
)

func (s *State) Completion(id int, uri string, pos lsp.Position) (lsp.CompletionResponse, error) {
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
		return s.completionResponse(id, []lsp.CompletionItem{}), nil
	}

	logger.Info.Println("Matching token: ", matchingTok.Literal)

	return s.completionResponse(id, []lsp.CompletionItem{}), nil
}

func (s *State) completionResponse(id int, items []lsp.CompletionItem) lsp.CompletionResponse {
	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
}
