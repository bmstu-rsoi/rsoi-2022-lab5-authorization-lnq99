package model

type ErrorDescription struct {
	Error *string `json:"error,omitempty"`
	Field *string `json:"field,omitempty"`
}

type ErrorResponse struct {
	// Информация об ошибке
	Message *string `json:"message,omitempty"`
}

type ValidationErrorResponse struct {
	// Массив полей с описанием ошибки
	Errors *[]ErrorDescription `json:"errors,omitempty"`

	// Информация об ошибке
	Message *string `json:"message,omitempty"`
}
