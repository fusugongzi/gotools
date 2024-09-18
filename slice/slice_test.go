package slice

import "testing"

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
