package models

// Gender is a custom type representing the gender of a user.
type Gender string

const (
  // Male represents male gender.
  Male Gender = "m"
  // Female represents female gender.
  Female Gender = "f"
  // Other represents non-binary or other gender.
  Other Gender = "o"
)

// The Genre type represents a music genre enum
type Genre string

// These constants represent the different music genres.
const (
	RythmAndBlues Genre = "rnb"
	Country       Genre = "country"
	Classic       Genre = "classic"
	Rock          Genre = "rock"
	Jazz          Genre = "jazz"
)
