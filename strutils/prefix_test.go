package strutils

import (
	"fmt"
	"testing"
)

func TestPrefixSearch_Search(t *testing.T) {
	ps := NewPrefixSearch([]string{"河北省", "北京市", "内蒙古自治区", "广西壮族自治区"})
	fmt.Println(ps.Search("河北省石家庄市区"))
	fmt.Println(ps.Search("内蒙古自治区"))
}
