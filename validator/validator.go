package validator

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate
var regxSpecialChar *regexp.Regexp
var regxXss *regexp.Regexp
var regxAlphaDash *regexp.Regexp
var regexAlphaSpace *regexp.Regexp
var regxDomain *regexp.Regexp

func init() {
	regxSpecialChar = regexp.MustCompile(`^[a-zA-Z0-9\s_.-]+$`)
	regxXss = regexp.MustCompile(`^[a-zA-Z0-9\s.,@+_=:/-]+$`)
	regxAlphaDash = regexp.MustCompile(`^[a-zA-Z-]+$`)
	regexAlphaSpace = regexp.MustCompile(`^[a-zA-Z\s]+$`)
	regxDomain = regexp.MustCompile(`^(?i)([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z]{2,63}$`)

	// init validator
	Validator = validator.New()

	// register custom validation
	_ = Validator.RegisterValidation("specialchar", specialchar)
	_ = Validator.RegisterValidation("xss_clean", xssClean)
	_ = Validator.RegisterValidation("passwdComplex", checkPasswordComplexity)
	_ = Validator.RegisterValidation("alphadash", alphaDash)
	_ = Validator.RegisterValidation("alphaSpace", alphaSpace)
	_ = Validator.RegisterValidation("domain", domain)
}

func specialchar(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	return regxSpecialChar.MatchString(fl.Field().String())
}

func xssClean(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	return regxXss.MatchString(fl.Field().String())
}

func alphaDash(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	return regxAlphaDash.MatchString(fl.Field().String())
}

func alphaSpace(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	return regexAlphaSpace.MatchString(fl.Field().String())
}

func domain(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	return regxDomain.MatchString(fl.Field().String())
}

// Function to map validation errors to user-friendly messages
func MapValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			// Generate a user-friendly error message based on the validation tag
			switch e.Tag() {
			case "required":
				errors[e.Field()] = "required field"
			case "email":
				errors[e.Field()] = "An error occurred, please try again."
			case "gte":
				errors[e.Field()] = fmt.Sprintf("field must be greater than or equal to %s", e.Param())
			case "lte":
				errors[e.Field()] = fmt.Sprintf("field must be less than or equal to %s", e.Param())
			case "min":
				errors[e.Field()] = fmt.Sprintf("field must be at least %s characters long", e.Param())
			case "max":
				errors[e.Field()] = fmt.Sprintf("field must be at most %s characters long", e.Param())
			case "len":
				errors[e.Field()] = fmt.Sprintf("field must be exactly %s characters long", e.Param())
			case "lowercase":
				errors[e.Field()] = "field must be in lowercase"
			case "uppercase":
				errors[e.Field()] = "field must be in uppercase"
			case "alphanum":
				errors[e.Field()] = "field must be alphanumeric"
			case "alpha":
				errors[e.Field()] = "field must only contain alphabetic characters"
			case "numeric":
				errors[e.Field()] = "field must be numeric"
			case "url":
				errors[e.Field()] = "field must be a valid URL"
			case "uuid":
				errors[e.Field()] = "field must be a valid UUID"
			case "oneof":
				errors[e.Field()] = fmt.Sprintf("field must be one of the following: %s", e.Param())
			case "unique":
				errors[e.Field()] = "field must be unique"
			case "xss_clean":
				errors[e.Field()] = "field contain not allowed character"
			case "specialchar":
				errors[e.Field()] = "field special character not allowed"
			case "alphadash":
				errors[e.Field()] = "field must only contain alphabetic characters and dashes"
			case "passwdComplex":
				errors[e.Field()] = "Your password is not secure. Please try again with a stronger one."
			case "alphaSpace":
				errors[e.Field()] = "field must only contain alphabetic characters and spaces"
			// Add cases for additional tags as needed
			default:
				errors[e.Field()] = "field is invalid"
			}
		}
	}
	return errors
}
func checkPasswordComplexity(fl validator.FieldLevel) bool {
	passwd := fl.Field().String()

	if len(passwd) < 12 {
		return false
	}

	var capitalFlag, lowerCaseFlag, numberFlag, specialCharFlag bool
	for _, c := range passwd {

		if unicode.IsUpper(c) {
			capitalFlag = true
		} else if unicode.IsLower(c) {
			lowerCaseFlag = true
		} else if unicode.IsDigit(c) {
			numberFlag = true
		} else if unicode.IsSymbol(c) || unicode.IsPunct(c) {
			specialCharFlag = true
		}
		if capitalFlag && lowerCaseFlag && numberFlag && specialCharFlag {
			return true
		}
	}
	return false
}
