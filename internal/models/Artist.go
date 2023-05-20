package models

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/leodahal4/artist-management-system-backend/internal/config"
	"github.com/leodahal4/artist-management-system-backend/utils"
)

// Artist represents a music artist.
type Artist struct{
  Timestamp

  // Name is the name of the artist.
  Name                string    `json:"artist_name" validate:"string,required"`

  // Dob is the date of birth of the artist.
  Dob                 time.Time `json:"dob"`

  // Gender is the gender of the artist.
  Gender              config.Gender    `json:"gender" validate:"string,gender" gorm:"type:gender"`

  // Address is the address of the artist.
  Address             string    `json:"address"`

  // FirstReleaseYear is the year when the artist released their first album.
  FirstReleaseYear    int       `json:"release_date" validate:"required,number"`

  // NoOfAlbumsReleased is the number of albums released by the artist.
  NoOfAlbumsReleased  int       `json:"no_of_albums_released" validate:"required,number" gorm:"default:0"`
}

func (artist *Artist) Validate(u *Artist) map[string][]string {
    invalidFields := make(map[string][]string)

    // Get the reflect.Type and reflect.Value of the input struct using the pointer
    artistType := reflect.TypeOf(*u)
    artistValue := reflect.ValueOf(*u)

    var wg sync.WaitGroup
    wg.Add(artistType.NumField())

    // Iterate through each field in the struct and validate if a "validate" tag is present
    for i := 0; i < artistType.NumField(); i++ {
        go func(i int){
            defer wg.Done()

            // Get the field, tag and value of the current field
            field := artistType.Field(i)
            tags := field.Tag.Get("validate")
            fieldValue := artistValue.Field(i)
            fmt.Printf("tags: %v\n", tags)

            // If a validate tag is present, validate the field against each tag
            if tags != "" {
                var invalid []string
                for _, tag := range strings.Split((field.Tag.Get("validate")), ","){
                    err := utils.ValidateThisTag(tag, fieldValue)
                    if err != "" {
                        invalid = append(invalid, err)
                    }
                }
                // If there are any invalid fields, add them to the invalidFields map
                if len(invalid) > 0 {
                    invalidFields[field.Name] = invalid
                }
            }
        }(i)
    }
    wg.Wait()

    return invalidFields
}
