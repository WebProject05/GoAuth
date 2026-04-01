package httpserver

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	// The below router does not have any default middlewares
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery()) // This will recovery from any panics like that

	r.GET("/health", health)
	return r
}