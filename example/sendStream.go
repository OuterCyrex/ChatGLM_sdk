package example

import (
	"fmt"
	"github.com/OuterCyrex/ChatGLM_sdk"
)

func main() {
	apiKey := "your-api-key"
	client := ChatGLM_sdk.NewClient(apiKey)
	ctx := ChatGLM_sdk.NewContext()

	// Send message
	resp := client.SendStream(ctx, "Hello, how are you?")

	for data := range resp {
		if data.Error != nil {
			fmt.Println(data.Error)
		}

		for _, v := range data.Message {
			fmt.Println(v)
		}
	}
}
