package utils

import (
	"fmt"
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
	res = GetIpFromUdpNetDial()
	return
}

/*
GetIpFromUdpNetDial
获取到当前机器的主IP，由于使用的是UDP协议，不会需要握手，所以不需要担心访问不到8.8.8.8:8
或者这样说，这里随便填一个IP地址都可以，反正不会去访问
*/
func GetIpFromUdpNetDial() (res string) {
	conn, err := net.Dial("udp", "8.8.8.8:8")
	if err != nil {
		fmt.Printf("net.Dial err: %v\n", err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	if localAddr.IP.To4() == nil {
		return
	}
	res = localAddr.IP.To4().String()
	return
}
