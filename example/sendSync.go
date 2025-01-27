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
