package utils

import (
	"catvod/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func SaveClassJson(path string, vodClass []model.VodClass) (err error) {
	// 判断文件是否存在，不存在则创建
	var (
		exist   bool
		filePtr *os.File
	)
	exist, err = PathExists(path)
	if err != nil {
		return err
	}
	if !exist {
		dir := filepath.Dir(path)
		exist, err = PathExists(dir)
		if err != nil {
			return err
		}
		if !exist {
			if err = os.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
		}
		filePtr, err = os.Create(path)
		if err != nil {
			fmt.Println("Create file failed", err.Error())
			return
		}
	}
	defer func() {
		err = filePtr.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	// 带JSON缩进格式写文件，这种写法需要更多硬盘空间，但易于阅读
	data, err := json.MarshalIndent(vodClass, "", "  ")
	if err != nil {
		fmt.Println("Encoder failed", err.Error())
		return
	}
	_, err = filePtr.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func LoadClassJson(path string) (vodClass []model.VodClass) {
	filePtr, err := os.Open(path)
	if err != nil {
		fmt.Printf("Open file failed [Err:%s]\n", err.Error())
		return
	}
	defer func() {
		err = filePtr.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	bytes, err := ioutil.ReadAll(filePtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(bytes, &vodClass)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("new: %+v\n", vodClass)
	}
	return
}

// json转结构化通用方法
func LoadJson(path string, v interface{}) (vJson interface{}) {
	filePtr, err := os.Open(path)
	if err != nil {
		fmt.Printf("Open file failed [Err:%s]\n", err.Error())
		return
	}
	defer func() {
		err = filePtr.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	bytes, err := ioutil.ReadAll(filePtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("new: %+v\n", v)
	}
	vJson = v
	return vJson
}
