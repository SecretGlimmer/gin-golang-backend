package routes

import "github.com/gin-gonic/gin"

func SetupRouter() {
	router := gin.Default()

	router.GET("/fx-rates")
}
