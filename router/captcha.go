package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ChatRoomRouter struct{}

func (c *ChatRoomRouter) InitChatRoomRouter(Router *gin.RouterGroup) {
	//chatRoomRouter := Router.Use(middleware.OperationRecord())
	captchaRouter := Router
	{
		fmt.Println("captcha router init...", captchaRouter)
	}
}
