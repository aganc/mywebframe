package utils

import (
	"reflect"
	"runtime"
)

// 获取调用者的函数名字
func GetCallerName(a interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(a).Pointer()).Name()
}

// //获取本函数的名字
func PrintMyFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// 获取调用者的函数名字
func PrintCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
