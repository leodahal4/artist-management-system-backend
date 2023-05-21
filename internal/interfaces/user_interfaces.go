package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/leodahal4/artist-management-system-backend/v1/account"
)

type AccountUseCase interface {
	LoginService(user *account.UserLoginRequest) (account.LoginResponse, error)
	LogoutService(c *fiber.Ctx) error
	CreateAccessToken(userID string) (string, error)
	CreateRefreshToken(userID string) (string, error)
  NewUser(account.UserRegisterRequest) (account.LoginResponse, error)
}

type AccountRepository interface {
	GetUserByEmail(email string) (*account.User, error)
	GetUserByID(id uuid.UUID) (*account.User, error)
  NewUser(account.UserRegisterRequest) error
}

