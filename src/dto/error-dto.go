package dto

type err struct {
	Error string `json:"error"`
}

func NewError(message string) *err {
	return &err{
		Error: message,
	}
}
