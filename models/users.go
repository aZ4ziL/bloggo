package models

import (
	"time"

	"github.com/aZ4ziL/bloggo/utils"
)

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	FirstName   string    `gorm:"size:100;null" json:"first_name"`
	LastName    string    `gorm:"size:100;null" json:"last_name"`
	Username    string    `gorm:"size:100;unique;index" json:"username"`
	Email       string    `gorm:"size:100;unique;index" json:"email"`
	Password    string    `gorm:"size:255"`
	IsSuperuser bool      `gorm:"default:0" json:"is_superuser"`
	IsStaff     bool      `gorm:"default:0" json:"is_staff"`
	IsActive    bool      `gorm:"default:1" json:"is_active"`
	LastLogin   time.Time `gorm:"null" json:"last_login"`
	DateJoined  time.Time `gorm:"autoCreateTime" json:"date_joined"`
	Articles    []Article `gorm:"foreignKey:AuthorID" json:"articles"`
}

// CreateUser create new user
func CreateUser(user *User) error {
	user.Password = utils.EncryptionPassword(user.Password)
	err := db.Create(user).Error
	return err
}

// GetUserByID get user by passing the ID
func GetUserByID(id uint) (User, error) {
	var user User
	err := db.Model(&User{}).Where("id = ?", id).Preload("Articles").First(&user).Error
	return user, err
}

// GetUserByUsername get user by passing the Username
func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.Model(&User{}).Where("username = ?", username).Preload("Articles").First(&user).Error
	return user, err
}
