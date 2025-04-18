package analysis

import (
	"strings"

	"github.com/textwire/lsp/lsp"
	"github.com/textwire/textwire/v2/lsp/completions"
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

	if textBeforeCursor == "@" {
		directives, err := completions.GetDirectives("en")
		if err != nil {
			return lsp.CompletionResponse{}, err
		}

		items := make([]lsp.CompletionItem, 0, len(directives))
		for _, dir := range directives {
			items = append(items, lsp.CompletionItem{
				Label:         dir.Label,
				InsertText:    dir.Insert,
				Documentation: dir.Documentation,
				LabelDetails: &lsp.CompletionItemLabelDetails{
					Kind: lsp.CIKSnippet,
				},
			})
		}

		return s.completionResponse(id, items), nil
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
