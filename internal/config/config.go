package config

import "strings"

// Config represents the configuration of the server.
type Config struct {
  // Project_name represents the name of the project.
  ProjectName string

  Env string // this determines whether the server is in production or development environment.

  PgHost string
  PgPort string
  PgUser string
  PgPass string
  PgDBName string

  Debug bool // if not found this will be set to false if on prod and true if on development
}

// Is_prod returns true if the server is in production environment.
func (c *Config) Is_prod() bool {
  if strings.ToLower(c.Env) == "production" || strings.ToLower(c.Env) == "prod" {
    return true
  } else{
    return false
  }
}

