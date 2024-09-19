package Set

import (
	"fmt"
	"testing"
)

func TestNewSet(t *testing.T) {
	intSet := NewSet[int, int](func(i int) int { return i })
	intSet.Add(1)
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)
	fmt.Println(intSet.Values())

	stringSet := NewSet[string, string](func(s string) string { return s })
	stringSet.Add("aaa")
	stringSet.Add("aaa")
	stringSet.Add("ccc")
	stringSet.Add("bbb")
	fmt.Println(stringSet.Values())

	type Student struct {
		Id   int
		Name string
	}
	studentSet := NewSet[int, Student](func(s Student) int { return s.Id })
	studentSet.Add(Student{
		Id:   1,
		Name: "a",
	})
	studentSet.Add(Student{
		Id:   2,
		Name: "b",
	})
	studentSet.Add(Student{
		Id:   1,
		Name: "c",
	})
	fmt.Println(studentSet.Values())
}
