package errors

import "net/http"

type ApplicationError struct {
	Code    int          `json:",omitempty"`
	Message string       `json:"message"`
	Error   ErrorMessage `json:"errors"`
}

func (e ApplicationError) Text() *ApplicationError {
	return &ApplicationError{
		Message: e.Message,
	}
}

func NewNotFoundError(messages []Message) *ApplicationError {
	errorMessage := ErrorMessage{
		Errors: messages,
	}

	return &ApplicationError{
		Code:  http.StatusNotFound,
		Error: errorMessage,
	}
}

func NewSingleNotFoundError(message string) *ApplicationError {
	errors := make([]Message, 0)
	errors = append(errors, NewNotFoundErrorMessage(message))

	errorMessage := ErrorMessage{
		Errors: errors,
	}

	return &ApplicationError{
		Code:  http.StatusNotFound,
		Error: errorMessage,
	}
}

func NewBadRequestError(messages []Message) *ApplicationError {
	errorMessage := ErrorMessage{
		Errors: messages,
	}

	return &ApplicationError{
		Code:  http.StatusBadRequest,
		Error: errorMessage,
	}
}

func NewSingleBadRequestError(message string) *ApplicationError {
	errors := make([]Message, 0)
	errors = append(errors, NewValidationErrorMessage(message))

	errorMessage := ErrorMessage{
		Errors: errors,
	}

	return &ApplicationError{
		Code:  http.StatusBadRequest,
		Error: errorMessage,
	}
}

func NewUnexpectedError(message string) *ApplicationError {
	errors := make([]Message, 0)
	errors = append(errors, NewUnexpectedErrorMessage(message))

	errorMessage := ErrorMessage{
		Errors: errors,
	}

	return &ApplicationError{
		Code:    http.StatusInternalServerError,
		Error:   errorMessage,
		Message: message,
	}
}
