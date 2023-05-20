package account

import (
	"time"

	"github.com/google/uuid"
)

type UserLoginRequest struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Refresh string            `json:"refresh"`
	Access  string            `json:"access"`
	User    UserLoginResponse `json:"user"`
}

type UserLoginResponse struct {
	ID          uuid.UUID `json:"id"`
  User
}

type UserRegisterRequest struct {
  User
  Password    string    `json:"password"`
}

type User struct {
	FirstName   string    `json:"firstName,omitempty"`
	LastName    string    `json:"lastName,omitempty"`
	Email       string    `json:"email,omitempty"`
	Phone       string    `json:"phone,omitempty"`
  DOB         time.Time `json:"dob"`
  Gender      string    `json:"gender"`
  Address     string    `json:"address"`
}
