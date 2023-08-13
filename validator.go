package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func Validate(Object interface{}) ([]string, bool) {

	var ValidationErrors []string

	// Create a universal translator using English locale
	english := en.New()
	uni := ut.New(english, english)
	translator, _ := uni.GetTranslator("en")

	// Create a new validator
	validate := validator.New()

	// Register English translations for the validation error messages
	en_translations.RegisterDefaultTranslations(validate, translator)

	// Validate the struct
	err := validate.Struct(Object)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				fmt.Println(e.Translate(translator))
				ValidationErrors = append(ValidationErrors, e.Translate(translator))
			}
			return ValidationErrors, false
		}
	}

	return nil, true
}
