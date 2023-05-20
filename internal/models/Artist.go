package models

import "time"


// Artist represents a music artist.
type Artist struct{
  Timestamp

  // Name is the name of the artist.
  Name                string    `json:"artist_name"`

  // Dob is the date of birth of the artist.
  Dob                 time.Time `json:"dob"`

  // Gender is the gender of the artist.
  Gender              Gender    `json:"gender" validate:"string,gender" gorm:"type:gender"`

  // Address is the address of the artist.
  Address             string    `json:"address"`

  // FirstReleaseYear is the year when the artist released their first album.
  FirstReleaseYear    int       `json:"release_date"`

  // NoOfAlbumsReleased is the number of albums released by the artist.
  NoOfAlbumsReleased  int       `json:"no_of_albums_released"`
}
