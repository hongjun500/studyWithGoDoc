package greetings

import (
	"regexp"
	"testing"
)

// 调用greetings的Hello方法检查是否有有效的返回值
func TestHelloName(t *testing.T) {
	name := "hongjun"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("hongjun")
	if !want.MatchString(name) || err != nil {
		t.Fatalf(`Hello("hongjun") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// 检查错误
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want"", err`, msg, err)
	}
}
