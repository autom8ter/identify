package log

import "fmt"

func Debugln(args ...interface{}) {
	fmt.Println(args...)
}

func Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
