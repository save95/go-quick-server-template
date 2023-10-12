package platform

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	ID uint `gorm:"type:INT(11) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`

	Genre  uint8 `gorm:"not null;uniqueIndex:uk_role"`
	UserID uint  `gorm:"not null;size:32;uniqueIndex:uk_role"`

	CreatedAt time.Time      `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"not null;default:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}
