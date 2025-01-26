package model

type ChatGLMResponse struct {
	Choices   []Choice `json:"choices"`
	Created   int64    `json:"created"`
	ID        string   `json:"id"`
	Model     string   `json:"model"`
	RequestID string   `json:"request_id"`
	Usage     Usage    `json:"usage"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
	Message      Message `json:"message"`
}

type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
