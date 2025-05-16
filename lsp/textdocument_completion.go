package lsp

type CompletionItemKind int

const (
	_ CompletionItemKind = iota
	CIKText
	CIKMethod
	CIKFunction
	CIKConstructor
	CIKField
	CIKVariable
	CIKClass
	CIKInterface
	CIKModule
	CIKProperty
	CIKUnit
	CIKValue
	CIKEnum
	CIKKeyword
	CIKSnippet
	CIKColor
	CIKFile
	CIKReference
	CIKFolder
	CIKEnumMember
	CIKConstant
	CIKStruct
	CIKEvent
	CIKOperator
	CIKTypeParameter
)

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

type CompletionItemLabelDetails struct {
	// Detail rendered less prominently directly after Label,
	// without any spacing. Should be used for function
	// signatures or type annotations.
	Detail string `json:"detail,omitempty"`

	// Description rendered less prominently after Detail.
	// Should be used for fully qualified names or file path.
	Description string `json:"description,omitempty"`

	// Kind of this completion item. Based of the kind
	// an icon is chosen by the editor. The standardized set
	// of available values is defined as`CompletionItemKind`
	// and all names of kinds start with CIK.
	Kind CompletionItemKind `json:"kind,omitempty"`
}

type CompletionItem struct {
	// Label property is also by default the text that
	// is inserted when selecting this completion.
	// If label details are provided the label itself should
	// be an unqualified name of the completion item.
	Label string `json:"label"`

	// A string that should be used when filtering a set of
	// completion items. When omitted the label is used as the
	// filter text for this item.
	FilterText string `json:"filterText,omitempty"`

	// LabelDetails is additional details for the label.
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails,omitempty"`

	// InsertText is a string that should be inserted into a document when selecting
	// this completion. When omitted the label is used as the insert text
	// for this item.
	InsertText string `json:"insertText,omitempty"`

	// Documentation is a human-readable string that represents a doc-comment.
	Documentation string `json:"documentation,omitempty"`
}

type CompletionOptions struct {
	// The additional characters, beyond the defaults provided by the client
	// (typically [a-zA-Z]), that should automatically trigger a completion
	// request. For example `.` in JavaScript represents the beginning of
	// an object property or method and is thus a good candidate for
	// triggering a completion request.
	//
	// Most tools trigger a completion request automatically without
	// explicitly requesting it using a keyboard shortcut (e.g. Ctrl+Space).
	// Typically they do so when the user starts to type an identifier.
	// For example if the user types `c` in a JavaScript file code complete
	// will automatically pop up present `console` besides others as a
	// completion item. Characters that make up identifiers don't need to
	// be listed here.
	TriggerCharacters []string `json:"triggerCharacters"`
}
