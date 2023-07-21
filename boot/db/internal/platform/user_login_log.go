package platform

import (
	"time"

	"gorm.io/gorm"
)

type UserLoginLog struct {
	ID uint `gorm:"primaryKey;autoIncrement;not null"`

	UserID    uint   `gorm:"no null;type:INT(11) UNSIGNED;index:idx_user"`
	UserAgent string `gorm:"no null;size:512"`
	IP        string `gorm:"no null;column:ip;size:24"`
	Referer   string `gorm:"no null;size:256"`

	CreatedAt time.Time      `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"not null;default:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}
