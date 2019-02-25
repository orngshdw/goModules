package utility

import (
	"fmt"
	"os"
	"runtime"
)

// CheckIfError should be used to naively panics if an error is not nil.
// https://github.com/src-d/go-git/blob/662e2c226e9b8352a90cd1951233fab30a4e5042/_examples/common.go
func CheckIfError(err error) {
	if err == nil {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\x1b[31;1m%s\n", fmt.Sprintf("error: line %d, %v", line, err))
	fmt.Printf("\t\x1b[31;1m%s\x1b[31;0m\n", fmt.Sprintf("file \"%s\"", file))
	os.Exit(1)
}

// GetType prints out a variable's type
func GetType(i interface{}) {
	switch t := i.(type) {
	case string:
		fmt.Println("Variable type: 'string'")
	default:
		fmt.Printf("Variable type: '%T'\n", t)
	}
}
