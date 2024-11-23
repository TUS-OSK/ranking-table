package models

import (
	"gorm.io/gorm"
)

// Service table model
type Service struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;unique"`
}

// Result table model
type Result struct {
	gorm.Model
	ServiceID uint   `json:"service_id" gorm:"not null;foreignKey:ServiceID;references:ID"`
	Score     int    `json:"score" gorm:"not null"`
	UserName  string `json:"user_name" gorm:"not null;unique"`
}

func ServiceSeed(db *gorm.DB) {
	wikipediaGolf := Service{Name: "wikipedia golf"}
	db.FirstOrCreate(&wikipediaGolf, wikipediaGolf)
	// test用service
	test := Service{Name: "qgomoku"}
	db.FirstOrCreate(&test, test)

}


func Get5thResultSortedByScore(id uint, db *gorm.DB) []Result {
	var results []Result
	db.Where("service_id = ? AND deleted_at IS NULL", id).Order("score asc").Order("created_at asc").Limit(5).Find(&results)
	return results
}

func DeleteAllResults(db *gorm.DB) {
	db.Delete(&Result{}, "1 = 1")
}