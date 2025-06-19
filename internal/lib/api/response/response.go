package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Response - структура для формирования ответа API.
type Response struct {
	Status string `json:"status"`          // Статус ответа, например "success" или "error"
	Error  string `json:"error,omitempty"` // Ошибка, если есть, иначе пустая строка
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

// OK формирует успешный ответ без ошибок.
func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

// Error формирует ответ с ошибкой.
func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

// ValidationError формирует ответ с ошибками валидации.
func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field '%s' is a required field", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("field '%s' is not valid URL", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field '%s' is not valid", err.Field()))

		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
