package model

type Response struct {
	Success    bool   `json:"success"`
	Data       any    `json:"data"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"status_code"`
}

type ErrorResponse struct {
	Success    bool   `json:"success"`
	Error      string `json:"message"`
	StatusCode int    `json:"status_code"`
}
