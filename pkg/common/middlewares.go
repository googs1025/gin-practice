package common

import "github.com/gin-gonic/gin"

// TODO: 中间件预计提供日志 和其他模块功能

// 使用中间件方式，捕获请求的错误。
// 这样只要错误，就能抛panic，让中间件来处理并且返回错误。
func ErrorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				err := e.(error)
				c.JSON(400, gin.H{"message": err.Error()})
			}
		}()
		c.Next()
	}
}
