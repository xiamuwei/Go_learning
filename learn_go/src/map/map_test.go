package map_test

import (
	"fmt"
	"testing"
)

func Test_map(t *testing.T) {
	// 创建map
	var map1 map[string]string
	// 使用之前一定要make初始化一下
	map1 = make(map[string]string, 5)
	map1["a"] = "A"
	map1["b"] = "B"
	map1["c"] = "C"

	// 遍历
	for k, v := range map1 {
		fmt.Printf("k = %v , v = %v\n", k, v)
	}

	// 增删改查
	// 增改
	map1["d"] = "D"
	map1["a"] = "A~"
	for k, v := range map1 {
		fmt.Printf("k = %v , v = %v\n", k, v)
	}

	// 删除
	delete(map1, "a")
	for k, v := range map1 {
		fmt.Printf("k = %v , v = %v\n", k, v)
	}

	// 查询
	res, ok := map1["b"]
	if !ok {
		fmt.Println("没有找到查询结果")
	} else {
		fmt.Println("查询结果 res = ", res)
	}
}
