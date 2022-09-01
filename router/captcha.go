package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CaptchaApiRouter struct{}

func (c *CaptchaApiRouter) InitCaptchaApiRouter(Router *gin.RouterGroup) {
	//captchaApiRouter := Router.Use(middleware.OperationRecord())
	captchaRouter := Router
	{
		fmt.Println("captcha router init...", captchaRouter)
	}
}
