package strutils

import "testing"

func TestIsBlank(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "blank",
			args: args{
				s: "   ",
			},
			want: true,
		},
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "not blank",
			args: args{
				s: " s ",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.args.s); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}
