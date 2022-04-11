package panic_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestForPanic(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover err = ", err)
		}
	}()
	var a = 10
	var b = 0
	var c = a / b
	fmt.Println("c = ", c)
	fmt.Println("this is a test")
}

func TestForError(t *testing.T) {
	err := readConf("config2.ini")
	if err != nil {
		panic(err)
	}
}

func readConf(str string) error {
	if str == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件失败...")
	}
}
