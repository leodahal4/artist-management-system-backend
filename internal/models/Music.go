package models

import (
	"github.com/google/uuid"
	"github.com/leodahal4/artist-management-system-backend/internal/config"
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
