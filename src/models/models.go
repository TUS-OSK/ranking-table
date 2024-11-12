package models

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;unique"`
}

type Result struct {
	gorm.Model
	ServiceID uint   `json:"service_id" gorm:"not null;foreignKey:ServiceID;references:ID"`
	Score     int    `json:"score" gorm:"not null"`
	UserName  string `json:"user_name" gorm:"not null;unique"`
}

func ServiceSeed(db *gorm.DB) {
	wikipediaGolf := Service{Name: "wikipedia golf"}
	//既に存在している場合は追加しない
	db.FirstOrCreate(&wikipediaGolf, wikipediaGolf)
}
