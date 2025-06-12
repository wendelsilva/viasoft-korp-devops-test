package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	current_datetime := time.Now().UTC()
	server := gin.Default()
	server.GET("/", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{
			"nome":    "Projeto korp",
			"horario": current_datetime,
		})
	})

	server.Run()
}
