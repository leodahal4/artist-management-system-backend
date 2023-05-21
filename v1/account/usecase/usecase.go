package usecase

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/leodahal4/artist-management-system-backend/internal/interfaces"
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
func (u *useCase) NewUser(account.UserRegisterRequest) (account.LoginResponse, error) {return account.LoginResponse{}, nil}
