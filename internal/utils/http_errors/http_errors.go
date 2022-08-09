package http_errors

import "net/http"

type RestErrors interface {
	Code() int
	Status() string
	Message() string
	Data() []interface{}
}

type restErrors struct {
	code    int           `json:"code"`
	status  string        `json:"status"`
	message string        `json:"message"`
	data    []interface{} `json:"data"`
}

func NewHTTPErrors(code int, status string, message string, data []interface{}) RestErrors {
	return &restErrors{
		code:    code,
		status:  status,
		message: message,
		data:    data,
	}
}

func (re restErrors) Code() int {
	return re.code
}
func (re restErrors) Status() string {
	return re.status
}
func (re restErrors) Message() string {
	return re.message
}
func (re restErrors) Data() []interface{} {
	return re.data
}

func NewGeneralError(code int, status string, message string, data []interface{}) restErrors {
	return restErrors{
		code:    code,
		status:  status,
		message: message,
		data:    data,
	}
}

func NewBadRequestError(message string, data []interface{}) restErrors {
	return restErrors{
		code:    http.StatusBadRequest,
		status:  "failed",
		message: message,
		data:    data,
	}
}

func NewInternalServerError(message string, data []interface{}) restErrors {
	return restErrors{
		code:    http.StatusInternalServerError,
		status:  "error",
		message: message,
		data:    data,
	}
}

func NewStatusNotFoundError(message string, data []interface{}) restErrors {
	return restErrors{
		code:    http.StatusNotFound,
		status:  "failed",
		message: message,
		data:    data,
	}
}
