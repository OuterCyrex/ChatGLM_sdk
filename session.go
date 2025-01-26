package ChatGLM_sdk

import "ChatGLM_sdk/model"

func NewContext() MessageContext {
	return make([]model.Message, 0)
}
