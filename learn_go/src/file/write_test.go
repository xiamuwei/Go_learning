package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func Test_writeDemo1(t *testing.T) {
	file, err := os.OpenFile("./write_test", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("os.OpenFile err = ", err)
		return
	}
	defer file.Close()

	file.Write([]byte("hello golang"))
	file.WriteString("hello test")
}

func Test_writeDemo2(t *testing.T) {
	// 打开文件
	file, err := os.OpenFile("./write_test", os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("os.OpenFile err = ", err)
	}
	defer file.Close()

	// 使用bufiosh生成writer
	writer := bufio.NewWriter(file)
	writer.WriteString("this is bufio writer")

	// 将缓存中的内容写入到文件
	writer.Flush()
}

func Test_writeDemo3(t *testing.T) {
	err := ioutil.WriteFile("./write_test", []byte("happy"), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func Test_CopyFile(t *testing.T) {
	copy_file("./test", "./copy_test")
}

func copy_file(srcName, dstName string) {
	// 打开源文件
	srcFile, err := os.OpenFile(srcName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(" srcName os.OpenFile err = ", err)
	}
	defer srcFile.Close()

	// 打开目标文件
	dstFile, err := os.OpenFile(dstName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(" dstName os.OpenFile err = ", err)
	}
	defer dstFile.Close()

	io.Copy(dstFile, srcFile)
}
