package ChatGLM_sdk

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OuterCyrex/ChatGLM_sdk/model"
	"io"
	"net/http"
)

// SendStream allows developers to communicate with GLM with the given
// dialog context and text. The return value of this API is a channel
// that continuously returns streaming results
func (client Client) SendStream(context *MessageContext, text string) <-chan Result {
	messageChannel := make(chan Result)

	*context = append(*context, model.Message{
		Role:    "user",
		Content: text,
	})

	var assistantResp string

	message := model.ChatGLMRequest{
		Model:       client.model,
		Messages:    *context,
		DoSample:    client.doSample,
		Stream:      true,
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
		messageChannel <- Result{
			Tokens:  0,
			Message: nil,
			Error:   err,
		}
		close(messageChannel)
		return messageChannel
	}

	req, err := http.NewRequest("POST", SyncUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		messageChannel <- Result{
			Tokens:  0,
			Message: nil,
			Error:   err,
		}
		close(messageChannel)
		return messageChannel
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.apiKey)

	go func() {
		respBody, err := http.DefaultClient.Do(req)
		if err != nil {
			messageChannel <- Result{
				Tokens:  0,
				Message: nil,
				Error:   err,
			}
			close(messageChannel)
			return
		}

		defer respBody.Body.Close()

		if respBody.StatusCode >= 400 {
			var errResp model.ErrorResponse
			body, _ := io.ReadAll(respBody.Body)
			err = json.Unmarshal(body, &errResp)

			if err != nil {
				messageChannel <- Result{
					Tokens:  0,
					Message: nil,
					Error:   err,
				}
				close(messageChannel)
				return
			}

			messageChannel <- Result{
				Tokens:  0,
				Message: nil,
				Error:   errors.New(errResp.Error.Message),
			}
			close(messageChannel)
			return
		}

		scanner := bufio.NewScanner(respBody.Body)
		for scanner.Scan() {
			line := scanner.Text()

			if len(line) > 0 {
				var resp model.ChatGLMStreamResponse
				if err = json.Unmarshal([]byte(line[6:]), &resp); err != nil {
					messageChannel <- Result{
						Tokens:  0,
						Message: nil,
						Error:   fmt.Errorf("无法解析JSON文件: %v. Line: %s", err, line),
					}
					close(messageChannel)
					return
				}

				for _, v := range resp.Choices {
					if v.FinishReason == "" {
						messageChannel <- Result{
							Tokens: 0,
							Message: []model.Message{
								{
									Role:    v.Delta.Role,
									Content: v.Delta.Content,
								}},
							Error: nil,
						}
						assistantResp += v.Delta.Content
					} else {
						messageChannel <- Result{
							Tokens:  int32(resp.Usage.TotalTokens),
							Message: nil,
							Error:   nil,
						}
						*context = append(*context, model.Message{
							Role:    v.Delta.Role,
							Content: assistantResp,
						})
						close(messageChannel)
						return
					}
				}
			}

			if err := scanner.Err(); err != nil {
				messageChannel <- Result{
					Tokens:  0,
					Message: nil,
					Error:   err,
				}
				close(messageChannel)
				return
			}
		}
	}()

	return messageChannel
}
