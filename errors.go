package v2

import "errors"

var (
	CreateRequestFailed         = errors.New("request creation failed")
	ExecRequestFailed           = errors.New("request execution failed")
	ReadBodyFailed              = errors.New("failed reading the response body")
	JsonError                   = errors.New("json marshalling/unmarshalling of response body failed")
	UnsupportedEnvironmentError = errors.New("the specified environment is not supported")
)
