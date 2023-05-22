package usecase

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/leodahal4/artist-management-system-backend/internal/interfaces"
	"github.com/leodahal4/artist-management-system-backend/internal/models"
	"github.com/leodahal4/artist-management-system-backend/v1/account"
)

type useCase struct {
	repo  interfaces.AccountRepository
	redisClient *redis.Client
}

func New(r interfaces.AccountRepository, rc *redis.Client) interfaces.AccountUseCase {
	return &useCase{
		repo:  r,
		redisClient: rc,
	}
}

func (u *useCase) LoginService(user *account.UserLoginRequest) (account.LoginResponse, error) { return account.LoginResponse{}, nil }
func (u *useCase) LogoutService(c *fiber.Ctx) error {return nil}
func (u *useCase) CreateAccessToken(userID string) (string, error) {return "", nil}
func (u *useCase) CreateRefreshToken(userID string) (string, error) {return "", nil}

func (u *useCase) NewUser(user *models.User) (account.LoginResponse, error) {
  // validate if the email is unique
  _, err := u.repo.GetUserByEmail(user.Email)
  if err == nil {
    // email is already registered with another user
    return account.LoginResponse{}, errors.New("user with that email already exists")
  }

  // validate if the phone is unique
  _, err = u.repo.GetUserByPhone(user.Phone)
  if err == nil {
    // phone is already registered with another user
    return account.LoginResponse{}, errors.New("user with that phone already exists")
  }

  // its now ok to proceed on
  err = u.repo.NewUser(user)
  return account.LoginResponse{}, err
}
