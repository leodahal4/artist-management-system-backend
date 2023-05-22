package http

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leodahal4/artist-management-system-backend/internal/interfaces"
	"github.com/leodahal4/artist-management-system-backend/internal/models"
)

func AccountRoutes(app fiber.Router, useCase interfaces.AccountUseCase) {
  app.Post("/register", RegisterationHandler(useCase))
}

func RegisterationHandler(useCase interfaces.AccountUseCase) fiber.Handler {
  return func(c *fiber.Ctx) error {
    var req models.User
    err := c.BodyParser(&req)
    if err != nil {
      c.Status(http.StatusBadRequest)
      return c.JSON(err)
    }

    fmt.Printf("req: %v\n", req)

    errors := req.Validate(&req)
    if len(errors) > 0 {
      return c.JSON(errors)
    }

    c.Status(http.StatusOK)
    return c.JSON("OK")
  }
}
