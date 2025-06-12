package analysis

import (
	"regexp"
	"strings"

	"github.com/textwire/lsp/internal/logger"
	"github.com/textwire/lsp/lsp"
	twLsp "github.com/textwire/textwire/v2/lsp"
	"github.com/textwire/textwire/v2/lsp/completions"
)

const locale = "en"

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

	var completionItems []completions.Completion
	var err error

	if directiveMatch != nil {
		completionItems, err = completions.GetDirectives(locale)
	} else if strings.HasSuffix(textBeforeCursor, "loop.") {
		doc = removeTrailingChar(doc, pos.Line, pos.Character, '.')

		isInsideLoop, errors := twLsp.IsInLoop(doc, uri, pos.Line, pos.Character)
		if len(errors) > 0 {
			logger.Error.Println(errors[0])
		}

		if isInsideLoop {
			completionItems, err = completions.GetLoopObjFields(locale)
		}
	}

	if err != nil {
		logger.Error.Println(err)
		return lsp.CompletionResponse{}, err
	}

	return s.completionResponse(id, s.makeCompletions(completionItems)), nil
}

func removeTrailingChar(doc string, line, col uint, char byte) string {
	lines := strings.Split(doc, "\n")
	if int(line) >= len(lines) {
		return doc
	}

	currentLine := lines[line]
	if col > 0 && currentLine[col-1] == char {
		currentLine = currentLine[:col-1] + currentLine[col:]
		lines[line] = currentLine
	}

	return strings.Join(lines, "\n")
}

func (s *State) makeCompletions(completionItems []completions.Completion) []lsp.CompletionItem {
	items := make([]lsp.CompletionItem, 0, len(completionItems))

	for _, item := range completionItems {
		items = append(items, lsp.CompletionItem{
			Label:      item.Label,
			FilterText: item.Insert,
			InsertText: item.Insert,
			Documentation: lsp.MarkupContent{
				Kind:  "markdown",
				Value: item.Documentation,
			},
			LabelDetails: &lsp.CompletionItemLabelDetails{
				Kind: lsp.CIKSnippet,
			},
		})
	}

	return items
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
