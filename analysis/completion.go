package analysis

import (
	"regexp"
	"slices"
	"strings"

	"github.com/textwire/lsp/internal/logger"
	"github.com/textwire/lsp/lsp"
	twLsp "github.com/textwire/textwire/v3/pkg/lsp"
	"github.com/textwire/textwire/v3/pkg/lsp/completions"
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
		completionItems, err = handleDirectivesAutocomplete(doc, pos, uri)
	} else if strings.HasSuffix(textBeforeCursor, "loop.") {
		completionItems, err = handleLoopObjectAutocomplete(doc, pos, uri)
	}

	if err != nil {
		logger.Error.Println(err)
		return lsp.CompletionResponse{}, err
	}

	return s.completionResponse(id, s.makeCompletions(completionItems)), nil
}

func handleDirectivesAutocomplete(doc string, pos lsp.Position, uri string) ([]completions.Completion, error) {
	dirs, err := completions.GetDirectives(locale)
	if err != nil {
		return []completions.Completion{}, err
	}

	isInsideLoop, errors := twLsp.IsInLoop(doc, uri, pos.Line, pos.Character)
	if len(errors) > 0 {
		logger.Error.Println(errors[0])
	}

	if isInsideLoop {
		return dirs, nil
	}

	filteredDirs := make([]completions.Completion, 0, len(dirs))
	loopDirs := []string{"@break", "@breakIf", "@continue", "@continueIf"}

	for _, dir := range dirs {
		if slices.Contains(loopDirs, dir.Label) {
			continue
		}

		filteredDirs = append(filteredDirs, dir)
	}

	return filteredDirs, nil
}

func handleLoopObjectAutocomplete(doc string, pos lsp.Position, uri string) ([]completions.Completion, error) {
	doc = removeTrailingChar(doc, pos.Line, pos.Character, '.')

	isInsideLoop, errors := twLsp.IsInLoop(doc, uri, pos.Line, pos.Character)
	if len(errors) > 0 {
		logger.Error.Println(errors[0])
	}

	if !isInsideLoop {
		return []completions.Completion{}, nil
	}

	fields, err := completions.GetLoopObjFields(locale)
	if err != nil {
		return []completions.Completion{}, err
	}

	return fields, nil
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
			Label:            item.Label,
			FilterText:       item.InsertText,
			InsertText:       item.InsertText,
			InsertTextFormat: item.InsertTextFormat,
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
