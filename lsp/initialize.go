package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// todo: There is a lot to go here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
