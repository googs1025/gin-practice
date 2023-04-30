package validators

import (
	"github.com/go-playground/validator/v10"
)

func init() {
	// 第一阶段：不封装校验函数
	//if err := myValid.RegisterValidation("UserName", VUserName); err != nil {
	//	log.Fatal("valid UserName error!")
	//}
	// 第二阶段：封装校验函数
	//if err := myValid.RegisterValidation("UserName", UserName("required,min=4").ToFunc()); err != nil {
	//	log.Fatal("valid UserName error!")
	//}
	// 第三阶段：封装注册valid函数
	registerValidation("UserName", UserName("required").ToFunc())

}

//// 校验函数
//var VUserName validator.Func = func(fl validator.FieldLevel) bool {
//	uname, ok := fl.Field().Interface().(string)
//	if ok && len(uname) >= 4 {
//		return true
//	} else {
//		return false
//	}
//}

// validate 自定义校验逻辑
func (this UserName) validate(v string) bool {
	err := myValid.Var(v, string(this))
	if err != nil {
		return false
	}

	// 这里可以自己扩展要用的校验，ex:限制最大长度
	//if len(v) > 256 {
	//	return false
	//}

	return true
}

type UserName string

func (this UserName) ToFunc() validator.Func {
	validatorError["UserName"] = "用户名必须长度大于4"
	return func(fl validator.FieldLevel) bool {
		uname, ok := fl.Field().Interface().(string)
		if ok {
			return this.validate(uname)
		}
		return false
	}
}
