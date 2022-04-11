package function

import (
	"fmt"
	"testing"
)

// 可变参数
func Test_varargus(t *testing.T) {
	variable_num(1, 2)
	variable_num(1, 2, 3, 4)
}

func variable_num(num ...int) {
	fmt.Printf("the type of num = %T \n", num)
	for k, v := range num {
		fmt.Printf("k = %v ,v = %v\n", k, v)
	}
}

// 返回值
func Test_calc(t *testing.T) {
	sum1, sub1 := calcDemo1(1, 2)
	fmt.Printf("sum1 = %v ,sub1 = %v\n", sum1, sub1)
	sum2, sub2 := calcDemo2(1, 2)
	fmt.Printf("sum2 = %v ,sub2 = %v\n", sum2, sub2)
}
func calcDemo1(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量
func calcDemo2(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

// defer
func Test_deferDemo1(t *testing.T) {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func Test_deferDemo2(t *testing.T) {
	defer fmt.Println(1)
	defer fmt.Println(2)
	panic("this is panic")
	defer fmt.Println(3)
	return
}

func Test_deferDemo3(t *testing.T) {
	a := 1
	defer fmt.Println(" a = ", a)
	a++
}

// 值传递 、引用传递
func Test_pass(t *testing.T) {
	// 值传递
	num := 1
	value_pass(num)
	fmt.Println("num = ", num)

	// 引用传递
	slice := []int{1, 2, 3}
	pointer_pass(slice)
	for k, v := range slice {
		fmt.Printf("k = %v ,v = %v\n", k, v)
	}
}

func value_pass(num int) {
	num++
	fmt.Println("num = ", num)
}

func pointer_pass(slice []int) {
	for k, v := range slice {
		v++
		slice[k] = v
		fmt.Printf("k = %v ,v = %v\n", k, v)
	}
}

// 匿名函数
func Test_anonymousFunction(t *testing.T) {
	func(a, b int) {
		fmt.Println("a + b = ", a+b)
	}(1, 2)

	var f = func() {
		fmt.Println("this is a anonymousFunction")
	}
	f()
}

// 闭包函数
func TestClosure(t *testing.T) {
	var f = adder()
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println(f(30))
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 递归函数
