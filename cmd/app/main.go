package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/leodahal4/artist-management-system-backend/internal/config"
	"github.com/leodahal4/artist-management-system-backend/internal"
	accountUseCase "github.com/leodahal4/artist-management-system-backend/v1/account/usecase"
	accountRepo "github.com/leodahal4/artist-management-system-backend/v1/account/repository"
	accountHandler "github.com/leodahal4/artist-management-system-backend/v1/account/handler/http"
	"gorm.io/gorm"
)

type server struct {
  dB *gorm.DB
  redisClient *redis.Client
  app *fiber.App
  router *fiber.Router
  config *config.Config
}

func main() {
  fmt.Println("[+] UNDER DEVELOPMENT [+]")
  var newServer server
  config, err := config.Read_config()
  if err != nil{
    log.Fatalf("Cannot read the .env file.\nERR: %s", err.Error())
  }

  newServer.config = &config
  // initialize the db
  newServer.dB = internal.DbConnect(&config)

  newServer.app = fiber.New()

  authRoute := newServer.app.Group("/api/v1/user")
  newServer.app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${time} ${status} - ${method} ${path} ${time}\n",
	}))

  auseCase := accountUseCase.New(accountRepo.NewRepo(newServer.dB), newServer.redisClient)
  accountHandler.AccountRoutes(authRoute, auseCase)
  
  log.Fatal(newServer.app.Listen(fmt.Sprintf(":%v", config.ServerPort)))
}
