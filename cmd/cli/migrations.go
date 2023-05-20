package main

import (
	"fmt"
  "github.com/leodahal4/artist-management-system-backend/internal"
  "github.com/leodahal4/artist-management-system-backend/internal/models"
)

func main() {
  fmt.Println("[+] Migrating the models [+]")
  db := internal.DbConnect()

  if err := db.AutoMigrate(&models.User{}); err != nil {
    println(err)
  }

  if err := db.AutoMigrate(&models.Artist{}); err != nil {
    println(err)
  }

  if err := db.AutoMigrate(&models.Music{}); err != nil {
    println(err)
  }

  fmt.Println("[+] Successfully migrated all the models [+]")
}
