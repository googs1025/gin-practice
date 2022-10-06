package validators

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

var myValid *validator.Validate
var validatorError map[string]string


func init() {
	validatorError = make(map[string]string)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		myValid = v
	} else {
		log.Fatal("error validator")
	}
}

func registerValidation(tag string, fn validator.Func) {
	err := myValid.RegisterValidation(tag, fn)
	if err != nil {
		log.Fatalf("validator %s error", tag)
	}
}

func CheckErrors(errors error) {
	if errs, ok := errors.(validator.ValidationErrors); ok {
		for _, err := range errs {
			if v, ok :=  validatorError[err.Tag()]; ok {
				panic(fmt.Errorf(v))
			}
		}
	}
}
