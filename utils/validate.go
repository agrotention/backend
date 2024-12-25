package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	english  = en.New()
	uni      = ut.New(english, english)
	trans, _ = uni.GetTranslator("en")
	_        = enTranslations.RegisterDefaultTranslations(Validate, trans)
)

func TranslateValidationError(err validator.ValidationErrors) HTTPError {
	var m []string
	for _, e := range err {
		translated := e.Translate(trans)
		m = append(m, translated)
	}
	return NewErrWithData[[]string](400, "validation error", m)
}
