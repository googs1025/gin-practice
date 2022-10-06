package result

import (
	"fmt"
	"gin-practice/src/validators"
)

type ErrorResult struct {
	data interface{}
	err  error
}

// 错误处理可以用这个抛panic，再由中间件捕获，统一抛错。
func (e *ErrorResult) UnWrap() interface{} {

	if e.err != nil {
		//fmt.Printf("%T\n", e.err)
		validators.CheckErrors(e.err)
		panic(e.err)
	}

	return e.data
}


// 仅能支持一个返回参数error
//func Result(err error) *ErrorResult {
//	return &ErrorResult{
//		err: err,
//	}
//}

func Result(values ...interface{}) *ErrorResult {
	if len(values) == 1{
		if values[0] == nil {
			return &ErrorResult{data: nil, err: nil}
		}

		if e, ok := values[0].(error); ok{
			return &ErrorResult{data: nil, err: e}
		}


	}

	if len(values) == 2 {
		if values[1] == nil {
			return &ErrorResult{data: values[0], err: nil}
		}

		if e, ok := values[1].(error); ok{
			return &ErrorResult{data: values[0], err: e}
		}
	}

	return &ErrorResult{data: nil, err: fmt.Errorf("传入的参数数量错误。")}


}
