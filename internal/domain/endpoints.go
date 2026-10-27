package domain

import "time"

type Endpoints struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	URL         string    `gorm:"type:text;not null"`
	Events      []string  `gorm:"type:jsonb;serializer:json"`
	Description *string   `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdateAt    time.Time `gorm:"autoUpdateTime"`
}
