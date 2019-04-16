package http

import (
	"github.com/Kamva/nautilus"
	"github.com/Kamva/orca"
	"github.com/Kamva/pantopoda"
	"github.com/Kamva/shark"
	"gopkg.in/go-playground/validator.v9"
)

// BaseRequest is base request struct contains default functionality of a request
type BaseRequest struct {
	nautilus.BaseTaggable
}

// Validate runs request data validation and returns validation error
func (r BaseRequest) Validate() pantopoda.ValidationError {
	validate := orca.GetValidator()
	validationError := pantopoda.ValidationError{}

	err := validate.Struct(r)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			validationError.ErrorType = pantopoda.BadRequest
		} else {
			errorBag := shark.NewErrorBag()

			for _, err := range err.(validator.ValidationErrors) {
				errorBag.Append(
					nautilus.ToSnake(err.StructField()),
					orca.GetTranslationKey(&r, err.StructField(), err.Tag()),
				)
			}

			validationError.ErrorType = pantopoda.RuleViolation
			validationError.ErrorBag = errorBag
		}
	}

	return validationError
}
