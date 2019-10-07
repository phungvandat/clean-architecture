package service

import (
	"reflect"
)

// Compose applies middlewares to Service.
// loop middlewares in reverse order to make them run in the order in which
// they are specified. ie.
// `Compose(s, logging, validation)` would first applies `logging` then `validation`.
func Compose(s interface{}, mws ...interface{}) interface{} {
	for i := len(mws) - 1; i >= 0; i-- {
		vv := reflect.ValueOf(mws[i]).Call([]reflect.Value{reflect.ValueOf(s)})
		s = vv[0].Interface()
	}
	return s
}
