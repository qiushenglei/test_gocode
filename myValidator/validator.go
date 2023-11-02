package myValidator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type DateTime struct {
	Time time.Time `validate:"datetime=2006-01-02 15:04:05"`
}

func Validate() {

	// 时间字符串
	str := "2023-01-01 12:00:00"

	// 创建一个实例
	dt := &DateTime{}

	// 使用 time.Parse 解析时间字符串
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("时间解析出错:", err)
		return
	}
	dt.Time = parsedTime

	// 参数校验
	v := validator.New()
	err = v.Struct(dt)
	if err != nil {
		fmt.Println("参数校验不通过:", err)
		return
	}

	fmt.Println("校验通过")

}
