package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

// JsonResult 统一返回的格式
type JsonResult struct {
	Message string `json:"message"`
	Code 	string `json:"code"`
	Result 	interface{} `json:"result"`
}

// 注：不用每次都用NewJsonResult的方式初始化，可以适当用Pool来进行存储。
// NewJsonResult 构建函数
func NewJsonResult(message string, code string, result interface{}) *JsonResult {
	return &JsonResult{
		Message: message,
		Code: code,
		Result: result,
	}
}


var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJsonResult("", "", nil)
		},
	}
}

type ResultFunc func(message string, code string, result interface{}) func(output OutFunc)
type OutFunc func(c *gin.Context, v interface{})

// 类似装饰器函数
func R(c *gin.Context) ResultFunc {
	return func(message string, code string, result interface{}) func(output OutFunc) {
		// 记得要断言成该类型！
		r := ResultPool.Get().(*JsonResult)
		defer ResultPool.Put(r)

		r.Message = message
		r.Code = code
		r.Result = result
		//c.JSON(200, r)
		return func(output OutFunc) {
			output(c, r)
		}
	}
}

// 分成OK Error两种

func OK(c *gin.Context, v interface{}) {
	c.JSON(200, v)
}

func Error(c *gin.Context, v interface{}) {
	c.JSON(400, v)
}

func OKWithString(c *gin.Context, v interface{}) {
	c.String(200, fmt.Sprintf("%v", v))
}


// 下列的方式会有很多冗余代码，ok跟error长得太像了。
//func OK(c *gin.Context) ResultFunc {
//	return func(message string, code string, result interface{}) {
//		r := ResultPool.Get().(*JsonResult)
//		defer ResultPool.Put(r)
//
//		r.Message = message
//		r.Code = code
//		r.Result = result
//		c.JSON(200, r)
//	}
//}

//func Error(c *gin.Context) ResultFunc {
//	return func(message string, code string, result interface{}) {
//		r := ResultPool.Get().(*JsonResult)
//		defer ResultPool.Put(r)
//
//		r.Message = message
//		r.Code = code
//		r.Result = result
//		c.JSON(400, r)
//	}
//}

