package ChatGLM_sdk

import (
	"ChatGLM_sdk/model"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type MessageContext []model.Message

type Result struct {
	Tokens  int32
	Message []model.Message
	Error   error
}

func (client Client) Send(context MessageContext, text string) Result {
	message := model.ChatGLMRequest{
		Model: client.Model,
		Messages: append(context, model.Message{
			Role:    "user",
			Content: text,
		}),
		DoSample:    client.DoSample,
		Stream:      false,
		Temperature: client.Temperature,
		TopP:        client.TopP,
		ResponseFormat: model.Format{
			Type: client.ResponseFormat,
		},
	}

	if client.MaxTokens != 0 {
		message.MaxTokens = client.MaxTokens
	}

	if client.Stop != "" {
		message.Stop = []string{client.Stop}
	}

	if client.UserID != "" {
		message.UserID = client.UserID
	}

	reqBody, err := json.Marshal(message)
	if err != nil {
		return Result{
			Tokens:  0,
			Message: nil,
			Error:   err,
		}
	}

	req, err := http.NewRequest("POST", client.Url, bytes.NewBuffer(reqBody))
	if err != nil {
		return Result{
			Tokens:  0,
			Message: nil,
			Error:   err,
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.APIKey)

	c := &http.Client{}
	respBody, err := c.Do(req)
	if err != nil {
		return Result{
			Tokens:  0,
			Message: nil,
			Error:   err,
		}
	}

	body, _ := io.ReadAll(respBody.Body)

	err = respBody.Body.Close()
	if err != nil {
		return Result{
			Tokens:  0,
			Message: nil,
			Error:   err,
		}
	}

	var resp model.ChatGLMResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return Result{
			Tokens:  0,
			Message: nil,
			Error:   err,
		}
	}

	var ms []model.Message

	stopErr := error(nil)

	for _, c := range resp.Choices {
		ms = append(ms, c.Message)
		if c.FinishReason != "stop" {
			stopErr = errors.New(c.FinishReason)
		}
	}

	return Result{
		Tokens:  int32(resp.Usage.TotalTokens),
		Message: ms,
		Error:   stopErr,
	}
}
