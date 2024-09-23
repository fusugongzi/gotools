package maps

import (
	"fmt"
	"testing"
)

func TestKeySet(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(KeySet(m1))

	m2 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	fmt.Println(KeySet(m2))
}

func TestValueSet(t *testing.T) {
	type Student struct {
		Id   int
		Name string
	}
	m := map[int]Student{
		1: Student{
			Id:   1,
			Name: "a",
		},
		2: Student{
			Id:   2,
			Name: "b",
		},
		3: Student{
			Id:   1,
			Name: "a",
		},
	}
	fmt.Println(ValueSet(m, func(s Student) int { return s.Id }))
}

func TestMerge(t *testing.T) {
	got := Merge(map[int]int{1: 1, 3: 3}, map[int]int{2: 2})
	fmt.Println(got)
}
