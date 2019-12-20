package errors

type Errors interface {
	Error()
}

type Err struct {
	ErrCode int    `json:"code"`
	Msg     string `json:"message"`
}

type ModelNoFoundError struct {
	Err
}

func (e *ModelNoFoundError) Error() {
	e.ErrCode = 10101
	e.Msg = "沒有任何符合資料"
}

type InsertErrors struct {
	Err
}

func (e *InsertErrors) Error()  {
	e.ErrCode = 10102
	e.Msg = "資料新增失敗"
}

type DeleteErrors struct {
	Err
}

func (e *DeleteErrors) Error()  {
	e.ErrCode = 10103
	e.Msg = "資料新增失敗"
}

type UpdateErrors struct {
	Err
}

func (e *UpdateErrors) Error()  {
	e.ErrCode = 10104
	e.Msg = "資料新增失敗"
}