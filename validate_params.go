package itools

import (
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"log"
)

var (
	msg string
)

// EnValidateParam validate params english
func EnValidateParam(param interface{}) string {
	validate := enBindValidate()
	err := validate.Struct(param)
	if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		for _, er := range errs {
			msg = er.Error()
		}
		return msg
	}
	return ""
}

// enBindValidate
// i18n
func enBindValidate() *validator.Validate {
	return validator.New()
}

// ZhValidateParam validate params english
func ZhValidateParam(param interface{}) string {
	translator, validate := zhBindValidate()
	err := validate.Struct(param)
	if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		for _, er := range errs {
			msg = er.Translate(translator)
		}
		return msg
	}
	return ""
}

// zhBindValidate
// i18n
func zhBindValidate() (ut.Translator, *validator.Validate) {
	validate := validator.New()
	e := en.New()
	uniTrans := ut.New(e, e, zh.New())
	translator, _ := uniTrans.GetTranslator("zh")
	err := zh2.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		log.Printf("i18n 国际化失败: %v", err)
	}
	return translator, validate
}
