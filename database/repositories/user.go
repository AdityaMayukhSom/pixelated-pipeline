package repositories

import (
	"time"

	"devstream.in/pixelated-pipeline/services/models"
)

type UserEntity struct {
	Name      string `gorm:"not null;size:64"`
	Email     string `gorm:"uniqueIndex;not null;size:255"`
	Username  string `gorm:"primaryKey;not null;unique"`
	Password  string `gorm:"not null;size:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts     []PostEntity
}

// TableName overrides the table name used by UserEntity, see Tabler interface of gorm
func (UserEntity) TableName() string {
	return "users"
}

type UserRepository interface {
	FindUserByUsername(username string)
	FindUserByEmail(email string)

	DeleteUserByUsername(username string)
	DeleteUserByEmail(email string)

	CreateUser(user models.User)
	UpdateUser(user models.User)
}