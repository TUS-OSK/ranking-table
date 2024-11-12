package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ranking-table/db"
	"ranking-table/models"

)

func main() {
	databaseInit()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func databaseInit() *gorm.DB {
	database := db.Connect()
	db.Migration(database)
	models.ServiceSeed(database)
	db.ShowAllTableHeader(database)
	return database
}
