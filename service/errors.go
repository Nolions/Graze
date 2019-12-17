package service

import (
	"gopkg.in/go-playground/validator.v9"
	"graze/models"
	"strings"
)

type ForumResp struct {
	Msg     string      `json:"message"`
	ErrCode int         `json:"code"`
	Errors  CommonError `json:"errors"`
}

type CommonError map[string]interface{}

func (f *ForumResp) Error(code int, msg string, err CommonError) {
	f.ErrCode = code
	f.Msg = msg
	f.Errors = err
}

func FieldValidatorError(err error, m models.ModelFieldTran) CommonError {
	res := CommonError{}
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		transtr := e.Translate(trans)
		f := strings.ToLower(e.Field())

		if rp, ok := m[e.Field()]; ok {
			res[f] = strings.Replace(transtr, e.Field(), rp, 1)
		} else {
			res[f] = transtr
		}
	}

	return res
}