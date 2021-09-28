package utils

import (
	"github.com/go-playground/validator"
	"regexp"
)

func Mobile(fl validator.FieldLevel) bool {

	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, fl.Field().String())

	return ok

}
