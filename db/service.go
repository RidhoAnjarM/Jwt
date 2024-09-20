package db

type Service struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	TechnicianID uint      `json:"technician_id" binding:"required"`
	ClientID     uint      `json:"client_id" binding:"required"`
	ACID         uint      `json:"ac_id" binding:"required"`
	Date         string    `json:"date" binding:"required"`
	Status       string    `json:"status" binding:"required"`
}

func (*Service) TableName() string {
	return "services"
}
