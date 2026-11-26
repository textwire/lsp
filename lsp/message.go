package lsp

const (
	JSONRPCVersion = "2.0"

	ErrCodeRequestCancelled = -32800
	ErrCodeInternalError    = -32603
)

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Response
	Error ResponseError `json:"error"`
}

func NewErrorResponse(id int, code int, msg string) ErrorResponse {
	return ErrorResponse{
		Response: Response{
			RPC: JSONRPCVersion,
			ID:  &id,
		},
		Error: ResponseError{
			Code:    code,
			Message: msg,
		},
	}
}
