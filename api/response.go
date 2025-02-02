package api

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type Response struct {
	Status  Status      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponse(status Status, data interface{}, message string) Response {
	return Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
}

func ErrorResponse(data interface{}, message string) Response {
	return NewResponse(StatusError, data, message)
}

func SuccessResponse(data interface{}, message string) Response {
	return NewResponse(StatusSuccess, data, message)
}
