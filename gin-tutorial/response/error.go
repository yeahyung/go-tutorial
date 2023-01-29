package response

import (
	"errors"
)

type ErrorResponse struct {
	Err ServerError `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	var cErr ServerError
	if !errors.As(err, &cErr) {
		cErr = ServerError{
			ErrCode: 500,
			Message: err.Error(),
			Err:     err,
		}
	}
	return ErrorResponse{
		Err: cErr,
	}
}
