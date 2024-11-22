package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ranking-table/db"
	"ranking-table/models"
)

func main() {
	database := databaseInit()
	db.ShowAllTableHeader(database) // development用

	router := gin.Default()

	// "/" ルートにアクセスすると {"message": "Hai"} を返す
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hai"})
	})

	router.GET("/osk", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"osk": true})
	})

	//　練習課題2.1
	router.POST("/", func(c *gin.Context) {
		var json struct {
			Message string `json:"message"`
		}
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": json.Message})
	})

	// 練習課題2.2
	router.GET("/tus/:id", func(c *gin.Context) {
		id := c.Param("id")

		if id == "1" {
			c.JSON(http.StatusOK, gin.H{"id": id})
			return
		}
		if id != "1" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
	})

	//本実装１
	router.POST("/:id", func(c *gin.Context) {
	
		serviceID := c.Param("id")
	
		var json struct {
			Score    int    `json:"score"`
			UserName string `json:"username"`
		}
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		newResult := models.Result{
			ServiceID: serviceID
			Score:     json.Score
			UserName:  json.UserName
		}
	
		// 失敗したときの実装だが、よくわからない
		result := database.Create(&newResult)
		// result.Errorはgormライブラリ内で定義されているらしい
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	
		// 成功したとき
		c.JSON(http.StatusCreated, gin.H{"message": "success"})
	})

	router.Run(":8080")

}

func databaseInit() *gorm.DB {
	database := db.Connect()
	db.Clean(database) // development用
	db.Migration(database)
	models.ServiceSeed(database) // development用を含む
	models.ResultSeed(database)  // development用
	return database
}
