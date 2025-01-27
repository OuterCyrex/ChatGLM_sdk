package model

type ChatGLMAsyncInfo struct {
	RequestId  string `json:"request_id"`
	ID         string `json:"id"`
	Model      string `json:"model"`
	TaskStatus string `json:"task_status"`
}

type ChatGLMAsyncResponse struct {
	ID         string   `json:"id"`
	RequestID  string   `json:"request_id"`
	Model      string   `json:"model"`
	TaskStatus string   `json:"task_status"`
	Choices    []Choice `json:"choices"`
	Usage      Usage    `json:"usage"`
}
