package _case

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type cusError struct {
	Code string
	Msg  string
	Time time.Time
}

func (err cusError) Error() string {
	return fmt.Sprintf("Code:%s,Msg:%s,Time:%s", err.Code, err.Msg, err.Time.Format("2006-01-02 15:04:05"))
}
func getCusError(code, msg string) error {
	return cusError{
		code,
		msg,
		time.Now(),
	}
}

func sum(a, b int) (int, error) {
	if a <= 0 && b <= 0 {
		return 0, getCusError("500", "两数求和不能小于等于0")
	}
	return a + b, nil
}
func ErrorsCase() {
	err := errors.New("程序发生了错误...")
	fmt.Println(err)

	var a, b = -1, -2
	res, err := sum(a, b)
	fmt.Println(res, err)
	//	断言
	if err != nil {
		log.Println(err)
		cusErr, ok := err.(cusError)
		if ok {
			fmt.Println("打印自定义错误信息:", cusErr.Code, cusErr.Msg, cusErr.Time)
		}
	}
}
