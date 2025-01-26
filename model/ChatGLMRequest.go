package model

type ChatGLMRequest struct {
	Model          string    `json:"model"`
	Messages       []Message `json:"messages"`
	RequestID      string    `json:"request_id,omitempty"`
	DoSample       bool      `json:"do_sample,omitempty"`
	Stream         bool      `json:"stream,omitempty"`
	Temperature    float64   `json:"temperature,omitempty"`
	TopP           float64   `json:"top_p,omitempty"`
	MaxTokens      int       `json:"max_tokens,omitempty"`
	ResponseFormat Format    `json:"response_format,omitempty"`
	Stop           []string  `json:"stop,omitempty"`
	Tools          []string  `json:"tools,omitempty"`
	ToolChoice     string    `json:"tool_choice,omitempty"`
	UserID         string    `json:"user_id,omitempty"`
}

type Format struct {
	Type string `json:"type"`
}
