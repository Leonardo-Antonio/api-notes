package helper

type Response struct {
	MessageType string
	Message string
	Error bool
	Data interface{}
}

func NewResponseJSON(mt, m string, err bool, data interface{}) *Response {
	return &Response{
		mt,
		m,
		err,
		data,
	}
}