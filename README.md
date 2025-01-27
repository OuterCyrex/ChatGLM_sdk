# ZhiPu-GLM-sdk for Go

- [中文文档](README_zh.md)

This is a ZhiPuQingYan **(智谱清言)** GLM SDK written in Go, used to interact with the ZhiPuGLM API. With this SDK, developers can easily send requests and handle responses.

## Installation

```sh
go get -u github.com/OuterCyrex/ChatGLM_sdk
```

## Usage

The ZhiPuGLM-sdk provides developers with **synchronous** and **asynchronous interfaces**:

| Interface       | Function                                                   |
| :-------------- | :--------------------------------------------------------- |
| SendSync        | Send a synchronous request and return the model's reply    |
| SendAsync       | Send an asynchronous request and return the message ID     |
| GetAsyncMessage | Obtain the asynchronous model reply through the message ID |

### Initialize Client and  Send Sync Request

```go
package main

import (
    "fmt"
    "github.com/OuterCyrex/ChatGLM_sdk"
)

func main() {
    apiKey := "your-api-key"
    client := ChatGLM_sdk.NewClient(apiKey)
    ctx := ChatGLM_sdk.NewContext()

    // Send message
    resp := client.SendSync(ctx, "Hello, how are you?")
    if resp.Error != nil {
        fmt.Println("Error:", resp.Error)
        return
    }

    // Print response
    for _, msg := range resp.Message {
        fmt.Printf("%s: %s\n", msg.Role, msg.Content)
    }
}
```

### Initialize Client and  Send Async Request

```go
package main

import (
	"fmt"
	"github.com/OuterCyrex/ChatGLM_sdk"
	"time"
)

func main() {
    apiKey := "your-api-key"
	client := ChatGLM_sdk.NewClient(apiKey)
	ctx := ChatGLM_sdk.NewContext()

    // send async message and get id
	id, err := client.SendAsync(ctx, "Hello, how are you?")

	if err != nil {
		fmt.Println("Error:", err)
        return
	}

    // wait for glm to generate
	time.Sleep(5 * time.Second)

    // get response by id returned by SendAsync
	resp := client.GetAsyncMessage(ctx, id)
	if resp.Error != nil {
		fmt.Println("Error:", resp.Error)
        return
	}

	for _, v := range resp.Message {
		fmt.Println(v)
	}
}
```

### Configuration Options

You can use the following options to customize the behavior of the client:

- `SetUserID(userID string)`: Set the user ID.
- `DoNotSample()`: Turn off sampling.
- `SetTemperature(temperature float64)`: Set the temperature parameter.
- `SetTopP(topP float64)`: Set the TopP parameter.
- `SetMaxToken(maxToken int)`: Set the maximum token count.
- `SetResponseFormatJSON()`: Set the response format to JSON.
- `SetStopWord(stopWord string)`: Set the stop word.

### Example

```go
client := ChatGLM_sdk.NewClient(apiKey, 
    ChatGLM_sdk.SetUserID("user123"),
    ChatGLM_sdk.DoNotSample(),
    ChatGLM_sdk.SetTemperature(0.8),
    ChatGLM_sdk.SetTopP(0.9),
    ChatGLM_sdk.SetMaxToken(2048),
    ChatGLM_sdk.SetResponseFormatJSON(),
    ChatGLM_sdk.SetStopWord("stop"),
)
```

## Notes

- Ensure your API key is valid.
- Ensure your network connection is normal and can access the ChatGLM API.
- If you encounter problems, please check the API key and network connection, and retry as appropriate.

## Contribution

Contributions and suggestions for improvement are welcome.

## License

This project follows the [MIT License](https://opensource.org/license/MIT).