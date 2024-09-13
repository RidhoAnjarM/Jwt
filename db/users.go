package db

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Phone     string    `json:"phone" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime,column:updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (*User) TableName() string {
	return "users"
}
