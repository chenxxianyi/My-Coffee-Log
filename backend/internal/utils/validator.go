package utils

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func init() {
	validate = validator.New()
	validate.RegisterValidation("flavor_score", flavorScoreValidation)
}

func flavorScoreValidation(fl validator.FieldLevel) bool {
	val := fl.Field().Int()
	return val >= 0 && val <= 5
}

func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var valErrs validator.ValidationErrors
	if !errors.As(err, &valErrs) {
		return map[string]string{"error": err.Error()}
	}

	errs := make(map[string]string)
	for _, e := range valErrs {
		field := e.Field()
		switch e.Tag() {
		case "required":
			errs[field] = fmt.Sprintf("%s is required", field)
		case "email":
			errs[field] = fmt.Sprintf("%s must be a valid email", field)
		case "min":
			errs[field] = fmt.Sprintf("%s must be at least %s characters", field, e.Param())
		case "max":
			errs[field] = fmt.Sprintf("%s must be at most %s characters", field, e.Param())
		case "flavor_score":
			errs[field] = fmt.Sprintf("%s must be between 0 and 5", field)
		default:
			errs[field] = fmt.Sprintf("%s is invalid", field)
		}
	}

	return errs
}

func IsEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsZero()
}
