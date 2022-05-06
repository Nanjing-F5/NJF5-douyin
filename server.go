package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// http://localhost:8080/api/v1/hello
		v1.GET("/hello", func(context *gin.Context) {
			context.String(200, "hello world!")
		})
		// http://localhost:8080/api/v1/ping
		v1.GET("/ping", func(context *gin.Context) {
			content := struct {
				Now string `json:"now"`
			}{Now: time.Now().Format(time.RFC3339)}
			context.JSON(200, content)
		})
	}
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
