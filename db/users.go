package db

type User struct {
	ID       uint    `json:"id" gorm:"primary key"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Phone    string `json:"phone" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Photo    string `json:"photo"`
	Address  string `json:"address"`
	RoleID   uint   `json:"role_id" gorm:"not null"`
	Role     Role   `gorm:"foreignKey:RoleID"`
}

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name" gorm:"unique"` 
	
}

func (*User) TableName() string {
	return "users"
}

func (*Role) TableName() string {
	return "roles"
}