package interface_test

import (
	"fmt"
	"testing"
)

// 定义一个接口
type Usb interface {
	start()
	stop()
}

type Phone struct{}

func (p Phone) start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct{}

func (c Camera) start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct{}

// 接收一个Usb类型的变量，只要实现Usb接口都可。usb变量会根据传入的实参，判断到底是Phone类型，还是Camera类型，再调用相应的方法实现
func (c Computer) Working(usb Usb) {
	usb.start()
	usb.stop()
}

func Test_interface(t *testing.T) {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)

	var a interface{}
	var num float64 = 3.14
	a = num
	if b, ok := a.(float64); ok {
		fmt.Println(b)
	} else {
		fmt.Println("转换错误...")
	}
}
