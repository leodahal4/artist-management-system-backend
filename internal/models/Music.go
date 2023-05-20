package models

import "github.com/google/uuid"

// The Music type represents a piece of music.
type Music struct {
	Timestamp

	// ArtistID is the ID of the artist who created the music.
	ArtistID uuid.UUID `json:"artist_id"`
	Artist Artist

	Title string `json:"title"`
	AlbumName string `json:"album_name"`

	// Genre is the genre of the music.
	Genre Genre `json:"genre" validate:"string,genre" gorm:"type:genre"`
}
