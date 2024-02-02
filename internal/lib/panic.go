package lib

import (
	"fmt"
	"runtime/debug"
)

func RecoverPanic() {
	if panicValue := recover(); panicValue != nil {
		fmt.Printf("Recovered from panic: %v\nStack: %v\n", panicValue, string(debug.Stack()))
	}
}
