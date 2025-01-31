package ChatGLM_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/OuterCyrex/ChatGLM_sdk/model"
	"io"
	"net/http"
)

// SendAsync allows developers to async communicate with GLM with the given
// dialog context and text. the response is an ID of type string, as well as an error
//
// this API is async but not stream, if you are looking for stream response
// Please check SendStream API
func (client Client) SendAsync(context *MessageContext, text string) (string, error) {
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
		return "", ErrSdkInternal
	}

	req, err := http.NewRequest("POST", AsyncUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", ErrHttpRequestTimeOut
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.apiKey)

	respBody, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", ErrHttpRequestTimeOut
	}

	body, _ := io.ReadAll(respBody.Body)

	err = respBody.Body.Close()
	if err != nil {
		return "", ErrHttpRequestTimeOut
	}

	if respBody.StatusCode >= 400 {
		switch respBody.StatusCode {
		case 404:
			return "", ErrNotFound
		default:
			return "", fmt.Errorf("%w: %q", ErrHttpBadRequest, http.StatusText(respBody.StatusCode))
		}
	}

	var resp model.ChatGLMAsyncInfo

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", ErrSdkInternal
	}

	return resp.ID, nil
}

// GetAsyncMessage use the ID returned by SendAsync to get messages from GLM.
// the response is a Result struct, containing tokens consumed,
// message returned from GLM, as well as Error message.
//
// this API is sync but not stream, if you are looking for stream response
// Please check SendStream API
func (client Client) GetAsyncMessage(context *MessageContext, ID string) Result {

	req, err := http.NewRequest("GET", getAsyncUrl+ID, nil)
	if err != nil {
		return Result{
			Tokens:       0,
			Message:      nil,
			Error:        ErrHttpRequestTimeOut,
			FinishReason: "",
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.apiKey)

	respBody, err := http.DefaultClient.Do(req)
	if err != nil {
		return Result{
			Tokens:       0,
			Message:      nil,
			Error:        ErrHttpRequestTimeOut,
			FinishReason: "",
		}
	}

	body, _ := io.ReadAll(respBody.Body)

	err = respBody.Body.Close()
	if err != nil {
		return Result{
			Tokens:       0,
			Message:      nil,
			Error:        ErrHttpRequestTimeOut,
			FinishReason: "",
		}
	}

	if respBody.StatusCode >= 400 {
		switch respBody.StatusCode {
		case 404:
			return Result{
				Tokens:       0,
				Message:      nil,
				Error:        ErrNotFound,
				FinishReason: "",
			}
		default:
			return Result{
				Tokens:       0,
				Message:      nil,
				Error:        fmt.Errorf("%w: %q", ErrHttpBadRequest, http.StatusText(respBody.StatusCode)),
				FinishReason: "",
			}
		}
	}

	var resp model.ChatGLMAsyncResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return Result{
			Tokens:       0,
			Message:      nil,
			Error:        ErrSdkInternal,
			FinishReason: "",
		}
	}

	switch resp.TaskStatus {
	case "PROCESSING":
		return Result{
			Tokens:       0,
			Message:      nil,
			Error:        ErrResultProcessing,
			FinishReason: "",
		}
	case "FAIL":
		return Result{
			Tokens:       0,
			Message:      nil,
			Error:        ErrGenerateFailed,
			FinishReason: "",
		}
	}

	var ms []model.Message

	var finishReason string

	for _, c := range resp.Choices {
		ms = append(ms, c.Message)
		*context = append(*context, model.Message{
			Role:    c.Message.Role,
			Content: c.Message.Content,
		})
		finishReason = c.FinishReason
	}

	return Result{
		Tokens:       int32(resp.Usage.TotalTokens),
		Message:      ms,
		Error:        nil,
		FinishReason: finishReason,
	}
}
