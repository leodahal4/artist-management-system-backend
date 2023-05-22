package repository

import (
	"github.com/google/uuid"
	"github.com/leodahal4/artist-management-system-backend/internal/interfaces"
	"github.com/leodahal4/artist-management-system-backend/internal/models"
	"github.com/leodahal4/artist-management-system-backend/v1/account"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) interfaces.AccountRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) GetUserByEmail(string) (*account.User, error) {
  return nil, nil
}

func (r *repo) GetUserByPhone(string) (*account.User, error) {
  return nil, nil
}

func (r *repo) GetUserByID(uuid.UUID) (*account.User, error) {
  return nil, nil
}

func (r *repo) NewUser(*models.User) (error) {
  return nil
}
