package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// "/" ルートにアクセスすると {"message": "Hai"} を返す
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hai"})
	})

	router.GET("/osk", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"osk": true})
	})

	router.GET("/tus/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": intID})
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

	router.Run(":8080")

}
