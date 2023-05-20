package utils

import (
	"reflect"
	"regexp"

	"github.com/leodahal4/artist-management-system-backend/internal/config"
)

// ValidateThisTag is a function that validates a given value based on a specified tag.
// The function takes in two arguments:
// - tag: a string representing the type of validation to be performed (e.g. "required", "email", "string", etc.)
// - value: a reflect.Value representing the value to be validated
// The function returns a string representing the error message if any error is found during the validation process. Otherwise, an empty string is returned.

func ValidateThisTag(tag string,value reflect.Value) string {
    switch tag {
        case "string":
            // Check if the value contains any numeric character
            numeric := regexp.MustCompile(`\d`).MatchString(value.String())
            if numeric {
                return "cannot contain number"
            }
        case "required":
            // Check if the field is an int and if so, check if it is set to 0
            if value.CanInt() {
                if value.Int() == 0 {
                    return "required"
                }
            }
            // Check if the field is an empty string
            if value.String() == ""{
                return "required"
            }
        case "email":
            // Check if the value is a valid email format
            email := regexp.MustCompile(`^[a-z0-9]+[\._]?[ a-z0-9]+[@]\w+[. ]\w{2,3}$`).MatchString(value.String())
            if !email {
                return "invalid email"
            }
        case "gender":
            // Check if the value is a valid gender
            gender := config.Gender(value.String())
            switch gender {
                case config.Male, config.Female, config.Other:
                    // Valid gender
                default:
                    return "not valid gender"
            }
        case "number":
            // Check if the value is of int type, if yes then ignore this check
            if value.CanInt() {
                return ""
            }
            // Check if the value contains any string character
            hasString := regexp.MustCompile(`[a-z]|[A-Z]`).MatchString(value.String())
            if hasString {
                return "cannot contain string"
            }
        case "genre":
            // Check if the value is a valid music genre
            genre := config.Genre(value.String())
            switch genre {
                case config.Classic, config.Country, config.Jazz, config.Rock, config.RythmAndBlues:
                    // Valid genre
                default:
                    return "not valid genre"
            }
    }
    // No error found during the validation process, return an empty string
    return ""
}
