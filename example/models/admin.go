package models

import "time"

type Admin struct {
	ID        uint      `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username  string    `gorm:"column:username;NOT NULL"`
	Mobile    string    `gorm:"column:mobile"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (m *Admin) TableName() string {
	return "admin"
}
