package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var wishlist = []string{"health", "love", "car", "flat", "cat"}

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()
	// Задаем путь 0.0.0.0:8080/ping  для вывода контекста
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main Website",
		})
	})

	r.GET("/wishes", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "My WishList",
			"wishlist": wishlist,
		})
	})

	// Добавление статических файлов
	r.Static("/image", "./resources")

	r.Run() // Прослушивает 0.0.0.0:8080

	log.Println("Server down")
}
