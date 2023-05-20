package models

import (
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/leodahal4/artist-management-system-backend/internal/config"
	"github.com/leodahal4/artist-management-system-backend/utils"

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
  Phone       string      `json:"phone" gorm:"index:user_email,unique" validate:"required,min=10,max=13,number"`
  Dob         time.Time   `json:"dob"`
  Gender      config.Gender      `json:"gender" validate:"string,gender" gorm:"type:gender"`
  Address     string      `json:"address"`
}

func (user *User) Validate(u *User) map[string][]string {
  invalidFields := make(map[string][]string)
  userType := reflect.TypeOf(*u)
  userValue := reflect.ValueOf(*u)
  var wg sync.WaitGroup
  wg.Add(userType.NumField())

  for i := 0; i < userType.NumField(); i++ {
    go func(i int){
      defer wg.Done()
      field := userType.Field(i)
      tags := field.Tag.Get("validate")
      fieldValue := userValue.Field(i)
      if tags != "" {
        var invalid []string
        for _, tag := range strings.Split((field.Tag.Get("validate")), ","){
          err := utils.ValidateThisTag(tag, fieldValue)
          if err != "" {
            invalid = append(invalid, err)
          }
        }
        if len(invalid) > 0 {
          invalidFields[field.Name] = invalid
        }
      }
    }(i)
  }
  wg.Wait()

  return invalidFields
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
