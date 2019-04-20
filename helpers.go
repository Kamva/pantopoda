package pantopoda

import (
	"github.com/Kamva/nautilus"
	"github.com/Kamva/orca"
	"github.com/Kamva/shark"
	"gopkg.in/go-playground/validator.v9"
)

// Validate runs request data validation and returns validation error
func Validate(r RequestData) ValidationError {
	validate := orca.GetValidator()
	validationError := ValidationError{}

	err := validate.Struct(r)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			validationError.ErrorType = BadRequest
		} else {
			errorBag := shark.NewErrorBag()

			t := r.(nautilus.Taggable)
			for _, err := range err.(validator.ValidationErrors) {
				errorBag.Append(
					nautilus.ToSnake(err.StructField()),
					orca.GetTranslationKey(t, err.StructField(), err.Tag()),
				)
			}

			validationError.ErrorType = RuleViolation
			validationError.ErrorBag = errorBag
		}
	}

	return validationError
}
