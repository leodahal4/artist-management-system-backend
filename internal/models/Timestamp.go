package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Timestamp represents a struct with fields to store information related to timestamps.
type Timestamp struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	
	CreatedAt time.Time      `gorm:"autoCreateTime:nano" json:"-"`
	UpdatedAt time.Time      `gorm:"autoCreateTime:nano" json:"-"`
  DeletedAt gorm.DeletedAt `gorm:"index:user_email,unique" json:"-"`
	
  // is_deleted represents whether the struct is deleted or not.
  IsDeleted bool           `gorm:"default:false"`
  // deleted_by represents the user who deleted the struct.
	DeletedBy string
}
