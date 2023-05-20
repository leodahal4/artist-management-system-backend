package models

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/leodahal4/artist-management-system-backend/internal/config"
	"github.com/leodahal4/artist-management-system-backend/utils"
)

// The Music type represents a piece of music.
type Music struct {
	Timestamp

	// ArtistID is the ID of the artist who created the music.
	ArtistID uuid.UUID `json:"artist_id"`
	Artist Artist

  Title string `json:"title" validate:"string"`
  AlbumName string `json:"album_name" validate:"string"`

	// Genre is the genre of the music.
	Genre config.Genre `json:"genre" validate:"string,genre" gorm:"type:genre"`
}

func (music *Music) Validate(u *Artist) map[string][]string {
    invalidFields := make(map[string][]string)

    // Get the reflect.Type and reflect.Value of the input struct using the pointer
    musicType := reflect.TypeOf(*u)
    musicValue := reflect.ValueOf(*u)

    var wg sync.WaitGroup
    wg.Add(musicType.NumField())

    // Iterate through each field in the struct and validate if a "validate" tag is present
    for i := 0; i < musicType.NumField(); i++ {
        go func(i int){
            defer wg.Done()

            // Get the field, tag and value of the current field
            field := musicType.Field(i)
            tags := field.Tag.Get("validate")
            fieldValue := musicValue.Field(i)
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
