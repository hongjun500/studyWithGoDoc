package greetings

import (
	"errors"
	"fmt"
)

// Hello 传入一个字符串参数返回另一个字符串和错误(如果有)
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("空名")
	}
	return fmt.Sprintf("Hi, %v. Welcome!", name), nil
}
