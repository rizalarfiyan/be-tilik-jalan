package validation

import (
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type customValidation struct{}

const (
	enumKey = "enum"
)

func (re customValidation) SetTagName() {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "" {
			return field.Name
		}

		if name == "-" {
			return ""
		}

		return name
	})
}

type EnumIsValid interface {
	IsValid() bool
}

func (re customValidation) SetEnum() {
	_ = validate.RegisterValidation(enumKey, func(fl validator.FieldLevel) bool {
		if enum, ok := fl.Field().Interface().(EnumIsValid); ok {
			return enum.IsValid()
		}

		return false
	})

	_ = validate.RegisterTranslation(enumKey, trans, func(ut ut.Translator) error {
		return ut.Add(enumKey, "{0} invalid enum format.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(enumKey, fe.Field())
		return t
	})
}
