package models

type Error struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Code    int    `json:"code"`
}

func NewError(message string, code int, details string) Error {
	return Error{Message: message, Code: code, Details: details}
}
