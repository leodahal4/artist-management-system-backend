package http

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leodahal4/artist-management-system-backend/internal/config"
	"github.com/leodahal4/artist-management-system-backend/internal/interfaces"
	"github.com/leodahal4/artist-management-system-backend/internal/models"
	"github.com/leodahal4/artist-management-system-backend/utils"
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

    // check if all the validation gets passed
    errors := req.Validate(&req)
    if len(errors) > 0 {
      return c.JSON(errors)
    } else {
      // check if the password is acceptable
      err := utils.CheckPasswordStrength(req.Password)
      if err != nil {
        errors := make(map[string][]string, 1)
        errors["password"] = []string{err.Error()}
        return c.Status(http.StatusBadRequest).JSON(errors)
      }
    }
    
    res, err := useCase.NewUser(&req)
    if err != nil {
      newError := config.ErrorResponse{Err: err}
      return c.Status(http.StatusBadRequest).JSON(newError.Errors())
    }

    c.Status(http.StatusOK)
    return c.JSON(res)
  }
}
