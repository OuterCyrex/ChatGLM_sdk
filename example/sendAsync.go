package example

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
