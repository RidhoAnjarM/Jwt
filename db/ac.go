package db

type AC struct {
	ID    uint    `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name" binding:"required"`
	Brand string  `json:"brand" binding:"required"`
	PK    float32 `json:"pk" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func (*AC) TableName() string {
	return "acs"
}
