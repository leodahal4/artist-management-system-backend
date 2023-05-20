package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents the model for a user in the system.
type User struct {
  Timestamp // Timestamp is embedding a struct to include its fields in User

  FirstName   string      `json:"first_name" validate:"string,min=3,required"`
  LastName    string      `json:"last_name" validate:"string,min=3"`

  // email will be used as unique together with deleted_at, and phone
  Email       string      `json:"email" gorm:"index:user_email,unique" validate:"email,required"`
  Password    string      `json:"passsord" validate:"required"`
  Phone       string      `json:"phone" gorm:"index:user_email,unique" validate:"required,min=10,max=13"`
  Dob         time.Time   `json:"dob"`
  Gender      Gender      `json:"gender" validate:"string,gender" gorm:"type:gender"`
  Address     string      `json:"address"`
}

// BeforeCreate
//
// This function will be responsible for creating a new UUID for every user
// that registers.
// This function is created, as sometimes the PgSQL i8s unable to execute,
// uuid_generate_v4 or gen_random_uuid extensions
func (u *User) BeforeCreate(_ *gorm.DB) error {
  u.ID = uuid.New()
  return nil
}
