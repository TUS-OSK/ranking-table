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

func Migration() {
	db := Connect()
	db.AutoMigrate(&models.Service{}, &models.Result{})
	fmt.Println("Migration is done.")
}

func ShowAllTables() {
	db := Connect()
	var tables []string
	db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables)
	fmt.Println("existing tables: ", tables)
}
