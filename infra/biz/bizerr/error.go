package bizerr

import "fmt"

type BizError struct {
	Code    int
	Message string
	Err     error
}

func (e *BizError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("code=%d, msg=%s, err=%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Message)
}

func (e *BizError) Unwrap() error {
	return e.Err
}

func NewBizError(code int, message string) *BizError {
	return &BizError{Code: code, Message: message}
}

func WrapBizError(code int, message string, err error) *BizError {
	return &BizError{Code: code, Message: message, Err: err}
}
