package repo

import (
	inter "github.com/ratheeshkumar25/opti_cut_userservice/pkg/repo/interfaces"
	"gorm.io/gorm"
)

// UserRepository for connecting db to UserRepoInter methods
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates an instance of user repo
func NewUserRepository(db *gorm.DB) inter.UserRepoInter {
	return &UserRepository{
		DB: db,
	}
}
