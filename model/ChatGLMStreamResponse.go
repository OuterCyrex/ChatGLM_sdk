package model

type ChatGLMStreamResponse struct {
	ID      string         `json:"id"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []StreamChoice `json:"choices"`
	Usage   Usage          `json:"usage,omitempty"`
}

type Delta struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type StreamChoice struct {
	Index        int    `json:"index"`
	FinishReason string `json:"finish_reason"`
	Delta        Delta  `json:"delta"`
}
