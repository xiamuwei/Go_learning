package struct_test

import (
	"fmt"
	"testing"
)

// 定义
type Person struct {
	name, city string
	age        int
}

func (p Person) SetAge(newAge int) {
	p.age = newAge
}

func (p *Person) SetAgePointer(newAge int) {
	p.age = newAge
}
func Test_struct(t *testing.T) {
	// 实例化 var 结构体实例 结构体类型
	var p1 Person
	// . 赋值
	p1.name = "jack"
	p1.city = "New York"
	p1.age = 18
	fmt.Println(p1)

	// 使用键值对初始化
	var p2 = Person{
		name: "candy",
		city: "Canada",
	}
	fmt.Println(p2)

	// 直接使用值对初始化简写，此时必须初始化结构体的所有字段，且初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	var p3 = Person{
		"candy",
		"Canada",
		18,
	}
	fmt.Println(p3)

	p1.SetAge(30)
	fmt.Println(p1)
	p1.SetAgePointer(30)
	fmt.Println(p1)
}

type MyInt int

func (i MyInt) sayHello() {
	fmt.Println("this is ", i)
}

func Test_int(t *testing.T) {
	var a MyInt = 1
	a.sayHello()
}

type Address struct {
	Province string
	City     string
}

type User struct {
	Name, Gender string
	Address
}

func Test_anonymous(t *testing.T) {
	user := User{
		Name:   "jack",
		Gender: "boy",
		Address: Address{
			Province: "New York",
			City:     "no",
		},
	}

	user.Address.Province = "山东" // 匿名字段默认使用类型名作为字段名
	user.City = "威海"             // 匿名字段可以省略
	fmt.Println(user)
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%v 会动\n", a.name)
}

type Dog struct {
	Feet    int
	*Animal // 通过匿名结构体实现继承
}

func (d *Dog) bark() {
	fmt.Printf("%v会汪汪汪~\n", d.name)
}

func Test_extends(t *testing.T) {
	dog := Dog{
		Feet: 4,
		Animal: &Animal{
			name: "little",
		},
	}

	dog.bark()
	dog.move()
}
