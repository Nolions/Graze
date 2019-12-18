package errors

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/go-playground/validator.v9/translations/zh_tw"
	"graze/models"
	"strings"
	"time"
)

type FieldErrorMsg map[string]interface{}

var (
	trans    ut.Translator
	Validate *validator.Validate
)

func init() {
	zhTw := zh_Hant_TW.New()
	uni := ut.New(zhTw)
	trans, _ = uni.GetTranslator("zh_tw")
	Validate = validator.New()
	Validate.RegisterValidation("datetime", TimeFormatValidator)
	zh_tw.RegisterDefaultTranslations(Validate, trans)
	Validate.RegisterTranslation("datetime", trans, dateTimeErrorRegister, dateTimeErrorTranslation)
}

// 註冊dateTime錯誤訊息
func dateTimeErrorRegister(ut ut.Translator) error {
	return ut.Add("datetime", "{0} 為無效日期或日期格式錯誤。", true)
}

// 翻譯dateTime錯誤訊息
func dateTimeErrorTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("datetime", fe.Field())
	return t
}

// 驗證日期格式
func TimeFormatValidator(fl validator.FieldLevel) bool {
	fmt.Println(fl.Field().String())
	if fl.Field().String() == "" {
		return true // 沒有輸入值時，直接返回true
	}

	t := fmt.Sprint(fl.Field())
	tt, err := time.Parse(time.RFC3339, t)
	if err != nil || tt.IsZero() {
		return false
	}

	return true
}

func FieldValidatorError(err error, m models.ModelFieldTran) FieldErrorMsg {
	res := FieldErrorMsg{}
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		transtr := e.Translate(trans)
		fmt.Println(transtr)
		f := strings.ToLower(e.Field())
		if rp, ok := m[e.Field()]; ok {
			res[f] = strings.Replace(transtr, e.Field(), rp, 1)
		} else {
			res[f] = transtr
		}
	}

	return res
}

type ValidatorError struct {
	Err
	Errors FieldErrorMsg `json:"errors"`
}

func (e *ValidatorError) Error() {
	e.ErrCode = 10001
	e.Msg = "Field Validate's Error"
}
