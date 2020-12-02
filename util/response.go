package util

const (
	InvalidParamsErr    = "INVALID_PARAMS"
	InvalidParamsErrMsg = "invalid parameters"

	RequestErr    = "REQUEST_ERR"
	RequestErrMsg = "wrong request"

	NotFound    = "NOT_FOUND"
	NotFoundMsg = "record not found"

	UnknownErr = "UNKNOWN_ERR"
)

type RestError interface {
	Type() string
	Message() string
}

type restError struct {
	ErrType    string `json:"type"`
	ErrMessage string `json:"message"`
}

func NewRestError(errType string, errMsg string) RestError {
	return restError{
		ErrType:    errType,
		ErrMessage: errMsg,
	}
}

func (r restError) Type() string {
	return r.ErrType
}

func (r restError) Message() string {
	return r.ErrMessage
}
