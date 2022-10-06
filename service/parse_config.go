package service

import (
	"catvod/global"
	"catvod/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func GetConfig(path string) (res model.ServerConfig) {
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
	}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		fmt.Println(err)
	} else {
		// 将Api里的地址替换成本机地址
		if global.BindAddr != "" {
			for i, v := range res.Sites {
				// 正则匹配ip:port
				regex := `(?P<ip>[\d\.]+):(?P<port>\d+)`
				//
				reg := regexp.MustCompile(regex)
				// 旧地址：IP:port
				var oldAddr string
				if reg.MatchString(v.Api) {
					match := reg.FindStringSubmatch(v.Api)
					oldAddr = match[0]
					if oldAddr == "103.222.188.156:99" {
						continue
					}
					newApi := strings.Replace(v.Api, oldAddr, global.BindAddr, -1)
					res.Sites[i].Api = newApi
				}
			}
		}
		fmt.Printf("%+v\n", res)
	}
	return
}
