package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()
	// Задаем путь 0.0.0.0:8080/ping  для вывода контекста
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // Прослушивает 0.0.0.0:8080

	log.Println("Server down")
}