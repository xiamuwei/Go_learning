package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func Test_fileDemo1(t *testing.T) {
	str := "./test"
	fileInfo, _ := os.Stat(str)
	fmt.Printf("name = %v\nSize = %v\nMode = %v\nIsDir = %v\n", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), fileInfo.IsDir())
	file, err := os.OpenFile(str, os.O_RDONLY, 0666)
	if err != nil {
		// fmt.Println(errors.New("cant open this file ..."))
		fmt.Println(" os.OpenFile err = ", err)
	}
	defer file.Close()

	buf := make([]byte, 16)
	var content []byte
	for {
		n, err1 := file.Read(buf)
		if err1 == io.EOF {
			// fmt.Println(errors.New("cant open this file ..."))
			fmt.Println("文件读取结束...")
			break
		}

		content = append(content, buf[:n]...)

	}

	fmt.Println(string(content))

}

func Test_fileDemo2(t *testing.T) {
	// 1、打开文件
	file, err := os.OpenFile("./test", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("os.OpenFile err = ", err)
		return
	}
	// 4、关闭文件
	defer file.Close()

	// 2、创建切片
	buf := make([]byte, 4)
	var content []byte
	for {
		// 3、 读取文件内容到buf
		n, err := file.Read(buf)
		if err == io.EOF {
			fmt.Println("读取文件结束...")
			break
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

func Test_fileDemo3(t *testing.T) {
	// 1 打开文件
	file, err := os.OpenFile("./test", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("os.OpenFile err = ", err)
	}

	// 记得及时关闭文件
	defer file.Close()

	//
	reader := bufio.NewReader(file)
	for {
		// reader.ReadString 读取直到第一次遇到\n ，返回一个包含已读取的数据和\n的字符串
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}

		fmt.Println(line)
	}
}

func Test_fileDemo4(t *testing.T) {

	content, err := ioutil.ReadFile("./test")
	if err != nil {
		fmt.Println("ioutil.ReadFile  err:", err)
	}
	fmt.Println(string(content))

}
