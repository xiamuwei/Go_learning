package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type MyInt int

func Test_reflect(t *testing.T) {
	var a int
	reflectType(a)

	var myint MyInt
	reflectType(myint)

	var b rune
	reflectType(b)

	var pointer *int
	reflectType(pointer)
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func Test_reflect_value(t *testing.T) {

	var num int = 2
	reflectSetValue(&num)
	fmt.Println("num = ", num)

	var a float32 = 3.14
	reflectValue(a)
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c)
}
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()

	switch k {
	case reflect.Int32:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int32(v.Int()))
	case reflect.Bool:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %v\n", bool(v.Bool()))
	default:
		fmt.Printf("cant know the type\n")
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}

}
