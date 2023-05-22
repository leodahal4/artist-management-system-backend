package config

import "fmt"

type EnvConfigError struct{
  Err error
}

func (e *EnvConfigError) Error () string {
  return fmt.Sprintf("unfortunately %v variable is not found on the .env file", e.Err.Error())
}


type ErrorResponse struct {
  Err error
}

func (e *ErrorResponse) Error() string {
  return e.Error()
}

func (e *ErrorResponse) Errors() map[string][]string {
  newError := make(map[string][]string, 1)
  newError["error"] = []string{e.Err.Error()}
  return newError
}
