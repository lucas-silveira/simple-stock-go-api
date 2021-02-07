package validator

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_BR_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

// Validate func validate a object
func Validate(aStruct interface{}) ([]string, error) {
	ptBR := pt_BR.New()
	uni := ut.New(ptBR, ptBR)
	trans, _ := uni.GetTranslator("pt_BR")
	validate = validator.New()
	pt_BR_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(aStruct)
	var errorsMessage []string

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorsMessage = append(errorsMessage, e.Translate(trans))
		}
	}

	return errorsMessage, err
}
