package utils

import (
	"bufio"
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteFileWithBuffer(path string, str string) error {
	//os.O_WRONLY|os.O_CREATE表示只写模式打开，如果打开不成功，就创建。
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
		return err
	}
	defer f.Close()
	//使用带缓存的writer
	writer := bufio.NewWriter(f)
	size, err := writer.WriteString(str) //可使用Write直接写[]byte，而无需转换成string
	if err != nil {
		fmt.Println("写入缓存失败")
		return err
	}
	fmt.Println("写入大小(bytes)：", size)
	err = writer.Flush() //将文件真正写到文件中
	if err != nil {
		fmt.Println("写入文件失败")
		return err
	}
	return nil
}
