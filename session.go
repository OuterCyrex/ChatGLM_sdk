package ChatGLM_sdk

import "github.com/OuterCyrex/ChatGLM_sdk/model"

const (
	SyncUrl     = "https://open.bigmodel.cn/api/paas/v4/chat/completions"
	AsyncUrl    = "https://open.bigmodel.cn/api/paas/v4/async/chat/completions"
	getAsyncUrl = "https://open.bigmodel.cn/api/paas/v4/async-result/"
)

// Result is the return value that contains
// tokens consumed, messages returned from GLM,
// as well as the error message.
type Result struct {
	Tokens  int32
	Message []model.Message
	Error   error
}

// MessageContext is the dialog context for GLM
// to recognize what was said before
type MessageContext []model.Message

// NewContext create a new background dialog context
func NewContext() *MessageContext {
	var ctx MessageContext
	return &ctx
}
