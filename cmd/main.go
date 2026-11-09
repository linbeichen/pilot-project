package main

import (
	"github.com/gin-gonic/gin"
	"shanyangcode.com/pilot/internal/controller"
)

func main() {

	// enginc := gin.Default()
	engine := gin.Default()
	/*
		engine.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	*/

	chatController := controller.NewChatController("gpt-3.5-turbo")

	api := engine.Group("/api/v1")
	{
		chat := api.Group("/chat")
		{
			chat.POST("", chatController.Chat)
		}
	}
	engine.Run(":8080")
}
