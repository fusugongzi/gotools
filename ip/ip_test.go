package ip

import (
	"fmt"
	"testing"
)

func TestGetLocalIps(t *testing.T) {
	fmt.Println(GetLocalIps())
}

func TestGetLocalIp(t *testing.T) {
	fmt.Println(GetLocalIp())
}
