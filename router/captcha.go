package router

import (
	"fmt"
	"gin-plugin-captcha/api"

	"github.com/gin-gonic/gin"
)

type CaptchaApiRouter struct{}

func (c *CaptchaApiRouter) InitCaptchaApiRouter(Router *gin.RouterGroup) {
	//captchaApiRouter := Router.Use(middleware.OperationRecord())
	captchaRouter := Router
	{
		fmt.Println("captcha router init...", captchaRouter)
		captchaRouter.GET("", api.ApiGroupApp.Get)
		captchaRouter.POST("", api.ApiGroupApp.Verify)
		captchaRouter.POST("library", api.ApiGroupApp.AddCaptcha)
	}
}
