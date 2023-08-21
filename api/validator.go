package api

import (
	"github.com/BerdiyorovAbrorjon/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrecy validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportCurrency(currency)
	}
	return false
}
