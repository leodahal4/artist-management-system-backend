package internal

import (
	"fmt"

	"github.com/leodahal4/artist-management-system-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect(conf *config.Config) *gorm.DB {
  dbURI := fmt.Sprintf(
    "host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
    conf.PgHost,
    conf.PgUser,
    conf.PgDBName,
    conf.PgPass,
    conf.PgPort,
  )
  db, err := gorm.Open(postgres.Open(dbURI))
  if err != nil {
    panic(err)
  }
  return db
}
