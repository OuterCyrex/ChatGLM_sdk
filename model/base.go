package model

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AsyncMessage struct {
	ID         string `json:"id"`
	TaskStatus string `json:"task_status"`
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
