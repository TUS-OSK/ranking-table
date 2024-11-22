package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ranking-table/db"
	"ranking-table/models"
	"strconv"
)

func main() {
	database := databaseInit()
	db.ShowAllTableHeader(database) // development用

	router := gin.Default()

	// ここからCorsの設定
	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{"*"},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{"*"},
	}))

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

		serviceID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

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
			ServiceID: uint(serviceID),
			Score:     json.Score,
			UserName:  json.UserName,
		}

		// 失敗したときの実装だが、よくわからない
		result := database.Create(&newResult)
		// result.Errorはgormライブラリ内で定義されているらしい
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// 成功したとき
		c.JSON(http.StatusCreated, newResult)
	})

	// 本実装2
	router.GET("/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		// 上位5件のスコアを取得する
		results := models.Get5thResultSortedByScore(uint(id), database)

		type responseUser struct {
			Name  string `json:"name"`
			Score int    `json:"score"`
		}

		var response []responseUser
		for _, result := range results {
			response = append(response, responseUser{Name: result.UserName, Score: result.Score})
		}

		c.JSON(http.StatusOK, response)
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
