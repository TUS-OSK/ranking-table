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


func CreateService(db *gorm.DB, name string) {
	service := Service{Name: name}
	db.Create(&service)
}

func FindServiceByID(db *gorm.DB, id int) Service {
	var service Service
	db.First(&service, id)
	return service
}

func CreateResult(db *gorm.DB, serviceID int, score int, userName string) {
	result := Result{ServiceID: uint(serviceID), Score: score, UserName: userName}
	db.Create(&result)
}

func GetDescSortedResult(db *gorm.DB, serviceID int) []Result {
	var results []Result
	db.Order("score desc").Where("service_id = ?", serviceID).Find(&results)
	return results
}


