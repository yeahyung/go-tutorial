package response

import "fmt"

type ServerError struct {
	ErrCode int    `json:"errorCode"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func NewServerError(errCode int32, message string, err error) error {
	return ServerError{
		ErrCode: int(errCode),
		Message: message,
		Err:     err,
	}
}

func (err ServerError) Error() string {
	if err.Err != nil {
		return fmt.Sprintf("err(%s), message(%s)", err.Err.Error(), err.Message)
	}
	return err.Message
}

func (err ServerError) Unwrap() error {
	return err.Err
}
