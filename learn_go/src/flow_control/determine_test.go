package flow_control

import (
	"fmt"
	"testing"
)

func Test_determine(t *testing.T) {
	var grade string
	var marks int = 90
	// if
	if marks == 90 {
		grade = "if A"
	} else if marks == 80 {
		grade = "if B"
	} else if marks == 50 || marks == 60 || marks == 70 {
		grade = "if C"
	} else {
		grade = "if D"
	}
	fmt.Printf("if grade = %v\n", grade)

	// switch
	switch marks {
	case 90:
		grade = "switch A"
	case 80:
		grade = "switch B"
	case 50, 60, 70:
		grade = "switch C"
	default:
		grade = "switch D"
	}

	fmt.Printf("switch grade = %v\n", grade)
}

func Test_circle(t *testing.T) {
	// for
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	// for-range
	var arr [5]string = [5]string{"hi", "hello", "world", "go", "!"}
	for k, v := range arr {
		fmt.Printf("k = %v , v = %v\n", k, v)
	}

	var i int
	// while
	for {
		if i > 10 {
			break
		}
		fmt.Println("while i = ", i)
		i++
	}

	var j int
	// do while
	for {
		fmt.Println("do while j = ", j)
		j++
		if j > 10 {
			break
		}
	}

}

// goto
func Test_goto(t *testing.T) {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Printf("i = %v,j = %v\n", i, j)
		}
	}

breakTag:
	fmt.Println("跳到goto flag")
}
