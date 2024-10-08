package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		items []int
		item  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "items nil",
			args: args{
				items: nil,
				item:  1,
			},
			want: false,
		},
		{
			name: "not contains",
			args: args{
				items: []int{1, 2},
				item:  3,
			},
			want: false,
		},
		{
			name: "contains",
			args: args{
				items: []int{1, 2, 5, 8, 10},
				item:  5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.items, tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	type args struct {
		items []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{
				items: nil,
			},
			want: false,
		},
		{
			name: "contain equal",
			args: args{
				items: []string{"a", "b", "c"},
				item:  "a",
			},
			want: true,
		},
		{
			name: "contain ignore case equal",
			args: args{
				items: []string{"abc", "bsA", "cde"},
				item:  "ABc",
			},
			want: true,
		},
		{
			name: "not contain ignore case",
			args: args{
				items: []string{"abc", "bsA", "cde"},
				item:  "ccc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsIgnoreCase(tt.args.items, tt.args.item); got != tt.want {
				t.Errorf("ContainsIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type Student struct {
		Name  string
		Score int
	}
	students := make([]Student, 0)
	students = append(students, Student{
		Name:  "aa",
		Score: 1,
	})
	students = append(students, Student{
		Name:  "bb",
		Score: 2,
	})
	m := ToMap(students, func(s Student) string {
		return s.Name
	})
	fmt.Println(m)
}

func TestFliter(t *testing.T) {
	s := []string{"abc", "a", "c", "a", "d"}
	fmt.Println(Fliter(s, func(v string) bool {
		return v != "a"
	}))
}

func TestDistinct(t *testing.T) {
	is := Distinct([]int{1, 2, 2, 3, 4, 3, 1})
	fmt.Println(is)

	ss := Distinct([]string{"aa", "bb", "cc", "bb"})
	fmt.Println(ss)
}

func TestDistinctI(t *testing.T) {
	type Student struct {
		Name  string
		Score int
	}
	students := make([]Student, 0)
	students = append(students, Student{
		Name:  "aa",
		Score: 1,
	})
	students = append(students, Student{
		Name:  "bb",
		Score: 2,
	})
	students = append(students, Student{
		Name:  "aa",
		Score: 3,
	})
	fmt.Println(DistinctI(students, func(s Student) string { return s.Name }))
}

func TestSearch(t *testing.T) {
	type args struct {
		items  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not exist",
			args: args{
				items:  []int{3, 1, 2, 8},
				target: 12,
			},
			want: -1,
		},
		{
			name: "exist",
			args: args{
				items:  []int{3, 1, 2, 8},
				target: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.items, tt.args.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	type Student struct {
		Name  string
		Score int
	}
	students := make([]Student, 0)
	students = append(students, Student{
		Name:  "aa",
		Score: 1,
	})
	students = append(students, Student{
		Name:  "bb",
		Score: 2,
	})
	students = append(students, Student{
		Name:  "aa",
		Score: 3,
	})
	ss := GroupBy(students, func(s Student) string {
		return s.Name
	})
	fmt.Println(ss)
}

func TestMerge(t *testing.T) {
	got := Merge([]int{1, 2, 4}, []int{3, 5})
	fmt.Println(got)
}

func TestRandom(t *testing.T) {
	fmt.Println(Random([]int{1, 2, 3, 4, 5}))
	fmt.Println(Random([]int{1, 2, 3, 4, 5}))
	fmt.Println(Random([]int{1, 2, 3, 4, 5}))
	fmt.Println(Random([]int{1, 2, 3, 4, 5}))
	fmt.Println(Random([]int{1, 2, 3, 4, 5}))
}

func TestRandoms(t *testing.T) {
	fmt.Println(Randoms([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(Randoms([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(Randoms([]int{1, 2, 3, 4, 5}, 3))
}

func TestMapToFixed(t *testing.T) {
	got := MapToFixed([]int{1, 2, 3}, "abc")
	fmt.Println(got)
}

func TestOffset(t *testing.T) {
	type args struct {
		items  []int
		value  int
		offset int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not exist",
			args: args{
				items:  []int{1, 2, 3, 4},
				value:  5,
				offset: 5,
			},
			want: 5,
		},
		{
			name: "exist",
			args: args{
				items:  []int{1, 2, 3, 4, 5},
				value:  2,
				offset: 2,
			},
			want: 4,
		},
		{
			name: "exist",
			args: args{
				items:  []int{1, 2, 3, 4, 5},
				value:  5,
				offset: 1,
			},
			want: 1,
		},
		{
			name: "exist",
			args: args{
				items:  []int{1, 2, 3, 4, 5},
				value:  5,
				offset: 21,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Offset(tt.args.items, tt.args.value, tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}
