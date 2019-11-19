package db

import "time"

type UserStatus struct {
	ID 				uint       `gorm:"primary_key"`
	Name 			string     `gorm:"column:name" json:"name"`
	CreatedAt    	*time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    	*time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    	*time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (UserStatus) TableName() string {
	return "user_status"
}
