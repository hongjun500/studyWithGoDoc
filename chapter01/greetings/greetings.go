package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello 传入一个字符串参数返回另一个字符串和错误(如果有)
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("空名")
	}
	//return fmt.Sprintf("Hi, %v. Welcome!", randomFormat()), nil
	return fmt.Sprintf(randomFormat(), name), nil
}

// init 初始化一个变量值
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {

	//定义一个切片
	formats := []string{
		"Hi, %v, Welcome",
		"Great to see you, %v",
		"Hail, %v, Well met !",
	}
	//返回随机一个在索引下标值范围内的数据
	return formats[rand.Intn(len(formats))]
}
