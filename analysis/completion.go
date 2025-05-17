package analysis

import (
	"regexp"
	"strings"

	"github.com/textwire/lsp/lsp"
	"github.com/textwire/textwire/v2/lsp/completions"
)

func (s *State) Completion(id int, uri string, pos lsp.Position) (lsp.CompletionResponse, error) {
	doc, ok := s.Documents[uri]

	if !ok {
		return lsp.CompletionResponse{}, nil
	}

	lines := strings.Split(doc, "\n")
	if int(pos.Line) >= len(lines) {
		return s.completionResponse(id, nil), nil
	}

	line := lines[pos.Line]
	cursorPos := int(pos.Character)
	textBeforeCursor := line[:cursorPos]

	directiveRegex := regexp.MustCompile(`(^|[^\\])@(\w*)$`)
	directiveMatch := directiveRegex.FindStringSubmatch(textBeforeCursor)

	if directiveMatch == nil {
		return s.completionResponse(id, []lsp.CompletionItem{}), nil
	}

	directives, err := completions.GetDirectives("en")
	if err != nil {
		return lsp.CompletionResponse{}, err
	}

	items := make([]lsp.CompletionItem, 0, len(directives))

	for _, dir := range directives {
		items = append(items, lsp.CompletionItem{
			Label:      dir.Label,
			FilterText: dir.Insert,
			InsertText: dir.Insert,
			Documentation: lsp.MarkupContent{
				Kind:  "markdown",
				Value: dir.Documentation,
			},
			LabelDetails: &lsp.CompletionItemLabelDetails{
				Kind: lsp.CIKSnippet,
			},
		})
	}

	return s.completionResponse(id, items), nil
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
