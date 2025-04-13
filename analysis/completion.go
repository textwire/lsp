package analysis

import (
	"strings"

	"github.com/textwire/lsp/lsp"
)

func (s *State) Completion(id int, uri string, pos lsp.Position) (lsp.CompletionResponse, error) {
	// TODO: check if document exists
	doc := s.Documents[uri]

	lines := strings.Split(doc, "\n")
	if int(pos.Line) >= len(lines) {
		return s.completionResponse(id, nil), nil
	}

	line := lines[pos.Line]
	cursorPos := int(pos.Character)
	textBeforeCursor := line[:cursorPos]

	// If we're at the start of a directive (just typed @)
	if textBeforeCursor == "@" {
		return s.completionResponse(id, []lsp.CompletionItem{
			{
				Label: "use",
				LabelDetails: &lsp.CompletionItemLabelDetails{
					Description: "Some desc",
					Kind:        lsp.CIKSnippet,
				},
			},
			{
				Label: "if",
				LabelDetails: &lsp.CompletionItemLabelDetails{
					Description: "Some desc",
					Kind:        lsp.CIKSnippet,
				},
			},
		}), nil
	}

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
