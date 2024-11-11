package models

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
}

type Result struct {
	gorm.Model
	ServiceID uint   `json:"service_id" gorm:"not null"`
	Score     int    `json:"score" gorm:"not null"`
}
