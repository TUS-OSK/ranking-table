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
	test := Service{Name: "test"}
	db.FirstOrCreate(&test, test)

}

// test用result
func ResultSeed(db *gorm.DB) {
	var results []Result
	results = append(results, Result{ServiceID: 1, Score: 10, UserName: "user_1"})
	results = append(results, Result{ServiceID: 1, Score: 20, UserName: "user_2"})
	results = append(results, Result{ServiceID: 2, Score: 30, UserName: "user_a"})
	results = append(results, Result{ServiceID: 1, Score: 40, UserName: "user_3"})
	results = append(results, Result{ServiceID: 1, Score: 40, UserName: "user_4"})
	results = append(results, Result{ServiceID: 1, Score: 60, UserName: "user_5"})
	results = append(results, Result{ServiceID: 1, Score: 70, UserName: "user_6"})

	for _, result := range results {
		db.FirstOrCreate(&result, result)
	}
}
