package ChatGLM_sdk

// Client contains basic parameter information for communicating with GLM
type Client struct {
	userID         string
	model          string
	apiKey         string
	doSample       bool
	temperature    float64
	topP           float64
	maxTokens      int
	responseFormat string
	stop           string
}

type Option func(client *Client) *Client

// SetUserID allows developers to set the userID information
// when calling NewClient
func SetUserID(userID string) Option {
	return func(client *Client) *Client {
		client.userID = userID
		return client
	}
}

// DoNotSample allows developers to turn off do_sample option
// when calling NewClient
func DoNotSample() Option {
	return func(client *Client) *Client {
		client.doSample = false
		return client
	}
}

// SetTemperature allows developers to set the temperature information
// when calling NewClient
func SetTemperature(temperature float64) Option {
	return func(client *Client) *Client {
		client.temperature = temperature
		return client
	}
}

// SetTopP allows developers to set the top_p information
// when calling NewClient
func SetTopP(topP float64) Option {
	return func(client *Client) *Client {
		client.topP = topP
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
			client.maxTokens = 4095
		case maxToken < 128:
			client.maxTokens = 128
		default:
			client.maxTokens = maxToken
		}
		return client
	}
}

// SetResponseFormatJSON allows developers to set the responseFormat
// to JSON format when calling NewClient. The default value is plain text.
func SetResponseFormatJSON() Option {
	return func(client *Client) *Client {
		client.responseFormat = "json_object"
		return client
	}
}

// SetStopWord allows developers to set the stop word
// when calling NewClient.
func SetStopWord(stopWord string) Option {
	return func(client *Client) *Client {
		client.stop = stopWord
		return client
	}
}

// NewClient creates a new client for developers to start
// a new session and connect to the GLM.
func NewClient(APIKey string, opt ...Option) *Client {
	client := &Client{
		model:          "glm-4",
		apiKey:         APIKey,
		userID:         "",
		doSample:       true,
		temperature:    0.95,
		topP:           0.70,
		maxTokens:      0,
		responseFormat: "text",
		stop:           "",
	}

	if len(opt) > 0 {
		for _, o := range opt {
			client = o(client)
		}
	}

	return client
}
