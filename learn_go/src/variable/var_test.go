package variable_test

import (
	"fmt"
	"testing"
)

const PI float32 = 3.14

// 全局变量
var v1 int = 1

// 自动类型推断
var v2 = "string"
var (
	v3 = 1
	v4 = 2
	v5 = 3
)

func Test_variable(t *testing.T) {
	// 局部变量还可以采取短变量声明方式
	v6 := 3.15
	fmt.Printf("v1= %v,v2 = %v,v3 = %v,v4 = %v,v5 = %v,v6 = %v\n", v1, v2, v3, v4, v5, v6)
	fmt.Println("this is test function")
}
