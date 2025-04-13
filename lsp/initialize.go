package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// TODO: There is a lot to go here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

// ServerCapabilities are features that a client and a server
// can negotiate to decide which ones to use
type ServerCapabilities struct {
	TextDocumentSync int `json:"textDocumentSync"`

	// Providers
	HoverProvider      bool              `json:"hoverProvider"`
	CompletionProvider CompletionOptions `json:"completionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
				HoverProvider:    true,
				CompletionProvider: CompletionOptions{
					TriggerCharacters: []string{".", "@", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
				},
			},
			ServerInfo: ServerInfo{
				Name:    "textwirelsp",
				Version: "1.0.0-beta1",
			},
		},
	}
}
