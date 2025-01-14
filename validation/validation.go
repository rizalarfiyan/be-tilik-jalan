package validation

import (
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rizalarfiyan/be-tilik-jalan/logger"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
	once     sync.Once
)

func Init() {
	once.Do(func() {
		log := logger.Get("validation")
		langEn := en.New()
		uni = ut.New(langEn)
		trans, _ = uni.GetTranslator("en")
		validate = validator.New()
		err := enTranslations.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to register default translations")
		}

		custom := customValidation{}
		custom.SetTagName()
		custom.SetEnum()
	})
}

func Get() *validator.Validate {
	return validate
}

func GetTranslator() ut.Translator {
	return trans
}
