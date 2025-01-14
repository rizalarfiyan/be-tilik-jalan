package response

type Base struct {
	Code    int         `json:"-"`
	Message string      `json:"message" example:"Message!"`
	Data    interface{} `json:"data"`
}

func (res *Base) Error() string {
	return res.Message
}

func NewErrorMessage(code int, message string, data interface{}) *Base {
	return &Base{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
