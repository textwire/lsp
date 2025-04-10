package lsp

type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

type CompletionParams struct {
	TextDocumentPositionParams
}

type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionResult struct {
	Contents string `json:"contents"`
}

type CompletionItem struct {
	// Label property is also by default the text that
	// is inserted when selecting this completion.
	//
	// If label details are provided the label itself should
	// be an unqualified name of the completion item.
	Label string `json:"label"`
}
