package errors

import (
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
	"graze/models"
	"strings"
)

var (
	trans    ut.Translator
	Validate *validator.Validate
)

func init() {
	zhTw := zh_Hant_TW.New()
	uni := ut.New(zhTw)
	trans, _ = uni.GetTranslator("zh_tw")
	Validate = validator.New()
	zh_tw_translations.RegisterDefaultTranslations(Validate, trans)
}

type Errors interface {
	Error()
}

type Err struct {
	Msg     string `json:"message"`
	ErrCode int    `json:"code"`
}

type ValidatorError struct {
	Err
	Errors FieldErrorMsg `json:"errors"`
}

func (e *ValidatorError) Error() {
	e.ErrCode = 10001
	e.Msg = "Field Validate's Error"
}

type FieldErrorMsg map[string]interface{}

func FieldValidatorError(err error, m models.ModelFieldTran) FieldErrorMsg {
	res := FieldErrorMsg{}
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
