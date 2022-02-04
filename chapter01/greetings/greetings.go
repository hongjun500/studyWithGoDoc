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

// Hellos 入参一个切片， 返回一个map和err
func Hellos(names []string) (map[string]string, error) {

	//定义一个名称和消息关联的map
	message := make(map[string]string)

	for _, name := range names {
		msg, err := Hello(name)
		//如果有错误
		if err != nil {
			return nil, err
		}

		//将名称和消息关联绑定
		message[name] = msg
	}
	return message, nil
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
