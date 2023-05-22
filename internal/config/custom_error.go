package config

import "fmt"

type EnvConfigError struct{
  Err error
}

func (e *EnvConfigError) Error () string {
  return fmt.Sprintf("unfortunately %v variable is not found on the .env file", e.Err.Error())
}
