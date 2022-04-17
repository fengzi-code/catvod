package utils

import (
	"github.com/go-resty/resty/v2"
	"net"
)

/*
IP接口：
1. https://api.ipify.org?format=json 或直接访问https://api.ipify.org
2. http://members.3322.org/dyndns/getip
*/
func GetExternalIp() (res string) {
	client := resty.New()
	resp, err := client.R().Get("http://members.3322.org/dyndns/getip")
	if err != nil {
		return
	}
	if resp.StatusCode() != 200 {
		return
	}
	// 使用正则判断是否是IP地址
	ip := net.ParseIP(resp.String())
	if ip == nil {
		return
	}
	res = resp.String()
	return
}

// 获取本机内网IP
func GetInternalIp() (res string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				res = ipnet.IP.String()
				return
			}
		}
	}
	return
}
