package config

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Read_config returns the configuration of the server.
func Read_config() (Config, error) {
  var env string
  var projectName string
  var pgHost string
  var pgPort string
  var pgUser string
  var pgPass string
  var pgDBName string
  var debugEnv string
  var debug bool
  var serverPort string

  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("cannot read .env file")
  }

  // if the env variable is not set than, default to DEV env
  if env = os.Getenv("env"); env == "" {
    env = "dev"
  }

  // if the project_name variable is not set than, default to `artist management` env
  if projectName = os.Getenv("project_name"); projectName == "" {
    projectName = "artist management"
  }
  if pgPort = os.Getenv("pg_port"); pgPort == "" {
    // if not found then, return error
    pgPort = "5432"
  }
  if serverPort = os.Getenv("port"); serverPort == "" {
    // if not found then, return error
    serverPort = "8080"
  }

  // check if the debug variable is found
  if debugEnv = os.Getenv("debug"); debugEnv == "" {
    debug = is_debug_enabled("", env)
  } else {
    debug = is_debug_enabled(debugEnv, "")
  }
  
  // all other obligatory variables on the .env
  if pgHost = os.Getenv("pg_host"); pgHost == "" {
    // if not found then, return error
    return Config{}, &EnvConfigError{errors.New("pg_host")}
  }
  if pgUser = os.Getenv("pg_user"); pgUser == "" {
    // if not found then, return error
    return Config{}, &EnvConfigError{errors.New("pg_user")}
  }

  if pgDBName = os.Getenv("pg_db_name"); pgDBName == "" {
    // if not found then, return error
    return Config{}, &EnvConfigError{errors.New("pg_db_name")}
  }

  if pgPass = os.Getenv("pg_pass"); pgPass == "" {
    // if not found then, return error
    return Config{}, &EnvConfigError{errors.New("pg_pass")}
  }

  return Config{
    Env: env,
    ProjectName: projectName,
    PgHost: pgHost,
    PgPort: pgPort,
    PgUser: pgUser,
    PgPass: pgPass,
    PgDBName: pgDBName,
    Debug: debug,
    ServerPort: serverPort,
  }, nil
}

func is_debug_enabled(debugVal string, debugEnv string) bool {
  if debugEnv != ""{
    if strings.ToLower(debugEnv) == "prod" || strings.ToLower(debugEnv) == "production" {
      return false
    }
    return true
  }
  if strings.ToLower(debugVal) == "true" {
    return true
  }
  return false
}
