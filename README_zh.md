# ChatGLM_sdk for Go

这是一个用Go语言编写的 **智谱清言**GLM SDK，用于与ZhiPuGLM API进行交互。通过这个SDK，开发者可以轻松地发送请求并处理响应。

目前支持开发者向GLM发送同步请求与异步请求。

## 安装

```sh
go get -u github.com/OuterCyrex/ChatGLM_sdk
```

## 使用方法

ChatGLM_sdk为开发者提供了**同步接口**与**异步接口**：

| 接口            | 作用                           |
| --------------- | ------------------------------ |
| SendSync        | 发送同步请求并返回模型回复     |
| SendAsync       | 发送异步请求并返回消息ID       |
| GetAsyncMessage | 通过消息ID获取异步的模型回复   |
| SendStream      | 发送流式请求并返回流式响应通道 |

### 初始化客户端并发送请求

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

    // 发送消息
    resp := client.SendSync(ctx, "Hello, how are you?")
    if resp.Error != nil {
        fmt.Println("Error:", resp.Error)
        return
    }

    // 打印响应
    for _, msg := range resp.Message {
        fmt.Printf("%s: %s\n", msg.Role, msg.Content)
    }
}
```

更多用法请参考：

- [代码示例](example)

### 配置选项

你可以使用以下选项来自定义客户端的行为：

- `SetUserID(userID string)`: 设置用户ID。
- `DoNotSample()`: 关闭采样功能。
- `SetTemperature(temperature float64)`: 设置温度参数。
- `SetTopP(topP float64)`: 设置TopP参数。
- `SetMaxToken(maxToken int)`: 设置最大令牌数。
- `SetResponseFormatJSON()`: 设置响应格式为JSON。
- `SetStopWord(stopWord string)`: 设置停止词。

### 示例

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

## 注意事项

- 确保你的API密钥是有效的。
- 确保你的网络连接正常，能够访问ChatGLM API。
- 如果遇到问题，请检查API密钥和网络连接，适当重试。

## 贡献

欢迎贡献代码和提出改进建议。

## 许可

本项目遵循 [MIT许可证](https://opensource.org/license/MIT)。