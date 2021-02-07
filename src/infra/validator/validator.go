package validator

import "github.com/go-playground/validator/v10"

// Validate func validate a object
func Validate(aStruct interface{}) ([]string, error) {
	var errorsMessage []string
	v := validator.New()
	err := v.Struct(aStruct)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorsMessage = append(errorsMessage, e.Error())
		}
	}

	return errorsMessage, err
}
