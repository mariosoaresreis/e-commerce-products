package errors

type ErrorMessage struct {
	Errors []Message `json:"errors"`
}

type Message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewNotFoundErrorMessage(message string) Message {
	return Message{
		Type:    "not found",
		Message: message,
	}
}

func NewValidationErrorMessage(message string) Message {
	return Message{
		Type:    "validation",
		Message: message,
	}
}

func NewUnexpectedErrorMessage(message string) Message {
	return Message{
		Type:    "unexpected",
		Message: message,
	}
}
