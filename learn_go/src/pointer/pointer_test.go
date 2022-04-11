package pointer_tests

import (
	"fmt"
	"testing"
)

func Test_pointer(t *testing.T) {
	var num int = 5
	// 使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针类型
	// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值
	var ptr *int = &num
	fmt.Println(*ptr)
}
