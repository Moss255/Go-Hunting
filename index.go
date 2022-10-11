package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Bug struct {
	ID   int
	Name string
	Type string
}

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("SQL_USER"), os.Getenv("SQL_PASS"), os.Getenv("SQL_HOST"), os.Getenv("SQL_PORT"), os.Getenv("SQL_DB"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Bug{})

	if err != nil {
		panic(err)
	}

	router := gin.Default()
	// router.LoadHTMLGlob("templates/*")

	router.Static("/assets", "www/assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.go.tmpl", gin.H{})
	})

	router.GET("/hunting/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		var bug = Bug{ID: id}
		db.First(&bug)
		c.JSON(http.StatusOK, bug)
	})

	router.GET("/hunting/search", func(c *gin.Context) {
		var out interface{}
		name := c.Query("name")
		db.Raw("SELECT * FROM bugs where name = ?", name).Scan(&out)
		c.JSON(http.StatusOK, &out)
	})

	router.POST("/hunting", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.PATCH("/hunting", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.DELETE("/hunting", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Atoi(id string) {
	panic("unimplemented")
}
