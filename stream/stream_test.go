package stream

import (
	"fmt"
	"testing"
)

func TestNewStream(t *testing.T) {
	is := []int{1, 2, 3, 4, 5}
	res := NewStream(is).
		Filter(func(orign int) bool {
			if orign%2 == 0 {
				return true
			}
			return false
		}).
		Map(func(origin int) int {
			return origin + 1
		}).
		Collect()
	fmt.Println(res)
}
