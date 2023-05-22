package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// saving as a local global for now,
// NOTE: NEED TO REMOVE THIS GLOBAL VARIABLE
var config *Config

// Read_config returns the configuration of the server.
func Read_config() (*Config, error) {
  var env string
  var projName string
  var pgHost string
  var pgPort string
  var pgUser string
  var pgPass string
  var pgDBName string
  var debugEnv string
  var debug bool
  var srvPort string
  var pEntropyEnv string
  var pEntropy float64

  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("cannot read .env file")
  }

  // if the env variable is not set than, default to DEV env
  if env = os.Getenv("env"); env == "" {
    env = "dev"
  }

  // if the project_name variable is not set than, default to `artist management` env
  if projName = os.Getenv("project_name"); projName == "" {
    projName = "artist management"
  }
  if pgPort = os.Getenv("pg_port"); pgPort == "" {
    // if not found then, return error
    pgPort = "5432"
  }
  if srvPort = os.Getenv("port"); srvPort == "" {
    // if not found then, return error
    srvPort = "8080"
  }
  if pEntropyEnv = os.Getenv("pass_entropy"); pEntropyEnv == "" {
    // if not found then, return error
    pEntropyEnv = "100"
  }
  // convert the string of password entropy to float64
  pEntropy, err = strconv.ParseFloat(pEntropyEnv, 10)
  if err != nil {
    fmt.Printf("err: %v\n", err)
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
    return nil, &EnvConfigError{errors.New("pg_host")}
  }
  if pgUser = os.Getenv("pg_user"); pgUser == "" {
    // if not found then, return error
    return nil, &EnvConfigError{errors.New("pg_user")}
  }

  if pgDBName = os.Getenv("pg_db_name"); pgDBName == "" {
    // if not found then, return error
    return nil, &EnvConfigError{errors.New("pg_db_name")}
  }

  if pgPass = os.Getenv("pg_pass"); pgPass == "" {
    // if not found then, return error
    return nil, &EnvConfigError{errors.New("pg_pass")}
  }

  config = &Config{
    Env: env,
    ProjectName: projName,
    PgHost: pgHost,
    PgPort: pgPort,
    PgUser: pgUser,
    PgPass: pgPass,
    PgDBName: pgDBName,
    Debug: debug,
    SrvPort: srvPort,
    PassEntropy: pEntropy,
  }

  return config, nil
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
