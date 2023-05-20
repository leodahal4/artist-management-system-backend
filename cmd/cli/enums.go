package main

import (
  "github.com/leodahal4/artist-management-system-backend/internal"
  "fmt"
)

func main() {
  fmt.Println("[+] Creating the enums [+]")
  db := internal.DbConnect()
  // gender enums
  err := db.Exec("CREATE TYPE gender AS ENUM('m','f','o');").Error
  if err != nil {
    panic(err)
  }

  err = db.Exec("CREATE TYPE genre AS ENUM('rnb','country','classic','rock', 'jazz');").Error
  if err != nil {
    panic(err)
  }
  fmt.Println("[+] Done [+]")
}
