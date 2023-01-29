package common

import (
	"reflect"
	"runtime"
)

// GetFunctionName returns given functions name in format NameOfPackage.FunctionName
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
