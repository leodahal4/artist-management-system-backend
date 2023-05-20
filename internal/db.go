package internal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect() *gorm.DB {
  dbURI := "host=localhost user=leo dbname=artist password=leo port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dbURI))
  if err != nil {
    panic(err)
  }
  return db
}
