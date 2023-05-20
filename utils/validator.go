package utils

import (
	"reflect"
	"regexp"

	"github.com/leodahal4/artist-management-system-backend/internal/config"
)

func ValidateThisTag(tag string,value reflect.Value) string {
    switch tag {
      case "string":
        numeric := regexp.MustCompile(`\d`).MatchString(value.String())
        if numeric {
          return "cannot contain number"
        }
      case "required":
        if value.String() == ""{
          return "must contain value"
        }
      case "email":
        email := regexp.MustCompile(`^[a-z0-9]+[\._]?[ a-z0-9]+[@]\w+[. ]\w{2,3}$`).MatchString(value.String())
        if !email {
          return "invalid email"
        }
      case "gender":
        gender := config.Gender(value.String())
        switch gender {
          case config.Male, config.Female, config.Other:
            //
          default:
            return "not valid gender"
        }
      case "number":
        hasString := regexp.MustCompile(`[a-z]|[A-Z]`).MatchString(value.String())
        if hasString {
          return "cannot contain string"
        }
      case "genre":
        genre := config.Genre(value.String())
        switch genre {
          case config.Classic, config.Country, config.Jazz, config.Rock, config.RythmAndBlues:
              // valid genre
          default:
              return "not valid genre"
        }
    }
  return ""
}
