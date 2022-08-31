package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: code, Data: data, Msg: msg})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, "ok", "操作成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "fail", "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}
