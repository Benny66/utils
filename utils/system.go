
package function

import (
	"net"
)

func GetLocalIp() (ip string) {
	ip = "127.0.0.1"

	addressList, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	//优先获取IPV4地址，存在多个IP则取第一个
	for _, address := range addressList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				return
			}
		}
	}

	//如果IPV4地址不存在，则获取IPV6地址，存在多个IP则取第一个
	for _, address := range addressList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To16() != nil {
				ip = ipNet.IP.String()
				return
			}
		}
	}

	//默认127.0.0.1
	return ip
}
