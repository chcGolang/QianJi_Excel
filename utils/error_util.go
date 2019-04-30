package utils

import "fmt"

func Errorf(format string, a ...interface{})error  {
	return fmt.Errorf(format,a...)
}
