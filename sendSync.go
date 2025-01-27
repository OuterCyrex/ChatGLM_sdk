package ChatGLM_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OuterCyrex/ChatGLM_sdk/model"
	"io"
	"net/http"
)

// SendSync allows developers to communicate with GLM with the given
// dialog context and text. the response is a Result struct, containing
// tokens consumed, message returned from GLM, as well as Error message.
//
// this API is sync but not stream, if you are looking for stream response
// Please check SendStream API
func (client Client) SendSync(context *MessageContext, text string) Result {
	*context = append(*context, model.Message{
		Role:    "user",
		Content: text,
	})

	message := model.ChatGLMRequest{
		Model:       client.model,
		Messages:    *context,
		DoSample:    client.doSample,
		Stream:      false,
		Temperature: client.temperature,
		TopP:        client.topP,
		ResponseFormat: model.Format{
			Type: client.responseFormat,
		},
	}

	if client.maxTokens != 0 {
		message.MaxTokens = client.maxTokens
	}

	if client.stop != "" {
		message.Stop = []string{client.stop}
	}

	if client.userID != "" {
		message.UserID = client.userID
	}

	reqBody, err := json.Marshal(message)
	if err != nil {
		return Result{
			Tokens:  0,
			Message: nil,
			Error:   err,
		}
	}

	req, err := http.NewRequest("POST", SyncUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return Result{
			Tokens:  0,
			Message: nil,
			Error:   err,
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.apiKey)

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

	if respBody.StatusCode >= 400 {
		var errResp model.ErrorResponse
		err = json.Unmarshal(body, &errResp)

		if err != nil {
			return Result{
				Tokens:  0,
				Message: nil,
				Error:   err,
			}
		}

		return Result{
			Tokens:  0,
			Message: nil,
			Error:   errors.New(errResp.Error.Message),
		}
	}

	var resp model.ChatGLMSyncResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return Result{
			Tokens:  0,
			Message: nil,
			Error:   fmt.Errorf("无法解析JSON文件: %v", err),
		}
	}

	var ms []model.Message

	stopErr := error(nil)

	for _, c := range resp.Choices {
		ms = append(ms, c.Message)
		*context = append(*context, model.Message{
			Role:    c.Message.Role,
			Content: c.Message.Content,
		})
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
