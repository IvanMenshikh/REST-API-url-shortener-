package response

type Response struct {
	Status string `json:"status"` // Статус ответа, например "success" или "error"
	Error  string `json:"error,omitempty"` // Ошибка, если есть, иначе пустая строка
}

const (
	StatusOk = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}