package logUtils

import (
	"fmt"
	"time"
)

//todo 日志信息
func Println(msg ...interface{}) {
	fmt.Println(append([]interface{}{time.Now().Format("2006-01-02 15:04:05")}, msg...)...)
}
