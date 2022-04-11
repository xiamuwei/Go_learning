package date_type

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_type(t *testing.T) {
	// 整型
	var int_var int32 = 52
	// 浮点型
	var float_var float64 = 3.14
	// 布尔类型
	var bool_var bool = false
	// 字符型 byte
	var char_var byte = 'C'
	fmt.Printf("int_var = %T\nfloat_var = %T\nbool_var = %T\nchar_var = %T\n", int_var, float_var, bool_var, char_var)

	// 字符串型
	str_var := "hello golang"
	fmt.Printf("str_var = %T\n", str_var)
}

// 类型转换,golang不支持类型自动转换，必须显式转换
// 基本数据类型之间的类型转换
func Test_dataTypeConversion(t *testing.T) {
	// 整型
	// var int_var int32 = 52
	// // 浮点型
	// var float_var float64 = 3.14
	// 布尔类型
	// var bool_var bool = false
	// 字符型 byte
	var char_var byte = 'C'

	v1 := int32(char_var)
	// v2 := int32(bool_var) cannot convert bool_var (type bool) to type int32
	fmt.Printf("v1 = %v\n", v1)
}

func Test_basicIntoString(t *testing.T) {
	// 整型
	var int_var int32 = 52
	// 浮点型
	var float_var float64 = 3.14
	// 布尔类型
	var bool_var bool = false
	// 字符型 byte
	// var char_var byte = 'C'

	// 方式一 ： 使用fmt.Sprintf
	// str_int := fmt.Sprintf("%v", int_var)
	// str_float := fmt.Sprintf("%v", float_var)
	// str_bool := fmt.Sprintf("%v", bool_var)
	// str_char := fmt.Sprintf("%v", char_var)

	// 方式二：使用strconv里的Format函数
	str_int := strconv.FormatInt(int64(int_var), 10)
	str_float := strconv.FormatFloat(float_var, 'e', 2, 64)
	str_bool := strconv.FormatBool(bool_var)
	fmt.Printf("str_int = %v\nstr_float = %v\nstr_bool = %v\n",
		str_int, str_float, str_bool)
	fmt.Printf("str_int = %T\nstr_float = %T\nstr_bool = %T\n",
		str_int, str_float, str_bool)

}

func Test_stringIntoBasic(t *testing.T) {
	str1 := "1"
	str2 := "3.14"
	str3 := "false"
	str_int, _ := strconv.ParseInt(str1, 10, 10)
	str_float, _ := strconv.ParseFloat(str2, 64)
	str_bool, _ := strconv.ParseBool(str3)
	fmt.Printf("str_int = %v\nstr_float = %v\nstr_bool = %v\n",
		str_int, str_float, str_bool)
	fmt.Printf("str_int = %T\nstr_float = %T\nstr_bool = %T\n",
		str_int, str_float, str_bool)
}
