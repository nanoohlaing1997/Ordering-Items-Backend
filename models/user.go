package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Address      string    `json:"address"`
	Status       int32     `json:"status"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type UserDB struct {
	db    *gorm.DB
	model *User
}

func UserManager(db *gorm.DB) *UserDB {
	return &UserDB{
		db:    db,
		model: &User{},
	}
}

func (udb *UserDB) CreateUser(user *User) (*User, error) {
	if res := udb.db.Create(&user); res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return user, nil
}

func (udb *UserDB) GetUser(email string) (*User, error) {
	var user User
	if res := udb.db.First(&user, "email = ?", email); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
