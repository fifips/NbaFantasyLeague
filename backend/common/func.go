package common

import (
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

// GetFunctionName returns given functions name in format NameOfPackage.FunctionName
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// randomString generates random string (of characters from a-zA-Z0-9) of given length.
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}
