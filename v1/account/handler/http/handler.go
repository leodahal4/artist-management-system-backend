package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leodahal4/artist-management-system-backend/internal/interfaces"
	"github.com/leodahal4/artist-management-system-backend/v1/account"
)

func AccountRoutes(app fiber.Router, useCase interfaces.AccountUseCase) {
  app.Post("/register", RegisterationHandler(useCase))
}

func RegisterationHandler(useCase interfaces.AccountUseCase) fiber.Handler {
  return func(c *fiber.Ctx) error {
    var req account.UserRegisterRequest
    
    err := c.BodyParser(req)
    if err != nil {
      c.Status(http.StatusBadRequest)
      return c.JSON(err)
    }
    
    c.Status(http.StatusOK)
    return c.JSON("OK")
  }
}
