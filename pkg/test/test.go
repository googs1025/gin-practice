package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 测试 UnWrap错误处理可支持多返回值，所写的函数
func GetInfo(id int) (gin.H, error) {
	if id > 10 {
		return gin.H{"message": "result"}, nil
	} else {
		return nil, fmt.Errorf("test error")
	}
}
