package ChatGLM_sdk

import (
	"errors"
)

var (
	ErrResultProcessing = errors.New("the return result is being generated")

	// ErrSdkInternal happens when some critical codes go wrong
	// this shall not happen, if it exists please push an issue
	ErrSdkInternal = errors.New("sdk internal error")

	ErrHttpRequestTimeOut = errors.New("http request to glm server failed")
	ErrNotFound           = errors.New("the result id is not found")
	ErrHttpBadRequest     = errors.New("http response not StatusOK")
	ErrGenerateFailed     = errors.New("glm generate failed")
)
