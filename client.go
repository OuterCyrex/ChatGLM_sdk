package ChatGLM_sdk

// Client contains basic parameter information for communicating with GLM
type Client struct {
	Url            string
	UserID         string
	Model          string
	APIKey         string
	DoSample       bool
	Temperature    float64
	TopP           float64
	MaxTokens      int
	ResponseFormat string
	Stop           string
}

const url = "https://open.bigmodel.cn/api/paas/v4/chat/completions"

type Option func(client *Client) *Client

// SetUserID allows developers to set the UserID information
// when calling NewClient
func SetUserID(userID string) Option {
	return func(client *Client) *Client {
		client.UserID = userID
		return client
	}
}

// DoNotSample allows developers to turn off do_sample option
// when calling NewClient
func DoNotSample() Option {
	return func(client *Client) *Client {
		client.DoSample = false
		return client
	}
}

// SetTemperature allows developers to set the Temperature information
// when calling NewClient
func SetTemperature(temperature float64) Option {
	return func(client *Client) *Client {
		client.Temperature = temperature
		return client
	}
}

// SetTopP allows developers to set the top_p information
// when calling NewClient
func SetTopP(topP float64) Option {
	return func(client *Client) *Client {
		client.TopP = topP
		return client
	}
}

// SetMaxToken allows developers to set the max_token
// when calling NewClient. We do not recommend setting
// max_token because the GLM will dynamically generate it
func SetMaxToken(maxToken int) Option {
	return func(client *Client) *Client {
		switch {
		case maxToken > 4095:
			client.MaxTokens = 4095
		case maxToken < 128:
			client.MaxTokens = 128
		default:
			client.MaxTokens = maxToken
		}
		return client
	}
}

// SetResponseFormatJSON allows developers to set the ResponseFormat
// to JSON format when calling NewClient. The default value is plain text.
func SetResponseFormatJSON() Option {
	return func(client *Client) *Client {
		client.ResponseFormat = "json_object"
		return client
	}
}

// SetStopWord allows developers to set the stop word
// when calling NewClient.
func SetStopWord(stopWord string) Option {
	return func(client *Client) *Client {
		client.Stop = stopWord
		return client
	}
}

// NewClient creates a new client for developers to start
// a new session and connect to the GLM.
func NewClient(APIKey string, opt ...Option) *Client {
	client := &Client{
		Url:            url,
		Model:          "glm-4",
		APIKey:         APIKey,
		UserID:         "",
		DoSample:       true,
		Temperature:    0.95,
		TopP:           0.70,
		MaxTokens:      0,
		ResponseFormat: "text",
		Stop:           "",
	}

	if len(opt) > 0 {
		for _, o := range opt {
			client = o(client)
		}
	}

	return client
}
