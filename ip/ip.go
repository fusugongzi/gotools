package ip

import (
	"errors"
	"fmt"
	"net"
	"sort"
)

func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	result := make([]string, 0)
	for _, addr := range addrs {
		// 检查地址是否为 IP 地址并排除环回地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			// 只显示 IPv4 地址
			if ipNet.IP.To4() != nil {
				result = append(result, ipNet.IP.To4().String())
			}
		}
	}
	if len(result) == 0 {
		return "", errors.New("no valid ip")
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result[0], nil
}

func GetLocalIps() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	result := make([]string, 0)
	for _, addr := range addrs {
		// 检查地址是否为 IP 地址并排除环回地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			// 只显示 IPv4 地址
			if ipNet.IP.To4() != nil {
				result = append(result, ipNet.IP.To4().String())
			}
		}
	}
	return result, nil
}
