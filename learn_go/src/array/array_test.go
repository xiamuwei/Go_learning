package array

import (
	"fmt"
	"testing"
)

func Test_arrray(t *testing.T) {
	// 数组声明
	arr1 := [5]int{1, 2, 3, 4, 5}
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}

	// for-range遍历
	for k, v := range arr1 {
		fmt.Printf("k = %v ,v = %v\n", k, v)
	}
	for k, v := range arr2 {
		fmt.Printf("k = %v ,v = %v\n", k, v)
	}

}
