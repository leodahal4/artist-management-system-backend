package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/leodahal4/artist-management-system-backend/internal/config"
	"gorm.io/gorm"
)

type server struct {
  DB *gorm.DB
  RedisClient *redis.Client
  App *fiber.App
  Router *fiber.Router
}

func main() {
  fmt.Println("[+] UNDER DEVELOPMENT [+]")
  config, err := config.Read_config()
  if err != nil{
    log.Fatalf("Cannot read the .env file.\nERR: %s", err.Error())
  }
  fmt.Printf("config: %v\n", config)
}
