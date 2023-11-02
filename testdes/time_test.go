package testdes_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T)  {
	timeString := "2018-01-02T15:04:05"
	dateTime, err := time.Parse(time.RFC3339, timeString)
	// 如果有错误就不转换直接返回
	if err != nil {
		panic(err)
	}
	dateTimeString := dateTime.Format("2006-01-02 15:04:05")
	fmt.Println(dateTimeString)
}
