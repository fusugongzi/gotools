package objutils

import (
	"fmt"
	"github.com/fusugongzi/gotools/conv"
	"testing"
)

func TestSameType(t *testing.T) {
	type student struct {
		Id int
	}
	s1 := &student{Id: 1}
	s2 := &student{}
	CopyProperties(s1, s2)
	fmt.Println(s2)
}

func TestDifferenceType(t *testing.T) {
	type source struct {
		Id int
	}
	type target struct {
		Id int
	}
	s1 := &source{Id: 1}
	s2 := &target{}
	CopyProperties(s1, s2)
	fmt.Println(s2)
}

func TestDifferenceField(t *testing.T) {
	type source struct {
		Id int
	}
	type target struct {
		Id1 int
	}
	s1 := &source{Id: 1}
	t1 := &target{}
	CopyProperties(s1, t1)
	fmt.Println(t1)
}

func TestPtrField(t *testing.T) {
	type source struct {
		Id *int
	}
	type target struct {
		Id *int
	}
	s1 := &source{Id: conv.IntPtr(1)}
	t1 := &target{}
	CopyProperties(s1, t1)
	s1.Id = conv.IntPtr(2)
	fmt.Println(conv.PtrInt(s1.Id, 0))
	fmt.Println(conv.PtrInt(t1.Id, 0))
}

func TestStructField(t *testing.T) {
	type c struct {
		Id int
	}
	type source struct {
		c c
	}
	type target struct {
		c c
	}
	s1 := &source{c: c{1}}
	t1 := &target{}
	CopyProperties(s1, t1)
	fmt.Println(s1.c.Id)
	fmt.Println(t1.c.Id)
}

func TestDifferentType(t *testing.T) {
	type source struct {
		Id int
	}
	type target struct {
		Id string
	}
	s1 := &source{Id: 1}
	t1 := &target{}
	CopyProperties(s1, t1)
	fmt.Println(s1.Id)
	fmt.Println(t1.Id)
}
