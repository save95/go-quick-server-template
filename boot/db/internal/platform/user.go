package platform

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement;not null"`

	Genre       uint8  `gorm:"not null;uniqueIndex:uk_account"`
	Account     string `gorm:"not null;size:32;uniqueIndex:uk_account"`
	Nickname    string `gorm:"size:32;default:''"`
	AvatarURL   string `gorm:"not null;size:128;default:''"`
	Password    string `gorm:"not null;size:128"`
	State       int8   `gorm:"not null;default:0"`
	LastLoginAt *time.Time
	LastLoginIP string `gorm:"size:32"`
	DriverNo    string `gorm:"not null;size:32;default:''"`

	CreatedAt time.Time      `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"not null;default:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}
