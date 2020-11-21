package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(StatusCode int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(StatusCode, Response{
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, "Success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, "Success", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, data, message, c)
}

func Fail(c *gin.Context) {
	Result(http.StatusInternalServerError, map[string]interface{}{}, "InternalServerError", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusInternalServerError, map[string]interface{}{}, message, c)
}

func FailWithCodeAndMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
