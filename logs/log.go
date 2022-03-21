package logs

import (
	"fmt"
	"os"
)

// LogError 打印 error 日志
func LogError(msg ...interface{}) {
	fmt.Sprintln(msg...)
	os.Exit(1)
}

// LogInfo 打印 info 日志
func LogInfo(msg ...interface{}) {
	fmt.Sprintln(msg...)
}

func LogErrorFormat(format string, msg ...interface{}) {
	fmt.Sprintf(format, msg...)
	os.Exit(1)
}
