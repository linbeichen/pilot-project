package controller

import (
	"net/http"

	"shanyangcode.com/pilot/internal/httpapi"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	agent string
}

func NewChatController(agent string) *ChatController {
	return &ChatController{
		agent: agent,
	}
}

type ChatRequest struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	Query     string `json:"query"`
}

func (cc *ChatController) Chat(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpapi.Response{
			Success: false,
			Code:    httpapi.CodeInvalidArgument,
			Message: "请求参数不合法",
		})
	}

	c.JSON(http.StatusOK, httpapi.Response{
		Success: true,
		Code:    httpapi.CodeOK,
		Message: "请求成功",
		Data:    req,
	})
}
