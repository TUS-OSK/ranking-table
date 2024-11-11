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
	ServiceID uint   `json:"service_id" gorm:"not null;foreignKey:ServiceID;references:ID"`
	Score     int    `json:"score" gorm:"not null"`
	UserName  string `json:"user_name" gorm:"not null;unique"`
}
