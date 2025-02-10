package lsp

type TextDocumentItem struct {
	// The text document's URI. Starts with file:///
	URI string `json:"uri"`

	// The text document's language identifier
	LanguageID string `json:"languageId"`

	// The version number of this document (it will
	// increase after each change, including undo/redo)
	Version int `json:"version"`

	// The content of the opened text document
	Text string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type VersionTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
