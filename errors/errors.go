package errors

type Errors interface {
	Error()
}

type Err struct {
	Msg     string `json:"message"`
	ErrCode int    `json:"code"`
}