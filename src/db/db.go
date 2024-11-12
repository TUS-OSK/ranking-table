package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ranking-table/models"
)

func Connect() *gorm.DB {
	dsn := "host=db user=user password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Clean(db *gorm.DB) {
	db.Exec("DELETE FROM results")
	db.Exec("DELETE FROM services")
	fmt.Println("Clean is done.")
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Service{}, &models.Result{})
	fmt.Println("Migration is done.")
}

func ShowAllTableHeader(db *gorm.DB) {
	var tables []string
	db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables)
	for _, table := range tables {
		var headers []map[string]interface{}
		db.Raw(fmt.Sprintf("SELECT * FROM %s", table)).Scan(&headers)
		fmt.Printf("Table: %s\n", table)
		for _, header := range headers {
			fmt.Println(header)
		}
	}
}
